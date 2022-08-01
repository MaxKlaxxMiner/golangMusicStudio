(module
(memory $mem 1)
(export "mem" (memory $mem))

;; version() int
(func (export "version") (result i32)
  i32.const 10004
)

;; noise(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "noise") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  loop $loop
    ;; buf[i] = int32(ofs<<(gmconst.SampleBits-gmconst.DynamicBits-4)) >> (gmconst.SampleBits - gmconst.DynamicBits)
    ;; buf[i] = int32(ofs<<4)>>8
    (i32.store                                                                 ;; buf[i] = int32(ofs<<4)>>8
     (local.get $buf)                                                          ;; buf[i]
     (i32.shr_s (i32.shl (local.get $ofs) (i32.const 4)) (i32.const 8))        ;; int32(ofs<<4)>>8
    )

    (local.set $buf (i32.add (local.get $buf) (i32.const 4)))                  ;; i++ / buf++

    ;; ofs = (ofs+incr)*214013 + 2531011
    (local.set $ofs (i32.add (i32.mul (i32.add (local.get $ofs) (local.get $incr)) (i32.const 214013)) (i32.const 2531011)))

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end
  (local.get $ofs)
)

;; squareLQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "squareLQ") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  loop $loop
    ;; buf[i] = (1 << (gmconst.DynamicBits - 1)) - int32(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits)
    ;; buf[i] = 8388608 - (ofs>>31<<24)
    (i32.store
      (local.get $buf)                                                         ;; buf[i]
      (i32.sub
        (i32.const 8388608)                                                    ;; 8388608 - ...
        (i32.shl (i32.shr_u (local.get $ofs) (i32.const 31)) (i32.const 24))   ;; ofs>>31<<24
      )
    )

    (local.set $buf (i32.add (local.get $buf) (i32.const 4)))                  ;; i++ / buf++

    ;; ofs += incr
    (local.set $ofs (i32.add (local.get $ofs) (local.get $incr)))

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end

  local.get $ofs
)

;; squareHQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "squareHQ") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  (local $curVal i32)                                                          ;; int curVal = 0;
  (local $incr1 i64)                                                           ;; long incr1 = 0;
  (local $nextOfs i32)
  (local $swapbit i32)
  (local $b23 i32)
  (local.set $b23 (i32.const 8388608))

  ;; curVal = int32(1<<(gmconst.DynamicBits-1) - int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits))
  ;; curVal = 8388608 - (ofs>>31<<24)
  (local.set $curVal
    (i32.sub
      (local.get $b23)                                                         ;; 8388608 - ...
      (i32.shl (local.tee $swapbit (i32.shr_u (local.get $ofs) (i32.const 31))) (i32.const 24)) ;; ofs>>31<<24
    )
  )
  ;; incr1 = (1 << (gmconst.DynamicBits + gmconst.DynamicBits)) / uint64(incr)
  ;; incr1 = 281474976710656 / uint64(incr)
  (local.set $incr1 (i64.div_u (i64.const 281474976710656) (i64.extend_i32_u (local.get $incr))))

  loop $loop
    ;; if ofs>>(gmconst.SampleBits-1) == (ofs+incr)>>(gmconst.SampleBits-1)
    ;; if ofs>>31 == (ofs+incr)>>31
    (i32.eq
      (local.get $swapbit)                                                     ;; ofs>>31
      (local.tee $swapbit (i32.shr_u (local.tee $nextOfs (i32.add (local.get $ofs) (local.get $incr))) (i32.const 31))) ;; swapbit = (nextOfs = (ofs+incr))>>31
    )
    (if
      (then
        ;; buf[i] = curVal
        (i32.store (local.get $buf) (local.get $curVal))
      )
      (else
        ;; curVal = int32((uint64(1<<(gmconst.SampleBits-1) - (ofs & (1<<(gmconst.SampleBits-1) - 1))) * incr1) >> gmconst.DynamicBits)
        ;; curVal = int32((uint64(2147483648 - (ofs & 2147483647)) * incr1) >> 24)
        (local.set $curVal
          (i32.wrap_i64                                                        ;; int32((uint64(2147483648 - (ofs & 2147483647)) * incr1) >> 24)
            (i64.shr_u                                                         ;; (uint64(2147483648 - (ofs & 2147483647)) * incr1) >> 24
              (i64.mul                                                         ;; uint64(2147483648 - (ofs & 2147483647)) * incr1
                (i64.extend_i32_u                                              ;; uint64(2147483648 - (ofs & 2147483647))
                  (i32.sub                                                     ;; 2147483648 - (ofs & 2147483647)
                    (i32.const 2147483648)                                     ;; 2147483648
                    ;;(i32.and (local.get $ofs) (i32.const 2147483647))        ;; ofs & 2147483647
                    (i32.shr_u (i32.shl (local.get $ofs) (i32.const 1)) (i32.const 1)) ;; ofs & 2147483647 --> uint32(ofs)<<1>>1
                  )
                )
                (local.get $incr1)
              )
              (i64.const 24)
            )
          )
        )

        ;; buf[i] = (ofs>>(gmconst.SampleBits-1) != 0)
        ;;          ? (curVal = 1<<(gmconst.DynamicBits-1) - curVal)
        ;;          : (curVal - 1<<(gmconst.DynamicBits-1))
        ;; buf[i] = ofs>>31 ? 8388608 - curVal : curVal - 8388608
        (i32.store (local.get $buf)
          (select
            (i32.sub (local.get $curVal) (local.get $b23))
            (i32.sub (local.get $b23) (local.get $curVal))
            (local.get $swapbit)
          )
        )

        ;; curVal = int32(1<<(gmconst.DynamicBits-1) - int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits))
        ;; curVal = 8388608 - (nextOfs>>31<<24)
        (local.set $curVal
          (i32.sub
            (local.get $b23)                                                   ;; 8388608 - ...
            (i32.shl (local.get $swapbit) (i32.const 24))                      ;; nextOfs>>31<<24
          )
        )
      )
    )

    ;; ofs += incr
    (local.set $ofs (local.get $nextOfs))

    (local.set $buf (i32.add (local.get $buf) (i32.const 4)))                  ;; i++ / buf++

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end

  local.get $ofs
)

;; convertIntSamplesToFloat32(floats *float, ints *int, sampleCount uint)
(func (export "convertIntSamplesToFloat32") (param $floats i32) (param $ints i32) (param $sampleCount i32)
  loop $loop
    (f32.store                                                                 ;; *floats = (float)(*ints) * (1.0 / 8388608.0)
      (local.get $floats)                                                      ;; *floats
      (f32.mul                                                                 ;; (float)(*ints) * (1.0 / 8388608.0)
        (f32.convert_i32_s (i32.load (local.get $ints)))                       ;; (float)(*ints)
        (f32.const 1.1920928955078125e-07)                                     ;; 1.0 / 8388608.0
      )
    )

    (local.set $floats (i32.add (local.get $floats) (i32.const 4)))            ;; floats++
    (local.set $ints (i32.add (local.get $ints) (i32.const 4)))                ;; ints++

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end
)

;; volumeUpdate(buf *int, sampleCount uint, vol int)
(func (export "volumeUpdate") (param $buf i32) (param $sampleCount i32) (param $vol i32)
  (local $vol64 i64)                                                           ;; long vol64 = 0;
  (local.set $vol64 (i64.extend_i32_s (local.get $vol)))                       ;; vol64 = (long)vol

  loop $loop
    ;; buf[i] = int32(((long)buf[i] * (long)vol) >> (gmconst.DynamicBits - 1))
    ;; buf[i] = int32(((long)buf[i] * (long)vol) >> 23)
    (i64.store32                                                               ;; buf[i] = int(((long)buf[i] * (long)vol) >> 23)
      (local.get $buf)                                                         ;; buf[i]
      (i64.shr_s                                                               ;; ((long)buf[i] * (long)vol) >> 23
        (i64.mul                                                               ;; (long)buf[i] * (long)vol
          (i64.load32_s (local.get $buf))                                      ;; (long)buf[i]
          (local.get $vol64)                                                   ;; (long)vol
        )
        (i64.const 23)
      )
    )

    (local.set $buf (i32.add (local.get $buf) (i32.const 4)))                  ;; i++ / buf++

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end
)

;; mix(buf *int, src1 *int, src2 *int, sampleCount uint)
(func (export "mix") (param $buf i32) (param $src1 i32) (param $src2 i32) (param $sampleCount i32)
  loop $loop
    ;; buf[i] = src1[i] + src2[i]
    (i32.store
      (local.get $buf)                                                         ;; buf[i]
      (i32.add                                                                 ;; src1[i] + src2[i]
        (i32.load (local.get $src1))                                           ;; src1[i]
        (i32.load (local.get $src2))                                           ;; src2[i]
      )
    )

    (local.set $buf (i32.add (local.get $buf) (i32.const 4)))                  ;; buf++
    (local.set $src1 (i32.add (local.get $src1) (i32.const 4)))                ;; src1++
    (local.set $src2 (i32.add (local.get $src2) (i32.const 4)))                ;; src2++

    (br_if $loop                                                               ;; if (sampleCount == 0) break;
      (local.tee $sampleCount                                                  ;; sampleCount--
        (i32.add (local.get $sampleCount) (i32.const -1))
      )
    )
  end
)

)

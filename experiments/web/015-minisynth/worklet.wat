(module
(memory $mem 1)
(export "mem" (memory $mem))

;; version() int
(func (export "version") (result i32)
  i32.const 10004
)

;; noise(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "noise") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  (local $i i32)                                                               ;; int i = 0;
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4

  loop $loop
    ;; buf[i] = int32(ofs<<(gmconst.SampleBits-gmconst.DynamicBits-4)) >> (gmconst.SampleBits - gmconst.DynamicBits)
    ;; buf[i] = int32(ofs<<4)>>8
    (i32.store
      (i32.add (local.get $buf) (local.get $i))                                ;; buf[i]
      (i32.shr_s (i32.shl (local.get $ofs) (i32.const 4)) (i32.const 8))       ;; int32(ofs<<4)>>8
    )

    ;; ofs = (ofs+incr)*214013 + 2531011
    (local.set $ofs (i32.add (i32.mul (i32.add (local.get $ofs) (local.get $incr)) (i32.const 214013)) (i32.const 2531011)))

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end

  local.get $ofs
)

;; squareLQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "squareLQ") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  (local $i i32)                                                               ;; int i = 0;
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4

  loop $loop
    ;; buf[i] = (1 << (gmconst.DynamicBits - 1)) - int32(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits)
    ;; buf[i] = 8388608 - (ofs>>31<<24)
    (i32.store
      (i32.add (local.get $buf) (local.get $i))                                ;; buf[i]
      (i32.sub
        (i32.const 8388608)                                                    ;; 8388608 - ...
        (i32.shl (i32.shr_u (local.get $ofs) (i32.const 31)) (i32.const 24))   ;; ofs>>31<<24
      )
    )

    ;; ofs += incr
    (local.set $ofs (i32.add (local.get $ofs) (local.get $incr)))

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end

  local.get $ofs
)

;; squareHQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
(func (export "squareHQ") (param $buf i32) (param $sampleCount i32) (param $incr i32) (param $ofs i32) (result i32)
  (local $i i32)                                                               ;; int i = 0;
  (local $curVal i32)                                                          ;; int curVal = 0;
  (local $incr1 i64)                                                           ;; long incr1 = 0;
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4

  ;; curVal = int32(1<<(gmconst.DynamicBits-1) - int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits))
  ;; curVal = 8388608 - (ofs>>31<<24)
  (local.set $curVal
    (i32.sub
      (i32.const 8388608)                                                      ;; 8388608 - ...
      (i32.shl (i32.shr_u (local.get $ofs) (i32.const 31)) (i32.const 24))     ;; ofs>>31<<24
    )
  )
  ;; incr1 = (1 << (gmconst.DynamicBits + gmconst.DynamicBits)) / uint64(incr)
  ;; incr1 = 281474976710656 / uint64(incr)
  (local.set $incr1 (i64.div_u (i64.const 281474976710656) (i64.extend_i32_u (local.get $incr))))

  loop $loop
    ;; if ofs>>(gmconst.SampleBits-1) == (ofs+incr)>>(gmconst.SampleBits-1)
    ;; if ofs>>31 == (ofs+incr)>>31
    (i32.eq
      (i32.shr_u (local.get $ofs) (i32.const 31))                              ;; ofs>>31
      (i32.shr_u (i32.add (local.get $ofs) (local.get $incr)) (i32.const 31))  ;; (ofs+incr)>>31
    )
    (if
      (then
        ;; buf[i] = curVal
        (i32.store (i32.add (local.get $buf) (local.get $i)) (local.get $curVal))
        ;; ofs += incr
        (local.set $ofs (i32.add (local.get $ofs) (local.get $incr)))
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
                    (i32.and (local.get $ofs) (i32.const 2147483647))          ;; ofs & 2147483647
                  )
                )
                (local.get $incr1)
              )
              (i64.const 24)
            )
          )
        )
        ;; if ofs>>(gmconst.SampleBits-1) != 0
        ;; if ofs>>31
        (i32.shr_u (local.get $ofs) (i32.const 31))
        (if
          (then
            ;; curVal = 1<<(gmconst.DynamicBits-1) - curVal
            ;; curVal = 8388608 - curVal
            (local.set $curVal (i32.sub (i32.const 8388608) (local.get $curVal)))
          )
          (else
            ;; curVal = curVal - 1<<(gmconst.DynamicBits-1)
            ;; curVal = curVal - 8388608
            (local.set $curVal (i32.sub (local.get $curVal) (i32.const 8388608)))
          )
        )
        ;; buf[i] = curVal
        (i32.store (i32.add (local.get $buf) (local.get $i)) (local.get $curVal))
        ;; ofs += incr
        (local.set $ofs (i32.add (local.get $ofs) (local.get $incr)))
        ;; curVal = int32(1<<(gmconst.DynamicBits-1) - int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits))
        ;; curVal = 8388608 - (ofs>>31<<24)
        (local.set $curVal
          (i32.sub
            (i32.const 8388608)                                                  ;; 8388608 - ...
            (i32.shl (i32.shr_u (local.get $ofs) (i32.const 31)) (i32.const 24)) ;; ofs>>31<<24
          )
        )
      )
    )

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end

  local.get $ofs
)

;; convertIntSamplesToFloat32(floats *float, ints *int, sampleCount uint)
(func (export "convertIntSamplesToFloat32") (param $floats i32) (param $ints i32) (param $sampleCount i32)
  (local $i i32)                                                               ;; int i = 0;
  (local $mul f32)
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4
  ;; $mul = 1.0 / float32(1<<(gmconst.DynamicBits-1))
  ;; $mul = 1.0 / 8388608.0
  (local.set $mul (f32.div (f32.const 1) (f32.const 8388608)))

  loop $loop
    (f32.store
      (i32.add (local.get $floats) (local.get $i))                             ;; floats[i]
      (f32.mul (f32.convert_i32_s (i32.load (i32.add (local.get $ints) (local.get $i)))) (local.get $mul)) ;; float32(ints[i]) * $mul
    )

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end
)

;; volumeUpdate(buf *int, sampleCount uint, vol int)
(func (export "volumeUpdate") (param $buf i32) (param $sampleCount i32) (param $vol i32)
  (local $i i32)                                                               ;; int i = 0;
  (local $v i64)
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4
  (local.set $v (i64.extend_i32_s (local.get $vol)))

  loop $loop
    ;; buf[i] = int32((int64(buf[i]) * int64(vol)) >> (gmconst.DynamicBits - 1))
    ;; buf[i] = int32((int64(buf[i]) * int64(vol)) >> 23)
    (i32.store
      (i32.add (local.get $buf) (local.get $i))                                ;; buf[i]
      (i32.wrap_i64                                                            ;; int32()
        (i64.shr_s
          (i64.mul
            (i64.extend_i32_s (i32.load (i32.add (local.get $buf) (local.get $i)))) ;; int64(buf[i])
            (local.get $v)                                                          ;; int64(vol)
           )
          (i64.const 23)
        )
      )
    )

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end
)

;; mix(buf *int, src1 *int, src2 *int, sampleCount uint)
(func (export "mix") (param $buf i32) (param $src1 i32) (param $src2 i32) (param $sampleCount i32)
  (local $i i32)                                                               ;; int i = 0;
  (local.set $sampleCount (i32.mul (local.get $sampleCount) (i32.const 4)))    ;; sampleCount *= 4

  loop $loop
    ;; buf[i] = src1[i] + src2[i]
    (i32.store
      (i32.add (local.get $buf) (local.get $i))                                ;; buf[i]
      (i32.add                                                                 ;; src1[i] + src2[i]
        (i32.load (i32.add (local.get $src1) (local.get $i)))                  ;; src1[i]
        (i32.load (i32.add (local.get $src2) (local.get $i)))                  ;; src2[i]
      )
    )

    (local.tee $i (i32.add (local.get $i) (i32.const 4)))                      ;; i += 4
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < $sampleCount * 4) loop
  end
)

)

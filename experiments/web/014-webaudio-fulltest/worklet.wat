(module
(memory $mem 1)
(export "mem" (memory $mem))

(func (export "active") (result i32)
  i32.const 1
)

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
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < 128 * 4) loop
  end

  local.get $ofs
)

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
    (br_if $loop (i32.lt_s (; $i ;) (local.get $sampleCount)))                 ;; if (i < 128 * 4) loop
  end
)

)

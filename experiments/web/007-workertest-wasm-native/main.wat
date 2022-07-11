(module

(func (export "calc") (result i32)
  (local $sum i32)
  (local $x i32)
  (loop $labelfor
    (local.set $sum (i32.add (local.get $sum) (local.get $x)))    ;; $sum += $x
    (local.tee $x (i32.add (local.get $x) (i32.const 1)))         ;; $x++        [ $x
    (br_if $labelfor (i32.lt_u (; $x ;) (i32.const 1000000000)))  ;; loop $x < 1000000000
  )
  local.get $sum
)

(func (export "calc2") (result i32)
  (local $sum i32)
  (local $x i32)
  (loop $labelfor
    (local.set $sum (i32.add (local.get $sum) (local.get $x)))    ;; $sum += $x
    (local.tee $x (i32.add (local.get $x) (i32.const 1)))         ;; $x++        [ $x
    (local.set $sum (i32.add (; $x ;) (local.get $sum)))          ;; $sum += $x
    (local.tee $x (i32.add (local.get $x) (i32.const 1)))         ;; $x++        [ $x
    (br_if $labelfor (i32.lt_u (; $x ;) (i32.const 1000000000)))  ;; loop $x < 1000000000
  )
  local.get $sum
)

(func (export "calc3") (param $num i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const -1
      i32.add
      i64.extend_i32_u
      local.get 0
      i32.const -2
      i32.add
      i64.extend_i32_u
      i64.mul
      i64.const 1
      i64.shr_u
      i32.wrap_i64
      local.get 0
      i32.add
      i32.const -1
      i32.add
      return
    end
    i32.const 0
)

)

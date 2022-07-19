(module
(memory $mem 1)
(export "mem" (memory $mem))

(func (export "calc") (result i32)
  (i32.add (i32.load (i32.const 0)) (i32.load (i32.const 4)))
)

)

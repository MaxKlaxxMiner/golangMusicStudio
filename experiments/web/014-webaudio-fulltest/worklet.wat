(module
(memory $mem 1)
(export "mem" (memory $mem))

(func (export "calc") (result i32)
  i32.const 12345
)

)

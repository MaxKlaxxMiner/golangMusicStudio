(module
(memory $mem 1)
(export "mem" (memory $mem))

(func (export "version") (result i32)
  i32.const 10001
)

)

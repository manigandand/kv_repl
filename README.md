## How to run

```shell
go build . && ./kv_repl
```

## Sample run

```shell
store initialized
K/V REPL started...

$ write name mani
$ read name
mani
$ start
$ write name manigandan
$ read name
manigandan
$ start
$ write name Manigandan Dharmalingam
$ start
$ write name mani aborting
$ abort
$ read name
Manigandan Dharmalingam
$ write newK delete
$ read newk
key not found: newk
$ read name
Manigandan Dharmalingam
$ read newK
delete
$ delete newK
$ read newK
key not found: newK
$ commit
$ read name
Manigandan Dharmalingam
$ commit
$ read name
Manigandan Dharmalingam
$ commit
no active transaction
$ read name
Manigandan Dharmalingam
$ quit
Exiting...
```

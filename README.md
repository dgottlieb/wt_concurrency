When changing include/wt_raii.h, the cppinclude.go file must be regenerated with go-bindata. Usage:
```
$ go-bindata -o cppinclude.go -pkg wt_concurrency ./include/
```

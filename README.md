## Go cyclic sample
This repository is for showing a sample for cyclic dependencies on Golang

## How to run
```bash
go run main.go
```
or
```bash
go build
```

Then you get an expected error!
```bash
package github.com/Shunpoco/go-cyclic-sample
        imports github.com/Shunpoco/go-cyclic-sample/moduleA
        imports github.com/Shunpoco/go-cyclic-sample/moduleB
        imports github.com/Shunpoco/go-cyclic-sample/moduleA: import cycle not allowed
```

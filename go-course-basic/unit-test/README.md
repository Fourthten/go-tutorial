# Unit Test

```bash
go test -v
go test -v -run=TestFuncName
go test -v ./...
go test -run=TestFuncName/SubTestName
go test -run=/SubTestName
go test -v -bench=.
go test -v -run=NotFuncTest -bench=BenchTest
go test -v -bench=. ./...
go test -v -bench=BenchName/SubName
```

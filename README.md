## OTUS HomeWork practice

Practical part of the GoLang OTUS courses:
- https://otus.ru/lessons/golang-professional/?int_source=courses_catalog&int_term=programming
- https://rutracker.org/forum/viewtopic.php?t=5863412

#### hw01
- Hello now()
- Завести Go репозиторий на GitHub, написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
- Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.
- On Win: ```"c:\Program Files\Git\bin\bash.exe"```
- Then ```GOPATH=`pwd`/gopath; go get; go run hw01.go; go get -t; go test -v```

#### hw02
- ```GOROOT``` - path to GOlang distributive
- ```GOPATH``` - path to modules/sources (where go get store everything)
- ```go get, go get -t, go get -v``` - get modules for app, for tests, verbose to the ```$GOPATH/src```
- ```go build, go build hw01.go``` - compile source file
- ``` GOARCH=386 GOOS=darwin go build,  GOOS=windows go build``` - cross-build!!!
- ```go.mod``` - file with module description: name, dependendend modules (with versions), required Go version. ```go.sum``` - checksums of the deps
- ```go mod tidy``` - auto-add deps for the 'main' package to go.mod
- ```go run hw01.go``` - build and run one executable
- ```go help``` - usage and help
- ```go fmt hw01.go``` - auto-format code (if you don't use IDE)
- ```go get golang.org/x/tools/cmd/goimports; $GOPATH/bin/goimports hw01.go``` - update 'imports ()' block in package, add skipped and remove unused
- ```go vet hw01.go``` - internal linter (static analyzer). For e.g. nothing found for hw01.go
- ```go get -u golang.org/x/lint/golint; $GOPATH/bin/golint -set_exit_status hw01.go``` - popular 3rdparty linter. For e.g. for hw01: ```hw01.go:32:4: should replace cnt += 1 with cnt++```)
- ```go get github.com/golangci/golangci-lint/cmd/golangci-lint``` - metalinter. Metaliners needed to run many specialized linters in one execution
- ```go test``` - test your package
- ```go get github.com/gorilla/mux; go test github.com/gorilla/mux``` - run tests for external module
- ```go test google.golang.org/grpc/...``` - test external module and all its submodules
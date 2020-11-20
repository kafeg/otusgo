## OTUS Golang HomeWork practice

Practical part of the Golang OTUS courses:
- https://otus.ru/lessons/golang-professional/?int_source=courses_catalog&int_term=programming
- https://rutracker.org/forum/viewtopic.php?t=5863412

#### hw01 Start
- Homework:
- - Hello now()
- - Create program which print current datetime using NTP library.
- - App should correctly handle NTP library errors: print them to STDERR and return non-zero exit code.
- On Win: ```"c:\Program Files\Git\bin\bash.exe"```
- Then ```GOPATH=`pwd`/gopath; go get; go get -t; go test -v -run MainHW01``` 'main' function removed to keep files in one dir

#### hw02 Tools, testing
- ```GOROOT``` - path to Golang distributive
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

#### hw03 Elementary data types (ints, strings, stdlib for strings)
- ```go test -v -run MainHW03``` 'main' function removed to keep files in one dir
- ```strings``` module functions:
- - Contains(s, substr string) bool
- - HasPrefix(s, prefix string) bool - like Starts with
- - Join(a []string, sep string) string
- - Split(s, sep string) []string
- - DecodeRuneInString(s string) (r rune, size int) - get first rune from the string and its length in bytes
- - RuneCountInString(s string) (n int) - get string length in bytes
- - ValidString(s string) bool - check is string valid
- ```https://golang.org/pkg/``` - internal modules
- ```http://golang-book.ru/``` - Golang book in russian
- Homework:
- - Unpack string
- - Create Go function, for unpack string, containing repeating symbols / runes, for e.g.:
- - - "a4bc2d5e" => "aaaabccddddde"
- - - "abcd" => "abcd"
- - - "45" => "" (invalid string)
- - Additional: escape-sequence supoport
- - - `qwe\4\5` => `qwe45` (*)
- - - `qwe\45` => `qwe44444` (*)

#### hw04 Slices, arrays
- ```go test -v -run MainHW04``` 'main' function removed to keep files in one dir
- Homework:
- - Frequency analysis
- - Write a function that receives text as input and returns
- - 10 most common words without word forms

#### hw05 Functions, Scopes, error handling
- ```go test -v -run MainHW05``` 'main' function removed to keep files in one dir
- Scopes:
- - universe block - global scope: bool, int32, int64, float64, …; nil; make(), new(), panic(), …; true or false
- - package block - full package scope
- - file block - current file scope
- - local block - just {}. E.g.: {}, for, if, switch, case, select, ...
- Functions can be simple, variadic, return multi-values, etc...
- Closures can be used to simulate local static vars in functions
- Structs can have a methods, similar to classes
- Possible to add methods to the basic types using something like 'type Age int'
- Error handling:
- - 'error' is an interface with 'Error() string' method
- - err := errors.New("Im an error")
- Defer, Panic, Recover

#### hw06 Structures
- ```go test -v -run MainHW06``` 'main' function removed to keep files in one dir
- Homework:
- - Двусвязный список
- - Цель: https://en.wikipedia.org/wiki/Doubly_linked_list?Ожидаемые типы (псевдокод):?
- - - List // тип контейнер 
- - - - Len() // длинна списка 
- - - - First() // первый Item
- - - - Last() // последний Item 
- - - - PushFront(v interface{}) // добавить значение в начало 
- - - - PushBack(v interface{}) // добавить значение в конец 
- - - - Remove(i Item) // удалить элемент
- - - Item // элемент списка 
- - - - Value() interface{} // возвращает значение 
- - - - Next() *Item // следующий Item 
- - - - Prev() *Item // предыдущий
- - Реализовать двусвязанный список на языке Go

#### hw07 Interfaces
- ```go test -v -run MainHW07``` 'main' function removed to keep files in one dir
- Interfaces could be realized implicity
- One type could realize many interfaces
- One interface could be realized by the many types
- Interface - the set of the methods, which need to implement to satisfy interface
- 'interface{}' - empty interface, something like void*

#### hw08 Goroutines and channels
- ```go test -v -run MainHW08``` 'main' function removed to keep files in one dir
- Goroutine - light thread, for one task for e.g.
- go f() - call goroutine
- Channels - used to exchange data between Goroutins and to sync their execution
- Channels can be unbuffered and buffered (store some values, cap() and len() present)

#### hw09 Synchronization primitives
- ```go test -v -run MainHW09``` 'main' function removed to keep files in one dir
- sync.WaitGroup - instrument to wait more than one Goroutine to finish. WaitGroup: Add, Done, Wait
- sync.Mutex - standard mutex with Lock, Unlock functions
- sync.RWMutex - RLock, RUnlock - block only on read, Lock, Unlock - block for r/w
- sync.Map - special Map with included sync.RWMutex for high load usage
- sync.Pool - storage for the data safety to use by some goroutines. Any element in the pool can be force deleted in random time by the Garbage Collector
- sync.Once - do something only once
- sync.Cond - goroutines will wait for event
- Golang contains Race detector ``` go test -race ...```
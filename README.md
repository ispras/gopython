# gopython
This lib let you work with python source code from golang. You can create objects, run theire methods, get results of function calls, etc

## install
```sh
go get github.com/davidBMSTU/gopython
```

## include in your code
```go
import preferred_name "github.com/davidBMSTU/gopython/gopython"
```

## before run 
You need to set path to the folder where python source code files are stored. Therefore execute:
```sh
export PYTHONPATH=/path/to/python/source/code/folder
```

## test
You can check correctness of gopython work(does gopython connect with your system python properly?) with test scenario. Run this after installation:
```sh
cd $GOPATH/src/github.com/davidBMSTU/gopython
export PYTHONPATH=$(pwd)/tests
go run tests/main.go
```
Everything is ok if you get this:
```
gopython is OK
```

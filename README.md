# gopython
This tool let you embed python code into golang applications. You can create objects, run theire methods, get results of function calls, etc.

## Installation
```sh
go get github.com/ispras/gopython
```

## Include into your code
```go
import gopython "github.com/ispras/gopython/src"
```

## Before run 
You need to set path to the folder where python source code files are stored. Therefore execute:
```sh
export PYTHONPATH=/path/to/python/source/code/folder
```

## Test installation
You can check correctness of gopython installation (does gopython connect with your system python properly?) with test scenario. Run this after installation:
```sh
cd $GOPATH/src/github.com/ispras/gopython
export PYTHONPATH=$(pwd)/test_installation
go run test_installation/test_inst.go
```
Everything is ok if you get this:
```
gopython is OK
```

## Examples

### Unique calculation

Let's consider that you have some unique python tool that calculate something for you. And you need to embed this into some go code. Nowadays there are a lot of unique Python libraries in many different areas.

For simplicity let's consider that our python library interface looks like this:
```python
import random

class unique_python_calc:
    def __init__(self):
        print("__init__() called")
        self.unique_number = random.randint(-10, 10)
    
    def calc(self, number):
        print("calc() called with arg:", number)
        
        # unique calc with unique python tool...
        res = number * self.unique_number

        print("calc() res =", res)
        return res
```

And the usage of the interface looks like this:
```python
from unique_calc import unique_python_calc

inputForCalc = 1.23

upc = unique_python_calc()
res = upc.calc(inputForCalc)

print("Py: Result from UniquePyCalc:", res)
```

And we need to call that 'calc' method in the go code and get the result. With the use of gopython one can do it like this:
```Go
package main

import (
	"fmt"

	gopython "github.com/ispras/gopython/src"
)

func main() {
	inputForUniquePyCalc := 1.23

	gopython.InitPythonInterpretetor()

	// 1. from unique_calc import unique_python_calc
	var pymodule gopython.PythonModule
	pymodule.SetModuleName("unique_calc")
	pymodule.MakeImport()

	upc_py, _ := pymodule.GetClass("unique_python_calc")

	// 2. upc = unique_python_calc()
	var initArgs gopython.PythonMethodArguments
	initArgs.SetArgCount(0)

	upc_py_obj, _ := upc_py.CreateObject(initArgs)

	// 3. res = upc.calc(inputForCalc)
	var calcArgs gopython.PythonMethodArguments
	calcArgs.SetArgCount(1)
	calcArgs.SetNextArgument(inputForUniquePyCalc)

	res_py, _ := upc_py_obj.CallMethod("calc", &calcArgs)
	res, _ := res_py[0].ToStandartGoType()

	gopython.FinalizePythonInterpretetor()

	fmt.Printf("Go: Result from UniquePyCalc: %f\n", res)
}
```

Every part of code above is marked with the comment with corresponding python line. Error handling omitted for clarity.

The source code of this example is [here](https://github.com/ispras/cotea/tree/main/examples/calculate). Don't forget to set PYTHONPATH env variable (it must be a path to python source file that you want to embed). In the case of *calculate* example you can do this like this:

```bash
export PYTHONPATH=$GOPATH/src/github.com/ispras/gopython/examples/calculate
```

<!--The main gopython abstractions are:-->
<!--- PythonModule-->
<!--- PythonClass-->
<!--- PythonObject-->
<!--- PythonMethodArguments-->

<!--PythonModule -->

A detailed overview of all interfaces is provided in [gopython documentation](https://github.com/ispras/gopython/blob/main/docs/gopython_docs.md)
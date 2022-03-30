package main

import (
	"fmt"

	gopython "github.com/ispras/gopython/src/gopython"
)

func main() {
	inputForUniquePyCalc := 1.23

	gopython.InitPythonInterpretetor()

	// from unique_calc import unique_python_calc
	var pymodule gopython.PythonModule
	pymodule.SetModuleName("unique_calc")
	pymodule.MakeImport()

	upc_py, _ := pymodule.GetClass("unique_python_calc")

	// upc = unique_python_calc()
	var initArgs gopython.PythonMethodArguments
	initArgs.SetArgCount(0)

	upc_py_obj, _ := upc_py.CreateObject(&initArgs)

	// res = upc.calc(inputForCalc)
	var calcArgs gopython.PythonMethodArguments
	calcArgs.SetArgCount(1)
	calcArgs.SetNextArgument(inputForUniquePyCalc)

	res_py, _ := upc_py_obj.CallMethod("calc", &calcArgs)
	res, _ := res_py[0].ToStandartGoType()

	gopython.FinalizePythonInterpretetor()

	fmt.Printf("Go: Result from UniquePyCalc: %f\n", res)
}

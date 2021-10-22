package main

import gopython "github.com/davidBMSTU/gopython/gopython"
import "fmt"
import "os"

func main() {
	InputToPyMethod := "It's 6 o'clock, man!"

	gopython.InitPythonInterpretetor()

	var pymodule gopython.PythonModule
	pymodule.SetModuleName("test")
	err := pymodule.MakeImport()
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopython is NOT OK")
		os.Exit(1)
	}

	var initArgs gopython.PythonMethodArguments
	initArgs.SetArgCount(1)
	initArgs.SetNextArgument(InputToPyMethod)

	var methodArgs gopython.PythonMethodArguments
	methodArgs.SetArgCount(0)

	watthetime, err := pymodule.GetClass("watthetime")
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopython is NOT OK")
		os.Exit(1)
	}

	watthetimeObj, err := watthetime.CreateObject(initArgs)
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopython is NOT OK")
		os.Exit(1)
	}

	res, err := watthetimeObj.CallMethod("say_it", &methodArgs)
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopython is NOT OK")
		os.Exit(1)
	}

	//isStandart, _ := res[0].IsStandartType()

	//fmt.Println(isStandart)

	toStd, err := res[0].ToStandartGoType()
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopython is NOT OK")
		os.Exit(1)
	}

	//fmt.Println(toStd)

	if toStd.(string) == InputToPyMethod {
		fmt.Println("gopython is OK")
	} else {
		fmt.Println("gopython is NOT OK")
	}

	gopython.FinalizePythonInterpretetor()
}

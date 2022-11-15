package python_class_test

import (
	"os"
	"testing"

	"github.com/ispras/gopython/src/gopython"
)

func TestCreateObject(t *testing.T) {
	os.Setenv("PYTHONPATH", ".")
	defer os.Unsetenv("PYTHONPATH")

	gopython.InitPythonInterpretetor()
	defer gopython.FinalizePythonInterpretetor()

	testingModuleName := "testmod"

	f, err := os.Create(testingModuleName + ".py")
	if err != nil {
		t.Fatalf("Can't create a file for test: %s", err)
	}

	className := "TestClass"

	TestClass := "class " + className + ":\n"
	TestClass += "\tdef __init__(self, str_arg):\n"
	TestClass += "\t\tif not isinstance(str_arg, str):\n"
	TestClass += "\t\t\traise ValueError('str_arg should be str')\n"

	_, err = f.Write([]byte(TestClass))
	if err != nil {
		t.Fatalf("Can't write to file for test: %s", err)
	}

	defer os.Remove(testingModuleName + ".py")

	t.Run("create obj of TestClass from testmod", func(t *testing.T) {
		testingModuleName := "testmod"

		var testModule gopython.PythonModule
		testModule.SetModuleName(testingModuleName)
		err = testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		testClassPy, err := testModule.GetClass(className)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgs gopython.PythonMethodArguments
		initArgs.SetArgCount(1)
		initArgs.SetNextArgument("this is string")

		_, err = testClassPy.CreateObject(&initArgs)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("create obj of TestClass from testmod with wrong arg", func(t *testing.T) {
		testingModuleName := "testmod"

		var testModule gopython.PythonModule
		testModule.SetModuleName(testingModuleName)
		err = testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		testClassPy, err := testModule.GetClass(className)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgs gopython.PythonMethodArguments
		initArgs.SetArgCount(1)
		initArgs.SetNextArgument(1) // passing not string

		_, err = testClassPy.CreateObject(&initArgs)
		if err == nil {
			t.Fatalf("Expected err not nil, but received nil")
		}
	})
}

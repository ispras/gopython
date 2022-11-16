package python_method_arguments_test

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
		t.Fatalf("Error during working with files for tests: %s", err)
	}

	classNameInt := "TestClassInt"
	TestClassInt := "class " + classNameInt + ":\n"
	TestClassInt += "\tdef __init__(self, arg):\n"
	TestClassInt += "\t\tif not isinstance(arg, int):\n"
	TestClassInt += "\t\t\traise ValueError('arg should be int')\n"

	classNameFloat := "TestClassFloat"
	TestClassFloat := "class " + classNameFloat + ":\n"
	TestClassFloat += "\tdef __init__(self, arg):\n"
	TestClassFloat += "\t\tif not isinstance(arg, float):\n"
	TestClassFloat += "\t\t\traise ValueError('arg should be float')\n"

	classNamePyobj := "TestClassPyobj"
	TestClassPyobj := "class " + classNamePyobj + ":\n"
	TestClassPyobj += "\tdef __init__(self, arg):\n"
	TestClassPyobj += "\t\tif not isinstance(arg, " + classNameInt + "):\n"
	TestClassPyobj += "\t\t\traise ValueError('arg should be " + classNameInt + "')\n"

	_, err = f.Write([]byte(TestClassInt + "\n" + TestClassFloat + "\n" + TestClassPyobj))
	if err != nil {
		t.Fatalf("Error during working with files for tests: %s", err)
	}

	var testModule gopython.PythonModule
	testModule.SetModuleName(testingModuleName)
	err = testModule.MakeImport()
	if err != nil {
		t.Fatalf("Expected err nil, but received: %s", err)
	}

	// defer os.Remove(testingModuleName + ".py")

	t.Run("passing int test", func(t *testing.T) {
		testClassPy, err := testModule.GetClass(classNameInt)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgs gopython.PythonMethodArguments
		initArgs.SetArgCount(1)
		initArgs.SetNextArgument(1)

		_, err = testClassPy.CreateObject(&initArgs)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("passing float test", func(t *testing.T) {
		testClassPy, err := testModule.GetClass(classNameFloat)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgs gopython.PythonMethodArguments
		initArgs.SetArgCount(1)

		var floatNum float64
		floatNum = 1.23
		initArgs.SetNextArgument(floatNum)

		_, err = testClassPy.CreateObject(&initArgs)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("passing other python object", func(t *testing.T) {
		testClassTmp, err := testModule.GetClass(classNameInt)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgsTmp gopython.PythonMethodArguments
		initArgsTmp.SetArgCount(1)
		initArgsTmp.SetNextArgument(1)

		pyobj, err := testClassTmp.CreateObject(&initArgsTmp)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		testClassPy, err := testModule.GetClass(classNamePyobj)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		var initArgs gopython.PythonMethodArguments
		initArgs.SetArgCount(1)
		initArgs.SetNextArgument(pyobj)

		_, err = testClassPy.CreateObject(&initArgs)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})
}

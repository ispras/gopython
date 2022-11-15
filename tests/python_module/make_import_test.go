package python_module_test

import (
	"os"
	"testing"

	"github.com/ispras/gopython/src/gopython"
)

func TestMakeImport(t *testing.T) {
	os.Setenv("PYTHONPATH", ".")
	defer os.Unsetenv("PYTHONPATH")

	gopython.InitPythonInterpretetor()
	defer gopython.FinalizePythonInterpretetor()

	t.Run("importing sys module", func(t *testing.T) {
		importingModuleName := "sys"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()

		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("importing non-existent mudule", func(t *testing.T) {
		importingModuleName := "module-that-doesnt-exist"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()

		if err == nil {
			t.Fatalf("Expected err not nil, but received: %s", err)
		}
	})

	t.Run("importing non-existent mudule", func(t *testing.T) {
		importingModuleName := "module-that-doesnt-exist"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()

		if err == nil {
			t.Fatalf("Expected err not nil, but received: %s", err)
		}
	})

	t.Run("importing custom module from PYTHONPATH", func(t *testing.T) {
		testingModuleName := "testmod"

		_, err := os.Create(testingModuleName + ".py")
		if err != nil {
			t.Fatalf("Can't create a file for test: %s", err)
		}

		var testModule gopython.PythonModule
		testModule.SetModuleName(testingModuleName)
		err = testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		err = os.Remove(testingModuleName + ".py")
		if err != nil {
			t.Fatalf("Can't delete a file for test: %s", err)
		}
	})
}

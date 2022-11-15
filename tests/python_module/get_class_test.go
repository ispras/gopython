package python_module_test

import (
	"os"
	"testing"

	"github.com/ispras/gopython/src/gopython"
)

func TestGetClass(t *testing.T) {
	os.Setenv("PYTHONPATH", ".")
	defer os.Unsetenv("PYTHONPATH")

	gopython.InitPythonInterpretetor()
	defer gopython.FinalizePythonInterpretetor()

	t.Run("getting Thread from threading", func(t *testing.T) {
		importingModuleName := "threading"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		className := "Thread"
		_, err = testModule.GetClass(className)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("getting non-existent class from threading", func(t *testing.T) {
		importingModuleName := "threading"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		className := "Threadd"
		_, err = testModule.GetClass(className)
		if err == nil {
			t.Fatalf("Expected err not nil, but received nil")
		}
	})
}

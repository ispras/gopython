package python_module_test

import (
	"os"
	"testing"

	"github.com/ispras/gopython/src/gopython"
)

func TestGetObject(t *testing.T) {
	os.Setenv("PYTHONPATH", ".")
	defer os.Unsetenv("PYTHONPATH")

	gopython.InitPythonInterpretetor()
	defer gopython.FinalizePythonInterpretetor()

	t.Run("getting non-existent object from sys", func(t *testing.T) {
		importingModuleName := "sys"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		objectName := "moduless"
		_, err = testModule.GetObject(objectName)
		if err == nil {
			t.Fatalf("Expected err not nil, but received nil")
		}
	})

	t.Run("getting non-callable object modules from sys", func(t *testing.T) {
		importingModuleName := "sys"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		objectName := "modules"
		_, err = testModule.GetObject(objectName)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})

	t.Run("getting callable object(function) settrace from sys", func(t *testing.T) {
		importingModuleName := "sys"

		var testModule gopython.PythonModule
		testModule.SetModuleName(importingModuleName)
		err := testModule.MakeImport()
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}

		objectName := "settrace"
		_, err = testModule.GetObject(objectName)
		if err != nil {
			t.Fatalf("Expected err nil, but received: %s", err)
		}
	})
}

package gopython

// #cgo pkg-config: python3
// #include <Python.h>
import "C"

// PythonModule ...
type PythonModule struct {
	moduleName string
	Module     *C.PyObject
}

func (pymod *PythonModule) SetModuleName(moduleName string) {
	pymod.moduleName = moduleName
}

func (pymod *PythonModule) MakeImport() error {
	moduleNameC := C.CString(pymod.moduleName)
	pythonModuleName := C.PyUnicode_DecodeFSDefault(moduleNameC)
	importResult := C.PyImport_Import(pythonModuleName)

	if importResult == nil {
		var e errors
		e.importError(pymod.moduleName)
		return &e
	}

	pymod.Module = importResult

	return nil
}

func (pymod *PythonModule) GetClass(className string) (*PythonClass, error) {
	if pymod.Module == nil {
		var e errors
		e.notImportedModule()
		return nil, &e
	}

	classNameC := C.CString(className)
	resultClass := C.PyObject_GetAttrString(pymod.Module, classNameC)

	if resultClass == nil || C.PyCallable_Check(resultClass) == 0 {
		var e errors
		e.gettingObjectFailed()
		return nil, &e
	}

	var res PythonClass
	res.classPointer = resultClass

	return &res, nil
}

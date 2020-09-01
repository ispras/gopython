package gopython

// #cgo pkg-config: python-3.5
// #include <python3.5m/Python.h>
import "C"

// PythonModule ...
type PythonModule35 struct {
	moduleName string
	Module     *C.PyObject
}

func (pymod *PythonModule35) SetModuleName(moduleName string) {
	pymod.moduleName = moduleName
}

func (pymod *PythonModule35) MakeImport() error {
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

func (pymod *PythonModule35) GetClass(className string) (PythonClass, error) {
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

	var res PythonClass35
	res.classPointer = resultClass

	return &res, nil
}

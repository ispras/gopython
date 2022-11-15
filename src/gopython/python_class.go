package gopython

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"

type PythonClass struct {
	classPointer *C.PyObject
}

func (pyclass *PythonClass) CreateObject(args *PythonMethodArguments) (*PythonObject, error) {
	if pyclass.classPointer == nil {
		var e errors
		e.classNotInited()
		return nil, &e
	}

	pObj := C.PyObject_CallObject(pyclass.classPointer, args.argumentsTurple)
	if pObj == nil {
		var e errors
		e.errorDuringMethodCall("__init__")
		return nil, &e
	}

	var resObj PythonObject
	resObj.ObjectPointer = pObj

	return &resObj, nil
}

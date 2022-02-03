package gopython

// #cgo pkg-config: python3
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

	// TODO: check, that pObj OK

	var resObj PythonObject
	resObj.ObjectPointer = pObj

	return &resObj, nil
}

package gopython

// #cgo pkg-config: python-3.5
// #include <python3.5m/Python.h>
import "C"

type PythonClass35 struct {
	classPointer *C.PyObject
}

func (pyclass *PythonClass35) CreateObject(args PythonMethodArguments) (PythonObject, error) {
	if pyclass.classPointer == nil {
		var e errors
		e.classNotInited()
		return nil, &e
	}

	argsTupleGoInterface := args.GetArgumentsTuple()
	argsTupleCPointer := argsTupleGoInterface.(*C.PyObject)
	pObj := C.PyObject_CallObject(pyclass.classPointer, argsTupleCPointer)

	// TODO: check, that pObj OK

	var resObj PythonObject35
	resObj.ObjectPointer = pObj

	return &resObj, nil
}

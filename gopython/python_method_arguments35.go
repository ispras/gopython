package gopython

// #cgo pkg-config: python3
// #include <Python.h>
import "C"

type PythonMethodArguments35 struct {
	argumentsTurple *C.PyObject
	argumentsCount  int
	curArgIndex     int
}

func (pyargs *PythonMethodArguments35) SetArgCount(count int) {
	pyargs.argumentsCount = count
	cLongLen := C.long(count)

	pyargs.argumentsTurple = C.PyTuple_New(cLongLen)
	pyargs.curArgIndex = 0
}

func (pyargs *PythonMethodArguments35) GetArgumentsTuple() interface{} {
	return pyargs.argumentsTurple
}

func (pyargs *PythonMethodArguments35) SetNextArgument(arg interface{}) {
	ind := C.long(pyargs.curArgIndex)

	switch v := arg.(type) {
	case int:
		tmp := C.long(v)
		intArg := C.PyLong_FromLong(tmp)
		C.PyTuple_SetItem(pyargs.argumentsTurple, ind, intArg)

	case float64:
		tmp := C.double(v)
		floatArg := C.PyFloat_FromDouble(tmp)
		C.PyTuple_SetItem(pyargs.argumentsTurple, ind, floatArg)

	case string:
		tmp := C.CString(v)
		stringArg := C.PyUnicode_DecodeFSDefault(tmp)
		C.PyTuple_SetItem(pyargs.argumentsTurple, ind, stringArg)

	case *PythonObject35:
		C.PyTuple_SetItem(pyargs.argumentsTurple, ind, v.ObjectPointer)
	}

	pyargs.curArgIndex++
}

package gopython

/*
#cgo pkg-config: python-3.6
#include <python3.6m/Python.h>

static PyObject *
null_error(void)
{
    if (!PyErr_Occurred())
        PyErr_SetString(PyExc_SystemError,
                        "null argument to internal routine");
    return NULL;
}

int PyTuple_CheckFunc(PyObject *p) {
	return PyTuple_Check(p);
}


PyAPI_FUNC(PyObject *) PyObject_CallMethodTupleArgs(PyObject *callable,
													PyObject *name,
													PyObject *turple_args)
{
	PyObject *res;

    if (callable == NULL || name == NULL)
        return null_error();

    callable = PyObject_GetAttr(callable, name);
    if (callable == NULL)
        return NULL;

    if (turple_args == NULL) {
        Py_DECREF(callable);
        return NULL;
	}

    res = PyObject_Call(callable, turple_args, NULL);
    Py_DECREF(turple_args);
    Py_DECREF(callable);

	return res;
}

const char* get_pyobject_type(PyObject *obj)
{
	const char *obj_type = Py_TYPE(obj)->tp_name;
	return obj_type;
}

char* string_from_pyobject(PyObject *obj)
{
	PyObject * temp_bytes = PyUnicode_AsEncodedString(obj, "UTF-8", "strict");
    char *res = PyBytes_AsString(temp_bytes);
	Py_DECREF(temp_bytes);
	return res;
}


*/
import "C"

type PythonObject35 struct {
	ObjectPointer *C.PyObject
}

func (pyobj *PythonObject35) CallMethod(mName string, args PythonMethodArguments) ([]PythonObject, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return nil, &e
	}

	mNameC := C.CString(mName)
	hasMethod := C.PyObject_HasAttrString(pyobj.ObjectPointer, mNameC)

	if hasMethod == 0 {
		var e errors
		e.noSuchAttr(mName)
		return nil, &e
	}

	mNamePy := C.PyUnicode_DecodeFSDefault(mNameC)
	argsTupleGoInterface := args.GetArgumentsTuple()
	argsTupleCPointer := argsTupleGoInterface.(*C.PyObject)
	pyResult := C.PyObject_CallMethodTupleArgs(pyobj.ObjectPointer, mNamePy, argsTupleCPointer)

	if pyResult == nil {
		var e errors
		e.errorDuringMethodCall(mName)
		return nil, &e
	}

	isTurple := C.PyTuple_CheckFunc(pyResult)
	var resultObjectsCount int

	if isTurple == 0 {
		resultObjectsCount = 1
	} else {
		tmp := C.PyTuple_Size(pyResult)
		resultObjectsCount = int(tmp)
	}

	res := make([]PythonObject, resultObjectsCount)

	if isTurple == 0 {
		res[0] = &PythonObject35{ObjectPointer: pyResult}
	} else {
		for i := 0; i < resultObjectsCount; i++ {
			tmpInd := C.long(i)
			tmpObjPointer := C.PyTuple_GetItem(pyResult, tmpInd)
			res[i] = &PythonObject35{ObjectPointer: tmpObjPointer}
		}
	}

	return res, nil
}

func (pyobj *PythonObject35) HasAttr(attrName string) (bool, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return false, &e
	}

	mNameC := C.CString(attrName)
	hasMethod := C.PyObject_HasAttrString(pyobj.ObjectPointer, mNameC)
	var res bool

	if hasMethod == 0 {
		res = false
	} else {
		res = true
	}

	return res, nil
}

func (pyobj *PythonObject35) GetAttr(attrName string) (PythonObject, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return nil, &e
	}

	hasAttr, _ := pyobj.HasAttr(attrName)
	if hasAttr == false {
		var e errors
		e.noSuchAttr(attrName)
		return nil, &e
	}

	cstring := C.CString(attrName)
	attrPointer := C.PyObject_GetAttrString(pyobj.ObjectPointer, cstring)

	// if attrPointer == nil ???

	var resObj PythonObject35
	resObj.ObjectPointer = attrPointer

	return &resObj, nil
}

func (pyobj *PythonObject35) GetType() (string, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return "", &e
	}

	objTypeC := C.get_pyobject_type(pyobj.ObjectPointer)
	objType := C.GoString(objTypeC)
	return objType, nil
}

func (pyobj *PythonObject35) IsStandartType() (bool, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return false, &e
	}

	objType, _ := pyobj.GetType()
	var res bool

	switch objType {
	case "int":
		res = true
	case "float":
		res = true
	case "str":
		res = true
	case "bool":
		res = true
	default:
		res = false
	}

	return res, nil
}

func (pyobj *PythonObject35) ToStandartGoType() (interface{}, error) {
	isStandart, _ := pyobj.IsStandartType()
	if isStandart == false {
		var e errors
		e.notStandartType()
		return nil, &e
	}

	var res interface{}
	objType, _ := pyobj.GetType()

	switch objType {
	case "int":
		clong := C.PyLong_AsLong(pyobj.ObjectPointer)
		res = int(clong)
	case "float":
		cdouble := C.PyFloat_AsDouble(pyobj.ObjectPointer)
		res = float64(cdouble)
	case "str":
		cstring := C.string_from_pyobject(pyobj.ObjectPointer)
		res = C.GoString(cstring)
	case "bool":
		boolVarC := C.PyObject_IsTrue(pyobj.ObjectPointer)
		tmp := int(boolVarC)
		if tmp == 1 {
			res = true
		} else {
			res = false
		}
	}

	return res, nil
}

// TODO: get object attr - 									 DONE
// TODO: hasAttr - 											 DONE
// TODO: conversesion to go type if PythonObject is standart
//       type(like int, float, string, bool) - 				 DONE
// TODO: isList(); isDict(); isTuple()
// TODO: if list/dict/tuple - convert to list/dict/tuple of
//		 pyobjects

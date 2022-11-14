package gopython

/*
#cgo pkg-config: python3-embed
#include <Python.h>

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

type PythonObject struct {
	ObjectPointer *C.PyObject
}

func (pyobj *PythonObject) CallMethod(mName string, args *PythonMethodArguments) ([]*PythonObject, error) {
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
	pyResult := C.PyObject_CallMethodTupleArgs(pyobj.ObjectPointer, mNamePy, args.argumentsTurple)

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

	res := make([]*PythonObject, resultObjectsCount)

	if isTurple == 0 {
		res[0] = &PythonObject{ObjectPointer: pyResult}
	} else {
		for i := 0; i < resultObjectsCount; i++ {
			tmpInd := C.long(i)
			tmpObjPointer := C.PyTuple_GetItem(pyResult, tmpInd)
			res[i] = &PythonObject{ObjectPointer: tmpObjPointer}
		}
	}

	return res, nil
}

func (pyobj *PythonObject) CallItself(args *PythonMethodArguments) ([]*PythonObject, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return nil, &e
	}

	pyResult := C.PyObject_CallObject(pyobj.ObjectPointer, args.argumentsTurple)

	if pyResult == nil {
		var e errors
		e.errorDuringMethodCall("Itself call")
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

	res := make([]*PythonObject, resultObjectsCount)

	if isTurple == 0 {
		res[0] = &PythonObject{ObjectPointer: pyResult}
	} else {
		for i := 0; i < resultObjectsCount; i++ {
			tmpInd := C.long(i)
			tmpObjPointer := C.PyTuple_GetItem(pyResult, tmpInd)
			res[i] = &PythonObject{ObjectPointer: tmpObjPointer}
		}
	}

	return res, nil
}

func (pyobj *PythonObject) HasAttr(attrName string) (bool, error) {
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

func (pyobj *PythonObject) GetAttr(attrName string) (*PythonObject, error) {
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

	var resObj PythonObject
	resObj.ObjectPointer = attrPointer

	return &resObj, nil
}

func (pyobj *PythonObject) GetType() (string, error) {
	if pyobj.ObjectPointer == nil {
		var e errors
		e.nilObjectPointer()
		return "", &e
	}

	objTypeC := C.get_pyobject_type(pyobj.ObjectPointer)
	objType := C.GoString(objTypeC)
	return objType, nil
}

func (pyobj *PythonObject) IsStandartType() (bool, error) {
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

func (pyobj *PythonObject) ToStandartGoType() (interface{}, error) {
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

func (pyobj *PythonObject) CreateFromGoSlice(goSlice interface{}) error {
	correctSliceType := true

	switch v := goSlice.(type) {
	case []int:
		elementsCount := len(v)
		cElementsCount := C.long(elementsCount)
		pyobj.ObjectPointer = C.PyList_New(cElementsCount)

		for i := 0; i < elementsCount; i += 1 {
			tmp := C.long(v[i])
			ind := C.long(i)
			intArg := C.PyLong_FromLong(tmp)

			C.PyList_SetItem(pyobj.ObjectPointer, ind, intArg)
		}

	case [][]int:
		rowsCount := len(v)
		cRowsCount := C.long(rowsCount)
		pyobj.ObjectPointer = C.PyList_New(cRowsCount)

		for i := 0; i < rowsCount; i += 1 {
			tmpRowLen := len(v[i])
			cTmpRowLen := C.long(tmpRowLen)
			tmpRowPy := C.PyList_New(cTmpRowLen)

			for j := 0; j < tmpRowLen; j += 1 {
				tmp := C.long(v[i][j])
				ind := C.long(j)
				intArg := C.PyLong_FromLong(tmp)

				C.PyList_SetItem(tmpRowPy, ind, intArg)
			}

			ind := C.long(i)
			C.PyList_SetItem(pyobj.ObjectPointer, ind, tmpRowPy)
		}

	case []float64:
		elementsCount := len(v)
		cElementsCount := C.long(elementsCount)
		pyobj.ObjectPointer = C.PyList_New(cElementsCount)

		for i := 0; i < elementsCount; i += 1 {
			tmp := C.double(v[i])
			ind := C.long(i)
			floatArg := C.PyFloat_FromDouble(tmp)

			C.PyList_SetItem(pyobj.ObjectPointer, ind, floatArg)
		}

	case []string:
		elementsCount := len(v)
		cElementsCount := C.long(elementsCount)
		pyobj.ObjectPointer = C.PyList_New(cElementsCount)

		for i := 0; i < elementsCount; i += 1 {
			tmp := C.CString(v[i])
			ind := C.long(i)
			stringArg := C.PyUnicode_DecodeFSDefault(tmp)

			C.PyList_SetItem(pyobj.ObjectPointer, ind, stringArg)
		}

	case []*PythonObject:
		elementsCount := len(v)
		cElementsCount := C.long(elementsCount)
		pyobj.ObjectPointer = C.PyList_New(cElementsCount)

		for i := 0; i < elementsCount; i += 1 {
			ind := C.long(i)

			C.PyList_SetItem(pyobj.ObjectPointer, ind, v[i].ObjectPointer)
		}

	default:
		correctSliceType = false
	}

	if !correctSliceType {
		var e errors
		e.notSupportedGoSlice()
		return &e
	}

	return nil
}

func (pyobj *PythonObject) GetPythonObjectsFromPyList() []*PythonObject {
	cListLen := C.PyList_Size(pyobj.ObjectPointer)
	listLen := int(cListLen)

	res := make([]*PythonObject, listLen)

	for i := 0; i < listLen; i += 1 {
		ind := C.long(i)
		listElem := C.PyList_GetItem(pyobj.ObjectPointer, ind)

		res[i] = &PythonObject{ObjectPointer: listElem}
	}

	return res
}

func (pyobj *PythonObject) GetGoSliceFromPyList() (interface{}, error) {
	pyObjectsOfList := pyobj.GetPythonObjectsFromPyList()

	listLen := len(pyObjectsOfList)

	if listLen < 1 {
		var e errors
		e.pyListEmpty()
		return nil, &e
	}

	isStandartType, err := pyObjectsOfList[0].IsStandartType()
	if err != nil {
		return nil, err
	}

	if !isStandartType {
		var e errors
		e.notSupportedPyList()
		return nil, &e
	}

	firstElemType, err := pyObjectsOfList[0].GetType()
	if err != nil {
		return nil, err
	}

	oneTypeList := true

	for i := 0; i < listLen; i += 1 {
		tmpType, err := pyObjectsOfList[i].GetType()
		if err != nil {
			return nil, err
		}

		if tmpType != firstElemType {
			oneTypeList = false
			break
		}
	}

	if !oneTypeList {
		var e errors
		e.pyListWithDifferentTypes()
		return nil, &e
	}

	switch firstElemType {
	case "int":
		resSlice := make([]int, listLen)

		for i := 0; i < listLen; i += 1 {
			tmpInterface, err := pyObjectsOfList[i].ToStandartGoType()
			if err != nil {
				return nil, err
			}

			resSlice[i] = tmpInterface.(int)
		}

		return resSlice, nil

	case "float":
		resSlice := make([]float64, listLen)

		for i := 0; i < listLen; i += 1 {
			tmpInterface, err := pyObjectsOfList[i].ToStandartGoType()
			if err != nil {
				return nil, err
			}

			resSlice[i] = tmpInterface.(float64)
		}

		return resSlice, nil

	}

	return nil, nil
}

func CreatePythonListFromGoSlice(goSlice interface{}) (*PythonObject, error) {
	var resPythonObject PythonObject

	err := resPythonObject.CreateFromGoSlice(goSlice)

	return &resPythonObject, err
}

// TODO: get object attr - 									 DONE
// TODO: hasAttr - 											 DONE
// TODO: conversesion to go type if PythonObject is standart
//       type(like int, float, string, bool) - 				 DONE
// TODO: isList(); isDict(); isTuple()
// TODO: if list/dict/tuple - convert to list/dict/tuple of
//		 pyobjects

package gopython

/*
#cgo pkg-config: python3-embed
#include <Python.h>

char* string_from_pyobject2(PyObject *obj)
{
	printf("START\n");
	PyObject * temp_bytes = PyUnicode_AsEncodedString(obj, "UTF-8", "strict");
	printf("RATATA\n");
    char *res = PyBytes_AsString(temp_bytes);
	Py_DECREF(temp_bytes);
	return res;
}

char* PyErr_GetErrorString()
{
	// getting exception parts
	PyObject *exception_cls, *err_msg, *traceback;
	PyErr_Fetch(&exception_cls, &err_msg, &traceback);

	// creating exception object (needed by format_exception function below)
	PyObject* exception_init_args = PyTuple_Pack(1, err_msg);
	PyObject* exception_obj = PyObject_Call(exception_cls, exception_init_args, NULL);

	printf("ZAEBUMBA\n");

	// getting traceback.format_exception
	PyObject* tracback_module = PyImport_ImportModule("traceback");
	PyObject* format_exception_name = PyUnicode_DecodeFSDefault("format_exception");
	PyObject* format_exception = PyObject_GetAttr(tracback_module, format_exception_name);

	// calling the format_exception function
	PyObject* format_exception_args = PyTuple_Pack(3, exception_cls, exception_obj, traceback);
	PyObject* formated_err_list = PyObject_Call(format_exception, format_exception_args, NULL);

	PyObject* formated_err = PyUnicode_FromString("");
	PyObject* formated_err_i;

	if (PyList_Check(formated_err_list))
	{
		int list_len = PyList_Size(formated_err_list);

		for (int i = 0; i < list_len; i++)
		{
			formated_err_i = PyList_GetItem(formated_err_list, i);

			// getting __add__ function of formated_err
			PyObject* add_name = PyUnicode_DecodeFSDefault("__add__");
			PyObject* add_func = PyObject_GetAttr(formated_err, add_name);

			// calling __add__ and saving the result
			PyObject* add_args = PyTuple_Pack(1, formated_err_i);
			formated_err = PyObject_Call(add_func, add_args, NULL);
		}
	}


	char *pStrErrorMessage = string_from_pyobject2(PyObject_Str(formated_err));

	// Py_DECREF(tracback_module);
	// Py_DECREF(format_exception_name);
	// //Py_DECREF(error_msg_py);
	// Py_DECREF(ptype);
	// Py_DECREF(pvalue);
	// Py_DECREF(traceback);
	// Py_DECREF(args);

	return pStrErrorMessage;
}

*/
import "C"

func GetPyErrorMsg() string {
	pyErrStringC := C.PyErr_GetErrorString()
	pyErrString := C.GoString(pyErrStringC)
	return pyErrString
}

// It will panic, if exception was occured
func HandlePossibleException() {
	errOccured := C.PyErr_Occurred()
	if errOccured == nil {
		// there was no exception
		return
	}

	errMsg := GetPyErrorMsg()

	panic("\n" + errMsg)
}

package gopython

/*
#cgo pkg-config: python3-embed
#include <Python.h>
int run_python_string(const char *command)
{
	return PyRun_SimpleString(command);
}
*/
import "C"

//import "fmt"

// InitPythonInterpretetor inits Python ...
func InitPythonInterpretetor() {
	intrprInited := C.Py_IsInitialized()

	if intrprInited == 0 {
		C.Py_Initialize()
		//fmt.Println("Interpreteror was inited")
		sysBugFixCommand := "import sys\nif not hasattr(sys, 'argv'):\n\tsys.argv  = ['']"
		sysBugFixCommandC := C.CString(sysBugFixCommand)
		C.run_python_string(sysBugFixCommandC)
	}
	//else {
	//	fmt.Println("Interpreteror has already inited")
	//}
}

// FinalizePythonInterpretetor ends work of Python interpretetor...
func FinalizePythonInterpretetor() {
	//fmt.Println("FinalizePythonInterpretetor was called")
	intrprInited := C.Py_IsInitialized()
	//fmt.Printf("isInitstatus = %d\n", intrprInited)

	if intrprInited == 1 {
		C.Py_Finalize()
		//fmt.Println("Interpreteror finilized his work")
	}
	//else {
	//	fmt.Println("Interpreteror has already finilized")
	//}
}

func RunPythonString(pycode string) {
	pycodeC := C.CString(pycode)
	C.run_python_string(pycodeC)
}

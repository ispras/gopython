package gopython

// #cgo pkg-config: python3
// #include <Python.h>
import "C"

//import "fmt"

// InitPythonInterpretetor inits Python ...
func InitPythonInterpretetor35() {
	intrprInited := C.Py_IsInitialized()

	if intrprInited == 0 {
		C.Py_Initialize()
		//fmt.Println("Interpreteror was inited")
	}
	//else {
	//	fmt.Println("Interpreteror has already inited")
	//}
}

// FinalizePythonInterpretetor ends work of Python interpretetor...
func FinalizePythonInterpretetor35() {
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

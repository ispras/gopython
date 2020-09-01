package gopython

// #cgo pkg-config: python-3.5
// #include <python3.5m/Python.h>
import "C"

//import "fmt"

// InitPythonInterpretetor inits Python ...
func InitPythonInterpretetor() {
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

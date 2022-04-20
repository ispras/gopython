package gopython

type errors struct {
	msg string
}

func (e *errors) Error() string {
	return e.msg
}

func (e *errors) notImportedModule() {
	e.msg = `Python module didn't imported.
	         Call makeImport method first`
}

func (e *errors) gettingObjectFailed() {
	e.msg = `Getting object from module was failed`
}

func (e *errors) importError(moduleName string) {
	e.msg = "Can't import module with name " + moduleName
}

func (e *errors) classNotInited() {
	e.msg = "Class didn't init"
}

func (e *errors) noSuchAttr(attrName string) {
	e.msg = "This object doesn't have attribute/method with name " + attrName
}

func (e *errors) nilObjectPointer() {
	e.msg = "Object pointer to PyObject is nil. Initialize it properly first"
}

func (e *errors) errorDuringMethodCall(methodName string) {
	e.msg = "During " + methodName + " call errors were occurred"
}

func (e *errors) notStandartType() {
	e.msg = `Type of pyobject is not standart. You can
			 check it with IsStandartType method`
}

func (e *errors) notSupportedGoSlice() {
	e.msg = `Type of go slice is not supported yet.`
}

func (e *errors) pyListEmpty() {
	e.msg = `Python list should has at least one element.`
}

func (e *errors) notSupportedPyList() {
	e.msg = `The list's elements type is not supported yet. (first element of the list was checked)`
}

func (e *errors) pyListWithDifferentTypes() {
	e.msg = `Python list has elements of different types`
}

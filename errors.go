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

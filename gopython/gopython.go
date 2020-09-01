package gopython

type PythonMethodArguments interface {
	SetArgCount(int)
	SetNextArgument(interface{})
	GetArgumentsTuple() interface{}
}

type PythonObject interface {
	CallMethod(string, PythonMethodArguments) ([]PythonObject, error)
	HasAttr(string) (bool, error)
	GetAttr(string) (PythonObject, error)
	GetType() (string, error)
	IsStandartType() (bool, error)
	ToStandartGoType() (interface{}, error)
}

type PythonClass interface {
	CreateObject(PythonMethodArguments) (PythonObject, error)
}

type PythonModule interface {
	SetModuleName(string)
	MakeImport() error
	GetClass(string) (PythonClass, error)
}

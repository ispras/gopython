package gopython

type PythonObjectsGate struct {
	version string
}

func (pog *PythonObjectsGate) SetVersion(version string) {
	pog.version = version
}

func (pog *PythonObjectsGate) ObtainSystemPythonVersion() error {
	// get system version
	// ...

	pog.version = "3.5"
	return nil
}

func (pog *PythonObjectsGate) InitPythonInterpretetor() error {
	switch pog.version {
	case "3.5":
		InitPythonInterpretetor()
		return nil
	}

	return nil
}

func (pog *PythonObjectsGate) FinalizePythonInterpretetor() error {
	switch pog.version {
	case "3.5":
		FinalizePythonInterpretetor()
		return nil
	}

	return nil
}

func (pog *PythonObjectsGate) GetModule() (PythonModule, error) {
	switch pog.version {
	case "3.5":
		var res PythonModule35
		return &res, nil
	}

	return nil, nil
}

func (pog *PythonObjectsGate) GetArguments() (PythonMethodArguments, error) {
	switch pog.version {
	case "3.5":
		var res PythonMethodArguments35
		return &res, nil
	}

	return nil, nil
}

func (pog *PythonObjectsGate) GetClass() (PythonClass, error) {
	switch pog.version {
	case "3.5":
		var res PythonClass35
		return &res, nil
	}

	return nil, nil
}

func (pog *PythonObjectsGate) GetObject() (PythonObject, error) {
	switch pog.version {
	case "3.5":
		var res PythonObject35
		return &res, nil
	}

	return nil, nil
}

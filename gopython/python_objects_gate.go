package gopython

//import (
//	"os/exec"
//	"strings"
//)

type PythonObjectsGate struct {
	version string
}

func (pog *PythonObjectsGate) SetVersion(version string) {
	pog.version = version
}

// ObtainSystemPythonVersion gets correct python version
func (pog *PythonObjectsGate) ObtainSystemPythonVersion() error {
	/*
		cmd := exec.Command("bash", "-c", "ls /usr/include/ | grep python")

		res, err := cmd.Output()
		if err == nil {
			pythons := strings.Split(string(res), "\n")
			for i := range pythons {
				if strings.Contains(pythons[i], "3.5m") {
					pog.version = "3.5"
					break
				}

				if strings.Contains(pythons[i], "3.6m") {
					pog.version = "3.6"
					break
				}

				if strings.Contains(pythons[i], "3.7m") {
					pog.version = "3.7"
					break
				}
			}
		}
	*/
	pog.version = "3.5"

	return nil
}

func (pog *PythonObjectsGate) InitPythonInterpretetor() error {
	switch pog.version {
	case "3.5":
		InitPythonInterpretetor35()
		return nil
	}

	return nil
}

func (pog *PythonObjectsGate) FinalizePythonInterpretetor() error {
	switch pog.version {
	case "3.5":
		FinalizePythonInterpretetor35()
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

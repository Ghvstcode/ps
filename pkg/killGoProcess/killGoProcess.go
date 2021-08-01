package killGoProcess

import (
	"errors"
	"fmt"
	"log"
	"os"
	goversion "rsc.io/goversion/version"

	"github.com/keybase/go-ps"
)

func isGo(pr ps.Process) (ok bool, err error) {
	if pr.Pid() == 0 {
		return
	}
	path, err := pr.Path()
	if err != nil {
		return false, err
	}
	_, err = goversion.ReadExe(path)
	return err == nil, err
}

func getGolangProcesses() ([]ps.Process, []error) {
	var (
		errGroup  []error
		processes []ps.Process
	)

	processList, err := ps.Processes()
	if err != nil {
		return processes, errGroup
	}

	for _, process := range processList {
		ok, err := isGo(process)
		if err != nil {
			errGroup = append(errGroup, err)
			continue
		} else if ok {
			processes = append(processes, process)
		}
	}

	return processes, errGroup
}

func KillGoProcess() error {
	processes, errs := getGolangProcesses()
	if len(processes) < 1 {
		fmt.Println(errs)
		return errors.New("no Go Processes found")
	}

	for _, process := range processes {
		if os.Getpid() == process.Pid() {
			continue
		}
		os_process, err := os.FindProcess(process.Pid())
		if err != nil {
			log.Println(err)
			return err
		}

		//sigChan := make(chan os.Signal)

		if err = os_process.Signal(os.Kill); err != nil {
			log.Println("69",err)
			continue
		}
	}
	fmt.Println("All 'em Go processes is ded. Ha!")
	return nil
}

package getprocess

import (
	ps "github.com/keybase/go-ps"
	"github.com/olekukonko/tablewriter"
	"os"
	goversion "rsc.io/goversion/version"
	"strconv"
)


//ListAllProcesses returns the PID of all currently running process.
func ListAllProcesses()error{
	var s []string
	var d [][]string

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for i, process := range processList {
		s = []string{
			strconv.Itoa(i),
			process.Executable(),
			strconv.Itoa(process.Pid()),
		}

		d = append(d, s)
	}
	
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S/N", "Process", "PID"})

	for _, v := range d {
		table.Append(v)
	}

	table.Render()
	return nil
}

func isGo(path string) (ok bool, err error) {
	_, err = goversion.ReadExe(path)
	return err == nil, err
}

func ListGoProcess()error{
	var s []string
	var d [][]string

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for i, process := range processList {
		path, err := process.Path()

		if err != nil {
			return err
		}

		ok,_ := isGo(path)
		if ok {
			s = []string{
				strconv.Itoa(i),
				process.Executable(),
				strconv.Itoa(process.Pid()),
			}

			d = append(d, s)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S/N", "Go-Processes", "PID"})

	for _, v := range d {
		table.Append(v)
	}

	table.Render()
	return nil
}
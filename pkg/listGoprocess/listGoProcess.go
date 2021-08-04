package listGoprocess

import (
	"os"
	"strconv"

	"github.com/keybase/go-ps"
	"github.com/olekukonko/tablewriter"
	goversion "rsc.io/goversion/version"
)

type d [][]string

var data d

func isGo(path string) (ok bool, err error) {
	_, err = goversion.ReadExe(path)
	return err == nil, err
}

func ListGoProcess() error {
	var s []string

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for i, process := range processList {
		path, err := process.Path()

		if err != nil {
			return err
		}

		ok, _ := isGo(path)
		if ok {
			s = []string{
				strconv.Itoa(i),
				process.Executable(),
				strconv.Itoa(process.Pid()),
			}

			data = append(data, s)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S/N", "Go-Processes", "PID"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}


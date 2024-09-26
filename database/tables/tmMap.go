package tables

import frmt "server/database/format"

type tmName struct {
	Name    string
	Version string
}

func CreateTmMap(machines []frmt.Machine) map[string][]tmName {
	tmMap := make(map[string][]tmName)

	for _, machine := range machines {
		tm := tmName{
			Name:    machine.TmNo,
			Version: machine.Version,
		}

		tmMap[machine.Move] = append(tmMap[machine.Move], tm)
	}
	return tmMap
}

package utils

import (
	"github.com/denisbrodbeck/machineid"
)

func GetMachineId(appName string) string {
	return MachineId(appName)
}

func MachineId(appKey string) string {
	m, err := machineid.ProtectedID(appKey)
	if err != nil {
		m = "notfound!!!"
	}
	return m
}

package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	TimeMicrossec                int64  // milliseconds
	MachineSource                string `gorm:"varchar(10)"`
	MachineTarget                string `gorm:"varchar(10)"`
	TokenSize                    int
	NumRoles                     int
	EncryptMethod                string
	TimeGeneratingTokenMicrossec int64
}

func NewRequest(
	timeMicrossec int64, machineSource string,
	machineTarget string, tokenSize int,
	numRoles int, encryptMethod string,
	timeGeneratingTokenMicrossec int64) Request {
	return Request{
		TimeMicrossec:                timeMicrossec,
		MachineSource:                machineSource,
		MachineTarget:                machineTarget,
		TokenSize:                    tokenSize,
		NumRoles:                     numRoles,
		EncryptMethod:                encryptMethod,
		TimeGeneratingTokenMicrossec: timeGeneratingTokenMicrossec,
	}
}

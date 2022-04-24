package model

import (
	"time"
)

const AdministratorStatusOk = 1
const AdministratorStatusForbid = 2

type Administrator struct {
	Id          int64
	Username    string
	Password    string
	Salt        string
	Mobile      string
	Nickname    string
	Avatar      string
	Status      int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt int64
}

func (Administrator) TableName() string {
	return "sys_administrator"
}

package db

import "github.com/Template7/common/structs"

const (
	sqlTempDb = "temp"
)

type ClientInterface interface {
	CreateAdmin(admin structs.Admin) (err error)
	InitDb(force bool)
}

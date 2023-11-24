package user

import (
	"cli-tool/internal/backend"
	"github.com/Template7/common/structs"
)

func (u *User) UpdateInfo(data structs.UserInfo) (err error) {
	log.Debug("update user info: ", u.Data.UserId)
	u.Data.BasicInfo = data
	err = backend.New().UpdateUser(u.Data, u.Token.AccessToken)
	return
}

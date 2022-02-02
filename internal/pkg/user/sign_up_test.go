package user

import (
	fakeDataGenerator "cli-tool/internal/pkg/fakeData"
	"github.com/spf13/viper"
	"testing"
)

func BenchmarkSignUp(b *testing.B) {
	viper.AddConfigPath("../../../configs")

	for i := 0; i < b.N; i++ {
		userData := fakeDataGenerator.RandomUser()
		fakeUser := User{
			Data: userData,
		}
		if err := fakeUser.SignUp(); err != nil {
			log.Error("fail to sign up user: ", fakeUser.Data.UserId)
			return
		}

		if err := fakeUser.UpdateInfo(userData.BasicInfo); err != nil {
			log.Error("fail to update user info: ", err.Error())
			return
		}
	}
}

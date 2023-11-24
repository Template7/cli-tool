package user

import (
	fakeDataGenerator "cli-tool/internal/fakeData"
	"github.com/spf13/viper"
	"testing"
)

func BenchmarkSignIn(b *testing.B) {
	viper.AddConfigPath("../../../configs")

	// sign up a user
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

	for i := 0; i < b.N; i++ {
		if err := fakeUser.SignIn(); err != nil {
			log.Fatal(err)
		}
	}
}

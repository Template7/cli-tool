package cli

import (
	fakeDataGenerator "cli-tool/internal/pkg/fakeData"
	"cli-tool/internal/pkg/user"
	"github.com/urfave/cli/v2"
)

var (
	stressTest = &cli.Command{
		Name:    "StressTest",
		Usage:   "Stress test",
		Aliases: []string{"st"},
		//Flags:   fakeDataFlag,
		Action: func(c *cli.Context) error {
			genFakeUser(c.Int("user"))
			return nil
		},
		Subcommands: []*cli.Command{
			batchSignUp,
			batchSignIn,
		},
	}

	batchSignUp = &cli.Command{
		Name:    "BatchSignUp",
		Usage:   "Batch sign up",
		Aliases: []string{"bsu"},
		Flags:   stressTestFlag,
		Action: func(c *cli.Context) error {
			doBatchSignUp(c.Int("user"))
			return nil
		},
	}

	batchSignIn = &cli.Command{
		Name:    "BatchSignIn",
		Usage:   "Batch sign in",
		Aliases: []string{"bsi"},
		Flags:   stressTestFlag,
		Action: func(c *cli.Context) error {
			doBatchSignIn(c.Int("user"))
			return nil
		},
	}
	stressTestFlag = []cli.Flag{
		&cli.IntFlag{
			Name:    "user",
			Aliases: []string{"n"},
			Usage:   "Number of user",
			Value:   10,
		},
		//&cli.IntFlag{
		//	Name:    "delay",
		//	Aliases: []string{"d"},
		//	Usage:   "Number of user",
		//	Value:   10,
		//},
	}
)

func doBatchSignUp(count int) {
	log.Debug("do batch sign up: ", count)

	for i := 0; i < count; i++ {
		userData := fakeDataGenerator.RandomUser()
		fakeUser := user.User{
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

	log.Debug("finish batch sign up")
}

func doBatchSignIn(count int) {
	log.Debug("do batch sign in: ", count)

	for i := 0; i < count; i++ {
		userData := fakeDataGenerator.RandomUser()
		fakeUser := user.User{
			Data: userData,
		}
		if err := fakeUser.SignIn(); err != nil {
			log.Error("fail to sign up user: ", fakeUser.Data.UserId)
			return
		}
	}

	log.Debug("finish batch sign in")
}

//func genFakeUser(number int) {
//	for i := 0; i < number; i++ {
//		fakeUser := fakeDataGenerator.RandomUser()
//		data := apiBody.CreateUserReq{
//			Mobile: fakeUser.Mobile,
//			Email:  fakeUser.Email,
//		}
//		if err := backend.New().CreateUser(data); err != nil {
//			log.Fatal(err)
//		}
//
//	}
//}

package cmd

import (
	fakeDataGenerator "cli-tool/internal/fakeData"
	"cli-tool/internal/user"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
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
		Flags:   batchSignUpFlag,
		Action: func(c *cli.Context) error {
			doBatchSignUp(c.Int("user"), c.String("out"))
			return nil
		},
	}

	batchSignIn = &cli.Command{
		Name:    "BatchSignIn",
		Usage:   "Batch sign in",
		Aliases: []string{"bsi"},
		Flags:   batchSignInFlag,
		Action: func(c *cli.Context) error {
			doBatchSignIn(c.String("file"))
			return nil
		},
	}

	batchSignUpFlag = []cli.Flag{
		&cli.IntFlag{
			Name:    "user",
			Usage:   "Number of user",
			Aliases: []string{"n"},
			Value:   10,
		},
		&cli.StringFlag{
			Name:    "out",
			Usage:   "Write sign up mobile list to yaml file",
			Aliases: []string{"o"},
		},
	}
	batchSignInFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Usage:   "Read mobile list from yaml file for batch sign in",
			Aliases: []string{"f"},
		},
	}
)

func doBatchSignUp(count int, out string) {
	log.Debug("do batch sign up: ", count)

	mobileList := make([]string, count)

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
		mobileList[i] = userData.Mobile
	}

	data, err := yaml.Marshal(&mobileList)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(out, data, 0666); err != nil {
		log.Fatal(err)
	}

	log.Debug("finish batch sign up")
}

func doBatchSignIn(file string) {
	log.Debug("do batch sign in")

	// read mobile list from file
	var mobileList []string
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &mobileList); err != nil {
		log.Fatal(err)
	}

	for _, mobile := range mobileList {
		tempUser := user.User{}
		tempUser.Data.Mobile = mobile
		if err := tempUser.SignIn(); err != nil {
			log.Error("fail to sign up user: ", tempUser.Data.UserId)
			return
		}

		if err := tempUser.GetInfo(); err != nil {
			log.Error("fail to get user info: ", err.Error())
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

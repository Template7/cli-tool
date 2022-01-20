package cli

import (
	"cli-tool/internal/pkg/backend"
	fakeDataGenerator "cli-tool/internal/pkg/fakeData"
	"github.com/Template7/backend/pkg/apiBody"
	"github.com/urfave/cli/v2"
)

var (
	fakeData = &cli.Command{
		Name:    "FakeData",
		Usage:   "Write some fake data to DB",
		Aliases: []string{"fd"},
		Flags:   fakeDataFlag,
		Action: func(c *cli.Context) error {
			genFakeUser(c.Int("user"))
			return nil
		},
	}

	fakeDataFlag = []cli.Flag{
		&cli.IntFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Number of fake user",
			Value:   10,
		},
	}
)

func genFakeUser(number int) {
	for i := 0; i < number; i++ {
		fakeUser := fakeDataGenerator.RandomUser()
		data := apiBody.CreateUserReq{
			Mobile: fakeUser.Mobile,
			Email:  fakeUser.Email,
		}
		if err := backend.New().CreateUser(data); err != nil {
			log.Fatal(err)
		}

	}
}

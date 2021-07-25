package fakeData

import (
	"cli-tool/internal/pkg/db/collection"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"time"
)

func RandomUser() collection.User {
	name := randomdata.SillyName()
	return collection.User{
		BasicInfo: collection.UserInfo{
			NickName: name,
			Gender:   Candidate.Gender[rand.Intn(len(Candidate.Gender))],
			Bio: randomdata.Paragraph(),
			//Birthday: randomdata.FullDate(),
		},
		Mobile: randomMobile(),
		Email:  fmt.Sprintf("%s@%s.com", name, Candidate.Mail[rand.Intn(len(Candidate.Mail))]),
		LoginClient: collection.LoginInfo{
			Device:  Candidate.LoginDevice[rand.Intn(len(Candidate.LoginDevice))],
			Channel: Candidate.LoginChannel[rand.Intn(len(Candidate.LoginChannel))],
		},
	}
}

func randomMobile() string {
	rand.Seed(time.Now().Unix())
	prefix := Candidate.MobilePrefix[rand.Intn(len(Candidate.MobilePrefix))]
	number := fmt.Sprintf("%010d", rand.Int63n(1e10))
	return prefix + number
}

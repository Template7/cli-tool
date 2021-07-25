package fakeData

import (
	"cli-tool/internal/pkg/db/collection"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"time"
)

const (
	avatar = "https://fakeimg.pl/128x128/?text=%s"
)

func RandomUser() collection.User {
	name := randomdata.SillyName()
	birthday, _ := time.Parse(randomdata.DateOutputLayout, randomdata.FullDateInRange("1990-01-01", "2010-01-01"))
	avatarUrl := fmt.Sprintf(avatar, name)
	return collection.User{
		BasicInfo: collection.UserInfo{
			NickName: name,
			Avatar:   avatarUrl,
			Gender:   Candidate.Gender[rand.Intn(len(Candidate.Gender))],
			Bio:      randomdata.Paragraph(),
			Birthday: birthday.Unix(),
		},
		Mobile: randomMobile(),
		Email:  fmt.Sprintf("%s@%s.com", name, Candidate.Mail[rand.Intn(len(Candidate.Mail))]),
		LoginClient: collection.LoginInfo{
			Os:      Candidate.LoginClientOs[rand.Intn(len(Candidate.LoginClientOs))],
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

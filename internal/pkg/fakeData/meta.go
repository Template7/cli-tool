package fakeData

import "cli-tool/internal/pkg/db/collection"

var (
	Candidate = struct {
		Mail         []string
		MobilePrefix []string
		Gender       []collection.Gender
		LoginChannel []collection.LoginChannel
		LoginDevice  []string
	}{
		Mail:         []string{"gmail", "icloud", "backend", "yahoo"},
		MobilePrefix: []string{"+886", "+1", "+86"},
		Gender:       []collection.Gender{collection.GenderUnknown, collection.GenderMale, collection.GenderFemale, collection.GenderGay, collection.GenderLesbian},
		LoginChannel: []collection.LoginChannel{collection.LoginChannelUnknown, collection.LoginChannelMobile, collection.LoginChannelFacebook, collection.LoginChannelGoogle},
		LoginDevice:  []string{"iPhone", "Pixel", "Samsung", "Android", "PC", "Unknown"},
	}
)

package fakeData

import "github.com/Template7/common/structs"

var (
	Candidate = struct {
		Mail          []string
		MobilePrefix  []string
		Gender        []structs.Gender
		LoginClientOs []structs.LoginClientOs
		LoginChannel  []structs.LoginChannel
		LoginDevice   []string
	}{
		Mail:          []string{"gmail", "icloud", "backend", "yahoo"},
		MobilePrefix:  []string{"+886", "+1", "+86"},
		Gender:        []structs.Gender{structs.GenderUnknown, structs.GenderMale, structs.GenderFemale, structs.GenderGay, structs.GenderLesbian},
		LoginClientOs: []structs.LoginClientOs{structs.LoginClientOsUnknown, structs.LoginClientOsIos, structs.LoginClientOsAndroid},
		LoginChannel:  []structs.LoginChannel{structs.LoginChannelUnknown, structs.LoginChannelMobile, structs.LoginChannelFacebook, structs.LoginChannelGoogle},
		LoginDevice:   []string{"iPhone", "Pixel", "Samsung", "PC", "Unknown"},
	}
)

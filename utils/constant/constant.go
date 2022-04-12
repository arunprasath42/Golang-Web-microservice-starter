package constant

import "regexp"

var (
	TestTextForSlack = "Test message Please ignore !!"
	UserNameForSlack = "Accuknox"
	Alphabets        = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	Digits           = regexp.MustCompile(`^[0-9]+$`).MatchString
	Email            = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString
)

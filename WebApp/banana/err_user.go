package banana

import "errors"

var (
	UserConflict= errors.New("user already exist!")
	SignUpFailed= errors.New("Sign Up Failed!")

)
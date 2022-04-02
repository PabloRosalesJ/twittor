package bd

import "golang.org/x/crypto/bcrypt"

func BcryptPwd(pwd string) (string, error) {
	cost := 8
	pwdE, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	return string(pwdE), err
}

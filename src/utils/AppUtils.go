package utils

import (
	"golang.org/x/crypto/bcrypt"
	"path"
	"runtime"
)

func EncodePassword(in string) (out string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	out = string(hash)
	return out, nil
}

func PasswordMatch(h1, h2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h1), []byte(h2))
	return err == nil
}

func GoFilePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

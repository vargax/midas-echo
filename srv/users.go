package srv

import (
	"errors"
	"github.com/vargax/midas-echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSrv struct {
	s midas.StorageSrv
}

func NewUserSrv(s midas.StorageSrv) *UserSrv {
	return &UserSrv{s: s}
}

func (us *UserSrv) New(u *midas.User) error {
	pass, err := encodePassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pass

	err = us.s.CreateUser(u)
	if errors.Is(err, midas.DuplicateKey) {
		return midas.UserAlreadyRegistered
	}
	return err
}

func (us *UserSrv) User(username string) (*midas.User, error) {
	u := midas.User{Username: username}
	err := us.s.SelectUser(&u)
	return &u, err
}

func (us *UserSrv) Authenticate(u *midas.User) error {
	pass, err := encodePassword(u.Password)
	if err != nil {
		return err
	}

	user := midas.User{Username: u.Username}

	err = us.s.SelectUser(&user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return midas.UserNotFound
	}
	if err != nil {
		return err
	}

	if !passwordMatch(pass, user.Password) {
		return midas.PasswordDontMatch
	}

	u = &user
	return nil
}

func encodePassword(in string) (out string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	out = string(hash)
	return out, nil
}

func passwordMatch(h1, h2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h1), []byte(h2))
	return err == nil
}

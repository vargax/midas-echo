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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

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
	pass := u.Password
	u.Password = ""

	err := us.s.SelectUser(u)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return midas.UserNotFound
	}
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return midas.PasswordDontMatch
	}

	return nil
}

package model

import (
	"errors"
	"strings"
	"time"
	"v1/monorepo/util/authentication"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Password struct {
	New string `json:"new"`
	Old string `json:"old"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required ")
	}

	if user.Nick == "" {
		return errors.New("nick is required ")
	}

	if user.Email == "" {
		return errors.New("email is required ")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is required and shouldn't be blank")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordWithHash, err := authentication.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}
	return nil
}

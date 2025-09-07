package service

import (
	"errors"
	"v1/monorepo/users/model"
	"v1/monorepo/users/repository"
)

type UserService interface {
	Create(user *model.User) (uint64, error)
	Get(nameOrNick string) ([]model.User, error)
	Update(userID uint64, userModel *model.User, tokenID uint64) error
	Follow(userId, followerID uint64) error
	Delete(userID uint64, tokenID uint64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *model.User) (uint64, error) {
	if err := user.Prepare("register"); err != nil {
		return 0, err
	}
	return s.repo.Create(*user)
}

func (s *userService) Get(nameOrNick string) ([]model.User, error) {
	if nameOrNick != "" {
		return s.repo.FindUserByID(nameOrNick)
	}

	return nil, errors.New("name or nick search null, nothing to find")
}

func (s *userService) Update(userID uint64, updated *model.User, tokenID uint64) error {
	if userID != tokenID {
		return errors.New("account doesn't match")
	}

	if err := updated.Prepare("edit"); err != nil {
		return err
	}
	return s.repo.Update(userID, *updated)
}

func (s *userService) Delete(userID uint64, tokenID uint64) error {
	if userID != tokenID {
		return errors.New("account doesn't match")
	}
	return s.repo.Delete(userID)
}

func (s *userService) Follow(userID, followerID uint64) error {
	if userID == followerID {
		return errors.New("you can't follow yourself")
	}
	// Implement repo.Follow(userID, followerID)
	return nil
}

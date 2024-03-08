package service

import (
	"auction/internal/user/dto"
	"auction/internal/user/model"
	"auction/internal/user/repository"
	"auction/pkg/jtoken"
	"auction/pkg/utils"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, int64, error)
	Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error)
	GetUserByID(ctx context.Context, id uint32) (*model.User, error)
	RefreshToken(ctx context.Context, userID uint32) (string, int64, error)
	ChangePassword(ctx context.Context, id uint32, req *dto.ChangePasswordReq) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, int64, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return nil, "", "", 0, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", 0, errors.New("invalid credentials")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"roles": user.Roles,
	}
	accessToken, expiresIn := jtoken.GenerateAccessToken(tokenData)
	refreshToken, _ := jtoken.GenerateRefreshToken(tokenData)
	return user, accessToken, refreshToken, expiresIn, nil
}

func (s *UserService) Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error) {
	var user model.User
	utils.Copy(&user, &req)
	err := s.repo.Create(ctx, &user)
	if err != nil {
		log.Printf("Register.Create fail, email: %s, error: %s\n", req.Email, err)
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint32) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("GetUserByID fail, id: %v, error: %s", id, err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) RefreshToken(ctx context.Context, userID uint32) (string, int64, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		log.Printf("RefreshToken.GetUserByID fail, id: %v, error: %s", userID, err)
		return "", 0, err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"roles": user.Roles,
	}
	accessToken, expiresIn := jtoken.GenerateAccessToken(tokenData)
	return accessToken, expiresIn, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uint32, req *dto.ChangePasswordReq) error {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("ChangePassword.GetUserByID fail, id: %v, error: %s", id, err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("wrong password")
	}

	user.Password = utils.HashPassword([]byte(req.NewPassword))
	err = s.repo.Update(ctx, user)
	if err != nil {
		log.Printf("ChangePassword.Update fail, id: %v, error: %s", id, err)
		return err
	}

	return nil
}

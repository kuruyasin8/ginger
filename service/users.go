package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kuruyasin8/ginger/config"
	"github.com/kuruyasin8/ginger/errors"
	"github.com/kuruyasin8/ginger/model"
	"github.com/kuruyasin8/ginger/stash"
	"go.mongodb.org/mongo-driver/bson"
)

type UserQuery struct {
	ID      uint   `param:"uid"`
	Page    uint   `query:"page"`
	PerPage uint   `query:"per_page"`
	Filter  string `query:"filter"`
}

func (s *Service) Register(ctx context.Context, user *model.User) error {
	salt := make([]byte, 32)
	rand.Read(salt)

	raw := append([]byte(user.Password), salt...)
	hash := sha256.Sum256(raw)

	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
	encodedSalt := base64.StdEncoding.EncodeToString(salt[:])

	user.Credentials = new(model.Credentials)
	user.Credentials.Hash = encodedHash
	user.Credentials.Salt = encodedSalt
	user.Credentials.Verified = true
	now := time.Now().UnixMilli()
	user.Credentials.CreatedAt = now
	user.Credentials.ModifiedAt = now

	if err := s.usersRepository.InsertSingleUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(ctx context.Context, payload *model.User) (interface{}, error) {
	filter := bson.M{"email": payload.Email}
	user, err := s.usersRepository.GetSingleUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	rawSalt := make([]byte, 32)
	base64.StdEncoding.Decode(rawSalt, []byte(user.Credentials.Salt))

	raw := append([]byte(payload.Password), rawSalt...)
	hash := sha256.Sum256(raw)

	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	if encodedHash != user.Credentials.Hash {
		return nil, errors.NewForbidden("authentication failed")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).UnixMilli(),
		"roles": []string{string(stash.Soy), string(stash.Pepper), string(stash.Salt)},
	})

	accessToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return nil, err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).UnixMilli(),
	})

	refreshToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken *model.Token) (interface{}, error) {
	verifiedRefreshToken, err := jwt.Parse(refreshToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return nil, errors.NewForbidden("invalid refresh token")
	}

	if err := verifiedRefreshToken.Method.Verify(verifiedRefreshToken.Raw, verifiedRefreshToken.Signature, []byte(config.Secret)); err != nil {
		return nil, errors.NewForbidden("invalid refresh token")
	}

	validUntil := verifiedRefreshToken.Claims.(jwt.MapClaims)["exp"].(float64)
	if time.Now().UnixMilli() > int64(validUntil) {
		return nil, errors.NewForbidden("refresh token expired")
	}

	filter := bson.M{"email": verifiedRefreshToken.Claims.(jwt.MapClaims)["email"].(string)}
	user, err := s.usersRepository.GetSingleUser(ctx, filter)
	if err != nil {
		return nil, errors.NewNotFound("user not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).UnixMilli(),
		"roles": []string{string(stash.Soy), string(stash.Pepper), string(stash.Salt)},
	})

	accessToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": accessToken,
	}, nil
}

func (s *Service) GetSingleUser(ctx context.Context, query *UserQuery) (*model.User, error) {
	filter := bson.M{"_id": query.ID}

	user, err := s.usersRepository.GetSingleUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetMultipleUsers(ctx context.Context, query *UserQuery) ([]*model.User, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query.Filter, "$options": "i"}},
		},
	}

	users, err := s.usersRepository.GetMultipleUsers(ctx, filter)
	if err != nil {
		return nil, err
	}

	if len(users) < int(query.Page*query.PerPage) {
		return users, nil
	}

	users = users[(query.Page-1)*query.PerPage : query.PerPage*query.Page]

	return users, nil
}

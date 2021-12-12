package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mby/minicom/auth/internal/auth/types"
	"github.com/mby/minicom/auth/internal/cfg"
	"github.com/mby/minicom/auth/internal/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

type IRepo interface {
	Cleanup()
	CreateUser(username, password string) error
	Login(username, password string) (string, error)
}

type Repo struct {
	cfg    cfg.Config
	client *mongo.Client
	users  *mongo.Collection
}

func NewRepo(cfg cfg.Config) IRepo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	users := client.Database(cfg.MongoDB).Collection(cfg.UsersCollection)

	return Repo{cfg, client, users}
}

func (r Repo) Cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (r Repo) CreateUser(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return errors.FailedHashingPassword
	}

	hashedPassword := string(bytes)
	_, err = r.users.InsertOne(ctx, types.User{
		Username: username,
		Password: hashedPassword,
	})
	if err != nil {
		fmt.Println(err)
		return errors.UserAlreadyExists
	}

	return nil
}

func (r Repo) Login(username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user types.User
	err := r.users.FindOne(ctx, types.User{Username: username}).Decode(&user)
	if err != nil {
		return "", errors.UserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.InvalidPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	tokenStr, err := token.SignedString([]byte(r.cfg.JWTSecret))
	if err != nil {
		return "", errors.FailedGeneratingToken
	}

	return tokenStr, nil
}
package models

import (
	"../utilities"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int64
}

func NewUser(username string, hash []byte, email Email, phone Phone, birth string, gender Gender) (*User, error) {
	var exists, err1 = redisClient.HExists(ctx, "user:by-username", username).Result()

	if err1 != nil {
		return nil, err1
	} else if exists {
		return nil, utilities.ErrUsernameTaken
	}

	var id, err2 = redisClient.Incr(ctx, "user:next-id").Result()
	if err2 != nil {
		return nil, err2
	}

	var (
		key = getUserKey(id)
		pipe = redisClient.Pipeline()
	)
	pipe.HSet(ctx, key, "id", id)
	pipe.HSet(ctx, key, "username", username)
	pipe.HSet(ctx, key, "email", email.address)
	pipe.HSet(ctx, key, "email:hide", email.isHide)
	pipe.HSet(ctx, key, "phone", phone.number)
	pipe.HSet(ctx, key, "phone:hide", phone.isHide)
	pipe.HSet(ctx, key, "birth", birth)
	pipe.HSet(ctx, key, "gender", gender.ToString())
	pipe.HSet(ctx, key, "hash", hash)
	pipe.HSet(ctx, "user:by-username", username, id)
	_, err2 = pipe.Exec(ctx)
	if err2 != nil {
		return nil, err2
	}

	return &User{id}, nil
}

func (user *User) GetUsername() (string, error) {
	return redisClient.HGet(ctx, getUserKey(user.id), "username").Result()
}

func (user *User) GetHash() ([]byte, error) {
	return redisClient.HGet(ctx, getUserKey(user.id), "hash").Bytes()
}

func (user *User) GetBirth() (string, error) {
	return redisClient.HGet(ctx, getUserKey(user.id), "birth").Result()
}

func (user *User) GetGender() (string, error) {
	return redisClient.HGet(ctx, getUserKey(user.id), "gender").Result()
}

func (user *User) GetId() (int64, error) {
	return user.id, nil
}

func (user *User) GetEmail() (string, error) {
	var hide, err1 = redisClient.HGet(ctx, getUserKey(user.id), "email:hide").Int()
	if err1 != nil {
		return "error", err1
	}

	var email, err2 = redisClient.HGet(ctx, getUserKey(user.id), "email").Result()
	if err2 != nil {
		return "error", err2
	} else if hide == 1 {
		var result = string(email[0])
		for i := 1; i < len(email) - 1; i++ {
			result += "*"
		}
		result += string(email[len(email) - 1])
		return result, nil
	} else {
		return email, nil
	}
}

func (user *User) GetPhone() (string, error) {
	var hide, err1 = redisClient.HGet(ctx, getUserKey(user.id), "phone:hide").Int()
	if err1 != nil {
		return "error", err1
	}

	var phone, err2 = redisClient.HGet(ctx, getUserKey(user.id), "phone").Result()
	if err2 != nil {
		return "error", err2
	} else if hide == 1 {
		var result = string(phone[0])
		result += string(phone[1])
		result += string(phone[2])
		for i := 3; i < len(phone) - 1; i++ {
			result += "*"
		}
		result += string(phone[len(phone) - 1])
		return result, nil
	} else {
		return phone, nil
	}
}

func GetUserByUsername(username string) (*User, error) {
	var id, err =  redisClient.HGet(ctx, "user:by-username", username).Int64()
	if err == redis.Nil {
		return nil, utilities.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return GetUserById(id)
}

func GetUserById(id int64) (*User, error) {
	return &User{id}, nil
}

func (user *User) Authenticate(password string) error {
	var hash, err = user.GetHash()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return utilities.ErrWrongPassword
	}
	return err
}

func RegisterUser(username, password string, email Email, phone Phone, birth string, gender Gender) error {
	if len(username) < 3 {
		return utilities.ErrTooShortUsername
	} else if len(password) < 3 {
		return utilities.ErrTooShortPassword
	} else if len(email.address) < 1 {
		return utilities.ErrEmptyEmail
	} else if len(phone.number) < 1 {
		return utilities.ErrEmptyPhone
	} else if len(birth) < 1 {
		return utilities.ErrEmptyBirth
	}
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
 	if err != nil {
		return err
	}
	_, err = NewUser(username, hash, email, phone, birth, gender)
	return err
}

func AuthenticateUser(username, password string) (*User, error) {
	var user, err = GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, user.Authenticate(password)
}

func getUserKey(id int64) string {
	return fmt.Sprintf("user:%d", id)
}
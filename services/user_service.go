package services

import (
	"encoding/json"
	"errors"
	"log"

	"go.etcd.io/bbolt"
	"golang.org/x/crypto/bcrypt"
)

const BUCKET_USERS = "users"

type StoredUser struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Hash     []byte `json:"hash"`
}

func InitUsers() {
	// Ensure users bucket exists
	_ = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_USERS))
		return err
	})

	// If no users, create default admin
	hasAny := false
	_ = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_USERS))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		k, _ := c.First()
		if k != nil {
			hasAny = true
		}
		return nil
	})

	if !hasAny {
		if err := CreateUser("admin", "admin123", "admin"); err != nil {
			log.Printf("failed to create default admin: %v", err)
		}
	}
}

func CreateUser(username, password, role string) error {
	if username == "" || password == "" {
		return errors.New("empty username or password")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	su := StoredUser{Username: username, Role: role, Hash: hash}
	b, err := json.Marshal(su)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_USERS))
		if bkt == nil {
			return errors.New("users bucket not found")
		}
		return bkt.Put([]byte(username), b)
	})
}

func ValidateCredentials(username, password string) (StoredUser, error) {
	var su StoredUser
	err := db.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_USERS))
		if bkt == nil {
			return errors.New("users bucket missing")
		}
		v := bkt.Get([]byte(username))
		if v == nil {
			return errors.New("user not found")
		}
		return json.Unmarshal(v, &su)
	})
	if err != nil {
		return StoredUser{}, err
	}
	if err := bcrypt.CompareHashAndPassword(su.Hash, []byte(password)); err != nil {
		return StoredUser{}, errors.New("invalid credentials")
	}
	return su, nil
}

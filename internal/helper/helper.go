package helper

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func CompareHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func ConvertStringToDate(data string) (time.Time, error) {
	layOut := "02-01-2006"
	date, err := time.Parse(layOut, data)
	if err != nil {
		return time.Time{}, fmt.Errorf("error while converting to time")
	}
	return date, nil
}

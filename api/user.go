package api

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"ID"`
	Flag      string    `gorm:"-:all" json:"flag"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) Validate(Flag string) error {
	match, _ := regexp.MatchString("20\\d{2}\\w{4}\\d{4}H\\b", u.ID)
	if !match {
		return errors.New("Invalid ID")
	}
	if u.Flag != Flag {
		return errors.New("Incorrect Flag")
	}
	return nil
}

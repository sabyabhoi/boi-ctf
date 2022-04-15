package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	u := User{
		ID:        "2020B3A72147H",
		Flag:      "boiCTF{TEST-AAA-CORRECT}",
		CreatedAt: time.Now(),
	}

	err := u.Validate("boiCTF{TEST-AAA-CORRECT}")
	assert.NoError(t, err)
}

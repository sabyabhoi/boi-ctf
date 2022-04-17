package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate1(t *testing.T) {
	u := User{
		ID:        "2020B3A72147H",
		Flag:      "boiCTF{TEST-AAA-CORRECT}",
		CreatedAt: time.Now(),
	}

	err := u.Validate("boiCTF{TEST-AAA-CORRECT}")
	assert.NoError(t, err)
}

func TestUser_Validate2(t *testing.T) {
	u := User{
		ID:        "2020B3A72147HA",
		Flag:      "boiCTF{TEST-AAA-CORRECT}",
		CreatedAt: time.Now(),
	}

	err := u.Validate("boiCTF{TEST-AAA-CORRECT}")
	assert.EqualError(t, err, "Invalid ID")
}

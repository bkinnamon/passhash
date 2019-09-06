package passhash

import (
	"testing"
)

func printError(e string, r string, t *testing.T) {
	t.Errorf("\nExpected: %s\nReceived: %s", e, r)
}

func TestGenerateFromPassword(t *testing.T) {
	password := "a password!#%"
	hash, err := generateFromPassword(password)
	if err != nil {
		t.Error(err)
	}

	result, err := verifyPassword(password, hash)
	if err != nil {
		t.Error(err)
	}
	if !result {
		printError("true", "false", t)
	}
}

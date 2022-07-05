package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedpassword, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)

	err = CheckPassword(hashedpassword, password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(hashedpassword, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}

package authtoken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"time"
)

// Validate -- called to validate the supplied token
func Validate(sharedSecret string, token string) bool {

	// Initialize a new instance of the standard claims
	claims := &jwt.StandardClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(sharedSecret), nil
	})

	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: JWT parse returns: %s", err.Error()))
		return false
	}

	if !tkn.Valid {
		logger.Log(fmt.Sprintf("ERROR: JWT is INVALID"))
		return false
	} else {
		logger.Log(fmt.Sprintf("INFO: token is valid, Expires %s", time.Unix(claims.ExpiresAt, 0)))
	}
	return true
}

//
// end of file
//

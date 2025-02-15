package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret" //secret necessary for signing tokens
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error{
	//jwt.Parse is used to check, with the anonymous key func, that the input token
	// was signed using the signing method that we specified when generating the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //checking if the input token was signed with hmac
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
		//returns secret key that is used by the jwt library parse function
	})

	if err != nil {
		return errors.New("Could not parse token.")
	}

	if !parsedToken.Valid {
		return errors.New("Token is not valid")

	}
	return nil
	// if !tokenIsValid {
	// 	return errors.New("Token is not valid")
	// }
	// claims, ok :=	parsedToken.Claims.(jwt.MapClaims) //type checking syntax, checks if Claims if of type jwt.MapClaims
	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }
	// email :=	claims["email"].(string)
	// userId :=	claims["userId"].(int64)
}

package middleware

import (
	"fmt"
	"meetup/utils/logging"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

type TokenUser struct {
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
}

var jwtKey = []byte("gtuw62jecs0023jncwyjj")

func GenerateJWTToken(mobile string, name string) string {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Name:   name,
		Mobile: mobile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "api",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		err_string := fmt.Sprintf("error during token.SignedString: %v\n", err)
		logging.Error(err_string)
	}
	return tokenString
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	// log.Printf(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	// fmt.Println(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*TokenUser, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// fmt.Println(claims)
	if ok && token.Valid {
		mobile, ok := claims["mobile"].(string)
		if !ok {
			return nil, err
		}
		usr_name, ok := claims["name"].(string)
		if !ok {
			return nil, err
		}
		return &TokenUser{
			Mobile: mobile,
			Name:   usr_name,
		}, nil
	}
	return nil, err
}

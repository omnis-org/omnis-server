package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/net"
	"github.com/omnis-org/omnis-server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	Username string
	jwt.StandardClaims
}

type UserToken struct {
	Id        int32      `json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Admin     bool       `json:"admin"`
	Token     string     `json:"token"`
	ExpireAt  *time.Time `json:"exprireAt"`
}

var InvalidTokenError error = errors.New("Invalid Token")
var AlreadyExistError error = errors.New("User already exist")
var InvalidCredentialError error = errors.New("Invalid Credentials")

func getPrivKey(jwtClaims *JWTClaims, token **jwt.Token) (interface{}, error) {
	var err error
	adminConf := config.GetConfig().Admin
	var key interface{}
	if adminConf.AuthKeyFile != "" && adminConf.AuthPubFile != "" { // rsa
		*token = jwt.NewWithClaims(jwt.SigningMethodRS512, jwtClaims)
		key, err = utils.ParsePrivKey(adminConf.AuthKeyFile)
		if err != nil {
			return "", fmt.Errorf("parsePrivKey failed <- %v", err)
		}
	} else if len(adminConf.AuthSimpleKey) != 0 { // key
		*token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwtClaims)
		key = adminConf.AuthSimpleKey
	} else {
		return "", errors.New("Invalid keys")
	}
	return key, nil
}

func createToken(id int32, username string, firstName string, lastName string, admin bool) (*UserToken, error) {
	var err error
	var token *jwt.Token
	adminConf := config.GetConfig().Admin
	expirationTokenTime := time.Now().Add(time.Duration(adminConf.ExpirationTokenTime) * time.Minute)

	jwtClaims := JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTokenTime.Unix(),
		},
	}

	key, err := getPrivKey(&jwtClaims, &token)
	if err != nil {
		return nil, fmt.Errorf("getPrivKey failed <- %v", err)
	}

	tokenSignedString, err := token.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("token.SignedString failed <- %v", err)
	}

	return &UserToken{Id: id, Username: username, FirstName: firstName,
		LastName: lastName, Admin: admin, Token: tokenSignedString, ExpireAt: &expirationTokenTime}, nil

}

func Login(user *model.User) (*UserToken, error) {
	userDB, err := net.GetUserByUsername(user.Username.String)
	if err != nil {
		return nil, fmt.Errorf("net.GetUserByUsername failed <- %v", err)
	}

	if !userDB.Valid() {
		return nil, InvalidCredentialError
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password.String), []byte(user.Password.String))
	if err != nil {
		return nil, InvalidCredentialError
	}
	// auth ok
	return createToken(user.Id.Int32, user.Username.String, user.FirstName.String,
		user.LastName.String, user.Admin.Bool)
}

func Register(user *model.User) error {

	userDB, err := net.GetUserByUsername(user.Username.String)
	if err != nil {
		return fmt.Errorf("net.GetUserByUsername failed <- %v", err)
	}

	if userDB.Valid() {
		return AlreadyExistError
	}

	enc, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword failed <- %v", err)
	}

	err = user.Password.Scan(enc)
	if err != nil {
		return fmt.Errorf("user.Password.Scan failed <- %v", err)
	}

	_, err = net.InsertUser(user)

	if err != nil {
		return fmt.Errorf("net.InsertUser failed <- %v", err)
	}

	return nil
}

func getPubKey(token *jwt.Token) (interface{}, error) {
	adminConf := config.GetConfig().Admin
	var key interface{}
	var err error
	if adminConf.AuthKeyFile != "" && adminConf.AuthPubFile != "" { // rsa
		key, err = utils.ParsePubKey(adminConf.AuthPubFile)
		if err != nil {
			return nil, fmt.Errorf("parsePubKey failed <- %v", err)
		}
	} else if len(adminConf.AuthSimpleKey) != 0 { // key
		key = adminConf.AuthSimpleKey
	} else {
		return nil, errors.New("Invalid keys")
	}
	return key, nil
}

func CheckToken(token string) (*UserToken, error) {
	jwtClaims := &JWTClaims{}
	tokenParse, err := jwt.ParseWithClaims(token, jwtClaims, getPubKey)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, InvalidTokenError
		}

		return nil, fmt.Errorf("jwt.ParseWithClaims failed <- %v", err)
	}

	if !tokenParse.Valid {
		return nil, InvalidTokenError
	}

	expire := time.Unix(jwtClaims.ExpiresAt, 0)
	return &UserToken{Username: jwtClaims.Username, Token: token, ExpireAt: &expire}, nil
}

func RefreshToken(token string) (*UserToken, error) {
	userToken, err := CheckToken(token)
	if err != nil {
		if err == InvalidTokenError {
			return nil, err
		}
		return nil, fmt.Errorf("CheckToken failed <- %v", err)
	}

	if userToken.ExpireAt.Sub(time.Now()) > 2*time.Minute {
		return userToken, nil
	}

	return createToken(userToken.Id, userToken.Username, userToken.FirstName, userToken.LastName, userToken.Admin)

}

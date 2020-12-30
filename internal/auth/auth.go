package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
	"github.com/omnis-org/omnis-server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims should have a comment.
type JWTClaims struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Admin     bool   `json:"admin"`
	jwt.StandardClaims
}

// UserToken should have a comment.
type UserToken struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}

// ErrInvalidToken should have a comment.
var ErrInvalidToken error = errors.New("Invalid Token")

// ErrAlreadyExist should have a comment.
var ErrAlreadyExist error = errors.New("User already exist")

// ErrNotExist should have a comment.
var ErrNotExist error = errors.New("User not exist")

// ErrInvalidCredential should have a comment.
var ErrInvalidCredential error = errors.New("Invalid Credentials")

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

func createToken(id int32, username string, firstName string, lastName string, admin bool) (*UserToken, error) {
	var err error
	var token *jwt.Token
	adminConf := config.GetConfig().Admin
	expirationTokenTime := time.Now().Add(time.Duration(adminConf.ExpirationTokenTime) * time.Minute)

	jwtClaims := JWTClaims{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Admin:     admin,
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

	return &UserToken{Token: tokenSignedString, ExpireAt: expirationTokenTime.Unix()}, nil
}

// Login should have a comment.
func Login(user *model.User) (*UserToken, error) {
	userDB, err := db.GetUserByUsername(user.Username.String)
	if err != nil {
		return nil, fmt.Errorf("net.GetUserByUsername failed <- %v", err)
	}

	if !userDB.Valid() {
		return nil, ErrInvalidCredential
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password.String), []byte(user.Password.String))
	if err != nil {
		return nil, ErrInvalidCredential
	}
	// auth ok
	return createToken(userDB.ID.Int32, userDB.Username.String, userDB.FirstName.String,
		userDB.LastName.String, userDB.Admin.Bool)
}

// Register should have a comment.
func Register(user *model.User) error {

	userDB, err := db.GetUserByUsername(user.Username.String)
	if err != nil {
		return fmt.Errorf("net.GetUserByUsername failed <- %v", err)
	}

	if userDB.Valid() {
		return ErrAlreadyExist
	}

	enc, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword failed <- %v", err)
	}

	err = user.Password.Scan(enc)
	if err != nil {
		return fmt.Errorf("user.Password.Scan failed <- %v", err)
	}

	_, err = db.InsertUser(user)

	if err != nil {
		return fmt.Errorf("net.InsertUser failed <- %v", err)
	}

	return nil
}

// Update should have a comment.
func Update(id int32, user *model.User) error {

	userDB, err := db.GetUser(id)
	if err != nil {
		return fmt.Errorf("net.GetUser failed <- %v", err)
	}

	if !userDB.Valid() {
		return ErrNotExist
	}

	if user.Password.Valid && user.Password.String != "" {
		enc, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("bcrypt.GenerateFromPassword failed <- %v", err)
		}

		err = user.Password.Scan(enc)
		if err != nil {
			return fmt.Errorf("user.Password.Scan failed <- %v", err)
		}
	}

	_, err = db.UpdateUser(id, user)

	if err != nil {
		return fmt.Errorf("db.UpdateUser failed <- %v", err)
	}

	return nil
}

// ParseToken should have a comment.
func ParseToken(token string) (*JWTClaims, error) {
	jwtClaims := JWTClaims{}
	tokenParse, err := jwt.ParseWithClaims(token, &jwtClaims, getPubKey)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrInvalidToken
		}
		return nil, fmt.Errorf("jwt.ParseWithClaims failed <- %v", err)
	}

	if !tokenParse.Valid {
		return nil, ErrInvalidToken
	}
	return &jwtClaims, nil
}

// RefreshToken should have a comment.
func RefreshToken(token string) (*UserToken, error) {
	jwtClaims, err := ParseToken(token)
	if err != nil {
		if err == ErrInvalidToken {
			return nil, err
		}
		return nil, fmt.Errorf("ParseToken failed <- %v", err)
	}

	expire := time.Unix(jwtClaims.ExpiresAt, 0)

	if expire.Sub(time.Now()) > 2*time.Minute {
		return &UserToken{Token: token, ExpireAt: expire.Unix()}, nil
	}

	return createToken(jwtClaims.ID, jwtClaims.Username, jwtClaims.FirstName, jwtClaims.LastName, jwtClaims.Admin)
}

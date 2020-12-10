package net

import (
	"fmt"

	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/config"
)

func InsertUser(user *model.User) (int32, error) {
	return insertObject(config.GetConfig().RestApi.AdminPath, user, "user")
}

func GetUsers() (model.Users, error) {
	users := model.Users{}

	err := getObjects(config.GetConfig().RestApi.AdminPath, "users", nil, &users)
	if err != nil {
		return nil, fmt.Errorf("getObjects failed <- %v", err)
	}
	return users, nil

}

func GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}

	err := getObject(config.GetConfig().RestApi.AdminPath, "user/username", username, &user)
	if err != nil {
		return nil, fmt.Errorf("getObject failed <- %v", err)
	}
	return &user, nil
}

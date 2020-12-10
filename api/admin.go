package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/internal/auth"
	"github.com/omnis-org/omnis-server/internal/net"
)

//////////////////	FUNCTIONS ADMIN	//////////////////

func getToken(r *http.Request) (string, error) {
	tokenBearer := r.Header.Get("Authorization")
	tokenBearerArray := strings.Split(tokenBearer, "Bearer ")
	if len(tokenBearerArray) != 2 {
		return "", auth.InvalidTokenError
	}
	return tokenBearerArray[1], nil
}

func (api *Api) validateToken(w http.ResponseWriter, r *http.Request) error {
	tokenValue, err := getToken(r)
	if err != nil {
		api.unauthorizedError(w, err)
		return err
	}

	_, err = auth.CheckToken(tokenValue)
	if err != nil {
		api.unauthorizedError(w, err)
		return err
	}
	return nil
}

func (api *Api) login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.internalError(w, err)
		return
	}

	token, err := auth.Login(&user)

	if err != nil {
		api.unauthorizedError(w, err)
		return
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		api.internalError(w, err)
	}

	api.sendJSON(w, jsonToken)
}

func (api *Api) refresh(w http.ResponseWriter, r *http.Request) {

	tokenValue, err := getToken(r)
	if err != nil {
		api.unauthorizedError(w, err)
		return
	}

	token, err := auth.RefreshToken(tokenValue)
	if err != nil {
		api.unauthorizedError(w, err)
		return
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		api.internalError(w, err)
	}

	api.sendJSON(w, jsonToken)

}

func (api *Api) register(w http.ResponseWriter, r *http.Request) {
	users, err := net.GetUsers()
	if len(users) != 0 {
		err := api.validateToken(w, r)
		if err != nil {
			return
		}
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.internalError(w, fmt.Errorf("json.NewDecoder(r.Body).Decode failed <- %v", err))
		return
	}

	err = auth.Register(&user)
	if err != nil {
		if err == auth.AlreadyExistError {
			api.badRequestError(w, err)
		} else {
			api.internalError(w, fmt.Errorf("auth.Register failed <- %v", err))
		}
		return
	}

	user.Password.String = ""
	jsonUser, err := json.Marshal(user)
	if err != nil {
		api.internalError(w, err)
	}

	api.sendJSON(w, jsonUser)
}

func (api *Api) admin(w http.ResponseWriter, r *http.Request) {
	err := api.validateToken(w, r)
	if err != nil {
		return
	}
	api.sendText(w, fmt.Sprintf("Welcome admin"))
}

func (api *Api) setupAdmin() {
	api.router.Methods("GET").Path("/admin/").HandlerFunc(api.admin)
	api.router.Methods("POST").Path("/admin/login").HandlerFunc(api.login)
	api.router.Methods("POST").Path("/admin/register").HandlerFunc(api.register)
	api.router.Methods("GET").Path("/admin/refresh").HandlerFunc(api.refresh)
}

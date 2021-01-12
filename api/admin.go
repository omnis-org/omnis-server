package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/auth"
	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
)

//////////////////	FUNCTIONS ADMIN	//////////////////

func getToken(r *http.Request) (string, error) {
	tokenBearer := r.Header.Get("Authorization")
	tokenBearerArray := strings.Split(tokenBearer, "Bearer ")
	if len(tokenBearerArray) != 2 {
		return "", auth.ErrInvalidToken
	}
	return tokenBearerArray[1], nil
}

func (api *API) listPendingMachine(w http.ResponseWriter, r *http.Request) {
	// Call la fonction de liste des machines et retourner la liste
	log.Debug("listPendingMachine")
	obj, err := db.GetPendingMachines()
	if err != nil {
		api.internalError(w, err)
		return
	}

	json, err := json.Marshal(obj)
	if err != nil {
		api.internalError(w, err)
		return
	}

	api.successReturnJSON(w, json)
}

func (api *API) doAuthorize(w http.ResponseWriter, r *http.Request, authorize bool) {
	idS := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idS)
	if err != nil {
		api.badRequestError(w, err)
		return
	}

	_, err = db.AuthorizeMachine(int32(id), authorize)
	if err != nil {
		api.internalError(w, err)
		return
	}

	api.successNoContent(w)
}

func (api *API) authorizeMachine(w http.ResponseWriter, r *http.Request) {
	// Call la fonction d'update de machine en passant authorize a true
	api.doAuthorize(w, r, true)
}

func (api *API) unauthorizeMachine(w http.ResponseWriter, r *http.Request) {
	api.doAuthorize(w, r, false)
}

func (api *API) login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.badRequestError(w, err)
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

	api.successReturnJSON(w, jsonToken)
}

func (api *API) refresh(w http.ResponseWriter, r *http.Request) {
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

	api.successReturnJSON(w, jsonToken)

}

func (api *API) register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.badRequestError(w, fmt.Errorf("json.NewDecoder(r.Body).Decode failed <- %v", err))
		return
	}

	err = auth.Register(&user)
	if err != nil {
		if err == auth.ErrAlreadyExist {
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

	api.successReturnJSON(w, jsonUser)
}

func (api *API) update(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idS)
	if err != nil {
		api.badRequestError(w, fmt.Errorf("strconv.Atoi failed <- %v", err))
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.badRequestError(w, fmt.Errorf("json.NewDecoder(r.Body).Decode failed <- %v", err))
		return
	}

	err = auth.Update(int32(id), &user)
	if err != nil {
		api.internalError(w, fmt.Errorf("auth.Update failed <- %v", err))
		return
	}

	user.Password.String = ""
	jsonUser, err := json.Marshal(user)
	if err != nil {
		api.internalError(w, err)
	}

	api.successReturnJSON(w, jsonUser)
}

// check if first connection
func (api *API) first(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	users, err := db.GetUsers()
	if err != nil {
		api.internalError(w, fmt.Errorf("db.GetUsers failed <- %v", err))
		return
	}
	if len(users) == 0 {
		api.successReturnJSON(w, []byte(`{"result": true}`))
	} else {
		api.successReturnJSON(w, []byte(`{"result": false}`))
	}
}

func (api *API) setupAdmin() {
	adminPath := config.GetConfig().Server.Admin
	// connection & token
	api.router.Methods("GET").Path(fmt.Sprintf("%s/first", adminPath)).HandlerFunc(api.first)
	api.router.Methods("POST").Path(fmt.Sprintf("%s/login", adminPath)).HandlerFunc(api.login)
	api.router.Methods("GET").Path(fmt.Sprintf("%s/refresh", adminPath)).HandlerFunc(api.refresh)
	// admin users
	api.router.Methods("POST").Path(fmt.Sprintf("%s/register", adminPath)).HandlerFunc(api.register)
	api.router.Methods("PATCH").Path(fmt.Sprintf("%s/update/{id:[0-9]+}", adminPath)).HandlerFunc(api.update)
	// pending machine
	api.router.Methods("GET").Path(fmt.Sprintf("%s/pending_machines/", adminPath)).HandlerFunc(api.listPendingMachine)
	api.router.Methods("PATCH").Path(fmt.Sprintf("%s/pending_machine/{id:[0-9]+}/authorize", adminPath)).HandlerFunc(api.authorizeMachine)
	api.router.Methods("PATCH").Path(fmt.Sprintf("%s/pending_machine/{id:[0-9]+}/unauthorize", adminPath)).HandlerFunc(api.unauthorizeMachine)
}

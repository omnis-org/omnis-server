package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/auth"
	"github.com/omnis-org/omnis-server/internal/db"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// API should have a comment.
type API struct {
	router *mux.Router
}

//////////////////			ERROR			//////////////////

func (api *API) badRequestError(w http.ResponseWriter, err error) {
	// 400
	log.Info(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Validation errors in your request"}`))
}

func (api *API) unauthorizedError(w http.ResponseWriter, err error) {
	// 401
	log.Info(err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Authentication credentials were missing or incorrect"}`))
}

func (api *API) forbiddenError(w http.ResponseWriter, err error) {
	// 403
	log.Info(err)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "The request is understood, but it has been refused or access is not allowed"}`))
}

func (api *API) notFoundError(w http.ResponseWriter) {
	// 404
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "The item does not exist"}`))
}

func (api *API) internalError(w http.ResponseWriter, err error) {
	// 500
	log.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Something is broken"}`))
}

//////////////////			SUCCESS			//////////////////

// func (api *API) success(w http.ResponseWriter) {
// 	w.WriteHeader(http.StatusOK)
// }

func (api *API) successReturnJSON(w http.ResponseWriter, json []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (api *API) successCreateItem(w http.ResponseWriter, location string) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", location)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "The item was created successfully"}`))
}

func (api *API) successNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func (api *API) sendText(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, text)
}

///// Router

func (api *API) root(w http.ResponseWriter, r *http.Request) {
	api.successReturnJSON(w, []byte(`{"message":"Welcome in OmnIS Server API"`))
}

func checkAccess(permissionsToCheck int, roleID int32, method string) error {
	role, err := db.GetRole(roleID)
	if err != nil {
		return err
	}

	// GET => 0
	// POST => 1
	// PATCH => 2
	// DELETE => 3
	var methodToCheck int = -1

	if method == "GET" {
		methodToCheck = 0
	} else if method == "POST" {
		methodToCheck = 1
	} else if method == "PATCH" {
		methodToCheck = 2
	} else if method == "DELETE" {
		methodToCheck = 3
	} else {
		return errors.New("Invalid method")
	}

	var permissions int32 = 0
	if permissionsToCheck == 1 {
		permissions = role.OmnisPermissions.Int32
	} else if permissionsToCheck == 2 {
		permissions = role.RolesPermissions.Int32
	} else if permissionsToCheck == 3 {
		permissions = role.UsersPermissions.Int32
	} else if permissionsToCheck == 4 {
		permissions = role.PendingMachinesPermissions.Int32
	}

	if permissions>>methodToCheck&1 == 1 {
		return nil
	} else {
		return fmt.Errorf("Unauthorize : %d >> %d & 1", permissions, methodToCheck)
	}

}

// permission to check
// OMNIS => 1
// ROLES => 2
// USERS => 3
// PENDING MACHINES => 4
func (api *API) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var permissionsToCheck int = 0

		if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.Client) {
			log.Info("No security check")
		} else if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.OmnisAPI) {
			permissionsToCheck = 1
		} else if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.AdminAPI) {
			if strings.HasPrefix(r.RequestURI, fmt.Sprintf("%s/user", config.GetConfig().Server.AdminAPI)) {
				permissionsToCheck = 3
			}
			if strings.HasPrefix(r.RequestURI, fmt.Sprintf("%s/role", config.GetConfig().Server.AdminAPI)) {
				permissionsToCheck = 2
			}
		} else if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.Admin) {
			if strings.HasPrefix(r.RequestURI, fmt.Sprintf("%s/pending_machine", config.GetConfig().Server.Admin)) {
				permissionsToCheck = 4
			} else if strings.HasPrefix(r.RequestURI, fmt.Sprintf("%s/update", config.GetConfig().Server.Admin)) {
				permissionsToCheck = 3
			} else if strings.HasPrefix(r.RequestURI, fmt.Sprintf("%s/register", config.GetConfig().Server.Admin)) {

				users, err := db.GetUsers()
				if err != nil {
					api.internalError(w, fmt.Errorf("db.GetUsers failed <- %v", err))
					return
				}

				if len(users) != 0 { // authorize register if no users
					permissionsToCheck = 3
				}
			}
		} else {
			api.unauthorizedError(w, errors.New("Invalid path"))
			return
		}

		if permissionsToCheck != 0 {
			tokenValue, err := getToken(r)
			if err != nil {
				api.unauthorizedError(w, err)
				return
			}

			jwtClaims, err := auth.ParseToken(tokenValue)
			if err != nil {
				api.unauthorizedError(w, err)
				return
			}

			err = checkAccess(permissionsToCheck, jwtClaims.RoleID, r.Method)
			if err != nil {
				api.forbiddenError(w, err)
				return
			}
		}

		log.Info("Authorize : ", r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func (api *API) setupAuthentication() {
	api.router.Use(api.middleware)
}

func (api *API) setupRouter() {
	api.router.Methods("GET").Path("/").HandlerFunc(api.root)
	api.setupClient()
	api.setupAdmin()
	api.setupRestAPI()
	api.setupAuthentication()
}

// Run should have a comment.
func Run() error {
	var err error
	// Init router
	router := mux.NewRouter().StrictSlash(true)

	// Init Api
	api := API{router}
	api.setupRouter()

	log.Info("Success SetupRouter")

	// Init Serve
	url := fmt.Sprintf("%s:%d", config.GetConfig().Server.IP, config.GetConfig().Server.Port)

	handler := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	fmt.Println("Listen and serve at url : ", url)

	if config.GetConfig().TLS.Activated {
		err = http.ListenAndServeTLS(url, config.GetConfig().TLS.ServerCrtFile, config.GetConfig().TLS.ServerKeyFile, handler(api.router))
	} else {
		err = http.ListenAndServe(url, handler(api.router))
	}

	if err != nil {
		log.Error("ListenAndServe failed : ", err)
	}

	return nil

}

package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/auth"

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
	log.Info(err)
	http.Error(w, http.StatusText(400), 400)
}

func (api *API) notFoundError(w http.ResponseWriter, err error) {
	log.Info(err)
	http.Error(w, http.StatusText(404), 404)
}

func (api *API) internalError(w http.ResponseWriter, err error) {
	log.Error(err)
	http.Error(w, http.StatusText(500), 500)
}

func (api *API) unauthorizedError(w http.ResponseWriter, err error) {
	log.Info(err)
	http.Error(w, http.StatusText(401), 401)
}

//////////////////			SUCCESS			//////////////////

func (api *API) success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success\n")
}

func (api *API) sendJSON(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (api *API) sendText(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, text)
}

func (api *API) sendNullJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	n := []uint8{110, 117, 108, 108}
	w.Write(n)
}

///// Router

func (api *API) root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome in OmnIS Server API\n")
}

func (api *API) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.AdminAPI) || strings.HasPrefix(r.RequestURI, config.GetConfig().Server.OmnisAPI) {
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

			if strings.HasPrefix(r.RequestURI, config.GetConfig().Server.AdminAPI) && !jwtClaims.Admin {
				api.unauthorizedError(w, err)
				return
			}
		}

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
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
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

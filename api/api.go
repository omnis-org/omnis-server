package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/omnis-org/omnis-server/internal/worker"

	"github.com/omnis-org/omnis-client/pkg/client_informations"

	"github.com/omnis-org/omnis-server/config"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Api struct {
	router *mux.Router
}

//////////////////			ERROR			//////////////////

func (api *Api) badRequestError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(400), 400)
}

func (api *Api) notFoundError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(404), 404)
}

func (api *Api) internalError(w http.ResponseWriter, err error) {
	log.Error(err)
	http.Error(w, http.StatusText(500), 500)
}

//////////////////			SUCCESS			//////////////////

func (api *Api) success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success\n")
}

func (api *Api) sendJSON(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//////////////////			FUNCTIONS			//////////////////

func (api *Api) informations(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.internalError(w, err)
		return
	}

	infos := client_informations.Informations{}

	err = json.Unmarshal(body, &infos)
	if err != nil {
		api.internalError(w, err)
		return
	}

	worker.AddWork(&worker.Work{worker.AnalyzeClientInformations, &infos})

	api.success(w)
}

///// Router

func (api *Api) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome in OmnIS Server API\n")
}

func (api *Api) setupRouter() {
	api.router.Methods("GET").Path("/api").HandlerFunc(api.home)
	api.router.Methods("POST").Path("/api/informations").HandlerFunc(api.informations)
}

func Run() error {
	// Init router
	router := mux.NewRouter().StrictSlash(true)

	// Init Api
	api := Api{router}
	api.setupRouter()

	log.Info("Success SetupRouter")

	// Init Serve
	url := fmt.Sprintf("%s:%d", config.GetConfig().Server.Ip, config.GetConfig().Server.Port)

	log.Info("ListenAndServer : ", url)
	err := http.ListenAndServe(url, api.router)

	if err != nil {
		log.Error("ListenAndServe failed : ", err)
	}

	return nil

}

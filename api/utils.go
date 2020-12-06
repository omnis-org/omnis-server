package api

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//////////////////			ERROR			//////////////////

func (api *Api) badRequestError(w http.ResponseWriter, err error) {
	log.Info(err)
	http.Error(w, http.StatusText(400), 400)
}

func (api *Api) notFoundError(w http.ResponseWriter, err error) {
	log.Info(err)
	http.Error(w, http.StatusText(404), 404)
}

func (api *Api) internalError(w http.ResponseWriter, err error) {
	log.Error(err)
	http.Error(w, http.StatusText(500), 500)
}

func (api *Api) unauthorizedError(w http.ResponseWriter, err error) {
	log.Info(err)
	http.Error(w, http.StatusText(401), 401)
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

func (api *Api) sendText(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, text)
}

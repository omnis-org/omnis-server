package api

import (
	"encoding/json"
	"net/http"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/worker"
)

//////////////////	FUNCTIONS CLIENT	//////////////////

func (api *Api) informations(w http.ResponseWriter, r *http.Request) {
	var infos client_informations.Informations
	err := json.NewDecoder(r.Body).Decode(&infos)
	if err != nil {
		api.internalError(w, err)
		return
	}

	worker.AddWork(&worker.Work{Job: worker.AnalyzeClientInformations, Handle: &infos})

	api.success(w)
}

func (api *Api) setupClient() {
	api.router.Methods("POST").Path("/client/informations").HandlerFunc(api.informations)
}

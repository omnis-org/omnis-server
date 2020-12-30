package api

import (
	"encoding/json"
	"net/http"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/internal/client"
)

//////////////////	FUNCTIONS CLIENT	//////////////////

func (api *API) informations(w http.ResponseWriter, r *http.Request) {
	var infos client_informations.Informations
	err := json.NewDecoder(r.Body).Decode(&infos)
	if err != nil {
		api.internalError(w, err)
		return
	}

	client.AnalyzeClientInformations(&infos)

	api.success(w)
}

func (api *API) setupClient() {
	api.router.Methods("POST").Path("/client/informations").HandlerFunc(api.informations)
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/client"
)

//////////////////	FUNCTIONS CLIENT	//////////////////

func (api *API) informations(w http.ResponseWriter, r *http.Request) {
	var infos client_informations.Informations
	err := json.NewDecoder(r.Body).Decode(&infos)
	if err != nil {
		api.badRequestError(w, err)
		return
	}

	/*
		machine
			[x] script SQL creation base
			[x] modification du model
			[x] modification des methodes
			[x] methode listing machines en attente de validation
				[x] creer nouvelle procedure SQL dans sql/create_procedure.sql
			[x] methode validation d'une machine
				utiliser la procedure deja existante d'update de machine
			[x] API listing machines
			[x] API validation machine
				utiliser endpoint d'update deja existant
			[ ] vue/template listing machines
			[ ] vue/template validation machine
	*/

	// MODIFIER DANS EXEMPLE INSERT DB - ADD UUID ET BOOL

	// Fonction de liste de machine en attente
	// Fonction savoir si UUID connu et valide
	// Fonction de validation d'une machine via son UUID

	// UI - faire un endpoint de listing : list des machines en attente de validation + endpoint valider machine par UUID
	// Faire une vue pour le listing et de validation

	// Checker si UUID est bien valide avant analyse - ou sont stockes les machines avec leur ID ?
	// Si oui analyser
	// Si non, mettre en attente de confirmation, api not authorised avec message uuid non confirmer
	// Dans panel admin, mettre en queue les confirmations des UUID (vue, api qui accepte, api qui refuse, api qui liste (queue))
	// Quand confirmer, lancer analyseClient qui mettra en BDD la machine avec l'id etc
	err = client.AnalyzeClientInformations(&infos)

	if err != nil {
		api.internalError(w, err)
	}

	api.successNoContent(w)
}

func (api *API) setupClient() {
	conf := config.GetConfig()
	clientPath := conf.Server.APIPath + conf.Server.ClientPath
	api.router.Methods("POST").Path(fmt.Sprintf("%s/informations", clientPath)).HandlerFunc(api.informations)
}

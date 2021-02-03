package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

//////////////////			SUCCESS			//////////////////

//////////////////		OMNIS	FUNCTIONS			//////////////////

func (api *API) getObjects(f func(bool) (model.Objects, error), automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getObjects")
		obj, err := f(automatic)
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
}

func (api *API) getObjectsByInt(f func(int32, bool) (model.Objects, error), s string, automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getObjectsByInt")
		idS := mux.Vars(r)[s]
		id, err := strconv.Atoi(idS)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		obj, err := f(int32(id), automatic)

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
}

func (api *API) getObjectsByString(f func(string, bool) (model.Objects, error), s string, automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getObjectsByString")
		p := bluemonday.UGCPolicy()
		if p == nil {
			api.internalError(w, errors.New("bluemonday.UGCPolicy failed"))
			return
		}

		sv := p.Sanitize(mux.Vars(r)[s])

		obj, err := f(sv, automatic)

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
}

func (api *API) getObject(f func(int32, bool) (model.Object, error), automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getObject")
		idS := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idS)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		obj, err := f(int32(id), automatic)

		if err != nil {
			api.internalError(w, err)
			return
		}

		if obj == nil {
			api.notFoundError(w)
			return
		}

		json, err := json.Marshal(obj)
		if err != nil {
			api.internalError(w, err)
			return
		}

		api.successReturnJSON(w, json)
	}
}

func (api *API) getObjectByString(f func(string, bool) (model.Object, error), s string, automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getObjectByString")
		p := bluemonday.UGCPolicy()
		if p == nil {
			api.internalError(w, errors.New("bluemonday.UGCPolicy failed"))
			return
		}

		sv := p.Sanitize(mux.Vars(r)[s])

		obj, err := f(sv, automatic)

		if err != nil {
			api.internalError(w, err)
			return
		}

		if obj == nil {
			api.notFoundError(w)
			return
		}

		json, err := json.Marshal(obj)
		if err != nil {
			api.internalError(w, err)
			return
		}

		api.successReturnJSON(w, json)
	}
}

func (api *API) getOutdatedObjects(f func(int) (model.Objects, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("getOutdatedObjects")
		outdatedDayS := mux.Vars(r)["outdated_day"]
		outdatedDay, err := strconv.Atoi(outdatedDayS)
		if err != nil {
			api.internalError(w, err)
			return
		}

		obj, err := f(outdatedDay)
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
}

func (api *API) insertObject(f func(*model.Object, bool) (int32, error), o *model.Object, apiPath string, objName string, automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("insertObject")

		obj := (*o).New()

		err := json.NewDecoder(r.Body).Decode(&obj)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		if !(obj.Valid()) {
			api.badRequestError(w, errors.New("Object invalid"))
			return
		}

		id, err := f(&obj, automatic)
		if err != nil {
			api.internalError(w, err)
			return
		}

		if id == 0 {
			api.badRequestError(w, errors.New("error insert id == 0"))
			return
		}

		api.successCreateItem(w, id)
	}
}

func (api *API) updateObject(f func(int32, *model.Object, bool) (int64, error), o *model.Object, automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("updateObject")

		obj := (*o).New()

		err := json.NewDecoder(r.Body).Decode(&obj)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		idS := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idS)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		_, err = f(int32(id), &obj, automatic)
		if err != nil {
			api.internalError(w, err)
			return
		}

		api.successNoContent(w)
	}
}

func (api *API) deleteObject(f func(int32) (int64, error), automatic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("deleteObject")
		idS := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idS)
		if err != nil {
			api.badRequestError(w, err)
			return
		}

		rowsAffected, err := f(int32(id))
		if err != nil {
			api.internalError(w, err)
			return
		}

		if rowsAffected == 0 {
			api.badRequestError(w, errors.New("No rows affected"))
			return
		}

		api.successNoContent(w)
	}
}

///// Router

func (api *API) setupBasicFunctions(apiPath string, getObjs func(bool) (model.Objects, error),
	getObj func(int32, bool) (model.Object, error),
	insertObj func(*model.Object, bool) (int32, error),
	updateObj func(int32, *model.Object, bool) (int64, error),
	deleteObj func(int32) (int64, error),
	outdatedObj func(int) (model.Objects, error),
	objName string,
	obj *model.Object) {

	apiPathAuto := fmt.Sprintf("%s/auto", apiPath)

	// getObjects
	if getObjs != nil {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss", apiPath, objName)).HandlerFunc(api.getObjects(getObjs, false))
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss", apiPathAuto, objName)).HandlerFunc(api.getObjects(getObjs, true))
	}

	// getObject
	if getObj != nil {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPath, objName)).HandlerFunc(api.getObject(getObj, false))
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPathAuto, objName)).HandlerFunc(api.getObject(getObj, true))
	}
	// insertObject
	if insertObj != nil && obj != nil {
		api.router.Methods("POST").Path(fmt.Sprintf("%s/%s", apiPath, objName)).HandlerFunc(api.insertObject(insertObj, obj, apiPath, objName, false))
		api.router.Methods("POST").Path(fmt.Sprintf("%s/%s", apiPathAuto, objName)).HandlerFunc(api.insertObject(insertObj, obj, apiPath, objName, true))
	}
	// updateObject
	if updateObj != nil && obj != nil {
		api.router.Methods("PATCH").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPath, objName)).HandlerFunc(api.updateObject(updateObj, obj, false))
		api.router.Methods("PATCH").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPathAuto, objName)).HandlerFunc(api.updateObject(updateObj, obj, true))
	}
	// delete
	if deleteObj != nil {
		api.router.Methods("DELETE").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPath, objName)).HandlerFunc(api.deleteObject(deleteObj, false))
		api.router.Methods("DELETE").Path(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiPathAuto, objName)).HandlerFunc(api.deleteObject(deleteObj, true))
	}

	if objName != "" {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss/outdated/{outdated_day:[0-9]+}", apiPath, objName)).HandlerFunc(api.getOutdatedObjects(outdatedObj))
	}

}

func (api *API) setupGetObjectsByString(apiPath string, f func(string, bool) (model.Objects, error), objName string, s string) {
	apiPathAuto := fmt.Sprintf("%s/auto", apiPath)

	if f != nil {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss/%s/{%s}", apiPath, objName, s, s)).HandlerFunc(api.getObjectsByString(f, s, false))
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss/%s/{%s}", apiPathAuto, objName, s, s)).HandlerFunc(api.getObjectsByString(f, s, true))
	}
}

func (api *API) setupGetObjectsByInt(apiPath string, f func(int32, bool) (model.Objects, error), objName string, s string) {
	apiPathAuto := fmt.Sprintf("%s/auto", apiPath)

	if f != nil {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss/%s/{%s:[0-9]+}", apiPath, objName, s, s)).HandlerFunc(api.getObjectsByInt(f, s, false))
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%ss/%s/{%s:[0-9]+}", apiPathAuto, objName, s, s)).HandlerFunc(api.getObjectsByInt(f, s, true))
	}
}

func (api *API) setupGetObjectByString(apiPath string, f func(string, bool) (model.Object, error), objName string, s string) {
	apiPathAuto := fmt.Sprintf("%s/auto", apiPath)

	if f != nil {
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%s/%s/{%s}", apiPath, objName, s, s)).HandlerFunc(api.getObjectByString(f, s, false))
		api.router.Methods("GET").Path(fmt.Sprintf("%s/%s/%s/{%s}", apiPathAuto, objName, s, s)).HandlerFunc(api.getObjectByString(f, s, true))
	}
}

///// API OMNIS

func (api *API) setupLocation(apiPath string) {
	var location model.Object = new(model.Location)
	api.setupBasicFunctions(apiPath, db.GetLocationsO, db.GetLocationO, db.InsertLocationO, db.UpdateLocationO, db.DeleteLocation, db.GetOutdatedLocationsO, "location", &location)
	api.setupGetObjectByString(apiPath, db.GetLocationByNameO, "location", "name")
}

func (api *API) setupPerimeter(apiPath string) {
	var perimeter model.Object = new(model.Perimeter)
	api.setupBasicFunctions(apiPath, db.GetPerimetersO, db.GetPerimeterO, db.InsertPerimeterO, db.UpdatePerimeterO, db.DeletePerimeter, db.GetOutdatedPerimetersO, "perimeter", &perimeter)
	api.setupGetObjectByString(apiPath, db.GetPerimeterByNameO, "perimeter", "name")
}

func (api *API) setupOperatingSystem(apiPath string) {
	var operatingSystem model.Object = new(model.OperatingSystem)
	api.setupBasicFunctions(apiPath, db.GetOperatingSystemsO, db.GetOperatingSystemO, db.InsertOperatingSystemO, db.UpdateOperatingSystemO, db.DeleteOperatingSystem, db.GetOutdatedOperatingSystemsO, "operatingSystem", &operatingSystem)
	api.setupGetObjectsByString(apiPath, db.GetOperatingSystemsByNameO, "operatingSystem", "name")
}

func (api *API) setupTag(apiPath string) {
	var tag model.Object = new(model.Tag)
	api.setupBasicFunctions(apiPath, db.GetTagsO, db.GetTagO, db.InsertTagO, db.UpdateTagO, db.DeleteTag, db.GetOutdatedTagsO, "tag", &tag)
}

func (api *API) setupSoftware(apiPath string) {
	var software model.Object = new(model.Software)
	api.setupBasicFunctions(apiPath, db.GetSoftwaresO, db.GetSoftwareO, db.InsertSoftwareO, db.UpdateSoftwareO, db.DeleteSoftware, db.GetOutdatedSoftwaresO, "software", &software)
}

func (api *API) setupMachine(apiPath string) {
	var machine model.Object = new(model.Machine)
	api.setupBasicFunctions(apiPath, db.GetMachinesO, db.GetMachineO, db.InsertMachineO, db.UpdateMachineO, db.DeleteMachine, db.GetOutdatedMachinesO, "machine", &machine)
}

func (api *API) setupInstalledSoftware(apiPath string) {
	var installedSoftware model.Object = new(model.InstalledSoftware)
	api.setupBasicFunctions(apiPath, db.GetInstalledSoftwaresO, db.GetInstalledSoftwareO, db.InsertInstalledSoftwareO, db.UpdateInstalledSoftwareO, db.DeleteInstalledSoftware, db.GetOutdatedInstalledSoftwaresO, "installedSoftware", &installedSoftware)
}

func (api *API) setupTaggedMachine(apiPath string) {
	var taggedMachine model.Object = new(model.TaggedMachine)
	api.setupBasicFunctions(apiPath, db.GetTaggedMachinesO, db.GetTaggedMachineO, db.InsertTaggedMachineO, db.UpdateTaggedMachineO, db.DeleteTaggedMachine, db.GetOutdatedTaggedMachinesO, "taggedMachine", &taggedMachine)
}

func (api *API) setupNetwork(apiPath string) {
	var network model.Object = new(model.Network)
	api.setupBasicFunctions(apiPath, db.GetNetworksO, db.GetNetworkO, db.InsertNetworkO, db.UpdateNetworkO, db.DeleteNetwork, db.GetOutdatedNetworksO, "network", &network)
	api.setupGetObjectsByString(apiPath, db.GetNetworksByIPO, "network", "ip")
}

func (api *API) setupInterface(apiPath string) {
	var interfaceO model.Object = new(model.InterfaceO)
	api.setupBasicFunctions(apiPath, db.GetInterfacesO, db.GetInterfaceO, db.InsertInterfaceO, db.UpdateInterfaceO, db.DeleteInterface, db.GetOutdatedInterfacesO, "interface", &interfaceO)
	api.setupGetObjectByString(apiPath, db.GetInterfaceByMacO, "interface", "mac")
	api.setupGetObjectsByInt(apiPath, db.GetInterfacesByMachineIDO, "interface", "machineId")
}

func (api *API) setupGateway(apiPath string) {
	var gateway model.Object = new(model.Gateway)
	api.setupBasicFunctions(apiPath, db.GetGatewaysO, db.GetGatewayO, db.InsertGatewayO, db.UpdateGatewayO, db.DeleteGateway, db.GetOutdatedGatewaysO, "gateway", &gateway)
	api.setupGetObjectsByInt(apiPath, db.GetGatewaysByInterfaceIDO, "gateway", "interfaceId")
}

func (api *API) setupOmnisAPI() {
	apiPath := config.GetConfig().Server.OmnisAPI
	api.router.Methods("GET").Path(apiPath).HandlerFunc(api.root)
	api.setupLocation(apiPath)
	api.setupPerimeter(apiPath)
	api.setupOperatingSystem(apiPath)
	api.setupTag(apiPath)
	api.setupSoftware(apiPath)
	api.setupMachine(apiPath)
	api.setupInstalledSoftware(apiPath)
	api.setupTaggedMachine(apiPath)
	api.setupNetwork(apiPath)
	api.setupInterface(apiPath)
	api.setupGateway(apiPath)
}

///// API ADMIN

func (api *API) setupUser(apiPath string) {
	var user model.Object = new(model.User)
	api.setupBasicFunctions(apiPath, db.GetUsersO, db.GetUserO, db.InsertUserO, db.UpdateUserO, db.DeleteUser, nil, "user", &user)
	api.setupGetObjectByString(apiPath, db.GetUserByUsernameO, "user", "username")
}

func (api *API) setupRole(apiPath string) {
	var role model.Object = new(model.Role)
	api.setupBasicFunctions(apiPath, db.GetRolesO, db.GetRoleO, db.InsertRoleO, db.UpdateRoleO, db.DeleteRole, nil, "role", &role)
}

func (api *API) setupAdminAPI() {
	apiPath := config.GetConfig().Server.AdminAPI
	api.setupUser(apiPath)
	api.setupRole(apiPath)
}

func (api *API) setupRestAPI() {
	api.setupAdminAPI()
	api.setupOmnisAPI()
}

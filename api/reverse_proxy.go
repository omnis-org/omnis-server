package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/auth"
)

func redirectRestApi(w http.ResponseWriter, r *http.Request) {
	url, _ := config.GetRestApiUrl()

	r.Host = url.Host
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func (api *Api) reverseProxy(w http.ResponseWriter, r *http.Request) {
	tokenValue, err := getToken(r)
	if err != nil {
		api.unauthorizedError(w, err)
		return
	}

	_, err = auth.CheckToken(tokenValue)
	if err != nil {
		api.unauthorizedError(w, err)
		return
	}

	redirectRestApi(w, r)
}

func (api *Api) setupProxy() {
	api.router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		match, _ := regexp.MatchString(fmt.Sprintf("%s.*", config.GetConfig().RestApi.OmnisPath), r.URL.Path)
		return match
	}).HandlerFunc(api.reverseProxy)
}

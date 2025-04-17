package routes

import (
	home "app/app/routes/home"
	sharedTypes "app/pkg/shared_types"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MetadataKey sharedTypes.MetadataKey
type RoutesMap map[MetadataKey]http.HandlerFunc

var routeRegistry = map[MetadataKey]http.HandlerFunc{}
var mainRouter *mux.Router = mux.NewRouter()

func LoadRoutes() {
	RegisterRoute(MetadataKey(home.UrlKey), MetadataKey(home.DisplayNameKey), home.Handler)
}

func RegisterRoute(url MetadataKey, displayName MetadataKey, handler http.HandlerFunc) {
	routeRegistry[url] = handler
	mainRouter.HandleFunc(string(url), handler)
	fmt.Println("<< Registered route:", displayName, "with url:", url, ">>")
}

func GetRoutes() RoutesMap {
	return routeRegistry
}

func GetRouter() *mux.Router {
	return mainRouter
}

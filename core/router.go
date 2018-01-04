package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route is type to define a uri to a handler function
type Route struct {
	Path string
	Fn   func(w http.ResponseWriter, r *http.Request)
	// Name    string
	// Method  string
	// Pattern string
	// Handler http.Handler
}

// Routes is a set of routes config
var Routes = []Route{}

//NewRouter is function to create a router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Fn)
		// r.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)
	}
	return r
}

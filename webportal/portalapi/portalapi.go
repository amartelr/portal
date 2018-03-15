package portalapi

import (
	//"github.com/amartelr/portal/webportal/portalapi"
	"net/http"

	"github.com/amartelr/portal/dblayer"

	"github.com/gorilla/mux"
)

// portal API :
// GET: Search /api/portal/nickname/rex (GetAnimalByNickname), or /api/portal/type/velociraptor (GetAnymalByType)
// POST: add or edit /api/portal/add (AddAnimal) /api/portal/edit (UpdateAnimal)
// the HTTP POST comes with a json body witch hosts the data to be added o modified
func RunApi(endpoint string, db dblayer.PortalDBHandler) error {
	//endpoint examples : localhost:8000
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db dblayer.PortalDBHandler) {
	handler := newPortalRESTAPIHandler(db)

	apirouter := r.PathPrefix("/api/portal").Subrouter()

	apirouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)
	
}

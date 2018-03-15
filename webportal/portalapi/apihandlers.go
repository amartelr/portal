package portalapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amartelr/portal/dblayer"

	"encoding/json"
	"strings"

	"github.com/gorilla/mux"
)

type PortalRESTAPIHandler struct {
	dbhandler dblayer.PortalDBHandler
}

func newPortalRESTAPIHandler(db dblayer.PortalDBHandler) *PortalRESTAPIHandler {
	return &PortalRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *PortalRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `No search criteria found, Options search:
			by nickname or by type`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `No search criteria found, Options search:
			by nickname or by type`)
		return
	}

	var animal dblayer.Animal
	var animals []dblayer.Animal
	var err error

	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetAnimalByNickname(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetAnymalByType(searchkey)
		if len(animals) > 0 {
			json.NewEncoder(w).Encode(animals)
			return
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error while querying animals %w", err)
		return
	}

	json.NewEncoder(w).Encode(animal)
}

func (handler *PortalRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["Operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `Operation was not privided:
			use add or edit`)
		return
	}

	var animal dblayer.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cound decode the request body to json %v", err)
		return
	}

	switch strings.ToLower(operation) {
	case "add":
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		// /api/portal/edit/rex

		nickname := r.RequestURI[len("/api/portal/edit/"):]
		log.Println("edit requested for nickname", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error ocurred while processing request %v", err)
	}
}

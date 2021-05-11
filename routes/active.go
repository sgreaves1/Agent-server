package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sgreaves1/Agent-server/helpers"
)

func GetActiveHandler(w http.ResponseWriter, r *http.Request) {
	activeElements := helpers.GetAssets()

	if activeElements != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(activeElements)
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

type Body struct {
	Location string
}

func AssetByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/asset/")

	switch r.Method {
	case "GET":
		activeAsset := helpers.GetAssetById(id)
		if activeAsset != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			j, _ := json.Marshal(activeAsset)
			w.Write(j)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	case "PATCH":
		var l Body
		err := json.NewDecoder(r.Body).Decode(&l)
		if err == nil {
			activeAsset := helpers.GetAssetById(id)
			if activeAsset != nil {
				helpers.UpdateAssetLocation(id, l.Location)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}


	default:
		fmt.Fprintf(w, "Sorry, only GET and PATCH methods are supported.")
	}
	//w.WriteHeader(http.StatusNoContent)
	//fmt.Fprint(w, "")
}

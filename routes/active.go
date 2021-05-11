package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sgreaves1/Agent-server/helpers"
)

func GetActiveHandler(w http.ResponseWriter, r *http.Request) {
	activeElements := helpers.GetActiveInfo()

	if activeElements != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(activeElements)
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
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
		fmt.Fprint(w, "Will update!")
	default:
		fmt.Fprintf(w, "Sorry, only GET and PATCH methods are supported.")
	}
	//w.WriteHeader(http.StatusNoContent)
	//fmt.Fprint(w, "")
}

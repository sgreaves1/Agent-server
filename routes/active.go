package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sgreaves1/Agent-server/helpers"
)

func GetActiveHandler(w http.ResponseWriter, r *http.Request) {
	activeElements := helpers.GetActiveInfo()

	if activeElements != nil {
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(activeElements)
		fmt.Fprint(w, string(j))
	}
}
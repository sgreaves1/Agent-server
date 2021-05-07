package routes

import (
	"fmt"
	"net/http"
)

func LivenessRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Agent server online")
}
package v1helloController

import (
	"encoding/json"
	"net/http"
)

func (h *V1hello) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hello")
}

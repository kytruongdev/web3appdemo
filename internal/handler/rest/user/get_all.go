package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userCtrl.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

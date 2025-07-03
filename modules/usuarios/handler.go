package usuarios

import (
	"encoding/json"
	"net/http"
)

type UsuarioHandler struct {
	Repo *UsuarioRepository
}

func NewUsuarioHandler(repo *UsuarioRepository) *UsuarioHandler {
	return &UsuarioHandler{Repo: repo}
}

func (h *UsuarioHandler) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.Repo.ListarUsuarios()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

package clientes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClienteHandler struct {
	Repo *ClienteRepository
}

func NewClienteHandler(repo *ClienteRepository) *ClienteHandler {
	return &ClienteHandler{Repo: repo}
}

func (h *ClienteHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var c Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Criar(&c); err != nil {
		http.Error(w, "Erro ao criar cliente", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *ClienteHandler) Editar(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var c Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Editar(id, &c); err != nil {
		http.Error(w, "Erro ao editar cliente", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.Repo.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar clientes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

func (h *ClienteHandler) Pesquisar(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Parâmetro de busca obrigatório", http.StatusBadRequest)
		return
	}
	clientes, err := h.Repo.BuscarPorNome(q)
	if err != nil {
		http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

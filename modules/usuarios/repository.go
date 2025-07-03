package usuarios

import (
	"database/sql"
)

type Usuario struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Login    string `json:"login"`
	PerfilID int64  `json:"perfil_id"`
	Ativo    bool   `json:"ativo"`
}

type UsuarioRepository struct {
	DB *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (r *UsuarioRepository) ListarUsuarios() ([]Usuario, error) {
	rows, err := r.DB.Query("SELECT id, nome, login, perfil_id, ativo FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Login, &u.PerfilID, &u.Ativo); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

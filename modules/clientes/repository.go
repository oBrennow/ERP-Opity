package clientes

import (
	"database/sql"
)

type ClienteRepository struct {
	DB *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{DB: db}
}

func (r *ClienteRepository) Criar(c *Cliente) error {
	query := `INSERT INTO clientes (nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id, criado_em, atualizado_em`
	return r.DB.QueryRow(query, c.Nome, c.TipoPessoa, c.CpfCnpj, c.Telefone, c.Email, c.EnderecoLogradouro, c.EnderecoNumero, c.EnderecoComplemento, c.EnderecoBairro, c.EnderecoCidade, c.EnderecoUf, c.EnderecoCep).Scan(&c.ID, &c.CriadoEm, &c.AtualizadoEm)
}

func (r *ClienteRepository) Editar(id int64, c *Cliente) error {
	query := `UPDATE clientes SET nome=$1, tipo_pessoa=$2, cpf_cnpj=$3, telefone=$4, email=$5, endereco_logradouro=$6, endereco_numero=$7, endereco_complemento=$8, endereco_bairro=$9, endereco_cidade=$10, endereco_uf=$11, endereco_cep=$12, atualizado_em=NOW() WHERE id=$13`
	_, err := r.DB.Exec(query, c.Nome, c.TipoPessoa, c.CpfCnpj, c.Telefone, c.Email, c.EnderecoLogradouro, c.EnderecoNumero, c.EnderecoComplemento, c.EnderecoBairro, c.EnderecoCidade, c.EnderecoUf, c.EnderecoCep, id)
	return err
}

func (r *ClienteRepository) Listar() ([]Cliente, error) {
	rows, err := r.DB.Query("SELECT id, nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep, criado_em, atualizado_em FROM clientes ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.ID, &c.Nome, &c.TipoPessoa, &c.CpfCnpj, &c.Telefone, &c.Email, &c.EnderecoLogradouro, &c.EnderecoNumero, &c.EnderecoComplemento, &c.EnderecoBairro, &c.EnderecoCidade, &c.EnderecoUf, &c.EnderecoCep, &c.CriadoEm, &c.AtualizadoEm); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}

func (r *ClienteRepository) BuscarPorNome(nome string) ([]Cliente, error) {
	rows, err := r.DB.Query("SELECT id, nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep, criado_em, atualizado_em FROM clientes WHERE nome ILIKE $1 ORDER BY nome", "%"+nome+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.ID, &c.Nome, &c.TipoPessoa, &c.CpfCnpj, &c.Telefone, &c.Email, &c.EnderecoLogradouro, &c.EnderecoNumero, &c.EnderecoComplemento, &c.EnderecoBairro, &c.EnderecoCidade, &c.EnderecoUf, &c.EnderecoCep, &c.CriadoEm, &c.AtualizadoEm); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}

func (r *ClienteRepository) BuscarPorID(id int64) (*Cliente, error) {
	var c Cliente
	query := "SELECT id, nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep, criado_em, atualizado_em FROM clientes WHERE id=$1"
	err := r.DB.QueryRow(query, id).Scan(&c.ID, &c.Nome, &c.TipoPessoa, &c.CpfCnpj, &c.Telefone, &c.Email, &c.EnderecoLogradouro, &c.EnderecoNumero, &c.EnderecoComplemento, &c.EnderecoBairro, &c.EnderecoCidade, &c.EnderecoUf, &c.EnderecoCep, &c.CriadoEm, &c.AtualizadoEm)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

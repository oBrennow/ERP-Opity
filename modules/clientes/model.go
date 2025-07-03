package clientes

type Cliente struct {
	ID                  int64  `json:"id"`
	Nome                string `json:"nome"`
	TipoPessoa          string `json:"tipo_pessoa"`
	CpfCnpj             string `json:"cpf_cnpj"`
	Telefone            string `json:"telefone"`
	Email               string `json:"email"`
	EnderecoLogradouro  string `json:"endereco_logradouro"`
	EnderecoNumero      string `json:"endereco_numero"`
	EnderecoComplemento string `json:"endereco_complemento"`
	EnderecoBairro      string `json:"endereco_bairro"`
	EnderecoCidade      string `json:"endereco_cidade"`
	EnderecoUf          string `json:"endereco_uf"`
	EnderecoCep         string `json:"endereco_cep"`
	CriadoEm            string `json:"criado_em"`
	AtualizadoEm        string `json:"atualizado_em"`
}

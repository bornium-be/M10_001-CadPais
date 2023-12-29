package infotabela_empresa

import (
	"net/http"

	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
)

var Rotas_TabEmpresa = []declaracao.Rota{
	{
		URI:                "/bornium/tabempresa/nomefantasia", //--> Parametro: CNPJ
		Metodo:             http.MethodGet,
		Funcao:             TabEmpresa_RecNomeFantasia,
		RequerAutenticacao: false,
	},
}

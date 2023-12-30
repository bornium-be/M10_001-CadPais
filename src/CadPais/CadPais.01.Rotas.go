package cadpais

import (
	"net/http"

	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
)

var Rotas_CadPais = []declaracao.Rota{
	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais", //--> Parametro: PaisNome
		Metodo:             http.MethodGet,
		Funcao:             buscarTodos,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",
		Metodo:             http.MethodGet,
		Funcao:             buscarPorId,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais.codNovoRegistro",
		Metodo:             http.MethodGet,
		Funcao:             codNovoRegistro,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",
		Metodo:             http.MethodPost,
		Funcao:             inserirRegistro,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",
		Metodo:             http.MethodPut,
		Funcao:             alterarRegistro,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",
		Metodo:             http.MethodDelete,
		Funcao:             deletar,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/cadPais.ibge/{IBGE_Id}",
		Metodo:             http.MethodGet,
		Funcao:             ibge_Buscar,
		RequerAutenticacao: false,
	},

	{
		URI:                "/bornium/v1/{EmpresaId}/abc/{IBGE_Id}",
		Metodo:             http.MethodGet,
		Funcao:             abc,
		RequerAutenticacao: false,
	},
}

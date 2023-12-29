package infotab_configuracao

import (
	"net/http"

	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
)

var Rotas_TabConfiguracao = []declaracao.Rota{
	{
		URI:                "/bornium/{codEmpresa}/tabconfiguracao/codregimetributario",
		Metodo:             http.MethodGet,
		Funcao:             TabConfiguracao_CodRegimeTributario,
		RequerAutenticacao: false,
	},
}

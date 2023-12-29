package infotab_configuracao

import (
	"net/http"
	"strconv"

	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
	respostas "github.com/bornium-be/M10_001-CadPais/src/Respostas"
	"github.com/gorilla/mux"
)

var repositorio TTabConfiguracao_Repositorio

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func TabConfiguracao_CodRegimeTributario(w http.ResponseWriter, r *http.Request) {

	//------------------------------------------------------------------

	//---> "/bornium/tabconfiguracao/codregimetributario",

	parametros := mux.Vars(r)

	codEmpresa, erro := strconv.ParseInt(parametros["codEmpresa"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//---------------------

	erro = repositorio.TabConfiguracao_RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.TabConfiguracao_RepositorioFechar()

	//---------------------

	regNomeFantasia, erro := repositorio.TabConfiguracao_CodRegimeTributario(codEmpresa)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var info declaracao.Resultado

	info.Result = append(info.Result, regNomeFantasia)

	respostas.JSON(w, http.StatusOK, info)

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

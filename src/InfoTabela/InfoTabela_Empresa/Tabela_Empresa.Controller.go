package infotabela_empresa

import (
	"net/http"

	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
	respostas "github.com/bornium-be/M10_001-CadPais/src/Respostas"
)

var repositorio TTabEmpresa_Repositorio

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func TabEmpresa_RecNomeFantasia(w http.ResponseWriter, r *http.Request) {

	//------------------------------------------------------------------

	//---> "/bornium/tabempresa/nomefantasia", //--> Parametro: CNPJ

	cnpj := r.URL.Query().Get("cnpj")

	//---------------------

	erro := repositorio.TabEmpresa_RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.TabEmpresa_RepositorioFechar()

	//---------------------

	regNomeFantasia, erro := repositorio.TabEmpresa_RecNomeFantasia(cnpj)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var info declaracao.Resultado

	if regNomeFantasia.NomeFantasia != "" {
		info.Result = append(info.Result, regNomeFantasia)
	}

	respostas.JSON(w, http.StatusOK, info)

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

package cadpais

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strconv"

	constantes "github.com/bornium-be/M10_001-CadPais/src/Constantes"
	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
	respostas "github.com/bornium-be/M10_001-CadPais/src/Respostas"
	"github.com/gorilla/mux"
)

var repositorio TCadPais_Repositorio

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func buscarTodos(w http.ResponseWriter, r *http.Request) {

	type TResultado struct {
		Result []TReg_PaisFiltro `json:"result"`
	}

	//------------------------------------------------------------------

	//---> "/bornium/v1/{EmpresaId}/cadPais", //--> Parametro: PaisNome

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//---------------------

	paisNome := r.URL.Query().Get("PaisNome")

	//---------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//---------------------

	var info TResultado

	info.Result, erro = repositorio.BuscarTodos(empresaId, paisNome)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, info)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func buscarPorId(w http.ResponseWriter, r *http.Request) {

	//-----> "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	paisIdEmp, erro := strconv.ParseInt(parametros["PaisIdEmp"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//---------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//---------------------

	registro, erro := repositorio.BuscarPorId(empresaId, paisIdEmp)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var info declaracao.Resultado

	if registro.PaisIdEmp > 0 {
		info.Result = append(info.Result, registro)
	}

	respostas.JSON(w, http.StatusOK, info)
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func codNovoRegistro(w http.ResponseWriter, r *http.Request) {

	//----->  "/bornium/v1/{EmpresaId}/cadPais.codNovoRegistro",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//----------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//----------------------

	var regCodigo declaracao.RegCodigo

	if regCodigo, erro = repositorio.CodNovoRegistro(empresaId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//---------------------

	var info declaracao.Resultado

	info.Result = append(info.Result, regCodigo)

	respostas.JSON(w, http.StatusOK, info)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func inserirRegistro(w http.ResponseWriter, r *http.Request) {

	//-----> "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	paisIdEmp, erro := strconv.ParseInt(parametros["PaisIdEmp"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//-------------

	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var registro TReg_Pais

	if erro = json.Unmarshal(corpoRequest, &registro); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		fmt.Println(erro.Error())
		return
	}

	if erro := registro.Preparar(constantes.CTE_CADASTRAR); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//----------------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//----------------------------

	erro = repositorio.InserirRegistro(empresaId, paisIdEmp, registro)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//---------------

	regCodigo := declaracao.RegCodigo{}

	regCodigo.Codigo, _ = repositorio.GetCodRegistro(empresaId, paisIdEmp)

	//---------------

	var info declaracao.Resultado

	info.Result = append(info.Result, regCodigo)

	respostas.JSON(w, http.StatusCreated, info)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func alterarRegistro(w http.ResponseWriter, r *http.Request) {

	//----->  "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	paisIdEmp, erro := strconv.ParseInt(parametros["PaisIdEmp"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// -------------------------------

	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var registro TReg_Pais

	if erro = json.Unmarshal(corpoRequest, &registro); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := registro.Preparar(constantes.CTE_ATUALIZAR); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//----------------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//----------------------------

	if erro = repositorio.AlterarRegistro(empresaId, paisIdEmp, registro); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func deletar(w http.ResponseWriter, r *http.Request) {

	//-----> "/bornium/v1/{EmpresaId}/cadPais/{PaisIdEmp}",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	empresaId, erro := strconv.ParseInt(parametros["EmpresaId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	paisIdEmp, erro := strconv.ParseInt(parametros["PaisIdEmp"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//----------------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//----------------------------

	if erro = repositorio.DeletarRegistro(empresaId, paisIdEmp); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func ibge_Buscar(w http.ResponseWriter, r *http.Request) {

	//-----> "/bornium/v1/{EmpresaId}/cadPais.ibge/{IBGE_Id}",

	//------------------------------------------------------------------

	parametros := mux.Vars(r)

	ibge_Id, erro := strconv.ParseInt(parametros["IBGE_Id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//----------------------------

	erro = repositorio.RepositorioAbrir()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer repositorio.RepositorioFechar()

	//----------------------------

	registro, erro := repositorio.BuscarIBGE(ibge_Id)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var info declaracao.Resultado

	if registro.PaisId > 0 {
		info.Result = append(info.Result, registro)
	}

	respostas.JSON(w, http.StatusOK, info)
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

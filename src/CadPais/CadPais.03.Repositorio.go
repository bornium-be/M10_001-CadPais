package cadpais

import (
	"fmt"
	"strings"

	atlaslog "github.com/bornium-be/M10_001-CadPais/src/AtlasLog"
	conexao "github.com/bornium-be/M10_001-CadPais/src/Conexao"
	conversor "github.com/bornium-be/M10_001-CadPais/src/Conversor"
	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
	infotabela "github.com/bornium-be/M10_001-CadPais/src/InfoTabela"
	sgdb "github.com/bornium-be/M10_001-CadPais/src/SGDB"
)

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

type TCadPais_Repositorio struct {
	conexao.TConexao
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorio TCadPais_Repositorio) BuscarTodos(empresaId int64, paisNome string) ([]TReg_PaisFiltro, error) {

	var conversor conversor.TConversor

	paisNome = strings.ToUpper(paisNome)

	strSQL := `  select TA.PaisId, 
	                    TA.PaisIdEmp, 
	                    TA.PaisNome 
	               from PAIS TA
				  Where TA.EmpresaId = ? 
				    And UPPER(TA.PaisNome) LIKE ?
			  ` //---> strSQL

	strSQL = conversor.C001(strSQL)

	linhas, erro := repositorio.BD().Query(strSQL, empresaId, paisNome)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var listaRegistro []TReg_PaisFiltro

	for linhas.Next() {
		var sqlRegistro TSQL_PaisFiltro

		if erro = linhas.Scan(
			&sqlRegistro.PaisId,
			&sqlRegistro.PaisIdEmp,
			&sqlRegistro.PaisNome,
		); erro != nil {
			return nil, erro
		}

		listaRegistro = append(listaRegistro, sqlRegistro.Converter())
	}

	return listaRegistro, nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) BuscarPorId(codEmpresa, codPaisEmp int64) (TReg_Pais, error) {

	var conversor conversor.TConversor

	strSQL := `select TA.EmpresaId, 
	                  TA.PaisId, 
					  TA.PaisIdEmp, 
					  TA.PaisNome, 
					  TA.IBGE_Id
				 from PAIS TA
				 Where TA.EmpresaId = ?
				   And TA.PaisIdEmp = ?
			  ` //--> Fim SQL

	strSQL = conversor.C001(strSQL)

	linhas, erro := repositorio.BD().Query(strSQL, codEmpresa, codPaisEmp)

	if erro != nil {
		return TReg_Pais{}, erro
	}
	defer linhas.Close()

	var sqlPais TSQL_RegPais
	var regPais TReg_Pais

	if linhas.Next() {
		if erro = linhas.Scan(
			&sqlPais.EmpresaId,
			&sqlPais.PaisId,
			&sqlPais.PaisIdEmp,
			&sqlPais.PaisNome,
			&sqlPais.IBGE_Id,
		); erro != nil {
			return TReg_Pais{}, erro
		}

		regPais = sqlPais.Converter()

		regIBGE, erro := repositorio.BuscarIBGE(regPais.IBGE_Id)

		regPais.InfoIBGE = regIBGE

		if erro != nil {
			return TReg_Pais{}, erro
		}
	}
	return regPais, nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) InserirRegistro(codEmpresa, codPaisEmp int64, regPais TReg_Pais) error {

	var conversor conversor.TConversor

	strSQL := `insert into PAIS(EmpresaId, PaisIdEmp, PaisNome, IBGE_Id) 
	           values(?, ?, ?, ?)`

	strSQL = conversor.C001(strSQL)

	statement, erro := repositorio.BD().Prepare(strSQL)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	sqlPais := regPais.Converter()

	_, erro = statement.Exec(
		sqlPais.EmpresaId,
		sqlPais.PaisIdEmp,
		sqlPais.PaisNome,
		sqlPais.IBGE_Id,
	)

	if erro != nil {
		return erro
	}

	nomeColecao := infotabela.CNPJ_Empresa(codEmpresa)

	atlaslog.Atlas_Inserir(nomeColecao, codEmpresa, regPais.RegistroLog)

	return nil

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) AlterarRegistro(codEmpresa, CodPaisEmp int64, regPais TReg_Pais) error {

	var conversor conversor.TConversor

	strSQL := `Update PAIS 
	           set PaisNome=?, IBGE_Id=? 
	    	   Where EmpresaId = ? 
			     and PaisIdEmp  = ?`

	strSQL = conversor.C001(strSQL)

	statement, erro := repositorio.BD().Prepare(strSQL)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	sqlPais := regPais.Converter()

	if _, erro = statement.Exec(
		sqlPais.PaisNome,
		sqlPais.IBGE_Id,
		sqlPais.EmpresaId,
		sqlPais.PaisIdEmp,
	); erro != nil {
		return erro
	}

	nomeColecao := infotabela.CNPJ_Empresa(codEmpresa)

	atlaslog.Atlas_Inserir(nomeColecao, codEmpresa, regPais.RegistroLog)

	return nil

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) DeletarRegistro(codEmpresa, codPaisEmp int64) error {

	var conversor conversor.TConversor

	strSQL := `Delete from PAIS 
	    	   Where EmpresaId = ?
			     and PaisIdEmp = ?`

	strSQL = conversor.C001(strSQL)

	statement, erro := repositorio.BD().Prepare(strSQL)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(codEmpresa, codPaisEmp); erro != nil {
		return erro
	}

	return nil

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) BuscarIBGE(IBGE_Id int64) (TReg_IBGE_Pais, error) {

	var conversor conversor.TConversor

	strSQL := `select PaisId, 
	                  PaisNome 
				 from IBGE_PAIS 
   				Where PaisId = ?`

	strSQL = conversor.C001(strSQL)

	linhas, erro := repositorio.BD().Query(strSQL, IBGE_Id)

	if erro != nil {
		return TReg_IBGE_Pais{}, erro
	}
	defer linhas.Close()

	var sql_IBGE TSQL_RegIBGE_Pais

	if linhas.Next() {
		if erro = linhas.Scan(
			&sql_IBGE.PaisId,
			&sql_IBGE.PaisNome,
		); erro != nil {
			return TReg_IBGE_Pais{}, erro
		}
	}

	return sql_IBGE.Converte(), nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) CodNovoRegistro(codEmpresa int64) (declaracao.RegCodigo, error) {

	var registro declaracao.RegCodigo

	filtro := fmt.Sprintf("EmpresaId = %d", codEmpresa)

	recInfo := sgdb.RecInformacao{}
	recInfo.DB = repositorio.BD()

	novoCodigo, erro := recInfo.CodNovoRegistro("PaisIdEmp", "PAIS", filtro)

	if erro != nil {
		return declaracao.RegCodigo{}, erro
	}

	registro.Codigo = novoCodigo

	return registro, erro
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorTio TCadPais_Repositorio) GetCodRegistro(codEmpresa, codPaisEmp int64) (int64, error) {

	filtro := fmt.Sprintf("(EmpresaId = %d) and (PaisIdEmp = %d)", codEmpresa, codPaisEmp)

	recInfo := sgdb.RecInformacao{}
	recInfo.DB = repositorio.BD()

	codigo, erro := recInfo.RecInfo_Int("PaisId", "PAIS", filtro)
	if erro != nil {
		return -1, erro
	}

	return codigo, nil

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

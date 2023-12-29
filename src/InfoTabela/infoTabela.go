package infotabela

import (
	"database/sql"

	sgdb "github.com/bornium-be/M10_001-CadPais/src/SGDB"
)

type InfoTabela_Repositorio struct {
	//----> R01 - Autidato em 29/11/2022
	db *sql.DB
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func IntoTabela_NovoRepositorio(db *sql.DB) *InfoTabela_Repositorio {
	return &InfoTabela_Repositorio{db}
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorio InfoTabela_Repositorio) infoTabela_CNPJ(codEmpresa int64) (string, error) {

	strSQL := `select TA.CNPJ 
				 from Empresa TA
				 Where TA.cdEmpresa = ?
			  ` // ---> strSQL
	linhas, erro := repositorio.db.Query(strSQL, codEmpresa)

	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var cnpjEmpresa string

	if linhas.Next() {
		if erro = linhas.Scan(
			&cnpjEmpresa,
		); erro != nil {
			return "", erro
		}

		//----------------

	}
	return cnpjEmpresa, nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func CNPJ_Empresa(empresaId int64) string {

	db, erro := sgdb.Conectar()

	if erro != nil {
		return ""
	}
	defer db.Close()

	repositorio := IntoTabela_NovoRepositorio(db)
	nomeColecao, _ := repositorio.infoTabela_CNPJ(empresaId)

	return nomeColecao

}

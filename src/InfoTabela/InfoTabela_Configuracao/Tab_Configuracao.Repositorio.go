package infotab_configuracao

import (
	"database/sql"

	sgdb "github.com/bornium-be/M10_001-CadPais/src/SGDB"
	util "github.com/bornium-be/M10_001-CadPais/src/Util"
)

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

type TTabConfiguracao_Repositorio struct {
	db *sql.DB
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorio *TTabConfiguracao_Repositorio) TabConfiguracao_RepositorioAbrir() error {

	var erro error

	repositorio.db, erro = sgdb.Conectar()

	if erro != nil {
		return erro
	}

	return nil

}

func (repositorio *TTabConfiguracao_Repositorio) TabConfiguracao_RepositorioFechar() {

	repositorio.db.Close()

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (repositorio TTabConfiguracao_Repositorio) TabConfiguracao_CodRegimeTributario(codEmpresa int64) (TRegCodRegimeTributario, error) {

	strSQL := `  select TA.codRegimeTributario 	 
	               from CONFIGURACAO TA
				  Where TA.cdConfiguracao = ?	
			  ` //---> strSQL

	linhas, erro := repositorio.db.Query(strSQL, codEmpresa)

	if erro != nil {
		return TRegCodRegimeTributario{}, erro
	}
	defer linhas.Close()

	var regEmpresa TRegCodRegimeTributario

	if linhas.Next() {

		if erro = linhas.Scan(
			&regEmpresa.iCodRegimeTributario,
		); erro != nil {
			return TRegCodRegimeTributario{}, erro
		}

		regEmpresa.CodRegimeTributario = util.ITF_toStr(regEmpresa.iCodRegimeTributario)

		if erro != nil {
			return TRegCodRegimeTributario{}, erro
		}
	}

	return regEmpresa, nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

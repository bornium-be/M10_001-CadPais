package cadpais

import (
	"errors"
	"strings"

	mongo_sgdb "github.com/bornium-be/M10_001-CadPais/src/Mongo_SGDB"
	util "github.com/bornium-be/M10_001-CadPais/src/Util"
)

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

type TReg_Pais struct {
	PaisId      int64               `json:"PaisId"`
	PaisIdEmp   int64               `json:"PaisIdEmp"`
	EmpresaId   int64               `json:"EmpresaId"`
	PaisNome    string              `json:"PaisNome"`
	IBGE_Id     int64               `json:"IBGE_Id"`
	InfoIBGE    TReg_IBGE_Pais      `json:"infoIBGE,omitempty"`
	RegistroLog mongo_sgdb.Registro `json:"registroLog,omitempty"`
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (reg *TReg_Pais) Converter() TSQL_RegPais {

	var novoReg TSQL_RegPais

	novoReg.PaisId = util.SQL_int64(reg.PaisId, 0)
	novoReg.PaisIdEmp = util.SQL_int64(reg.PaisIdEmp, 0)
	novoReg.EmpresaId = util.SQL_int64(reg.EmpresaId, 0)
	novoReg.PaisNome = util.SQL_String(reg.PaisNome)
	novoReg.IBGE_Id = util.SQL_int64(reg.IBGE_Id, 0)
	novoReg.InfoIBGE = reg.InfoIBGE
	novoReg.RegistroLog = reg.RegistroLog

	return novoReg
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (u *TReg_Pais) validar(etapa int) error {

	if u.PaisNome == "" {
		return errors.New("nome do pais é obrigatório e não pode ser nulo")
	}

	return nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (u *TReg_Pais) formatar(etapa int) error {
	u.PaisNome = strings.TrimSpace(u.PaisNome)

	return nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (u *TReg_Pais) Preparar(etapa int) error {

	if erro := u.formatar(etapa); erro != nil {
		return erro
	}

	if erro := u.validar(etapa); erro != nil {
		return erro
	}

	return nil
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

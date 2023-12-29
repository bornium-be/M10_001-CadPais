package cadpais

import (
	"database/sql"

	mongo_sgdb "github.com/bornium-be/M10_001-CadPais/src/Mongo_SGDB"
)

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

type TSQL_RegPais struct {
	PaisId     sql.NullInt64
	PaisIdEmp  sql.NullInt64
	EmpresaId  sql.NullInt64
	PaisNome    sql.NullString
	IBGE_Id     sql.NullInt64
	InfoIBGE    TReg_IBGE_Pais
	RegistroLog mongo_sgdb.Registro
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (reg *TSQL_RegPais) Converter() TReg_Pais {

	var novoReg TReg_Pais

	novoReg.PaisId = reg.PaisId.Int64
	novoReg.PaisIdEmp = reg.PaisIdEmp.Int64
	novoReg.EmpresaId = reg.EmpresaId.Int64
	novoReg.PaisNome = reg.PaisNome.String
	novoReg.IBGE_Id = reg.IBGE_Id.Int64
	novoReg.InfoIBGE = reg.InfoIBGE
	novoReg.RegistroLog = reg.RegistroLog

	return novoReg
}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
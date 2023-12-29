package sgdb

import (
	"database/sql"

	configSistema "github.com/bornium-be/M10_001-CadPais/src/ConfigSistema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/nakagami/firebirdsql"
)

func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("firebirdsql", configSistema.StrConexaoBanco_FB)

	// db, erro := sql.Open("mysql", configSistema.StrConexaoBanco_MySQL)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}

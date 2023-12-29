package configSistema

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	
)

var (
	StrConexaoBanco_FB    = ""
	StrConexaoBanco_MySQL = ""
	Porta                 = 0
	SecureKey             []byte
)

func Carregar() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Porta = 9000
	}

	StrConexaoBanco_FB = fmt.Sprintf("%s:%s@%s/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_SERVIDOR"),
		os.Getenv("DB_NOME_ARQUIVO"),
	)

	StrConexaoBanco_MySQL = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_SERVIDOR"),
		os.Getenv("DB_NOME_ARQUIVO"),
	)

}

package configRotas

import (
	cadpais "github.com/bornium-be/M10_001-CadPais/src/CadPais"
	declaracao "github.com/bornium-be/M10_001-CadPais/src/Declaracao"
	"github.com/gorilla/mux"
)

func Configurar(r *mux.Router) *mux.Router {

	rotas := make([]declaracao.Rota, 1)

	rotas = append(rotas, cadpais.Rotas_CadPais...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r

}

package mongo_sgdb

type Registro struct {
	//ID           primitive.ObjectID `json:"_id,omitempty"     		bson:"_id,omitempty"`
	//ID           string `json:"_id,omitempty"  	        bson:"_id,omitempty"`
	EmpresaCodigo int64       `json:"empresaCodigo,omitempty" bson:"procesoCodigo,omitempty"`
	EmpresaCNPJ   string      `json:"empresaCNPJ,omitempty" bson:"empresaCNPJ,omitempty"`
	ProcessoCod   int64       `json:"processoCod,omitempty" bson:"processoCod,omitempty"`
	ProcessoNome  string      `json:"processoNome,omitempty" bson:"processoNome,omitempty"`
	TabelaNome    string      `json:"tabelaNome,omitempty" bson:"tabelaNome,omitempty"`
	Data          float64     `json:"data,omitempty" bson:"data,omitempty"`
	UsuraioCod    int64       `json:"usuarioCod,omitempty" bson:"usuarioCod,omitempty"`
	UsaurioNome   string      `json:"UsuarioNome,omitempty" bson:"usuarioNome,omitempty"`
	RegCodEmp     string      `json:"regCodEmp,omitempty" bson:"regCodEmp,omitempty"`
	RegCodItem    string      `json:"regCodItem,omitempty" bson:"regCodItem,omitempty"`
	Producao      string      `json:"producao,omitempty" bson:"producao,omitempty"`
	AtributoLog   []CampoInfo `json:"atributoLog,omitempty" bson:"atributoLog,omitempty"`
}





type CampoInfo struct {
	CodItem  string `json:"codItem,omitempty" bson:"codItem,omitempty"`
	Atributo string `json:"atributo,omitempty" bson:"atributo,omitempty"`
	Antigo   string `json:"antigo,omitempty" bson:"antigo,omitempty"`
	Novo     string `json:"novo,omitempty" bson:"novo,omitempty"`
}

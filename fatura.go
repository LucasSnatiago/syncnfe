package syncnfe

// Fatura Estrutura que armazena todos os elementos de uma fatura
type Fatura struct {
	ID1                  int
	ID2                  int
	DescricaoServico     string
	ValorUnitario        int
	ICMS                 string
	AliquotaReducao      string
	Unidade              string
	QuantidadeContratada int
	QuantidadeFornecida  int
	AliquotaICMS         int
	CodigoClassificacao  int
	// A partir daqui todos os itens são opcionais
	BC                          string
	ValoresIsentos              string
	OutrosValores               string
	Desconto                    int
	ValorAproxTributosFederais  int
	ValorAproxTributosEstadual  int
	ValorAproxTributosMunicipal int
}

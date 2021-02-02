package syncnfe

// NewFatura Retornar uma nova fatura
func NewFatura() *Fatura {
	return &Fatura{
		ID1:                         0,
		ID2:                         0,
		ValorUnitario:               0,
		QuantidadeContratada:        0,
		QuantidadeFornecida:         0,
		AliquotaICMS:                0,
		AliquotaReducao:             0,
		BC:                          0,
		ValoresIsentos:              0,
		OutrosValores:               0,
		Desconto:                    0,
		ValorAproxTributosFederais:  0,
		ValorAproxTributosEstadual:  0,
		ValorAproxTributosMunicipal: 0,
	}
}

// Fatura Estrutura que armazena todos os elementos de uma fatura
type Fatura struct {
	ID1                  int
	ID2                  int
	DescricaoServico     string
	ValorUnitario        int
	ICMS                 string
	AliquotaReducao      int
	Unidade              string
	QuantidadeContratada int
	QuantidadeFornecida  int
	AliquotaICMS         int
	CodigoClassificacao  int
	// A partir daqui todos os itens são opcionais
	BC                          int
	ValoresIsentos              int
	OutrosValores               int
	Desconto                    int
	ValorAproxTributosFederais  int
	ValorAproxTributosEstadual  int
	ValorAproxTributosMunicipal int
}

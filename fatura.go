package syncnfe

import "fmt"

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

// String retornar a Fatura formatada
func (f Fatura) String() string {
	return fmt.Sprintf("%d|%d|%s|%d|%s|%d|%s|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|", f.ID1, f.ID2,
		f.DescricaoServico, f.ValorUnitario, f.ICMS, f.AliquotaReducao, f.Unidade, f.QuantidadeContratada,
		f.QuantidadeFornecida, f.AliquotaICMS, f.CodigoClassificacao, f.BC, f.ValoresIsentos, f.OutrosValores,
		f.Desconto, f.ValorAproxTributosFederais, f.ValorAproxTributosEstadual, f.ValorAproxTributosMunicipal)
}

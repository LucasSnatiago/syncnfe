package syncnfe

// SyncNFe Interface que contem os métodos e estruturas para a geração de uma nota fiscal eletrônica pela SyncNFe
type SyncNFe interface {
	NewCliente() *Cliente
	NewFatura() *Fatura
}

package bpa

import "errors"

var (
	// ErrValidacaoBpa indica erro de validacao nos dados de entrada.
	ErrValidacaoBpa = errors.New("bpa: erro de validacao")

	// ErrTipoMisturado indica tentativa de misturar BPA-C e BPA-I no mesmo arquivo.
	ErrTipoMisturado = errors.New("bpa: nao e permitido misturar BPA-C e BPA-I no mesmo arquivo")

	// ErrPaginacaoBpa indica erro de paginacao.
	ErrPaginacaoBpa = errors.New("bpa: erro de paginacao")

	// ErrCampoControleBpa indica erro no campo de controle.
	ErrCampoControleBpa = errors.New("bpa: erro no campo de controle")
)

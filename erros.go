package bpa

import "errors"

var (
	// ErrValidacao indica erro de validacao nos dados de entrada.
	ErrValidacao = errors.New("bpa: erro de validacao")

	// ErrTipoMisturado indica tentativa de misturar BPA-C e BPA-I no mesmo arquivo.
	ErrTipoMisturado = errors.New("bpa: nao e permitido misturar BPA-C e BPA-I no mesmo arquivo")
)

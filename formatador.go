package bpa

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// formatarNumerico alinha o valor a direita e preenche com zeros a esquerda.
func formatarNumerico(valor string, tamanho int) string {
	if len(valor) >= tamanho {
		return valor[:tamanho]
	}
	return strings.Repeat("0", tamanho-len(valor)) + valor
}

// formatarAlfanumerico alinha o valor a esquerda e preenche com espacos a direita.
func formatarAlfanumerico(valor string, tamanho int) string {
	if len(valor) >= tamanho {
		return valor[:tamanho]
	}
	return valor + strings.Repeat(" ", tamanho-len(valor))
}

// normalizarTexto converte para maiusculas e remove diacriticos (acentos).
func normalizarTexto(texto string) string {
	texto = strings.ToUpper(texto)
	// Decompoe em NFD para separar caracteres base de diacriticos
	resultado := norm.NFD.String(texto)
	// Remove caracteres de combinacao (diacriticos)
	var sb strings.Builder
	for _, r := range resultado {
		if !unicode.Is(unicode.Mn, r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

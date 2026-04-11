package bpa

import "testing"

func TestFormatarNumerico(t *testing.T) {
	testes := []struct {
		nome     string
		valor    string
		tamanho  int
		esperado string
	}{
		{"preenche com zeros", "123", 6, "000123"},
		{"valor exato", "123456", 6, "123456"},
		{"valor maior trunca", "1234567", 6, "123456"},
		{"valor vazio", "", 4, "0000"},
	}
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := formatarNumerico(tt.valor, tt.tamanho)
			if resultado != tt.esperado {
				t.Errorf("formatarNumerico(%q, %d) = %q, esperado %q", tt.valor, tt.tamanho, resultado, tt.esperado)
			}
		})
	}
}

func TestFormatarAlfanumerico(t *testing.T) {
	testes := []struct {
		nome     string
		valor    string
		tamanho  int
		esperado string
	}{
		{"preenche com espacos", "ABC", 6, "ABC   "},
		{"valor exato", "ABCDEF", 6, "ABCDEF"},
		{"valor maior trunca", "ABCDEFG", 6, "ABCDEF"},
		{"valor vazio", "", 3, "   "},
	}
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := formatarAlfanumerico(tt.valor, tt.tamanho)
			if resultado != tt.esperado {
				t.Errorf("formatarAlfanumerico(%q, %d) = %q, esperado %q", tt.valor, tt.tamanho, resultado, tt.esperado)
			}
		})
	}
}

func TestNormalizarTexto(t *testing.T) {
	testes := []struct {
		nome     string
		texto    string
		esperado string
	}{
		{"acentos e maiusculas", "João José", "JOAO JOSE"},
		{"ja maiusculo", "ABC", "ABC"},
		{"cedilha", "Ação", "ACAO"},
		{"til", "São Paulo", "SAO PAULO"},
		{"vazio", "", ""},
	}
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := normalizarTexto(tt.texto)
			if resultado != tt.esperado {
				t.Errorf("normalizarTexto(%q) = %q, esperado %q", tt.texto, resultado, tt.esperado)
			}
		})
	}
}

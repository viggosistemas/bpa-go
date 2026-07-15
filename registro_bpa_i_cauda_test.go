package bpa

import (
	"errors"
	"strings"
	"testing"
)

// linhaBpaI constroi um arquivo com um unico registro BPA-I e devolve a linha
// de conteudo (338 bytes, sem o CRLF): bytes [132 : 132+338] do arquivo.
func linhaBpaI(t *testing.T, reg RegistroBpaIEntrada) string {
	t.Helper()
	c := NovoConstrutor(cabecalhoValido())
	if err := c.AdicionarRegistroBpaI(reg); err != nil {
		t.Fatalf("erro ao adicionar registro: %v", err)
	}
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro ao construir: %v", err)
	}
	// cabecalho = 132 bytes; registro BPA-I = 340 bytes (338 + CRLF).
	linha := string(dados[132 : 132+338])
	if len(linha) != 338 {
		t.Fatalf("conteudo da linha BPA-I = %d bytes, esperado 338", len(linha))
	}
	return linha
}

// registroICaudaCompleta devolve um registro com TODOS os campos da cauda
// preenchidos, com valores de tamanho exato para checar as posicoes.
func registroICaudaCompleta() RegistroBpaIEntrada {
	r := registroIValido()
	r.RacaCor = "05"
	r.Etnia = "0042"
	r.Nacionalidade = "010"
	r.Servico = "121"
	r.Classificacao = "005"
	r.EquipeSeq = "00000123"
	r.EquipeArea = "0044"
	r.Cnpj = "12345678000199"
	r.Cep = "01310100"
	r.CodLogradouro = "008"
	r.Endereco = "AVENIDA PAULISTA"
	r.Complemento = "APTO 12"
	r.Numero = "1578"
	r.Bairro = "BELA VISTA"
	r.Telefone = "11987654321"
	r.Email = "PACIENTE@EXEMPLO.COM"
	r.Ine = "0000123456"
	return r
}

// TestBpaICaudaPosicoes garante que cada campo da cauda (151-338) cai no
// intervalo de bytes EXATO do layout DATASUS. Posicao 1-indexada P -> [P-1:Fim].
func TestBpaICaudaPosicoes(t *testing.T) {
	linha := linhaBpaI(t, registroICaudaCompleta())

	casos := []struct {
		nome     string
		ini, fim int // posicoes 1-indexadas do layout (inclusivas)
		esperado string
	}{
		{"raca", 151, 152, "05"},
		{"etnia", 153, 156, "0042"},
		{"nacionalidade", 157, 159, "010"},
		{"servico", 160, 162, "121"},
		{"classificacao", 163, 165, "005"},
		{"equipe_seq", 166, 173, "00000123"},
		{"equipe_area", 174, 177, "0044"},
		{"cnpj", 178, 191, "12345678000199"},
		{"cep", 192, 199, "01310100"},
		{"logradouro", 200, 202, "008"},
		{"endereco", 203, 232, "AVENIDA PAULISTA" + strings.Repeat(" ", 14)},
		{"complemento", 233, 242, "APTO 12" + strings.Repeat(" ", 3)},
		{"numero", 243, 247, "1578 "},
		{"bairro", 248, 277, "BELA VISTA" + strings.Repeat(" ", 20)},
		{"telefone", 278, 288, "11987654321"},
		{"email", 289, 328, "PACIENTE@EXEMPLO.COM" + strings.Repeat(" ", 20)},
		{"ine", 329, 338, "0000123456"},
	}
	for _, c := range casos {
		got := linha[c.ini-1 : c.fim]
		if got != c.esperado {
			t.Errorf("%s [%d-%d] = %q, esperado %q", c.nome, c.ini, c.fim, got, c.esperado)
		}
	}
}

// TestBpaICaudaVaziaEBranca garante retrocompatibilidade: sem campos da cauda,
// as posicoes 151-338 sao TODAS espacos (mesmo comportamento do reservado(188)
// das versoes <= v1.0.1) e a linha continua com 338 bytes de conteudo.
func TestBpaICaudaVaziaEBranca(t *testing.T) {
	linha := linhaBpaI(t, registroIValido())
	cauda := linha[150:338] // posicoes 151-338
	if cauda != strings.Repeat(" ", 188) {
		t.Errorf("cauda sem dados deveria ser 188 brancos, obteve %q", cauda)
	}
}

// TestBpaICaudaEtniaSoIndigena documenta a regra de negocio: quando raca != 05,
// a etnia deve vir em branco. A lib nao aplica a regra (responsabilidade do
// caller), mas garante que etnia vazia gera 4 brancos na posicao correta.
func TestBpaICaudaEtniaBrancoQuandoVazia(t *testing.T) {
	r := registroICaudaCompleta()
	r.RacaCor = "01" // branca
	r.Etnia = ""     // caller nao deve preencher etnia p/ nao-indigena
	linha := linhaBpaI(t, r)
	if got := linha[152:156]; got != "    " {
		t.Errorf("etnia vazia [153-156] = %q, esperado 4 brancos", got)
	}
	if got := linha[150:152]; got != "01" {
		t.Errorf("raca [151-152] = %q, esperado 01", got)
	}
}

// TestBpaICaudaValidacaoNumerica garante que campo numerico com nao-digito e
// rejeitado (ErrValidacaoBpa) e nao gera arquivo silenciosamente corrompido.
func TestBpaICaudaValidacaoNumerica(t *testing.T) {
	r := registroICaudaCompleta()
	r.RacaCor = "XX"
	c := NovoConstrutor(cabecalhoValido())
	err := c.AdicionarRegistroBpaI(r)
	if !errors.Is(err, ErrValidacaoBpa) {
		t.Errorf("esperava ErrValidacaoBpa para raca nao-numerica, obteve %v", err)
	}
}

// TestBpaICaudaTamanhoTotal reforca que preencher a cauda NAO altera o tamanho
// de 340 bytes da linha (a cauda substitui exatamente os 188 bytes reservados).
func TestBpaICaudaTamanhoTotal(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	if err := c.AdicionarRegistroBpaI(registroICaudaCompleta()); err != nil {
		t.Fatalf("erro ao adicionar: %v", err)
	}
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro ao construir: %v", err)
	}
	if resto := dados[132:]; len(resto) != 340 {
		t.Errorf("linha BPA-I com cauda cheia = %d bytes, esperado 340", len(resto))
	}
}

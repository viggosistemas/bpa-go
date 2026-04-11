package bpa

import (
	"errors"
	"strings"
	"testing"
)

func cabecalhoValido() CabecalhoEntrada {
	return CabecalhoEntrada{
		Competencia:      "202604",
		OrgaoResponsavel: "SMS SAO PAULO",
		SiglaOrgao:       "SMS",
		CnpjCpf:          "12345678901234",
		OrgaoDestino:     "SES SAO PAULO",
		TipoDestino:      "E",
		VersaoSistema:    "1.0",
	}
}

func registroCValido() RegistroBpaCEntrada {
	return RegistroBpaCEntrada{
		Cnes:         "1234567",
		Competencia:  "202604",
		Cbo:          "225142",
		Procedimento: "0301010064",
		Idade:        30,
		Quantidade:   1,
		Origem:       "BPA",
	}
}

func registroIValido() RegistroBpaIEntrada {
	return RegistroBpaIEntrada{
		Cnes:            "1234567",
		Competencia:     "202604",
		CnsProfissional: "123456789012345",
		Cbo:             "225142",
		DataAtendimento: "20260401",
		Procedimento:    "0301010064",
		CnsPaciente:     "987654321012345",
		Sexo:            "M",
		MunicipioIbge:   "355030",
		Cid:             "J060",
		Idade:           30,
		Quantidade:      1,
		Origem:          "BPA",
		NomePaciente:    "JOAO DA SILVA",
		DataNascimento:  "19960115",
	}
}

func TestTamanhoCabecalho(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaC(registroCValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	linhas := strings.Split(string(dados), "\r\n")
	// A primeira linha e o cabecalho (sem contar o CRLF final)
	cabLinha := linhas[0]
	// 132 bytes = conteudo sem CRLF
	if len(cabLinha) != 130 {
		t.Errorf("tamanho do cabecalho = %d, esperado 130 (132 com CRLF)", len(cabLinha))
	}
}

func TestTamanhoBpaC(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaC(registroCValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	linhas := strings.Split(string(dados), "\r\n")
	// linhas[1] e o registro BPA-C (sem CRLF)
	regLinha := linhas[1]
	// 50 bytes = conteudo sem CRLF = 48
	if len(regLinha) != 48 {
		t.Errorf("tamanho do registro BPA-C = %d, esperado 48 (50 com CRLF)", len(regLinha))
	}
}

func TestTamanhoBpaI(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaI(registroIValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	linhas := strings.Split(string(dados), "\r\n")
	regLinha := linhas[1]
	// 340 bytes = conteudo sem CRLF = 338
	if len(regLinha) != 338 {
		t.Errorf("tamanho do registro BPA-I = %d, esperado 338 (340 com CRLF)", len(regLinha))
	}
}

func TestMisturaProibida(t *testing.T) {
	t.Run("BPA-C depois BPA-I", func(t *testing.T) {
		c := NovoConstrutor(cabecalhoValido())
		_ = c.AdicionarRegistroBpaC(registroCValido())
		err := c.AdicionarRegistroBpaI(registroIValido())
		if !errors.Is(err, ErrTipoMisturado) {
			t.Errorf("esperava ErrTipoMisturado, obteve %v", err)
		}
	})

	t.Run("BPA-I depois BPA-C", func(t *testing.T) {
		c := NovoConstrutor(cabecalhoValido())
		_ = c.AdicionarRegistroBpaI(registroIValido())
		err := c.AdicionarRegistroBpaC(registroCValido())
		if !errors.Is(err, ErrTipoMisturado) {
			t.Errorf("esperava ErrTipoMisturado, obteve %v", err)
		}
	})
}

func TestConstruirSemRegistros(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_, err := c.Construir()
	if !errors.Is(err, ErrValidacaoBpa) {
		t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
	}
}

func TestPaginacao(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	for i := 0; i < 21; i++ {
		reg := registroCValido()
		reg.Quantidade = i + 1
		if err := c.AdicionarRegistroBpaC(reg); err != nil {
			t.Fatalf("erro ao adicionar registro %d: %v", i, err)
		}
	}
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	conteudo := string(dados)
	// O cabecalho deve conter total_folhas = "000002"
	cabecalho := strings.Split(conteudo, "\r\n")[0]
	// total_folhas esta na posicao 2+5+6+6 = 19 ate 19+6 = 25
	totalFolhas := cabecalho[19:25]
	if totalFolhas != "000002" {
		t.Errorf("total_folhas = %q, esperado %q", totalFolhas, "000002")
	}
}

func TestArquivoCompletoBpaC(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	n := 3
	for i := 0; i < n; i++ {
		_ = c.AdicionarRegistroBpaC(registroCValido())
	}
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	// cabecalho(132) + n * registro(50) = 132 + 3*50 = 282
	esperado := 132 + n*50
	if len(dados) != esperado {
		t.Errorf("tamanho total = %d, esperado %d", len(dados), esperado)
	}
}

func TestArquivoCompletoBpaI(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	n := 2
	for i := 0; i < n; i++ {
		_ = c.AdicionarRegistroBpaI(registroIValido())
	}
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	// cabecalho(132) + n * registro(340) = 132 + 2*340 = 812
	esperado := 132 + n*340
	if len(dados) != esperado {
		t.Errorf("tamanho total = %d, esperado %d", len(dados), esperado)
	}
}

func TestCabecalhoTamanhoExato132(t *testing.T) {
	// Testa que o cabecalho tem exatamente 132 bytes (incluindo CRLF)
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaC(registroCValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	// O primeiro \r\n ocorre na posicao 130-131 (0-indexed)
	// Entao os primeiros 132 bytes sao o cabecalho completo
	primeiroCRLF := strings.Index(string(dados), "\r\n")
	tamanhoCabecalhoComCRLF := primeiroCRLF + 2
	if tamanhoCabecalhoComCRLF != 132 {
		t.Errorf("tamanho cabecalho com CRLF = %d, esperado 132", tamanhoCabecalhoComCRLF)
	}
}

func TestRegistroBpaCTamanhoExato50(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaC(registroCValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	// Pegar bytes apos o cabecalho (132 bytes)
	resto := dados[132:]
	// Deve ter exatamente 50 bytes
	if len(resto) != 50 {
		t.Errorf("tamanho registro BPA-C com CRLF = %d, esperado 50", len(resto))
	}
}

func TestRegistroBpaITamanhoExato340(t *testing.T) {
	c := NovoConstrutor(cabecalhoValido())
	_ = c.AdicionarRegistroBpaI(registroIValido())
	dados, err := c.Construir()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	resto := dados[132:]
	if len(resto) != 340 {
		t.Errorf("tamanho registro BPA-I com CRLF = %d, esperado 340", len(resto))
	}
}

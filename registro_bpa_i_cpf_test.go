package bpa

import (
	"errors"
	"strings"
	"testing"
)

// TestBpaIBloco339a353Posicoes garante que CPF do paciente e situacao de rua
// caem nos bytes EXATOS medidos no arquivo do BPA-Magnetico do municipio:
// cpf 339-349, reservado 350, sit_rua 351, reservado 352-353.
func TestBpaIBloco339a353Posicoes(t *testing.T) {
	r := registroIValido()
	r.CnsPaciente = ""
	r.CpfPaciente = "12345678901"
	r.SituacaoRua = "S"
	r.CaraterAtendimento = "02"
	linha := linhaBpaI(t, r)

	casos := []struct {
		nome     string
		ini, fim int // 1-indexado, inclusivo
		esperado string
	}{
		{"cns_paciente em branco", 60, 74, strings.Repeat(" ", 15)},
		{"carater_atendimento", 95, 96, "02"},
		{"naut reservado", 97, 109, strings.Repeat(" ", 13)},
		{"cpf_paciente", 339, 349, "12345678901"},
		{"reservado 350", 350, 350, " "},
		{"sit_rua", 351, 351, "S"},
		{"reservado 352-353", 352, 353, "  "},
	}
	for _, c := range casos {
		if got := linha[c.ini-1 : c.fim]; got != c.esperado {
			t.Errorf("%s [%d-%d] = %q, esperado %q", c.nome, c.ini, c.fim, got, c.esperado)
		}
	}
}

// TestBpaISemCpfSaiBrancoESitRuaN cobre o default: sem CPF as posicoes
// 339-349 saem em branco e a situacao de rua sai "N" (nunca em branco, como
// no arquivo do municipio). O CNS informado continua em 60-74.
func TestBpaISemCpfSaiBrancoESitRuaN(t *testing.T) {
	linha := linhaBpaI(t, registroIValido())
	if got := linha[59:74]; got != "987654321012345" {
		t.Errorf("cns_paciente [60-74] = %q", got)
	}
	if got := linha[338:349]; got != strings.Repeat(" ", 11) {
		t.Errorf("cpf_paciente [339-349] = %q, esperado brancos", got)
	}
	if got := linha[94:96]; got != "  " {
		t.Errorf("carater_atendimento [95-96] = %q, esperado brancos", got)
	}
	if got := linha[350:351]; got != "N" {
		t.Errorf("sit_rua [351] = %q, esperado N", got)
	}
}

func TestBpaIValidacaoCpfESituacaoRua(t *testing.T) {
	casos := map[string]func(*RegistroBpaIEntrada){
		"cpf curto":             func(r *RegistroBpaIEntrada) { r.CpfPaciente = "1234567890" },
		"cpf longo":             func(r *RegistroBpaIEntrada) { r.CpfPaciente = "123456789012" },
		"cpf com letras":        func(r *RegistroBpaIEntrada) { r.CpfPaciente = "1234567890A" },
		"cpf com mascara":       func(r *RegistroBpaIEntrada) { r.CpfPaciente = "123.456.789-01" },
		"situacao rua invalida": func(r *RegistroBpaIEntrada) { r.SituacaoRua = "X" },
		"caten invalido":        func(r *RegistroBpaIEntrada) { r.CaraterAtendimento = "03" },
	}
	for nome, mutar := range casos {
		t.Run(nome, func(t *testing.T) {
			r := registroIValido()
			mutar(&r)
			if err := validarRegistroBpaI(r); !errors.Is(err, ErrValidacaoBpa) {
				t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
			}
		})
	}
	for _, v := range []string{"", "S", "N"} {
		r := registroIValido()
		r.SituacaoRua = v
		if err := validarRegistroBpaI(r); err != nil {
			t.Errorf("situacao_rua %q deveria ser valida: %v", v, err)
		}
	}
}

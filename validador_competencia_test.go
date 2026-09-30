package bpa

import (
	"errors"
	"strings"
	"testing"
)

// Competencia AAAAMM: o mes (2 ultimos digitos) deve ficar entre 01 e 12.
// Espelha a regra D-15 do bpa-ts.
func TestCompetenciaMesValido(t *testing.T) {
	cabecalho := CabecalhoEntrada{
		Competencia:      "202604",
		OrgaoResponsavel: "SMS",
		SiglaOrgao:       "SMS",
		CnpjCpf:          "12345678901234",
		OrgaoDestino:     "SES",
		TipoDestino:      "E",
		VersaoSistema:    "1.0",
	}
	bpaC := RegistroBpaCEntrada{
		Cnes:         "1234567",
		Cbo:          "225142",
		Procedimento: "0301010064",
		Idade:        30,
		Quantidade:   1,
		Origem:       "BPA",
	}
	bpaI := RegistroBpaIEntrada{
		Cnes:            "1234567",
		CnsProfissional: "123456789012345",
		Cbo:             "225142",
		DataAtendimento: "20260401",
		Procedimento:    "0301010064",
		CnsPaciente:     "987654321012345",
		Sexo:            "M",
		MunicipioIbge:   "355030",
		Cid:             "J06 ",
		Idade:           30,
		Quantidade:      1,
		Origem:          "BPA",
		NomePaciente:    "JOAO DA SILVA",
		DataNascimento:  "19960115",
	}

	pontos := []struct {
		nome    string
		validar func(competencia string) error
	}{
		{"cabecalho", func(c string) error {
			e := cabecalho
			e.Competencia = c
			return validarCabecalho(e)
		}},
		{"bpa-c", func(c string) error {
			e := bpaC
			e.Competencia = c
			return validarRegistroBpaC(e)
		}},
		{"bpa-i", func(c string) error {
			e := bpaI
			e.Competencia = c
			return validarRegistroBpaI(e)
		}},
	}

	for _, p := range pontos {
		for _, comp := range []string{"202601", "202604", "202612"} {
			t.Run(p.nome+"/aceita "+comp, func(t *testing.T) {
				if err := p.validar(comp); err != nil {
					t.Errorf("esperava nil, obteve %v", err)
				}
			})
		}
		for _, comp := range []string{"202613", "202600", "202699"} {
			t.Run(p.nome+"/rejeita "+comp, func(t *testing.T) {
				err := p.validar(comp)
				if !errors.Is(err, ErrValidacaoBpa) {
					t.Fatalf("esperava ErrValidacaoBpa, obteve %v", err)
				}
				if !strings.Contains(err.Error(), "mes entre 01 e 12") {
					t.Errorf("mensagem sem regra do mes: %v", err)
				}
			})
		}
	}
}

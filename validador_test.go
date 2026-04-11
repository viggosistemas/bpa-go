package bpa

import (
	"errors"
	"testing"
)

func TestValidarCabecalho(t *testing.T) {
	valido := CabecalhoEntrada{
		Competencia:      "202604",
		OrgaoResponsavel: "SMS",
		SiglaOrgao:       "SMS",
		CnpjCpf:          "12345678901234",
		OrgaoDestino:     "SES",
		TipoDestino:      "E",
		VersaoSistema:    "1.0",
	}

	t.Run("cabecalho valido", func(t *testing.T) {
		if err := validarCabecalho(valido); err != nil {
			t.Errorf("esperava nil, obteve %v", err)
		}
	})

	t.Run("competencia invalida", func(t *testing.T) {
		c := valido
		c.Competencia = "2026"
		err := validarCabecalho(c)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("tipo destino invalido", func(t *testing.T) {
		c := valido
		c.TipoDestino = "X"
		err := validarCabecalho(c)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})
}

func TestValidarRegistroBpaC(t *testing.T) {
	valido := RegistroBpaCEntrada{
		Cnes:         "1234567",
		Competencia:  "202604",
		Cbo:          "225142",
		Procedimento: "0301010064",
		Idade:        30,
		Quantidade:   1,
		Origem:       "BPA",
	}

	t.Run("registro valido", func(t *testing.T) {
		if err := validarRegistroBpaC(valido); err != nil {
			t.Errorf("esperava nil, obteve %v", err)
		}
	})

	t.Run("cnes maior que 7", func(t *testing.T) {
		r := valido
		r.Cnes = "12345678"
		err := validarRegistroBpaC(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("idade maior que 130", func(t *testing.T) {
		r := valido
		r.Idade = 131
		err := validarRegistroBpaC(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("quantidade zero", func(t *testing.T) {
		r := valido
		r.Quantidade = 0
		err := validarRegistroBpaC(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})
}

func TestValidarRegistroBpaI(t *testing.T) {
	valido := RegistroBpaIEntrada{
		Cnes:            "1234567",
		Competencia:     "202604",
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

	t.Run("registro valido", func(t *testing.T) {
		if err := validarRegistroBpaI(valido); err != nil {
			t.Errorf("esperava nil, obteve %v", err)
		}
	})

	t.Run("cbo obrigatorio", func(t *testing.T) {
		r := valido
		r.Cbo = ""
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("sexo invalido", func(t *testing.T) {
		r := valido
		r.Sexo = "X"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("data atendimento com letras", func(t *testing.T) {
		r := valido
		r.DataAtendimento = "2026040A"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("data atendimento tamanho invalido", func(t *testing.T) {
		r := valido
		r.DataAtendimento = "202604"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("data nascimento com letras", func(t *testing.T) {
		r := valido
		r.DataNascimento = "1996011X"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("data nascimento tamanho invalido", func(t *testing.T) {
		r := valido
		r.DataNascimento = "1996"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("cns profissional com letras", func(t *testing.T) {
		r := valido
		r.CnsProfissional = "12345678901234A"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("cns paciente com letras", func(t *testing.T) {
		r := valido
		r.CnsPaciente = "98765432101234B"
		err := validarRegistroBpaI(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})
}

func TestCampoNumericoComLetras(t *testing.T) {
	t.Run("cnes com letras no BPA-C", func(t *testing.T) {
		r := RegistroBpaCEntrada{
			Cnes:         "123ABC7",
			Competencia:  "202604",
			Cbo:          "225142",
			Procedimento: "0301010064",
			Idade:        30,
			Quantidade:   1,
			Origem:       "BPA",
		}
		err := validarRegistroBpaC(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("procedimento com letras no BPA-C", func(t *testing.T) {
		r := RegistroBpaCEntrada{
			Cnes:         "1234567",
			Competencia:  "202604",
			Cbo:          "225142",
			Procedimento: "030101ABC4",
			Idade:        30,
			Quantidade:   1,
			Origem:       "BPA",
		}
		err := validarRegistroBpaC(r)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})

	t.Run("cnpj_cpf com letras no cabecalho", func(t *testing.T) {
		c := CabecalhoEntrada{
			Competencia:      "202604",
			OrgaoResponsavel: "SMS",
			SiglaOrgao:       "SMS",
			CnpjCpf:          "1234567890ABCD",
			OrgaoDestino:     "SES",
			TipoDestino:      "E",
			VersaoSistema:    "1.0",
		}
		err := validarCabecalho(c)
		if !errors.Is(err, ErrValidacaoBpa) {
			t.Errorf("esperava ErrValidacaoBpa, obteve %v", err)
		}
	})
}

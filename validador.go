package bpa

import (
	"fmt"
	"regexp"
)

var reDigitos = regexp.MustCompile(`^\d+$`)

// validarCabecalho valida os campos da entrada do cabecalho.
func validarCabecalho(c CabecalhoEntrada) error {
	if len(c.Competencia) != 6 || !reDigitos.MatchString(c.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacao)
	}
	if c.TipoDestino != "E" && c.TipoDestino != "M" {
		return fmt.Errorf("%w: tipo_destino deve ser 'E' ou 'M'", ErrValidacao)
	}
	if len(c.OrgaoResponsavel) > 30 {
		return fmt.Errorf("%w: orgao_responsavel deve ter no maximo 30 caracteres", ErrValidacao)
	}
	if len(c.SiglaOrgao) > 6 {
		return fmt.Errorf("%w: sigla_orgao deve ter no maximo 6 caracteres", ErrValidacao)
	}
	if len(c.CnpjCpf) > 14 {
		return fmt.Errorf("%w: cnpj_cpf deve ter no maximo 14 digitos", ErrValidacao)
	}
	if len(c.OrgaoDestino) > 40 {
		return fmt.Errorf("%w: orgao_destino deve ter no maximo 40 caracteres", ErrValidacao)
	}
	if len(c.VersaoSistema) > 10 {
		return fmt.Errorf("%w: versao_sistema deve ter no maximo 10 caracteres", ErrValidacao)
	}
	return nil
}

// validarRegistroBpaC valida os campos de um registro BPA-C.
func validarRegistroBpaC(r RegistroBpaCEntrada) error {
	if len(r.Cnes) > 7 {
		return fmt.Errorf("%w: cnes deve ter no maximo 7 digitos", ErrValidacao)
	}
	if len(r.Competencia) != 6 || !reDigitos.MatchString(r.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacao)
	}
	if len(r.Cbo) > 6 {
		return fmt.Errorf("%w: cbo deve ter no maximo 6 caracteres", ErrValidacao)
	}
	if len(r.Procedimento) > 10 {
		return fmt.Errorf("%w: procedimento deve ter no maximo 10 digitos", ErrValidacao)
	}
	if r.Idade < 0 || r.Idade > 130 {
		return fmt.Errorf("%w: idade deve estar entre 0 e 130", ErrValidacao)
	}
	if r.Quantidade < 1 || r.Quantidade > 999999 {
		return fmt.Errorf("%w: quantidade deve estar entre 1 e 999999", ErrValidacao)
	}
	if len(r.Origem) > 3 {
		return fmt.Errorf("%w: origem deve ter no maximo 3 caracteres", ErrValidacao)
	}
	return nil
}

// validarRegistroBpaI valida os campos de um registro BPA-I.
func validarRegistroBpaI(r RegistroBpaIEntrada) error {
	if len(r.Cnes) > 7 {
		return fmt.Errorf("%w: cnes deve ter no maximo 7 digitos", ErrValidacao)
	}
	if len(r.Competencia) != 6 || !reDigitos.MatchString(r.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacao)
	}
	if r.Cbo == "" {
		return fmt.Errorf("%w: cbo e obrigatorio para BPA-I", ErrValidacao)
	}
	if len(r.Cbo) > 6 {
		return fmt.Errorf("%w: cbo deve ter no maximo 6 caracteres", ErrValidacao)
	}
	if r.Sexo != "M" && r.Sexo != "F" {
		return fmt.Errorf("%w: sexo deve ser 'M' ou 'F'", ErrValidacao)
	}
	if len(r.Procedimento) > 10 {
		return fmt.Errorf("%w: procedimento deve ter no maximo 10 digitos", ErrValidacao)
	}
	if r.Idade < 0 || r.Idade > 130 {
		return fmt.Errorf("%w: idade deve estar entre 0 e 130", ErrValidacao)
	}
	if r.Quantidade < 1 || r.Quantidade > 999999 {
		return fmt.Errorf("%w: quantidade deve estar entre 1 e 999999", ErrValidacao)
	}
	if len(r.Origem) > 3 {
		return fmt.Errorf("%w: origem deve ter no maximo 3 caracteres", ErrValidacao)
	}
	if len(r.CnsProfissional) > 15 {
		return fmt.Errorf("%w: cns_profissional deve ter no maximo 15 digitos", ErrValidacao)
	}
	if len(r.CnsPaciente) > 15 {
		return fmt.Errorf("%w: cns_paciente deve ter no maximo 15 digitos", ErrValidacao)
	}
	if len(r.MunicipioIbge) > 6 {
		return fmt.Errorf("%w: municipio_ibge deve ter no maximo 6 digitos", ErrValidacao)
	}
	if len(r.Cid) > 4 {
		return fmt.Errorf("%w: cid deve ter no maximo 4 caracteres", ErrValidacao)
	}
	if len(r.NomePaciente) > 30 {
		return fmt.Errorf("%w: nome_paciente deve ter no maximo 30 caracteres", ErrValidacao)
	}
	return nil
}

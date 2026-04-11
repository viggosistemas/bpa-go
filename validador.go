package bpa

import (
	"fmt"
	"regexp"
)

var reDigitos = regexp.MustCompile(`^\d+$`)

// apenasDigitos retorna true se a string contiver apenas digitos e nao for vazia.
func apenasDigitos(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

// validarCabecalho valida os campos da entrada do cabecalho.
func validarCabecalho(c CabecalhoEntrada) error {
	if len(c.Competencia) != 6 || !reDigitos.MatchString(c.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacaoBpa)
	}
	if c.TipoDestino != "E" && c.TipoDestino != "M" {
		return fmt.Errorf("%w: tipo_destino deve ser 'E' ou 'M'", ErrValidacaoBpa)
	}
	if len(c.OrgaoResponsavel) > 30 {
		return fmt.Errorf("%w: orgao_responsavel deve ter no maximo 30 caracteres", ErrValidacaoBpa)
	}
	if len(c.SiglaOrgao) > 6 {
		return fmt.Errorf("%w: sigla_orgao deve ter no maximo 6 caracteres", ErrValidacaoBpa)
	}
	if len(c.CnpjCpf) > 14 {
		return fmt.Errorf("%w: cnpj_cpf deve ter no maximo 14 digitos", ErrValidacaoBpa)
	}
	if c.CnpjCpf != "" && !apenasDigitos(c.CnpjCpf) {
		return fmt.Errorf("%w: cnpj_cpf deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(c.OrgaoDestino) > 40 {
		return fmt.Errorf("%w: orgao_destino deve ter no maximo 40 caracteres", ErrValidacaoBpa)
	}
	if len(c.VersaoSistema) > 10 {
		return fmt.Errorf("%w: versao_sistema deve ter no maximo 10 caracteres", ErrValidacaoBpa)
	}
	return nil
}

// validarRegistroBpaC valida os campos de um registro BPA-C.
func validarRegistroBpaC(r RegistroBpaCEntrada) error {
	if len(r.Cnes) > 7 {
		return fmt.Errorf("%w: cnes deve ter no maximo 7 digitos", ErrValidacaoBpa)
	}
	if r.Cnes != "" && !apenasDigitos(r.Cnes) {
		return fmt.Errorf("%w: cnes deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.Competencia) != 6 || !reDigitos.MatchString(r.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacaoBpa)
	}
	if len(r.Cbo) > 6 {
		return fmt.Errorf("%w: cbo deve ter no maximo 6 caracteres", ErrValidacaoBpa)
	}
	if len(r.Procedimento) > 10 {
		return fmt.Errorf("%w: procedimento deve ter no maximo 10 digitos", ErrValidacaoBpa)
	}
	if r.Procedimento != "" && !apenasDigitos(r.Procedimento) {
		return fmt.Errorf("%w: procedimento deve conter apenas digitos", ErrValidacaoBpa)
	}
	if r.Idade < 0 || r.Idade > 130 {
		return fmt.Errorf("%w: idade deve estar entre 0 e 130", ErrValidacaoBpa)
	}
	if r.Quantidade < 1 || r.Quantidade > 999999 {
		return fmt.Errorf("%w: quantidade deve estar entre 1 e 999999", ErrValidacaoBpa)
	}
	if len(r.Origem) > 3 {
		return fmt.Errorf("%w: origem deve ter no maximo 3 caracteres", ErrValidacaoBpa)
	}
	return nil
}

// validarRegistroBpaI valida os campos de um registro BPA-I.
func validarRegistroBpaI(r RegistroBpaIEntrada) error {
	if len(r.Cnes) > 7 {
		return fmt.Errorf("%w: cnes deve ter no maximo 7 digitos", ErrValidacaoBpa)
	}
	if r.Cnes != "" && !apenasDigitos(r.Cnes) {
		return fmt.Errorf("%w: cnes deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.Competencia) != 6 || !reDigitos.MatchString(r.Competencia) {
		return fmt.Errorf("%w: competencia deve ter 6 digitos numericos", ErrValidacaoBpa)
	}
	if r.Cbo == "" {
		return fmt.Errorf("%w: cbo e obrigatorio para BPA-I", ErrValidacaoBpa)
	}
	if len(r.Cbo) > 6 {
		return fmt.Errorf("%w: cbo deve ter no maximo 6 caracteres", ErrValidacaoBpa)
	}
	if r.Sexo != "M" && r.Sexo != "F" {
		return fmt.Errorf("%w: sexo deve ser 'M' ou 'F'", ErrValidacaoBpa)
	}
	if len(r.DataAtendimento) != 8 {
		return fmt.Errorf("%w: data_atendimento deve ter exatamente 8 digitos", ErrValidacaoBpa)
	}
	if !apenasDigitos(r.DataAtendimento) {
		return fmt.Errorf("%w: data_atendimento deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.Procedimento) > 10 {
		return fmt.Errorf("%w: procedimento deve ter no maximo 10 digitos", ErrValidacaoBpa)
	}
	if r.Procedimento != "" && !apenasDigitos(r.Procedimento) {
		return fmt.Errorf("%w: procedimento deve conter apenas digitos", ErrValidacaoBpa)
	}
	if r.Idade < 0 || r.Idade > 130 {
		return fmt.Errorf("%w: idade deve estar entre 0 e 130", ErrValidacaoBpa)
	}
	if r.Quantidade < 1 || r.Quantidade > 999999 {
		return fmt.Errorf("%w: quantidade deve estar entre 1 e 999999", ErrValidacaoBpa)
	}
	if len(r.Origem) > 3 {
		return fmt.Errorf("%w: origem deve ter no maximo 3 caracteres", ErrValidacaoBpa)
	}
	if len(r.CnsProfissional) > 15 {
		return fmt.Errorf("%w: cns_profissional deve ter no maximo 15 digitos", ErrValidacaoBpa)
	}
	if r.CnsProfissional != "" && !apenasDigitos(r.CnsProfissional) {
		return fmt.Errorf("%w: cns_profissional deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.CnsPaciente) > 15 {
		return fmt.Errorf("%w: cns_paciente deve ter no maximo 15 digitos", ErrValidacaoBpa)
	}
	if r.CnsPaciente != "" && !apenasDigitos(r.CnsPaciente) {
		return fmt.Errorf("%w: cns_paciente deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.DataNascimento) != 8 {
		return fmt.Errorf("%w: data_nascimento deve ter exatamente 8 digitos", ErrValidacaoBpa)
	}
	if !apenasDigitos(r.DataNascimento) {
		return fmt.Errorf("%w: data_nascimento deve conter apenas digitos", ErrValidacaoBpa)
	}
	if len(r.MunicipioIbge) > 6 {
		return fmt.Errorf("%w: municipio_ibge deve ter no maximo 6 digitos", ErrValidacaoBpa)
	}
	if len(r.Cid) > 4 {
		return fmt.Errorf("%w: cid deve ter no maximo 4 caracteres", ErrValidacaoBpa)
	}
	if len(r.NomePaciente) > 30 {
		return fmt.Errorf("%w: nome_paciente deve ter no maximo 30 caracteres", ErrValidacaoBpa)
	}
	return nil
}

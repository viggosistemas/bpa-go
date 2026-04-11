package bpa

import (
	"fmt"
	"strconv"
)

// tipoBpa indica o tipo de registros no construtor.
type tipoBpa int

const (
	tipoBpaNenhum tipoBpa = iota
	tipoBpaC
	tipoBpaI
)

const registrosPorFolha = 20

// Construtor e o builder principal para gerar arquivos BPA.
type Construtor struct {
	cabecalho   CabecalhoEntrada
	registrosC  []RegistroBpaCEntrada
	registrosI  []RegistroBpaIEntrada
	tipoAtual   tipoBpa
}

// NovoConstrutor cria um novo construtor com o cabecalho informado.
func NovoConstrutor(cabecalho CabecalhoEntrada) *Construtor {
	return &Construtor{
		cabecalho: cabecalho,
		tipoAtual: tipoBpaNenhum,
	}
}

// AdicionarRegistroBpaC adiciona um registro BPA-C ao construtor.
func (c *Construtor) AdicionarRegistroBpaC(reg RegistroBpaCEntrada) error {
	if c.tipoAtual == tipoBpaI {
		return ErrTipoMisturado
	}
	if err := validarRegistroBpaC(reg); err != nil {
		return err
	}
	c.tipoAtual = tipoBpaC
	c.registrosC = append(c.registrosC, reg)
	return nil
}

// AdicionarRegistroBpaI adiciona um registro BPA-I ao construtor.
func (c *Construtor) AdicionarRegistroBpaI(reg RegistroBpaIEntrada) error {
	if c.tipoAtual == tipoBpaC {
		return ErrTipoMisturado
	}
	if err := validarRegistroBpaI(reg); err != nil {
		return err
	}
	c.tipoAtual = tipoBpaI
	c.registrosI = append(c.registrosI, reg)
	return nil
}

// Construir gera o arquivo BPA completo como bytes.
func (c *Construtor) Construir() ([]byte, error) {
	if err := validarCabecalho(c.cabecalho); err != nil {
		return nil, err
	}

	totalRegistros := len(c.registrosC) + len(c.registrosI)
	if totalRegistros == 0 {
		return nil, fmt.Errorf("%w: nenhum registro adicionado", ErrValidacao)
	}

	totalFolhas := (totalRegistros-1)/registrosPorFolha + 1
	campoControle := calcularChecksum(c)

	var conteudo string

	// cabecalho com total_linhas = total de registros (sem contar cabecalho)
	conteudo += gerarCabecalho(c.cabecalho, totalRegistros, totalFolhas, campoControle)

	for i, reg := range c.registrosC {
		folha := i/registrosPorFolha + 1
		sequencia := i%registrosPorFolha + 1
		conteudo += gerarRegistroBpaC(reg, folha, sequencia)
	}

	for i, reg := range c.registrosI {
		folha := i/registrosPorFolha + 1
		sequencia := i%registrosPorFolha + 1
		conteudo += gerarRegistroBpaI(reg, folha, sequencia)
	}

	return []byte(conteudo), nil
}

// calcularChecksum calcula o campo de controle do cabecalho.
// soma = soma(procedimento_como_int64 + quantidade), campo_controle = (soma % 1111) + 1111
func calcularChecksum(c *Construtor) int {
	var soma int64
	for _, reg := range c.registrosC {
		procVal, _ := strconv.ParseInt(reg.Procedimento, 10, 64)
		soma += procVal + int64(reg.Quantidade)
	}
	for _, reg := range c.registrosI {
		procVal, _ := strconv.ParseInt(reg.Procedimento, 10, 64)
		soma += procVal + int64(reg.Quantidade)
	}
	return int(soma%1111) + 1111
}

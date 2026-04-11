package bpa

import (
	"fmt"
	"strings"
)

// gerarRegistroBpaI gera uma linha BPA-I (tipo 03) com exatamente 340 bytes + CRLF.
// Layout: 03(2) + cnes(7) + competencia(6) + cns_prof(15) + cbo(6) + data_atend(8) +
//
//	folha(3) + seq(2) + procedimento(10) + cns_pac(15) + sexo(1) + municipio(6) +
//	cid(4) + idade(3) + quantidade(6) + reservado(15) + origem(3) + nome_pac(30) +
//	data_nasc(8) + reservado(188) + CRLF(2)
func gerarRegistroBpaI(r RegistroBpaIEntrada, folha int, sequencia int) string {
	linha := "03"
	linha += formatarNumerico(r.Cnes, 7)
	linha += formatarNumerico(r.Competencia, 6)
	linha += formatarNumerico(r.CnsProfissional, 15)
	linha += formatarAlfanumerico(r.Cbo, 6)
	linha += formatarNumerico(r.DataAtendimento, 8)
	linha += formatarNumerico(fmt.Sprintf("%d", folha), 3)
	linha += formatarNumerico(fmt.Sprintf("%d", sequencia), 2)
	linha += formatarNumerico(r.Procedimento, 10)
	linha += formatarNumerico(r.CnsPaciente, 15)
	linha += formatarAlfanumerico(r.Sexo, 1)
	linha += formatarNumerico(r.MunicipioIbge, 6)
	linha += formatarAlfanumerico(r.Cid, 4)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Idade), 3)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Quantidade), 6)
	linha += strings.Repeat(" ", 15) // reservado
	linha += formatarAlfanumerico(r.Origem, 3)
	linha += formatarAlfanumerico(normalizarTexto(r.NomePaciente), 30)
	linha += formatarNumerico(r.DataNascimento, 8)
	linha += strings.Repeat(" ", 188) // reservado
	linha += "\r\n"
	return linha
}

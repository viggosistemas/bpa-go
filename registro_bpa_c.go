package bpa

import "fmt"

// gerarRegistroBpaC gera uma linha BPA-C (tipo 02) com exatamente 50 bytes + CRLF.
// Layout: 02(2) + cnes(7) + competencia(6) + cbo(6) + folha(3) + sequencia(2) +
//
//	procedimento(10) + idade(3) + quantidade(6) + origem(3) + CRLF(2)
func gerarRegistroBpaC(r RegistroBpaCEntrada, folha int, sequencia int) string {
	linha := "02"
	linha += formatarNumerico(r.Cnes, 7)
	linha += formatarNumerico(r.Competencia, 6)
	linha += formatarAlfanumerico(r.Cbo, 6)
	linha += formatarNumerico(fmt.Sprintf("%d", folha), 3)
	linha += formatarNumerico(fmt.Sprintf("%d", sequencia), 2)
	linha += formatarNumerico(r.Procedimento, 10)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Idade), 3)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Quantidade), 6)
	linha += formatarAlfanumerico(r.Origem, 3)
	linha += "\r\n"
	return linha
}

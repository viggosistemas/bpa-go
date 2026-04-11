package bpa

import "fmt"

// gerarCabecalho gera a linha de cabecalho (tipo 01) com exatamente 132 bytes + CRLF.
// Layout: 01(2) + #BPA#(5) + competencia(6) + total_linhas(6) + total_folhas(6) +
//
//	campo_controle(4) + orgao_responsavel(30) + sigla(6) + cnpj_cpf(14) +
//	orgao_destino(40) + tipo_destino(1) + versao(10) + CRLF(2)
func gerarCabecalho(c CabecalhoEntrada, totalLinhas int, totalFolhas int, campoControle int) string {
	linha := "01"
	linha += "#BPA#"
	linha += formatarNumerico(c.Competencia, 6)
	linha += formatarNumerico(fmt.Sprintf("%d", totalLinhas), 6)
	linha += formatarNumerico(fmt.Sprintf("%d", totalFolhas), 6)
	linha += formatarNumerico(fmt.Sprintf("%d", campoControle), 4)
	linha += formatarAlfanumerico(normalizarTexto(c.OrgaoResponsavel), 30)
	linha += formatarAlfanumerico(normalizarTexto(c.SiglaOrgao), 6)
	linha += formatarNumerico(c.CnpjCpf, 14)
	linha += formatarAlfanumerico(normalizarTexto(c.OrgaoDestino), 40)
	linha += formatarAlfanumerico(c.TipoDestino, 1)
	linha += formatarAlfanumerico(c.VersaoSistema, 10)
	linha += "\r\n"
	return linha
}

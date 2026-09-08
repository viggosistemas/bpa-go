package bpa

import (
	"fmt"
	"strings"
)

// gerarRegistroBpaI gera uma linha BPA-I (tipo 03) com exatamente 353 bytes + CRLF.
//
// Layout DATASUS vigente (355 bytes com CRLF), conferido byte a byte contra o
// arquivo do BPA-Magnetico de um municipio. Posicoes 1-indexadas:
//
//	03(2) 001-002 | cnes(7) 003-009 | competencia(6) 010-015 | cns_prof(15) 016-030 |
//	cbo(6) 031-036 | data_atend(8) 037-044 | folha(3) 045-047 | seq(2) 048-049 |
//	procedimento(10) 050-059 | cns_pac(15) 060-074 | sexo(1) 075 | municipio(6) 076-081 |
//	cid(4) 082-085 | idade(3) 086-088 | quantidade(6) 089-094 |
//	caten(2) 095-096 | naut(13) 097-109 (branco) | origem(3) 110-112 |
//	nome_pac(30) 113-142 | data_nasc(8) 143-150 |
//	raca(2) 151-152 | etnia(4) 153-156 | nacionalidade(3) 157-159 |
//	servico(3) 160-162 | classificacao(3) 163-165 | equipe_seq(8) 166-173 |
//	equipe_area(4) 174-177 | cnpj(14) 178-191 | cep(8) 192-199 | logradouro(3) 200-202 |
//	endereco(30) 203-232 | complemento(10) 233-242 | numero(5) 243-247 |
//	bairro(30) 248-277 | telefone(11) 278-288 | email(40) 289-328 | ine(10) 329-338 |
//	cpf_pac(11) 339-349 | reservado(1) 350 | sit_rua(1) 351 | reservado(2) 352-353 |
//	CRLF(2) 354-355
//
// A cauda 151-338 substitui o antigo bloco reservado(188): campos vazios viram
// brancos, mantendo compatibilidade com o comportamento das versoes <= v1.0.1.
// O bloco 339-353 (v1.2.0) e' aditivo: sem CPF sai em brancos e sit_rua sai "N".
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
	// CNS do paciente vazio sai em BRANCO (v1.2.0): e o que o arquivo do
	// municipio faz quando o paciente e identificado pelo CPF em 339-349.
	// Ate a v1.1.0 saia zero-padded ("000000000000000").
	linha += formatarNumericoOuBranco(r.CnsPaciente, 15)
	linha += formatarAlfanumerico(r.Sexo, 1)
	linha += formatarNumerico(r.MunicipioIbge, 6)
	linha += formatarAlfanumerico(r.Cid, 4)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Idade), 3)
	linha += formatarNumerico(fmt.Sprintf("%d", r.Quantidade), 6)
	linha += formatarNumericoOuBranco(r.CaraterAtendimento, 2) // 095-096 PRD_CATEN (v1.2.0)
	linha += strings.Repeat(" ", 13)                           // 097-109 PRD_NAUT (reservado)
	linha += formatarAlfanumerico(r.Origem, 3)
	linha += formatarAlfanumerico(normalizarTexto(r.NomePaciente), 30)
	linha += formatarNumerico(r.DataNascimento, 8)
	// Cauda 151-338 (antigo reservado). Numericos opcionais -> brancos quando vazio.
	linha += formatarNumericoOuBranco(r.RacaCor, 2)                   // 151-152 PRD_RACA
	linha += formatarNumericoOuBranco(r.Etnia, 4)                     // 153-156 PRD_ETNIA
	linha += formatarNumericoOuBranco(r.Nacionalidade, 3)             // 157-159 PRD_NAC
	linha += formatarNumericoOuBranco(r.Servico, 3)                   // 160-162 PRD_SRV
	linha += formatarNumericoOuBranco(r.Classificacao, 3)             // 163-165 PRD_CLF
	linha += formatarNumericoOuBranco(r.EquipeSeq, 8)                 // 166-173 PRD_EQUIPE_SEQ
	linha += formatarNumericoOuBranco(r.EquipeArea, 4)                // 174-177 PRD_EQUIPE_AREA
	linha += formatarNumericoOuBranco(r.Cnpj, 14)                     // 178-191 PRD_CNPJ
	linha += formatarNumericoOuBranco(r.Cep, 8)                       // 192-199 PRD_CEP
	linha += formatarNumericoOuBranco(r.CodLogradouro, 3)             // 200-202 PRD_LOGRAD
	linha += formatarAlfanumerico(normalizarTexto(r.Endereco), 30)    // 203-232 PRD_END
	linha += formatarAlfanumerico(normalizarTexto(r.Complemento), 10) // 233-242 PRD_COMPL
	linha += formatarAlfanumerico(normalizarTexto(r.Numero), 5)       // 243-247 PRD_NUM
	linha += formatarAlfanumerico(normalizarTexto(r.Bairro), 30)      // 248-277 PRD_BAIRRO
	linha += formatarAlfanumerico(r.Telefone, 11)                     // 278-288 PRD_DDTEL
	linha += formatarAlfanumerico(normalizarTexto(r.Email), 40)       // 289-328 PRD_EMAIL
	linha += formatarNumericoOuBranco(r.Ine, 10)                      // 329-338 PRD_INE
	// Bloco 339-353: CPF do paciente em campo proprio (nao no CNS), situacao de
	// rua e reservados em branco — layout medido no arquivo do municipio.
	linha += formatarNumericoOuBranco(r.CpfPaciente, 11) // 339-349 PRD_CPF_PCNTE
	linha += " "                                         // 350 reservado
	linha += formatarSituacaoRua(r.SituacaoRua)          // 351 PRD_SIT_RUA
	linha += "  "                                        // 352-353 reservado
	linha += "\r\n"
	return linha
}

// formatarSituacaoRua devolve o byte de PRD_SIT_RUA: "S" quando informado,
// senao "N" — o arquivo do municipio nunca deixa a posicao em branco.
func formatarSituacaoRua(valor string) string {
	if valor == "S" {
		return "S"
	}
	return "N"
}

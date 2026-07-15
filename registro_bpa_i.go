package bpa

import (
	"fmt"
	"strings"
)

// gerarRegistroBpaI gera uma linha BPA-I (tipo 03) com exatamente 340 bytes + CRLF.
//
// Layout oficial DATASUS (novo layout, 340 bytes). Posicoes 1-indexadas:
//
//	03(2) 001-002 | cnes(7) 003-009 | competencia(6) 010-015 | cns_prof(15) 016-030 |
//	cbo(6) 031-036 | data_atend(8) 037-044 | folha(3) 045-047 | seq(2) 048-049 |
//	procedimento(10) 050-059 | cns_pac(15) 060-074 | sexo(1) 075 | municipio(6) 076-081 |
//	cid(4) 082-085 | idade(3) 086-088 | quantidade(6) 089-094 |
//	reservado(15) 095-109 [PRD_CATEN(2)+PRD_NAUT(13)] | origem(3) 110-112 |
//	nome_pac(30) 113-142 | data_nasc(8) 143-150 |
//	raca(2) 151-152 | etnia(4) 153-156 | nacionalidade(3) 157-159 |
//	servico(3) 160-162 | classificacao(3) 163-165 | equipe_seq(8) 166-173 |
//	equipe_area(4) 174-177 | cnpj(14) 178-191 | cep(8) 192-199 | logradouro(3) 200-202 |
//	endereco(30) 203-232 | complemento(10) 233-242 | numero(5) 243-247 |
//	bairro(30) 248-277 | telefone(11) 278-288 | email(40) 289-328 | ine(10) 329-338 |
//	CRLF(2) 339-340
//
// A cauda 151-338 substitui o antigo bloco reservado(188): campos vazios viram
// brancos, mantendo compatibilidade com o comportamento das versoes <= v1.0.1.
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
	linha += strings.Repeat(" ", 15) // reservado: PRD_CATEN(2) + PRD_NAUT(13)
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
	linha += "\r\n"
	return linha
}

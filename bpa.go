// Package bpa implementa a geracao de arquivos BPA (Boletim de Producao Ambulatorial)
// no formato fixed-width do SUS/DATASUS.
package bpa

// CabecalhoEntrada contem os dados para gerar a linha de cabecalho (tipo 01).
type CabecalhoEntrada struct {
	Competencia      string // AAAAMM
	OrgaoResponsavel string // max 30 chars
	SiglaOrgao       string // max 6 chars
	CnpjCpf          string // max 14 digitos
	OrgaoDestino     string // max 40 chars
	TipoDestino      string // "E" ou "M"
	VersaoSistema    string // max 10 chars
}

// RegistroBpaCEntrada contem os dados para gerar uma linha BPA-C (tipo 02).
type RegistroBpaCEntrada struct {
	Cnes         string // 7 digitos
	Competencia  string // AAAAMM
	Cbo          string // 6 chars (opcional)
	Procedimento string // 10 digitos SIGTAP
	Idade        int    // 0-130
	Quantidade   int    // 1-999999
	Origem       string // max 3 chars
}

// RegistroBpaIEntrada contem os dados para gerar uma linha BPA-I (tipo 03).
type RegistroBpaIEntrada struct {
	Cnes            string // 7 digitos
	Competencia     string // AAAAMM
	CnsProfissional string // 15 digitos
	Cbo             string // 6 chars (obrigatorio)
	DataAtendimento string // AAAAMMDD
	Procedimento    string // 10 digitos SIGTAP
	CnsPaciente     string // 15 digitos
	Sexo            string // "M" ou "F"
	MunicipioIbge   string // 6 digitos
	Cid             string // 4 chars
	Idade           int    // 0-130
	Quantidade      int    // 1-999999
	Origem          string // max 3 chars
	NomePaciente    string // max 30 chars
	DataNascimento  string // AAAAMMDD
}

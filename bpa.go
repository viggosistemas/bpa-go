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

	// Cauda do layout BPA-I de 340 bytes (posicoes 151-338). Todos OPCIONAIS:
	// quando vazios, aquelas posicoes sao preenchidas com brancos, preservando
	// o comportamento do bloco reservado das versoes <= v1.0.1 (aditivo, sem
	// breaking change). Campos numericos aceitam apenas digitos quando
	// preenchidos; alfanumericos sao normalizados (maiusculas, sem acento) e
	// truncados ao tamanho. Enviar digitos ja limpos (sem mascara) nos campos
	// numericos e no Telefone.
	RacaCor       string // 2 digitos - PRD_RACA (01 Branca,02 Preta,03 Parda,04 Amarela,05 Indigena,99 Sem info)
	Etnia         string // 4 digitos - PRD_ETNIA (tabela de etnias indigenas); so quando RacaCor=05, senao branco
	Nacionalidade string // 3 digitos - PRD_NAC (010 Brasileiro,020 Naturalizado,030 Estrangeiro)
	Servico       string // 3 digitos - PRD_SRV (quando o procedimento exigir)
	Classificacao string // 3 digitos - PRD_CLF (quando o procedimento exigir)
	EquipeSeq     string // 8 digitos - PRD_EQUIPE_SEQ
	EquipeArea    string // 4 digitos - PRD_EQUIPE_AREA
	Cnpj          string // 14 digitos - PRD_CNPJ
	Cep           string // 8 digitos - PRD_CEP_PCNTE
	CodLogradouro string // 3 digitos - PRD_LOGRAD_PCNTE (codigo DNE do tipo de logradouro)
	Endereco      string // max 30 - PRD_END_PCNTE
	Complemento   string // max 10 - PRD_COMPL_PCNTE
	Numero        string // max 5 - PRD_NUM_PCNTE (default "SN" quando sem numero)
	Bairro        string // max 30 - PRD_BAIRRO_PCNTE
	Telefone      string // max 11 digitos - PRD_DDTEL_PCNTE (com DDD, sem mascara)
	Email         string // max 40 - PRD_EMAIL_PCNTE
	Ine           string // 10 digitos - PRD_INE (Identificacao Nacional de Equipes)
}

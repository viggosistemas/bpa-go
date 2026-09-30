# QA prático — competência com mês fora de 01–12 (2026-09-30)

Fluxo habitual (NovoConstrutor → AdicionarRegistroBpaC → Construir, pela API pública): `202601`, `202612` → 182 bytes, sem erro.

Tentar quebrar (cada valor no cabeçalho e no registro):

| entrada | resultado |
|---|---|
| `202613`, `202600` | `ErrValidacaoBpa`: "competencia deve ter mes entre 01 e 12" |
| `20261`, `2026123`, `""` | "competencia deve ter 6 digitos numericos" |
| `２０２６０４` (dígitos largos), `20260a`, ` 20260`, `-20261`, `202612\n` | "competencia deve ter 6 digitos numericos" |
| `000001`, `999912` | aceitos (ano não é validado; igual ao bpa-ts D-15; PLAUSIBLE da volta 1 no ledger) |

Achado aberto: nenhum bloqueante. Ano de competência sem faixa = decisão fora do escopo (#276).

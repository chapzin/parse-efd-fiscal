# Plano Executivo de Implementação Contínua

Data: 2026-05-06
Base: `docs/roadmap/SPED_MASTER_PLAN.md`

## Política de execução

- Trabalhar por marcos pequenos e commitáveis.
- Cada feature deve ter validação local antes de commit.
- Sempre que uma etapa ficar longa, criar auto-poke para retomar automaticamente.
- Não commitar dados fiscais reais, `.env`, XMLs privados, SPEDs reais, planilhas sensíveis ou credenciais.
- Para pesquisa externa, priorizar fontes oficiais do SPED/RFB, Guia Prático, Notas Técnicas, PVA e documentação de DF-e.

## Sequência ativa

### Marco 0: base, segurança e governança

1. GitHub Actions para `go test ./...`, `go build ./...` e `go test -race ./...`.
2. `.env.example` com chaves sem valores sensíveis.
3. `SECURITY.md` real com política de vulnerabilidade e dados fiscais.
4. `.gitignore` para `.env`, SPED/XML/planilhas geradas e bancos locais.
5. Fixtures sintéticas e aviso LGPD.
6. Atualizar README com validações e segurança.

### Marco 1: núcleo independente de banco

1. `pkg/efd/types` com tipos fiscais fortes.
2. `pkg/efd/diagnostic`.
3. `pkg/efd/parser` streaming inicial.
4. Golden tests sintéticos.

### Marco 2: catálogo versionado

1. Schema YAML inicial.
2. Catálogo mínimo dos registros atuais e bloco 9.
3. Gerador inicial `cmd/efdgen`.

## Auto-pokes

- Primeiro auto-poke: retomar Marco 0 se a sessão parar.
- Segundo auto-poke: iniciar Marco 1 após Marco 0.

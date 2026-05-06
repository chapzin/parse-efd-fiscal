# Política de Segurança

## Versões suportadas

Enquanto o projeto estiver antes de `v1.0.0`, somente a branch `master` recebe correções de segurança.

| Versão | Suporte |
| --- | --- |
| `master` | Sim |
| releases antigas | Não garantido |

## Como reportar vulnerabilidades

Abra um aviso privado pelo GitHub Security Advisories quando disponível, ou entre em contato com os mantenedores do repositório antes de publicar detalhes.

Inclua, quando possível:

- versão/commit afetado;
- passos mínimos de reprodução;
- impacto esperado;
- se o problema envolve dados fiscais, envie apenas amostras anonimizadas ou sintéticas.

## Dados sensíveis

Este projeto processa SPED Fiscal, XMLs de documentos fiscais, inventário, dados de clientes/fornecedores e credenciais de banco. Esses dados podem conter informações protegidas por LGPD, sigilo fiscal ou contrato.

Nunca commite:

- `.env` real;
- credenciais de banco, tokens, certificados ou chaves privadas;
- SPED/XML fiscal real;
- planilhas de inventário reais;
- dumps de banco;
- logs com SQL ou valores fiscais identificáveis.

Use fixtures sintéticas em testes e issues públicos.

## Escopo de segurança técnico

Áreas críticas:

- parsing de arquivos grandes ou malformados;
- exposição acidental de valores em logs;
- importação em banco de produção;
- comandos destrutivos de schema;
- dependências Go com vulnerabilidades conhecidas;
- geração de relatórios contendo dados pessoais/fiscais.

## Boas práticas para contribuidores

- Rode `go test ./...` e `go build ./...` antes de enviar alterações.
- Para mudanças de concorrência, rode `go test -race ./...`.
- Não adicione dependências sem necessidade clara.
- Não use dados reais em testes.
- Documente qualquer nova leitura/escrita de arquivos sensíveis.

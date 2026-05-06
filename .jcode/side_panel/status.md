# Project Status Panel

## Current goal

Jcode init swarm analysis completed for `parse-efd-fiscal`.

## Detected stack

- Go CLI
- Go modules
- MySQL through GORM v1
- Docker Compose for MySQL/phpMyAdmin

## Swarm status

- architect: completed
- qa: completed
- documenter: report received
- tooling-security: report received
- Final synthesis: `.jcode/init/SWARM_ANALYSIS_REPORT.md`

## Validation

- `go test ./...` currently fails in `tools/SpedConvert_test.go::TestFloatToString`.
- `go build` is documented in `README.md` and should be run after tests pass.
- DB/runtime commands require a reviewed `.env` and should use disposable or non-sensitive data.

## Evidence-backed commands

```bash
go test ./...
go build
docker-compose -f docker/docker-compose.yml --env-file .env up -d
parse-efd-fiscal -schema
parse-efd-fiscal -importar-sped
parse-efd-fiscal -importar-xml
parse-efd-fiscal -inventario -anoInicial=2012 -anoFinal=2016
parse-efd-fiscal -excel
```

## Risks

- `.env`, SPED/XML fiscal files, generated spreadsheets, and DB contents are sensitive.
- Schema and inventory workflows can be destructive.
- `db.LogMode(true)` can expose SQL/log data.
- README inventory flag docs are stale versus `main.go`.
- No CI config was found despite README Travis badge.

## Open questions

See `.jcode/INIT_QUESTIONS.md` and `.jcode/init/SWARM_ANALYSIS_REPORT.md`.

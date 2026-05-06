# Skills Plan

## Recommended initial skills

- karpathy-guidelines
- optimization
- llmwiki-memory

## Project-specific routing

- Go CLI/database work: load coding and testing skills only for the task at hand; validate with `go test ./...` and `go build` after the current `tools` test mismatch is fixed.
- Fiscal parser/import work: prefer small fixture-based tests and avoid using real SPED/XML fiscal data unless explicitly marked safe.
- Documentation work: use `llmwiki-memory` only for durable project decisions and provenance, not for `.env` values, credentials, fiscal records, XML payloads, spreadsheets, or generated inventory outputs.
- Performance/concurrency work: use optimization skills around `pkg/worker`, recursive readers, and inventory processing only with measurable fixtures.

## Evidence-backed project context

- Stack: Go modules, MySQL, GORM v1, Docker Compose for local DB/phpMyAdmin.
- Current validation issue: `go test ./...` fails in `tools/SpedConvert_test.go` because expected `FloatToString(3.5)` output differs from implementation.
- High-sensitivity paths/data: `.env`, `SPEDS_PATH` contents, SPED/XML fiscal files, generated spreadsheets, and database contents.

## Notes

- Built-in skills are available offline.
- Project-local skills can override built-ins under `.jcode/skills/<name>/SKILL.md`.
- Do not inject every full skill by default. Route skills by task.
- Do not sync secrets, credentials, `.env` values, provider tokens, deployment secrets, database credentials, fiscal records, XML payloads, or spreadsheet contents into wiki memory.

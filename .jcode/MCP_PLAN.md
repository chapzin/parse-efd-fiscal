# MCP Plan

MCP setup is intentionally review-first. This init command does not download or install MCP servers automatically.

## Evidence-backed project context

- Package manager: Go modules via `go.mod` and `go.sum`.
- Runtime systems: MySQL through GORM v1, configured from `.env`.
- Local automation: `docker/docker-compose.yml` starts MySQL and phpMyAdmin using `.env` values.
- CI evidence: README has a Travis badge, but no `.travis.yml` or `.github` workflow was found.

## Candidate server categories

- Filesystem/code search: usually already covered by native Jcode tools. Keep this as the default.
- GitHub/GitLab: useful for issues/PRs if this repo uses remote review workflows; requires token review and must not run in CI without explicit approval.
- Database: high risk. Disable by default unless connected only to a disposable or read-only database with documented credentials boundary. Never expose `.env` values.
- Browser/Playwright: low priority for this CLI project, but can help inspect generated docs or reports if needed.
- Docs/search: useful for Go, GORM v1, MySQL, Docker, or SPED reference research if network access is approved.

## Required review steps before enabling MCP

1. Identify the exact external system and why native Jcode tools are insufficient.
2. Document credentials required, without storing values.
3. Define allowed read/write scope, especially for database and GitHub tools.
4. Prefer local/offline servers and least privilege.
5. Add reviewed server definitions to `.jcode/mcp.json` only after approval.
6. Validate with `jcode` after reviewing permissions.

## Explicit boundaries

- Do not configure MCP to read or persist `.env` values, database credentials, provider tokens, private keys, fiscal XML/SPED payloads, generated spreadsheets, or production database contents.
- Do not auto-install remote MCP servers without human review.

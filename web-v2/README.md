# web-v2

New Fern Platform SPA per **RFC-004** (Frontend Modernization). Lives
alongside the legacy `web/` directory during the strangler migration
described in [docs/specs/frontend-modernization/](../docs/specs/frontend-modernization/).

At Phase 5 cutover, the legacy `web/` is deleted and this directory is
renamed to `web/`.

## Stack

React 18 · TypeScript (strict) · Vite 5 · TanStack Query/Router ·
GraphQL Codegen · Tailwind + shadcn/ui · Vitest · Playwright · size-limit.

## Development

```bash
pnpm install              # first time only
make web-dev              # or: pnpm dev
```

Vite runs on `:5173` and proxies `/api`, `/graphql`, `/auth` to the Go
server on `:8080`.

## Build

```bash
make web-build            # typecheck, lint, test, vite build → dist/
```

`dist/` is copied into `internal/web/dist/` so the Go binary embeds it
via `//go:embed`.

## Tests

```bash
pnpm test                 # vitest watch
pnpm test:run             # vitest single run
pnpm test:coverage        # with coverage
pnpm playwright           # E2E (needs deployed stack)
```

## Bundle budgets

Enforced by `size-limit`:

- Initial bundle ≤ 150 KB gzipped
- Total app ≤ 600 KB gzipped

See [docs/specs/frontend-modernization/design.md](../docs/specs/frontend-modernization/design.md)
for the full design contract.

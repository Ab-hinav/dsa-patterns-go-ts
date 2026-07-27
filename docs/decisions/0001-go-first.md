# Decision 0001: Use Go First

## Status

Accepted.

## Context

The learning goal is to build interview fluency without maintaining every solution in multiple languages. Go is the default interview language, while TypeScript remains useful for comparison and JavaScript-oriented roles.

## Decision

- Write and test the primary implementation in Go.
- Add a TypeScript implementation only when it teaches a meaningful language or API difference.
- Keep explanations focused on the reusable algorithmic pattern rather than syntax.

## Consequences

- The repository stays small and easier to revise.
- Go testing provides a consistent verification path.
- TypeScript coverage will intentionally be incomplete until a comparison is valuable.

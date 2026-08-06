# DSA Patterns in Go and TypeScript

[![Go tests](https://github.com/Ab-hinav/dsa-patterns-go-ts/actions/workflows/test.yml/badge.svg)](https://github.com/Ab-hinav/dsa-patterns-go-ts/actions/workflows/test.yml)

A curated coding-interview practice repository focused on reusable problem-solving patterns, readable implementations, tests, complexity analysis, and mistake-driven revision.

## Status

**Practice evidence.** This repository is not yet a pinned resume project.

## Current Coverage

### Arrays and Hashing

| Problem | Pattern | Go | TypeScript | Review status |
|---|---|---:|---:|---|
| [Two Sum](https://leetcode.com/problems/two-sum/) | Complement lookup | Yes | Planned | Reviewed |
| [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Hash-set membership | Yes | Planned | Solved clean |
| [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Canonical frequency signature | Yes | Planned | Solved clean |
| [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Frequency map and bucket ordering | Yes | Planned | Solved clean |

### Two Pointers

| Problem | Pattern | Go | TypeScript | Review status |
|---|---|---:|---:|---|
| [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Opposite-direction pointers | Yes | Planned | Solved clean |
| [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Ordered convergence | Yes | Planned | Solved clean |

Only links and original notes are included. Problem statements remain on their original platforms.

## Repository Structure

```text
arrays-hashing/
  contains_duplicate.go
  contains_duplicate_test.go
  group_anagrams.go
  group_anagrams_test.go
  top_k_frequent.go
  top_k_frequent_test.go
  two_sum.go
  two_sum_test.go
two-pointers/
  two_sum_sorted.go
  two_sum_sorted_test.go
  valid_palindrome.go
  valid_palindrome_test.go
docs/
  decisions/
  patterns/
```

TypeScript implementations will be added selectively when comparing language tradeoffs is useful.

## Run Locally

Requirements:

- Go 1.24 or later

Run all tests:

```bash
go test ./...
```

Run the arrays-and-hashing tests with detailed output:

```bash
go test -v ./arrays-hashing
```

Run the two-pointers tests with detailed output:

```bash
go test -v ./two-pointers
```

## Learning Method

Each problem follows the same review loop:

1. Clarify constraints and edge cases.
2. Explain a brute-force approach.
3. Identify repeated work and the reusable pattern.
4. Implement the optimized approach.
5. State time and space complexity.
6. Add targeted tests.
7. Record mistakes and schedule weak problems for revision.

## Evidence

- Tests cover empty inputs, duplicate values, missing results, repeated words, order-independent anagram groups, frequency-bucket selection, pointer skipping, mixed case, punctuation, digits, sorted boundaries, and distinct-index handling.
- GitHub Actions runs the complete Go test suite on each push and pull request.
- Pattern notes explain when the technique applies, why it works, and its limitations.

## Roadmap

- Continue the two-pointers pattern drill with greedy pointer movement.
- Add the sliding-window pattern group after two pointers.
- Add selected TypeScript comparisons.
- Add benchmark exercises only when performance differences are meaningful.
- Complete a 25-, 45-, and 60-minute mock-interview cycle.

## Attribution

Problem names and links refer to their respective original platforms. The explanations, implementations, tests, and learning notes in this repository are original practice material.

## License

No reuse license is granted yet. See [LICENSE](LICENSE).

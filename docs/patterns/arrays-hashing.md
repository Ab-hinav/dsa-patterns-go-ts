# Arrays and Hashing

## Complement Lookup

Use a hash map when a current value tells you exactly which earlier value would complete the answer.

For Two Sum, store each value with its index. Before storing the current value, check whether `target - currentValue` has already been seen.

- Time: `O(n)`
- Space: `O(n)`
- Important edge case: equal values must come from different indices.

## Hash-Set Membership

Use a set when the question only asks whether a value has already appeared.

For Contains Duplicate, scan from left to right. Return immediately when a value is already in the set.

- Time: `O(n)` expected
- Space: `O(n)`
- Go representation: `map[int]struct{}` models membership without storing a second boolean value.

## Canonical Frequency Signature

Use a canonical key when multiple inputs should belong to the same group.

For lowercase English anagrams, a `[26]int` letter-frequency array is comparable in Go and can be used directly as a map key.

- Time: `O(n * k)`, where `k` is the average word length
- Space: `O(n * k)` for the output and grouped strings
- Limitation: the fixed 26-entry signature assumes lowercase English letters.

## Review Questions

- What constraints make a hash map appropriate?
- Is the key a value, complement, frequency, or canonical representation?
- Does the solution depend on expected constant-time hashing?
- Which input assumptions should be validated or documented?

# Two Pointers

## Opposite-Direction Pointers

Use opposite-direction pointers when a sequence must be checked or combined
from both ends and each comparison can safely shrink the unprocessed region.

For Valid Palindrome, place one pointer at each end of the string. Skip
non-alphanumeric ASCII bytes. When both pointers reference meaningful
characters, compare them without case sensitivity and then move both inward.

The loop invariant is:

> Every meaningful character pair outside the pointer boundaries has already
> matched.

This invariant makes an early mismatch decisive. If the pointers meet or cross
without a mismatch, every required pair has matched.

- Time: `O(n)`
- Extra space: `O(1)`
- Movement proof: the left pointer moves only right and the right pointer moves
  only left, so their combined movement is at most `2n`.
- Limitation: byte indexing and the helper intentionally follow the problem's
  printable ASCII constraint rather than general Unicode text.

## Ordered Convergence

Use ordered convergence when the input is sorted and comparing the two
boundary values reveals which boundary cannot participate in a valid answer.

For Two Sum II, place one pointer at each end of the sorted slice. If their sum
is below the target, the left value cannot work: even pairing it with the
largest available value was too small. Move the left pointer right. If the sum
is above the target, the right value cannot work: even pairing it with the
smallest available value was too large. Move the right pointer left.

The loop invariant is:

> If a valid pair remains, both of its indices are inside the current pointer
> boundaries.

Each comparison safely discards one impossible boundary while preserving any
possible answer. Using `left < right` also guarantees that one index is never
paired with itself.

- Time: `O(n)`
- Extra space: `O(1)`
- Movement proof: each pointer moves in only one direction, so the search
  region can shrink at most `n - 1` times.
- Constraint dependency: the movement proof relies on the input being sorted.
- Review note: return `nil` instead of invalid one-indexed positions if the
  problem's exactly-one-solution guarantee is absent.

## Review Questions

- What property makes examining both ends useful?
- What part of the input is already verified before each iteration?
- Which pointer should move when one side is irrelevant?
- What ordering property proves that a discarded boundary cannot work?
- Does every iteration make progress toward termination?
- Are the input characters bytes, Unicode code points, or another unit?

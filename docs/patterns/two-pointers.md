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

## Review Questions

- What property makes examining both ends useful?
- What part of the input is already verified before each iteration?
- Which pointer should move when one side is irrelevant?
- Does every iteration make progress toward termination?
- Are the input characters bytes, Unicode code points, or another unit?

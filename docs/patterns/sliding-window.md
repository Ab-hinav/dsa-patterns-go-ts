# Sliding Window

## Positive-Sum Threshold

For [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/),
keep a contiguous window and its running sum. Expand the right edge once per
element. Whenever the sum reaches the target, record the current length and
shrink from the left until the sum is below the target again.

The invariant is:

> Before each right-edge expansion, the current window is contiguous, its sum
> is known, and no shorter qualifying window ending at the previous right edge
> remains unchecked.

All values must be positive. Removing a left value then strictly decreases the
sum; once the sum falls below the target, further shrinking cannot make that
same window qualify. With negative values, that argument fails.

- Time: `O(n)` because each edge advances at most `n` times.
- Extra space: `O(1)`.
- No solution: initialize the minimum to `n + 1` and return `0` if it never
  changes.
- Review note: if the right edge is exclusive, the window length is
  `right - left`, not `right - left + 1`. The implementation uses an inclusive
  `right` from `range`, so its length is `right - left + 1`.

## Review Questions

- What condition causes expansion, and what condition causes shrinking?
- Why is shrinking safe only with positive values?
- Why can both nested loops still take only linear total time?

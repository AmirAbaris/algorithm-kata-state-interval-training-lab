### Track 09 — Merge Intervals

**Tier C** · Prerequisite tracks: 05–08

---

### Problem: Merge Intervals

Given a list of intervals, merge all overlapping or touching intervals.

Definition:
Two intervals overlap if:
`a.Start <= b.End AND b.Start <= a.End`

Two intervals are touching if one ends exactly where the other starts (e.g. `[1,3]` and `[3,5]`). Touching intervals must be merged into one.

Input:
- `intervals`: a slice of `Interval` values
- Each interval has `Start` and `End` where `Start <= End`
- Input order is not guaranteed

Output:
- A slice of merged intervals, sorted by `Start` ascending
- No two output intervals overlap or touch
- Return an empty slice for empty input

Constraints:
- `0 <= len(intervals) <= 10_000`
- `-1_000_000 <= Start <= End <= 1_000_000`

Examples:

**Example 1**
Input: `[[1,3], [2,6], [8,10], [15,18]]`
Output: `[[1,6], [8,10], [15,18]]`
Explanation: `[1,3]` and `[2,6]` overlap, so they merge to `[1,6]`. The others do not overlap with anything.

**Example 2**
Input: `[[1,4], [4,5]]`
Output: `[[1,5]]`
Explanation: The intervals touch at `4`, so they merge into one interval.

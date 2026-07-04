### Track 10 — Insert Interval

**Tier C** · Prerequisite tracks: 05–09

---

### Problem: Insert Interval

Given a list of non-overlapping intervals sorted by `Start`, insert a new interval and merge any overlaps or touching intervals.

The existing intervals do not overlap each other and are already sorted by `Start`.
After insertion, return a sorted list where no intervals overlap or touch.

Input:
- `intervals`: sorted, non-overlapping intervals (`Start` ascending)
- `newInterval`: the interval to insert
- `Start <= End` for all intervals

Output:
- Updated interval list, sorted by `Start`
- No overlaps or touching intervals in the result

Constraints:
- `0 <= len(intervals) <= 10_000`
- `-1_000_000 <= Start <= End <= 1_000_000`

Examples:

**Example 1**
Input:
- intervals: `[[1,3], [6,9]]`
- newInterval: `[2,5]`

Output: `[[1,5], [6,9]]`
Explanation: `[2,5]` overlaps `[1,3]`, merging into `[1,5]`. `[6,9]` is unchanged.

**Example 2**
Input:
- intervals: `[[1,2], [3,5], [6,7], [8,10], [12,16]]`
- newInterval: `[4,8]`

Output: `[[1,2], [3,10], [12,16]]`
Explanation: `[4,8]` overlaps `[3,5]`, touches `[6,7]`, and overlaps `[8,10]`, merging into `[3,10]`.

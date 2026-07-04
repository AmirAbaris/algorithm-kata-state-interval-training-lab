### Track 11 — Subtract Intervals

**Tier C** · Prerequisite tracks: 05–08 (especially 08)

---

### Problem: Subtract Intervals

Given a set of base intervals and a set of intervals to remove, return the portions of the base intervals that are not covered by any remove interval.

Definition:
A point `t` is covered by an interval if `Start <= t <= End`.
A base interval is reduced by cutting out every point covered by any remove interval.

Input:
- `base`: intervals representing occupied/covered regions
- `toRemove`: intervals representing regions to cut away
- Intervals in `base` may overlap each other (treat their union as the region to keep)
- Intervals in `toRemove` may overlap each other (treat their union as the region to cut)
- Order is not guaranteed for either slice

Output:
- A slice of remaining intervals after subtraction
- Sorted by `Start` ascending
- No overlaps or touching intervals in the output (merge adjacent fragments)
- Return an empty slice if nothing remains or if `base` is empty

Constraints:
- `0 <= len(base), len(toRemove) <= 10_000`
- `-1_000_000 <= Start <= End <= 1_000_000`

Examples:

**Example 1**
Input:
- base: `[[1,10]]`
- toRemove: `[[3,5], [7,8]]`

Output: `[[1,2], [6,6], [9,10]]`
Explanation: `[3,5]` and `[7,8]` are cut out of `[1,10]`, leaving three fragments.

**Example 2**
Input:
- base: `[[1,5], [8,12]]`
- toRemove: `[[0,20]]`

Output: `[]`
Explanation: The remove interval fully covers both base intervals, so nothing remains.

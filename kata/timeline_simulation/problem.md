### Track 12 — Timeline State Simulation

**Tier C** · Prerequisite tracks: 03, 05–09

---

### Problem: Timeline State Simulation

Simulate how a numeric state changes over time given a list of events.

Each event applies a `Delta` to the state for every integer time unit in `[Start, End]` (inclusive on both ends).
Events may overlap; overlapping regions stack their deltas.

The timeline starts at state `0` everywhere. Only time points affected by at least one event appear in the output.

Input:
- `events`: a slice of `Event` values, each with `Start`, `End`, and `Delta`
- `Start <= End` for every event
- Order is not guaranteed

Output:
- A slice of `StateSegment` values: `{Interval, State}`
- Sorted by `Interval.Start` ascending
- Segments are maximal: state is constant within each segment and changes at segment boundaries
- Adjacent segments with the same state must be merged into one
- Return an empty slice if `events` is empty

Constraints:
- `0 <= len(events) <= 10_000`
- `-1_000_000 <= Start <= End <= 1_000_000`
- `-1_000_000 <= Delta <= 1_000_000`

Examples:

**Example 1**
Input: `[{Start: 1, End: 3, Delta: 2}, {Start: 2, End: 4, Delta: 3}]`

Output:
- `[1,1]` → state `2`
- `[2,3]` → state `5` (2+3)
- `[4,4]` → state `3`

Explanation: At time 1 only the first event applies (+2). Times 2–3 have both events (+2+3=5). Time 4 has only the second event (+3).

**Example 2**
Input: `[{Start: 5, End: 5, Delta: 10}, {Start: 5, End: 7, Delta: -3}]`

Output:
- `[5,5]` → state `7` (10-3)
- `[6,7]` → state `-3`

Explanation: At time 5 both events apply. Times 6–7 only the second event applies.

### Track 07 — Merge Two Intervals

**Tier B** · Merge intervals with only two items — learn the running merge state

---

## Plain English

Given two intervals, merge them if they overlap or touch. Otherwise return both, sorted.

## DATA

- `a`, `b`: two intervals

## OUTPUT

- 1 interval if they overlap/touch
- 2 intervals sorted by `Start` if they don't

## STATE (when merging)

| Variable | Meaning |
|----------|---------|
| `merged.Start` | min of starts |
| `merged.End` | max of ends |

## PROCESS

1. If `Overlaps(a, b)` (track 05): return one interval `[min starts, max ends]`
2. Else: return `[a, b]` sorted by Start

## Examples

- `[1,3]` + `[2,5]` → `[1,5]`
- `[1,3]` + `[4,5]` → `[1,3], [4,5]`
- `[1,4]` + `[4,5]` → `[1,5]` (touch)

## Connection

`interval_merge` loop body is this decision + a **running** merged interval carried to the next item.

### Track 06 — Sort Intervals

**Tier B** · Ordering is not optional for merge algorithms

---

## Plain English

Return the same intervals sorted by `Start` ascending. When `Start` ties, smaller `End` first.

Do **not** merge. Only sort.

## DATA

- `intervals`: slice, may be empty, order not guaranteed

## OUTPUT

- New sorted slice (you may sort in place or copy — tests compare values)
- Empty → empty

## STATE

Typical sort: compare adjacent items while ordering. Use `sort.Slice` in Go if you want — focus on **what** you're sorting by, not reinventing quicksort.

## PROCESS

Sort by `(Start, End)` lexicographically.

## Examples

Input: `[[8,10], [1,3], [2,6]]`  
Output: `[[1,3], [2,6], [8,10]]`

Input: `[[1,5], [1,2]]`  
Output: `[[1,2], [1,5]]`

## Connection

`interval_merge` step 1 is always sort. Isolate it here so merge tests don't fail for two reasons at once.

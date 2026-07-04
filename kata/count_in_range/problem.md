### Track 04 — Count In Range

**Tier A** · Data vs process — input is static, loop asks a question per item

---

## Plain English

Count how many numbers in `nums` fall in the closed range `[low, high]` (inclusive).

## DATA

- `nums`, `low`, `high`
- `low <= high` always
- `nums` may be empty

## OUTPUT

- Count as `int`
- Empty `nums` → `0`

## STATE

| Variable | Meaning | Start |
|----------|---------|-------|
| `count` | Matches so far | `0` |
| `x` | Current number | loop variable |

## PROCESS

For each `x` in `nums`: if `low <= x <= high`, increment `count`.

**Data** = the slice. **Process** = the question "is x in range?"

## State table

`nums = [1, 5, 3, 8, 2]`, `low=2`, `high=5`

| Step | x | in range? | count after |
|------|---|-----------|-------------|
| 1 | 1 | no | 0 |
| 2 | 5 | | |
| 3 | 3 | | |
| 4 | 8 | | |
| 5 | 2 | | |

Answer: `3`

## Examples

- All outside → `0`
- All inside → `len(nums)`
- Single point range `low=high=4` counts only 4s

## Connection

Later, `find_available_slots` filters intervals by **duration** — same "does this item pass a condition?" pattern.

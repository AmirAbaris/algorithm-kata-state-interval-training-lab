### Track 02 — Running Max

**Tier A** · First real "state while iterating"

---

## Plain English

Walk left to right. At each position, record the **largest value seen so far** (including the current item).

## DATA

- `nums`: `[]int`, may be empty

## OUTPUT

- `[]int` same length as `nums`
- `result[i]` = max of `nums[0..i]`
- Empty input → empty slice (not `nil` required, but length 0)

## STATE

| Variable | Meaning | Start |
|----------|---------|-------|
| `best` | Max seen so far | first element (or handle empty) |
| `i` | Current index / item | loop |

## PROCESS

For each index `i`, update `best` if `nums[i]` is larger, append `best` to output.

## State table

Input: `[3, 1, 4, 2]`

| Step | i | nums[i] | best before | best after | append |
|------|---|---------|-------------|------------|--------|
| 1 | 0 | 3 | 3 | 3 | 3 |
| 2 | 1 | 1 | | | |
| 3 | 2 | 4 | | | |
| 4 | 3 | 2 | | | |

Output: `[3, 3, 4, 4]`

## Examples

- `[5]` → `[5]`
- `[2, 2, 2]` → `[2, 2, 2]`
- `[1, 5, 3, 5, 2]` → `[1, 5, 5, 5, 5]`

## Connection

`best` is **running state**. Timeline simulation later tracks running **state sum** the same way — but on a time axis instead of an array index.

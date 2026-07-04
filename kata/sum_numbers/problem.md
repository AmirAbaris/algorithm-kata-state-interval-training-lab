### Track 01 — Sum Numbers

**Tier A** · Loops · No running state beyond one accumulator

---

## Plain English

Add every number in a slice. That's it.

## DATA (what you are given)

- `nums`: `[]int` — may be empty
- Order does not matter

## OUTPUT (what you return)

- One `int`: the total
- Empty input → `0`

## STATE (what to remember while looping)

| Variable | Meaning | Start value |
|----------|---------|-------------|
| `total` | Sum so far | `0` |

## PROCESS (one step)

For each number in `nums`, add it to `total`. Return `total`.

## State table (fill this before coding)

Input: `[3, 1, 4]`

| Step | Current item | `total` before | Action | `total` after |
|------|--------------|----------------|--------|---------------|
| 0 | — | 0 | — | 0 |
| 1 | 3 | | | |
| 2 | 1 | | | |
| 3 | 4 | | | |

## Examples

**Empty:** `[]` → `0`

**Single:** `[7]` → `7`

**Mixed signs:** `[10, -3, 2]` → `9`

## Connection

Next track (`running_max`) uses the same loop shape but **stores a slice** of state history instead of one final number.

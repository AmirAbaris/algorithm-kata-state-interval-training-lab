### Track 03 — Apply Deltas (Trace)

**Tier A** · Sequential state transformations · Direct prep for timeline simulation

---

## Plain English

You start with a number `start`. Apply each delta in order. After **every** step, record the new value.

This is "bank balance after each transaction" or "capacity after each booking event."

## DATA

- `start`: initial `int`
- `deltas`: `[]int` — changes to apply in order (may be empty)

## OUTPUT

- `[]int` of length `len(deltas)`
- `result[i]` = state **after** applying `deltas[i]`
- Empty `deltas` → empty slice

## STATE

| Variable | Meaning | Start |
|----------|---------|-------|
| `current` | Value right now | `start` |

## PROCESS

For each delta: `current += delta`, append `current`.

## State table

`start = 10`, `deltas = [+3, -5, +2]`

| Step | delta | current before | current after | append |
|------|-------|----------------|---------------|--------|
| 1 | +3 | 10 | | |
| 2 | -5 | | | |
| 3 | +2 | | | |

Output: `[13, 8, 10]`

## Examples

- `start=0`, `[]` → `[]`
- `start=0`, `[5]` → `[5]`
- `start=100`, `[-30, -30]` → `[70, 40]`

## Connection

`timeline_simulation` does this on a **time line** where many events overlap. Master the 1D "current number" version first.

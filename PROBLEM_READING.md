# Problem Reading Worksheet

Copy this block into your notes (or under the examples in each `problem.md`) **before** coding.

```text
KATA: _______________________

1. DATA (what I am given)
   - Type:
   - Can it be empty?
   - Is order guaranteed?

2. OUTPUT (what I return)
   - Type:
   - Sorted? Unique? Merged?

3. STATE (what I remember during the loop)
   - Variable name:
   - Initial value:
   - When does it change?

4. CURRENT ITEM (what one iteration looks at)
   - Loop variable:
   - What question do I ask about it?

5. SMALLEST FIRST TEST
   - Test name:
   - Why start here?

6. ONE-LINE PROCESS (no code)
   - "For each ___, I ___ until ___"
```

## Data vs process — quick check

| Data (nouns) | Process (verbs) |
|--------------|-----------------|
| input slice | iterate |
| interval | compare |
| index `i` | update accumulator |
| empty list | return early |

If you're hunting for "missing variables," you probably need a **state** row, not more inputs.

## State table template

Use for any kata with iteration:

| Step | Current item | State before | Action | State after | Output so far |
|------|--------------|--------------|--------|-------------|---------------|
| 0 | — | | | | |
| 1 | | | | | |

Fill at least 3 rows from the problem examples before writing `for`.

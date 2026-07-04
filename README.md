# Algorithm Kata — State & Interval Training Lab

A **progressive** Go kata track built for one goal: understand the problem *before* you write code.

If merge intervals, booking logic, or "what is my loop variable even doing?" has made you rage-quit — start at **Track 01**, not Track 10.

## How to use this repo

1. Open `CURRICULUM.md` — follow the order strictly.
2. For each kata, read `problem.md` **fully** before opening `solution.go`.
3. Fill in the **State table** in `problem.md` (with pen or in the file) using the examples.
4. Write the smallest code that passes the first test, then the next.
5. When stuck >10 minutes: read `WHEN_STUCK.md`. Do not spiral.

```bash
# Run one kata's tests
go test ./kata/sum_numbers/...

# Run everything (foundation katas will fail until you implement them)
go test ./...
```

## Track overview

| Track | Kata | Skill |
|------:|------|-------|
| 01 | `sum_numbers` | Loop + accumulator (no state yet) |
| 02 | `running_max` | Running state while iterating |
| 03 | `apply_deltas_trace` | State after each transformation step |
| 04 | `count_in_range` | Loop + condition (data vs process) |
| 05 | `interval_overlaps` | One interval question, no lists |
| 06 | `sort_intervals` | Ordering before merge algorithms |
| 07 | `merge_two_intervals` | Two-item merge (merge intervals in miniature) |
| 08 | `split_interval` | Cut one interval (subtract in miniature) |
| 09 | `interval_merge` | Full merge |
| 10 | `interval_insert` | Insert + merge on sorted list |
| 11 | `interval_subtract` | Remove coverage from base |
| 12 | `timeline_simulation` | Overlapping events → state segments |
| 13 | `find_available_slots` | Booking bridge (compose subtract + filter) |

Tracks 09–12 were the original project. Tracks 01–08 exist because those skills are **prerequisites**, not optional.

## Problem-reading ritual (do this every time)

Before any code:

1. **Data** — What am I given? (types, order, empty cases)
2. **Output** — What shape do I return?
3. **State** — What do I need to remember while the loop runs?
4. **One step** — What happens for a *single* item?
5. **First test** — Which test is the easiest? Start there.

See `PROBLEM_READING.md` for a fill-in template.

## When you're stuck

Read `WHEN_STUCK.md`. Confusion is almost always missing structure, not missing ability.

# Curriculum

Complete tracks in order. Do **not** skip to interval merge until Tracks 01–08 feel boring.

## Tier A — Loops without panic (Tracks 01–04)

**Targets:** Go syntax noise, data vs process, first taste of "running state"

| Track | Folder | You learn |
|------:|--------|-----------|
| 01 | `kata/sum_numbers` | `for`, accumulator, empty input |
| 02 | `kata/running_max` | State that updates each iteration |
| 03 | `kata/apply_deltas_trace` | Sequential transforms; output history of state |
| 04 | `kata/count_in_range` | Filter condition inside a loop |

**Exit criteria:** You can explain what your loop variable is *right now* vs what you're *building*.

## Tier B — Intervals as objects (Tracks 05–08)

**Targets:** Understand intervals before algorithms on lists of intervals

| Track | Folder | You learn |
|------:|--------|-----------|
| 05 | `kata/interval_overlaps` | Overlap definition (booking collision) |
| 06 | `kata/sort_intervals` | Sort by `Start` — prerequisite for merge |
| 07 | `kata/merge_two_intervals` | Merge logic with only two items |
| 08 | `kata/split_interval` | Remove a middle chunk from one interval |

**Exit criteria:** You can draw `[1,10]` minus `[3,5]` on paper without code.

## Tier C — Interval algorithms (Tracks 09–12)

**Targets:** Multi-step state on lists; booking engine primitives

| Track | Folder | You learn |
|------:|--------|-----------|
| 09 | `kata/interval_merge` | Sort + sweep merge |
| 10 | `kata/interval_insert` | Three phases: before / merge / after |
| 11 | `kata/interval_subtract` | Union base, union remove, cut |
| 12 | `kata/timeline_simulation` | Event sweep → state segments |

**Exit criteria:** You can say which track maps to "user books 2pm–3pm" for each.

## Tier D — Composition (Track 13)

| Track | Folder | You learn |
|------:|--------|-----------|
| 13 | `kata/find_available_slots` | Availability = working hours − busy, then filter by duration |

**Exit criteria:** You see slot-engine / Cal.com as wiring, not magic.

## Weekly pace (suggested)

- **Week 1:** Tracks 01–04 (one per day, repeat if needed)
- **Week 2:** Tracks 05–08
- **Week 3:** Tracks 09–10
- **Week 4:** Tracks 11–13

Adjust speed to confidence, not ego.

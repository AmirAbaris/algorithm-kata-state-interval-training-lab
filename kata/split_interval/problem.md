### Track 08 — Split Interval

**Tier B** · Subtract one region from one interval · Mini version of track 11

---

## Plain English

You have one interval `base` and one `cut` interval. Remove the overlap of `cut` from `base`. Return the remaining pieces (0, 1, or 2 intervals).

Assume `base` and `cut` may or may not overlap. No merging of output pieces needed unless they would touch (they won't from one cut).

## DATA

- `base`, `cut`: single intervals

## OUTPUT

- 0 intervals if `cut` fully covers `base`
- 1 interval if no overlap or `cut` clips only one side
- 2 intervals if `cut` punches a hole in the middle

## PROCESS (draw first)

```
base:  [=======]
cut:      [--]
left:  [==]     right: [==]
```

Cases:
1. No overlap → `[base]`
2. `cut` covers all of `base` → `[]`
3. `cut` overlaps left part → `[cut.End+1, base.End]` (watch boundaries)
4. `cut` overlaps right part → `[base.Start, cut.Start-1]`
5. `cut` in middle → two fragments

Use inclusive integer points: point `t` is inside iff `Start <= t <= End`.

## Examples

- base `[1,10]`, cut `[3,5]` → `[1,2], [6,10]`
- base `[1,5]`, cut `[6,8]` → `[1,5]`
- base `[1,5]`, cut `[1,5]` → `[]`
- base `[1,10]`, cut `[0,20]` → `[]`

## Connection

`interval_subtract` applies this logic across many base and cut intervals. Master one base, one cut on paper first.

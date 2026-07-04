### Track 13 — Find Available Slots

**Tier D** · Composition · Cal.com availability in one function

---

## Plain English

Given when a host **can** work and when they're **busy**, return free slots long enough for a meeting.

This is what slot-engine and booking systems actually call at the end of the pipeline.

## DATA

- `workingHours`: when the host is available (may overlap — treat as union)
- `busy`: existing bookings / blocked time (may overlap — treat as union)
- `minDuration`: minimum inclusive length a slot must have

**Inclusive length:** `[2,5]` has duration `5 - 2 + 1 = 4` time units.

## OUTPUT

- Sorted, non-overlapping free intervals
- Only intervals with `duration >= minDuration`
- Empty if nothing fits

## PROCESS (decompose — do not skip steps)

1. **Subtract** busy from working hours → raw free time (track 11)
2. **Filter** each interval: keep if `End - Start + 1 >= minDuration` (track 04 pattern)

You may import `interval_subtract.Subtract` once track 11 is done, or inline the logic.

## STATE

Depends on how you implement subtract. If composing:

- Step 1 output = intermediate **data**
- Step 2 loop = **process** over that data

## Examples

Working: `[9,17]` (9am–5pm as integers), Busy: `[11,12], [14,15]`, minDuration `2`  
Free after subtract: `[9,10], [13,13], [16,17]` — all length ≥ 2?  
- `[9,10]` len 2 ✓  
- `[13,13]` len 1 ✗  
- `[16,17]` len 2 ✓  

Output: `[9,10], [16,17]`

## Connection

You just built a booking availability query without HTTP, DB, or timezones. Next step in real life: map integers to `time.Time` and add IANA zones.

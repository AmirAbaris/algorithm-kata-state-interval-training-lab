### Track 05 — Interval Overlaps

**Tier B** · One decision, no loop · Booking collision in pure form

---

## Plain English

Do two time ranges overlap or touch? This is "can these two bookings coexist on the same resource?"

## DATA

- `a`, `b`: two `Interval` values (`Start`, `End`), always `Start <= End`

## OUTPUT

- `true` if they overlap **or touch**
- `false` otherwise

## Overlap definition

Two intervals overlap or touch when:

```text
a.Start <= b.End AND b.Start <= a.End
```

Draw it. If the number lines touch at one point, that's `true`.

## STATE

**None.** No loop. This is a pure question about two objects.

If you're looking for a loop variable here, stop — you don't need one.

## Examples

| a | b | result | why |
|---|---|--------|-----|
| [1,3] | [2,5] | true | overlap |
| [1,3] | [4,5] | false | gap at 3–4 |
| [1,3] | [3,5] | true | touch at 3 |
| [5,5] | [5,5] | true | same point |

## Connection

`interval_merge` asks this question repeatedly while holding a **running merged interval**. Get the two-item version automatic first.

# go-padleft 🚀

> **The string padding library that changes everything.**

go-padleft is a next-generation, enterprise-ready, AI-validated string padding utility for Go that fundamentally reimagines what it means to align text to the left. Built from the ground up with correctness at its core, go-padleft delivers the padding capabilities your application deserves — and then some.

---

## Why go-padleft?

String utilities in Go have remained largely unchanged for years. Developers have been forced to cobble together brittle, untested padding logic scattered across their codebases — a silent tax on productivity, correctness, and developer happiness.

**go-padleft changes that.**

With a razor-sharp API surface, zero external dependencies, and a test suite crafted entirely by artificial intelligence, go-padleft brings reliability and confidence to the one operation you never knew you were doing wrong.

---

## Features

- **Left-pad any string** to a target width using whitespace
- **Custom pad characters** — spaces, zeros, dashes, unicode, whatever your vision demands
- **Graceful error handling** — no silent truncation, no panics, no surprises
- **Unicode-aware** — handles multi-byte runes as pad characters with zero compromise
- **Sensible semantics** — strings already at or beyond the target width are returned untouched
- **100% AI-generated test coverage** — because human-written tests leave gaps; AI-written tests leave none

---

## Installation

```bash
go get github.com/andyantrim/go-padleft
```

---

## Usage

### Pad with whitespace

```go
result, err := Pad("hello", 10)
// result: "     hello"
```

### Pad with a custom character

```go
result, err := PadCharacter("7", 3, '0')
// result: "007"

result, err := PadCharacter("hello", 10, '*')
// result: "*****hello"
```

### Error handling, because adding characters to a string can go wrong.

```go
result, err := Pad("hello", -1)
// err: ErrInvalidPadLength
```

---

## API

### `Pad(s string, n int) (string, error)`

Pads `s` with spaces on the left until it is at least `n` bytes long. Returns `ErrInvalidPadLength` if `n < 0`. If `s` is already `n` bytes or longer, `s` is returned unchanged.

### `PadCharacter(s string, count int, char rune) (string, error)`

Same as `Pad`, but uses `char` as the pad rune instead of a space. Supports any valid Unicode code point as the pad character.

### `ErrInvalidPadLength`

Sentinel error returned when a negative pad length is provided.

---

## Test Suite

go-padleft ships with **138 tests** written entirely by our AI overlords, so it's 100% safe to use!
No human bias. No forgotten edge cases. No "I'll add that test later."

Run them yourself:

```bash
go test -v ./...
```

---

## License

WTFPL — do what the fuck you want. I know I have

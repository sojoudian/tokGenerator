# tokGenerator

A tiny, dependency-free Go utility for generating **cryptographically secure**, **URL-safe**, fixed-length tokens.  
It uses `crypto/rand` for entropy and `encoding/base64` (URL variant, no padding) for compact, interoperable encoding.

> **Why?** Security comes from **entropy** (unpredictable random bits), not the text encoding. This tool focuses on generating high-entropy tokens you can safely place in URLs, headers, and database fields.

---

## Features

- 🔐 **CSPRNG**: Uses `crypto/rand` (not math/rand).
- 🔗 **URL-safe**: Base64 URL encoding (`-` and `_`), **no padding**.
- 📏 **Fixed length**: Default token is **64 characters** (from 48 random bytes → 64 chars).
- 🧩 **Zero external deps**: Standard library only.

---

## Quick start

### Run as a CLI

```bash
# clone and run
git clone https://github.com/sojoudian/tokGenerator.git
cd tokGenerator
go run .
```

Example output:

```
2vEJ4m4r0l3mN7bM5U0k6s5J_4b6s0…   # exactly 64 URL-safe chars
```

> The default build prints a single 64-character token to stdout.


## How it works (in one minute)

- **Entropy source:** `crypto/rand.Read` fills a byte slice with cryptographically secure random bytes from the OS.
- **Encoding:** `base64.RawURLEncoding.EncodeToString` converts those bytes to URL-safe Base64:
  - Uses `-` and `_` instead of `+` and `/`
  - **No** trailing `=` padding
- **Length math:** Base64 encodes every 3 bytes into 4 characters.  
  `48 bytes × (4 / 3) = 64 characters` → neat, fixed length, no padding.

---

## Good practices (if you store/verify tokens)

- **Store only a hash** of the token (e.g., SHA-256) and show the plaintext once.
- **Compare in constant time** (`hmac.Compare`) when verifying.
- **Add prefixes** (e.g., `tnl_`) to identify type/version.
- **Support rotation & revocation** (track created-at, last-used, scopes).

---

## Requirements

- Go **1.24.6** (as specified in `go.mod`)

---

## Module path note

Your `go.mod` currently declares:

```
module gtihub.com/sojoudian/tokGenerator
```

If you want consumers to import via GitHub, change it to:

```
module github.com/sojoudian/tokGenerator
```

Then run:

```bash
go mod tidy
```

---

## License

MIT (or your preferred license). Add a `LICENSE` file if you want others to reuse it.

---

## Roadmap (optional ideas)

- Flags for **length** and **count** (e.g., `-len 64 -n 5`)
- Alternate encodings: **base62** (alphanumeric only), **hex** (debug-friendly)
- Optional **prefix/suffix** (e.g., `tnl_…`)
- Output formats: plain, JSON, newline-delimited

---

Happy generating! If you want, I can add a tiny flag parser so `go run . -len 64 -n 10` prints 10 tokens on demand.


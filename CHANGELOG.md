# Studio engine releases

## 0.12.0-studio.1 — 2026-10-12

First Studio fork release, based on upstream `2fc4c893` following danser-go 0.12.0.

- Studio CLI manifest protocol with replay SHA-256 verification and fixed identity-bound RGB colors.
- Configured FFmpeg executable passed without shell argument construction.
- Background reads confined to the selected map directory.
- Bounded map/replay reads, LZMA expansion/dictionaries/frame counts, decoded images and textures.
- Windows native regressions and synthetic video/screenshot integration validated; Windows junction containment passes. File-symlink regression remains privilege-dependent.

Complete source assets include the modified engine and Go dependency sources. Original runtime resources and BASS are obtained from their original distribution and retain separate terms. Intended use is the local single-user Studio application; this release does not provide an OS process sandbox.

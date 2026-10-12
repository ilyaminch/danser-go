# Danser Studio engine fork

This is an unofficial fork of [Wieku/danser-go](https://github.com/Wieku/danser-go), used by [Danser Studio WebUI](https://github.com/ilyaminch/danser-go-webui). Original authorship, history, GPLv3 license, asset credits and third-party notices are preserved. Studio modifications were introduced in October 2026; this separation is dated 2026-10-12.

Base: upstream `2fc4c8931ff4446a411b0094d7d7b4ede82c61c4`, following release `0.12.0`. `main` preserves upstream 0.12.0; `develop` is the integration branch. Studio work is reviewed through feature/fix branches. The WebUI pins a specific engine commit, not a moving branch.

## Studio integration

- CLI-only build entry point: `tools/studio-cli`.
- `-studio-version` prints `danser-studio-manifest:1`.
- `-studio-manifest <file>` loads selected replay paths, SHA256 identities and fixed RGB colors. It rejects changed inputs rather than silently misassigning colors or skipping a selected attempt.
- Colors follow replay identities after score sorting and optional addition of the danser cursor.
- `DANSER_STUDIO_FFMPEG` selects the owner-configured FFmpeg executable.

Example manifest:

```json
{"Version":1,"Replays":[{"Path":"attempt.osr","SHA256":"<64 lowercase hex characters>","Color":"#ff5588"}]}
```

Stock danser flags/settings remain available. The upstream desktop launcher is unchanged; the WebUI uses the Studio CLI. Build instructions and runtime preparation are maintained in the WebUI repository's `Build-Engine.ps1` and README. No compiled runtime, replays, credentials or output media are distributed in this source repository.

## License and assets

Studio engine modifications are distributed under GPLv3, as part of this covered work. See [LICENSE](LICENSE) and [CREDITS.md](CREDITS.md). A modified binary must be accompanied by a GPL-compliant way to obtain its corresponding source, including these changes and build instructions. Assets and BASS have separate conditions; do not assume the code license covers all runtime resources.

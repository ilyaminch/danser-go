# danser-go — Studio engine fork

Unofficial fork of [Wieku/danser-go](https://github.com/Wieku/danser-go), used by [Danser Studio WebUI](https://github.com/ilyaminch/danser-go-webui). Original history, authorship, GPLv3 license and asset credits are preserved. Studio modifications were introduced in October 2026.

![Danser Studio](https://raw.githubusercontent.com/ilyaminch/danser-go-webui/main/docs/screenshots/workspace.png)

## Which download should I use?

For the web application, download the [Danser Studio Windows package](https://github.com/ilyaminch/danser-go-webui/releases/latest). It includes this engine and launch/setup scripts. This repository's release contains the CLI engine and corresponding sources for developers, not a complete standalone WebUI.

The Studio CLI requires original danser-go 0.12.0 runtime resources/DLLs and FFmpeg. The WebUI installer downloads these from their original distributors with checksum verification; the engine release does not redistribute BASS or default skin artwork. See [CREDITS.md](CREDITS.md) and the original runtime distribution's notices.

## Studio changes

- Explicit replay manifest with SHA-256 identities and individual RGB colors; colors remain bound to attempts after native score sorting.
- CLI entry point in `tools/studio-cli`; `-studio-version` reports `danser-studio-manifest:1`.
- Owner-selected FFmpeg via `DANSER_STUDIO_FFMPEG`.
- Map-directory background containment at file opening, bounded map/replay reads, LZMA output/dictionary/frame limits and decoded image/texture allocation budgets.

See [STUDIO.md](STUDIO.md) for the integration contract, limits and validation scope. Stock CLI flags and engine settings remain available. The upstream desktop launcher is retained in source; Studio uses the CLI build.

## Build and validation

The WebUI pins an exact engine commit as a submodule. Clone WebUI with `--recurse-submodules`, then follow its README and `Build-Engine.ps1`: tested toolchain is Go 1.26.1 and WinLibs GCC 16.2.0 x64 POSIX/MSVCRT. The script builds this CLI, stamps `STUDIO_VERSION`, and runs settings, dance, map-directory, native image/texture and replay-decoder regressions. The complete source release includes Go dependency sources under `vendor/`.

Windows native regressions and synthetic WebUI video/screenshot integration pass for this release. This is not an OS process sandbox: conditional FFmpeg secondary-resource and cross-account access hypotheses still require disposable-environment validation. Map videos are disabled by default in Studio; the application is intended for one trusted local OS user.

`main` contains released Studio work; `develop` is the integration branch. Changes use Conventional Commits and GitFlow with merge commits. Original upstream documentation is retained in [UPSTREAM_README.md](UPSTREAM_README.md); the fork does not claim upstream endorsement.

## License

Engine and Studio modifications are [GPL-3.0](LICENSE). Release binaries are accompanied by corresponding source assets and build instructions. Libraries, BASS and artwork retain their own conditions; the code license does not cover every runtime resource. Original credits and third-party notices are preserved in [CREDITS.md](CREDITS.md) and source dependencies.

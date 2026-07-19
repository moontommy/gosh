# Gosh

A simple Shell written in Go.

## Building

```bash
go build -o gosh .
```

## Releases

Binaries are automatically built and uploaded to [GitHub Releases](https://github.com/tommymoonz/gosh/releases) for each tagged version.

### Download

Get pre-built binaries for Linux, macOS, and Windows from the [releases page](https://github.com/tommymoonz/gosh/releases).

### Create a Release

Tag a commit and push it:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The GitHub Actions workflow will automatically build and release the binaries.

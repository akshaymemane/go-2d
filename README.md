# README.md

# go-2d (M0)

A Go‑native 2D engine with LÖVE‑style ergonomics. This is **Milestone 0**: window, update/draw loop, input, and a simple draw API backed by Ebiten.

## Quick start

```bash
# From the repo root where go.mod lives
go run ./examples/hello-sprite
```

Use ← → to move the square.

## Next steps (M1)

- Implement `gfx.Device.Clear` properly (currently the backend clears in the ebiten frame).
- Add image decode + `LoadImage(path)` using `image/*` and `ebiten.NewImageFromImage`.
- Add basic text (truetype + ebiten/text).
- Add a small snapshot test harness that renders the example offscreen and image-hashes the frame.

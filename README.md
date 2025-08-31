# luvgo (M1)

A Go‑native 2D engine with LÖVE‑style ergonomics. Currently at **Milestone 1**: window, update/draw loop, input, images, and text rendering backed by Ebiten.

## Features (so far)

- Core loop: `Run(Game, Options)`
- Input: keyboard (Left/Right/Up/Down/Space/Esc)
- Graphics:
  - `NewImage(w,h)` → create blank image
  - `LoadImage(path)` → load PNG/JPG from disk
  - `Draw(img, opts)` → draw with position/scale/rotation/tint
  - `Clear(color)` → clear background
- Text:
  - `NewFont(ttfBytes, size)` → load TTF font
  - `NewText(font, "string")` → prepare text object
  - `DrawText(text, x, y)` → draw text on screen

## Examples

### Hello Sprite

```bash
go run ./examples/hello-sprite
```

Use ← → keys to move the square.

### Text and Image

Place assets:

```
examples/text-and-image/assets/ship.png
examples/text-and-image/assets/Inter-Regular.ttf
```

Run:

```bash
go run ./examples/text-and-image
```

You’ll see “Hello, luvgo!” rendered with your font and the PNG drawn below.

## Project layout

```
/core        – engine entrypoint (Run, Game, Options, Context)
/gfx         – drawing API
/input       – keyboard input API
/internal/backend/ebitengl – Ebiten-based backend implementation
/examples    – demo programs
```

## Next milestones

- **M2 Audio**: sound/music playback
- **M3 Graphics**: Canvas (offscreen), SpriteBatch, Scissor, BlendMode, Mouse API

## License

Same spirit as LÖVE: permissive (zlib/libpng).

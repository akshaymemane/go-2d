package ebitengl

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"image/color"
	"testing"

	"go-2d/gfx"

	"github.com/hajimehoshi/ebiten/v2"
)

func hashImage(img *ebiten.Image) string {
	w, h := img.Size()
	pix := make([]uint8, 4*w*h)
	img.ReadPixels(pix)
	sum := sha256.Sum256(pix)
	return hex.EncodeToString(sum[:])
}

type snapGame struct {
	got    string
	sent   bool
	inited bool
}

func (g *snapGame) Update() error {
	if g.inited {
		// Exit after we’ve rendered one frame
		return errors.New("done")
	}
	g.inited = true
	return nil
}

func (g *snapGame) Draw(screen *ebiten.Image) {
	// Use our backend’s draw pipeline but target the provided screen
	currentScreen = screen
	defer func() { currentScreen = nil }()

	dev := newGfx()
	// Background
	dev.Clear(gfx.Color{R: 0.1, G: 0.1, B: 0.12, A: 1})
	// 64x64 square
	square, _ := dev.NewSolid(64, 64, color.RGBA{200, 100, 240, 255})
	dev.Draw(square, &gfx.DrawOptions{X: 96, Y: 96})

	if !g.sent {
		g.got = hashImage(screen)
		g.sent = true
	}
}

func (g *snapGame) Layout(outsideW, outsideH int) (int, int) { return 256, 256 }

func Test_DrawSquareSnapshot(t *testing.T) {
	// Run headless so no window pops up on CI/macOS
	t.Setenv("EBITEN_HEADLESS", "1")

	game := &snapGame{}
	err := ebiten.RunGame(game)
	// We expect to exit via the sentinel error from Update after one frame
	if err == nil || err.Error() != "done" {
		t.Fatalf("unexpected error: %v", err)
	}

	got := game.got

	// First run: log the hash and set it below.
	// t.Logf("snapshot hash: %s", got)

	const want = "" // <-- paste the logged hash here after first run
	if want != "" && got != want {
		t.Fatalf("snapshot mismatch\n got: %s\nwant: %s", got, want)
	}
}

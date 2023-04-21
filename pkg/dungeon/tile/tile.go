package tile

import "github.com/veandco/go-sdl2/sdl"

const IN_RANGE_MODIFIER = "_IN_RANGE"

type Tile struct {
	tileType   string
	lit        bool
	wasJustLit bool
	walkable   bool
}

func (t Tile) GetType() string {
	return t.tileType
}

func (t Tile) IsWalkable() bool {
	return t.walkable
}

func (t Tile) DrawTile(ren *sdl.Renderer, x int, y int) string {
	if !t.Visible() {
		return " "
	}

	tileName := t.tileType

	if t.wasJustLit {
		tile, ok := TILE_CHARACTERS_MAP[tileName+IN_RANGE_MODIFIER]
		if ok {
			return tile
		}
	}

	tile, _ := TILE_CHARACTERS_MAP[t.tileType]

	var tileRect sdl.Rect

	tileRect.X = int32(y) * 8
	tileRect.Y = int32(x) * 8
	tileRect.W = 8
	tileRect.H = 8

	if t.lit {
		if t.tileType == "FLOOR" {
			ren.SetDrawColor(0, 0, 100, 255)
			ren.DrawRect(&tileRect)
		} else if t.tileType == "WALL" {
			ren.SetDrawColor(0, 0, 255, 255)
			ren.DrawRect(&tileRect)
		}
	}

	return tile
}

func (t Tile) Visible() bool {
	return t.lit
}

func (t *Tile) LightUp() {
	t.lit = true
	t.wasJustLit = true
}

func (t *Tile) MarkDrawnTile() {
	t.wasJustLit = false
}

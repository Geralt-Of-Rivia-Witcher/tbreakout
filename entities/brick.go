package entities

import "math/rand"

type Brick struct {
	X      int
	Y      int
	Width  int
	height int
	Alive  bool
}

const (
	brickWidth = 5
	spacing    = 3
)

func newBrick(X int, Y int) *Brick {
	return &Brick{
		X:      X,
		Y:      Y,
		Width:  5,
		height: 1,
		Alive:  true,
	}
}

func GenerateClassicGrid(
	rows, cols, screenWidth, startY int,
) []*Brick {
	var bricks []*Brick

	for row := 0; row < rows; row++ {
		startX := centeredRowStartX(cols, screenWidth)
		y := startY + row*2

		for col := 0; col < cols; col++ {
			x := startX + col*(brickWidth+spacing) + brickWidth/2
			bricks = append(bricks, newBrick(x, y))
		}
	}
	return bricks
}

func GenerateOffsetGrid(
	rows, cols, screenWidth, startY int,
) []*Brick {
	var bricks []*Brick

	for row := 0; row < rows; row++ {
		startX := centeredRowStartX(cols, screenWidth)

		if row%2 == 1 {
			startX += (brickWidth + spacing) / 2
		}

		y := startY + row*2

		for col := 0; col < cols; col++ {
			x := startX + col*(brickWidth+spacing) + brickWidth/2
			bricks = append(bricks, newBrick(x, y))
		}
	}
	return bricks
}

func GeneratePyramid(
	rows, screenWidth, startY int,
) []*Brick {
	var bricks []*Brick

	for row := 0; row < rows; row++ {
		cols := rows - row
		startX := centeredRowStartX(cols, screenWidth)
		y := startY + row*2

		for col := 0; col < cols; col++ {
			x := startX + col*(brickWidth+spacing) + brickWidth/2
			bricks = append(bricks, newBrick(x, y))
		}
	}
	return bricks
}

func GenerateCheckerboard(
	rows, cols, screenWidth, startY int,
) []*Brick {
	var bricks []*Brick

	for row := 0; row < rows; row++ {
		startX := centeredRowStartX(cols, screenWidth)
		y := startY + row*2

		for col := 0; col < cols; col++ {
			if (row+col)%2 == 0 {
				x := startX + col*(brickWidth+spacing) + brickWidth/2
				bricks = append(bricks, newBrick(x, y))
			}
		}
	}
	return bricks
}

func GenerateDiamond(
	rows, screenWidth, startY int,
) []*Brick {
	var bricks []*Brick
	mid := rows / 2

	for row := 0; row < rows; row++ {
		var cols int
		if row <= mid {
			cols = row + 1
		} else {
			cols = rows - row
		}

		startX := centeredRowStartX(cols, screenWidth)
		y := startY + row*2

		for col := 0; col < cols; col++ {
			x := startX + col*(brickWidth+spacing) + brickWidth/2
			bricks = append(bricks, newBrick(x, y))
		}
	}
	return bricks
}

func GenerateRandomLayout(
	screenWidth, startY int,
) []*Brick {
	switch rand.Intn(5) {
	case 0:
		return GenerateClassicGrid(7, 12, screenWidth, startY)
	case 1:
		return GenerateOffsetGrid(8, 12, screenWidth, startY)
	case 2:
		return GeneratePyramid(9, screenWidth, startY)
	case 3:
		return GenerateCheckerboard(9, 12, screenWidth, startY)
	default:
		return GenerateDiamond(12, screenWidth, startY)
	}
}

func centeredRowStartX(cols, screenWidth int) int {
	rowWidth := (cols * brickWidth) + (cols-1)*spacing
	return (screenWidth - rowWidth) / 2
}

func AreAllBricksDead(bricks []*Brick) bool {
	for _, brick := range bricks {
		if brick.Alive {
			return false
		}
	}
	return true
}

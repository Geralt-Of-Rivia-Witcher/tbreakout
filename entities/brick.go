package entities

type Brick struct {
	X      int
	Y      int
	Width  int
	height int
	Alive  bool
}

func newBrick(X int, Y int) *Brick {
	return &Brick{
		X:      X,
		Y:      Y,
		Width:  5,
		height: 1,
		Alive:  true,
	}
}

func GenerateBricks(rows int, cols int, screenWidth int, startY int) []*Brick {
	var bricks []*Brick
	const brickWidth = 5
	const spacing = 1

	emptySpaceAroungBrickArea := screenWidth - ((cols * brickWidth) + (cols-1)*spacing)
	for i := 1; i <= rows; i++ {
		startX := emptySpaceAroungBrickArea / 2
		for j := 1; j <= cols; j++ {
			bricks = append(bricks, newBrick(startX+(brickWidth/2), startY+i+i))
			startX += (brickWidth + 1)
		}
	}
	return bricks
}

func AreAllBricksDead(bricks []*Brick) bool {
	for _, brick := range bricks {
		if brick.Alive {
			return false
		}
	}
	return true
}

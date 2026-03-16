package game

import (
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/constants"
)

type ScoreEvent int

const (
	BrickHitEvent ScoreEvent = iota
	LevelClearedEvent
)

func (game *Game) updateScore(event ScoreEvent, combo int) {
	switch event {
	case BrickHitEvent:
		game.runningGameEntities.score += (constants.ScoreForHittingBrick * combo)
	case LevelClearedEvent:
		game.runningGameEntities.score += constants.ScoreForClearingLevel
	}
}

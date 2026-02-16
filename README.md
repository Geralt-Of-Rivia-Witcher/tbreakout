# tbreakout

A terminal Breakout clone written in Go using `tcell`.

## Features

- Playable in a terminal UI
- Randomized brick layout each run
- Score + lives HUD
- Title and game-over screens

## Controls

- `Enter`: Start game from title screen
- `Left Arrow`: Move paddle left
- `Right Arrow`: Move paddle right
- `R`: Restart after game over
- `Esc` or `Ctrl+C`: Quit

## Install (macOS via Homebrew)

```bash
brew tap Geralt-Of-Rivia-Witcher/tbreakout
brew install tbreakout
tbreakout
```

## Run From Source

### Prerequisites

- Go `1.25.6` or newer

### Run

```bash
go run .
```

### Build

```bash
go build -o tbreakout
./tbreakout
```

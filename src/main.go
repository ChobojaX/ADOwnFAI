package main

import (
	"github.com/ChobojaX/ADOwnFAI/src/theme"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// https://adofai.gg/api/v1/levels/{id}

func main() {
	Theme := theme.Theme{
		BackGround: rl.DarkGray,
		Primary:    rl.Color{},
		Secondary:  rl.Color{},
		Text:       rl.Color{},
		SubText:    rl.Color{},
	}

	rl.InitWindow(1920, 1080, "ADOwnFAI")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(Theme.BackGround)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

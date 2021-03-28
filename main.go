package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/Newmi1988/gol/game"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	game            *game.Life
	hz, generations int
}

func NewGame(width, height int, fill float32, hz, generations int) *Runner {

	return &Runner{
		game:        game.GameOfLife(width, height, fill),
		hz:          hz,
		generations: generations,
	}
}

func (r *Runner) Run() {
	for i := 0; i < r.generations; i++ {
		r.game.Step()
		fmt.Print("\x0c", r.game) // Clear screen and print field.
		time.Sleep(time.Second / time.Duration(r.hz))
	}
}

func main() {
	app := &cli.App{
		Name:     "Go Game of Life",
		Version:  "0.1.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Newmi1988",
				Email: "newmi1988@gmail.com",
			},
		},
		HelpName:             "gol",
		Usage:                "Game of Life in GO",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "width",
				Usage:       "Set the width of the game field",
				Value:       100,
				DefaultText: "width",
			},
			&cli.IntFlag{
				Name:        "height",
				Usage:       "Set the height of the game field",
				Value:       50,
				DefaultText: "height",
			},
			&cli.IntFlag{
				Name:        "hz",
				Usage:       "set update rate in hz",
				Value:       50,
				DefaultText: "width of the game field",
			},
			&cli.IntFlag{
				Name:        "gens",
				Usage:       "set update rate in hz",
				Value:       500,
				DefaultText: "width of the game field",
			},
			&cli.Float64Flag{
				Name:        "fill",
				Usage:       "Percent of values filled on the board",
				Value:       0.33,
				DefaultText: "0.33",
			},
		},
		Action: func(ctx *cli.Context) error {

			fmt.Printf("height %v \n", ctx.Int("height"))

			r := NewGame(ctx.Int("width"), ctx.Int("height"), float32(ctx.Float64("fill")), ctx.Int("hz"), ctx.Int("gens"))
			r.Run()
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

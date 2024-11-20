package main

import (
	"flag"
	"fmt"

	"github.com/EvaLLLLL/ghcld/config"
	"github.com/EvaLLLLL/ghcld/draw"
	"github.com/EvaLLLLL/ghcld/fetch"
	"github.com/EvaLLLLL/ghcld/symbol"
	"github.com/EvaLLLLL/ghcld/types"
)

func main() {

	fmt.Printf("%v", len("❄"))

	initFlag := flag.Bool("init", false, "Initialize configuration")
	flag.Parse()

	var c *types.Config
	var err error
	if *initFlag {
		c, err = config.InitConfig()
		if err != nil {
			panic("Error initializing configuration")
		}
	} else {
		c, err = config.CheckConfigValue()
		if err != nil {
			panic("Failed to check configuration")
		}

		symbol, _ := symbol.GetSymbol()

		c.SYMBOL = symbol
	}

	weeks := fetch.FetchGithubCalendar(c)

	fmt.Println()

	draw.DrawCalendar(weeks, c)
}

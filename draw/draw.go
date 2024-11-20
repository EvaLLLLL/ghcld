package draw

import (
	"fmt"

	"github.com/EvaLLLLL/ghcld/types"
)

type Rectangle struct {
	Color string
}

func hexToRGB(hex string) (r, g, b int) {
	var color int
	fmt.Sscanf(hex, "#%x", &color)
	r = (color >> 16) & 0xFF
	g = (color >> 8) & 0xFF
	b = color & 0xFF
	return
}

func drawRectangle(rect Rectangle, config *types.Config) {
	r, g, b := hexToRGB(rect.Color)

	fmt.Printf("\033[38;2;%d;%d;%dm", r, g, b)

	fmt.Print(config.SYMBOL)
	fmt.Print(" ")
	fmt.Print("\033[0m")
}

func mapDataToRectangles(data *[]types.Week) []Rectangle {
	rectangles := []Rectangle{}

	for _, value := range *data {
		for _, day := range value.ContributionDays {
			rectangles = append(rectangles, Rectangle{
				Color: day.Color,
			})
		}
	}
	return rectangles
}

func DrawCalendar(data *[]types.Week, config *types.Config) {
	rectangles := mapDataToRectangles(data)

	rows := make([][]Rectangle, 7)

	for i, rect := range rectangles {
		for j := 0; j < 7; j++ {
			if (i-j)%7 == 0 {
				rows[j] = append(rows[j], rect)
			}
		}
	}

	for i, row := range rows {
		for j, day := range row {
			if j == 0 {
				fmt.Print("    ")
			}

			drawRectangle(day, config)
		}

		if i < 7 {
			fmt.Println()
		}
	}
}

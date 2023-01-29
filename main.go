package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

func hexTo256(hex string) {
	// hex
	color.HEXStyle("#ffffff", hex).Printf(hex)

	fmt.Printf(" ") // margin

	// rgb
	rgb := color.HexToRGB(hex)
	r, g, b := uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])
	color.NewRGBStyle(
		color.RGBColor{255, 255, 255, 0},
		color.RGBColor{r, g, b, 0},
	).Printf(fmt.Sprintf("%3d %3d %3d", r, g, b))

	fmt.Printf(" ") // margin

	// 256
	v256 := color.RgbTo256(r, g, b)
	color.S256(
		255,
		uint8(v256),
	).Println(fmt.Sprintf("%3d", v256))
}

func header() {
	color.S256(255, 32).Printf(" # hex ")
	fmt.Printf(" ") // margin
	color.S256(255, 32).Printf(" r   g   b ")
	fmt.Printf(" ") // margin
	color.S256(255, 32).Println("256")
	fmt.Println()
}

var example_hexes = []string{
	"#268bd2",
	"#2aa198",
	"#586e75",
	"#657b83",
	"#6c71c4",
	"#859900",
	"#93a1a1",
	"#b58900",
	"#cb4b16",
	"#dc322f",
	"#eee8d5",
}

func main() {
	var hexes []string

	if len(os.Args) > 1 {
		arg := os.Args[1]

		dat, err := os.ReadFile(arg)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("So will print examples instead")
			hexes = example_hexes
		} else {
			hexes = strings.Split(string(dat), "\n")
		}
	} else {
		fmt.Println("No file is provided")
		fmt.Println("So will print examples instead")
		hexes = example_hexes
	}

	header()

	for _, hex := range hexes {
		if hex != "" {
			hexTo256(hex)
		}
	}
}

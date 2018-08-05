package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"math/rand"
	"syscall/js"
)

var target js.Value
var colors []string
var offset int

func main() {
	for i := 0; i < 32; i++ {
		base := 255.0 * rand.Float64()
		color := colorful.Color{base, base, base}
		colors = append(colors, color.Hex())
	}
	target = js.Global().Get("document").Call("getElementById", "svg")
	animate := js.Global().Get("requestAnimationFrame")
	animate.Invoke(js.NewCallback(animationTick))
	select {}
}

func animationTick(args []js.Value) {
	target.Set("innerHTML", Tiger())
	animate := js.Global().Get("requestAnimationFrame")
	animate.Invoke(js.NewCallback(animationTick))
}

func randomColor() string {
	color := colors[offset]
	newOffset := offset + 1
	if newOffset == 32 {
		newOffset = 0
	}
	offset = newOffset
	return color
}

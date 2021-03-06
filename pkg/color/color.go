package color

import "github.com/fatih/color"

var (
	// White color
	White = color.New(color.FgWhite).SprintFunc()
	// WhiteBold bold
	WhiteBold = color.New(color.FgWhite & color.Bold).SprintFunc()
	// Green color
	Green = color.New(color.FgGreen).SprintFunc()
	// Red color
	Red = color.New(color.FgRed).SprintFunc()
	// Cyan color
	Cyan = color.New(color.FgCyan).SprintFunc()
)

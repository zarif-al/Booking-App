package print

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func PrintError(err string) {
	fmt.Println(color.Ize(color.Red, err))
}

func PrintWarning(warning string) {
	fmt.Println(color.Ize(color.Yellow, warning))
}

func PrintSuccess(success string) {
	fmt.Println(color.Ize(color.Green, success))
}

func PrintInputMsg(msg string) {
	fmt.Print(color.Ize(color.Cyan, msg))
}

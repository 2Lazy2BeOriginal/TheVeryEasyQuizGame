package col

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	REVERSE    = "\033[7m"
	GREEN      = "\033[32m"
	WHITE      = "\033[37m"
	RESETCOLOR = "\033[0m"
	YELLOW     = "\033[33m"
	GREY       = "\033[37m"
	BOLD       = "\033[1m"
	BLUE       = "\033[34m"
	RED        = "\033[31m"
	BLACK      = "\033[90m"
	MAGENTA    = "\033[35m"
	CYAN       = "\033[36m"
)

func Clearscreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Highlight(s ...interface{}) string {
	text := BOLD + REVERSE
	for _, i := range s {
		text += fmt.Sprint(i)
	}
	text += RESETCOLOR
	return text
}

// round to 2 places because why not
func Round(num float64) float64 {
	ans := fmt.Sprintf("%.2f", num)
	rounded, _ := strconv.ParseFloat(ans, 64)
	return rounded
}

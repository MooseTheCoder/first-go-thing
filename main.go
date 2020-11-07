package main

// Import Necessary packages
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Declare current menu render iteration
var menuIteration int64

// Declare menu items map
var MenuMap = make(map[string]int64)

// Draw menu and await user input
func main() {
	renderMenu(1)
	awaitInput()
}

// Read stdin
func awaitInput() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	// Convert rune to string, string to int
	CharInt, Error := strconv.ParseInt(string(char), 10, 64)
	if Error != nil {
		renderMenu(1)
	}
	// Pass int to renderMenu
	renderMenu(CharInt)
}

// Draw the menu
func renderMenu(Option int64) {
	menuIteration = 0
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("-----MENU------")
	drawMenuItem(Option, "Item 1")
	drawMenuItem(Option, "Item 2")
	drawMenuItem(Option, "Item 3")
	drawMenuItem(Option, "Item 4")
	drawMenuItem(Option, "Item 5")
	awaitInput()
}

func drawMenuItem(Option int64, Title string) {
	menuIteration = menuIteration + 1
	MenuMap[Title] = menuIteration
	fmt.Print(Title)
	if MenuMap[Title] == Option {
		fmt.Print(" <--")
	}
	fmt.Print("\n")
}

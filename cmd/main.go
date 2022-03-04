package main

import (
	"bufio"
	"os"

	"github.com/fatih/color"

	"github.com/sravindra99/Go-Stand-Up/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	infoText := color.New(color.FgYellow, color.Bold)
	infoText.Println("Enter participants (space separated): ")
	text, _ := reader.ReadString('\n')
	utils.Standup(text)
}

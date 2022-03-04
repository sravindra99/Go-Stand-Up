package main

import (
	"bufio"
	"github.com/fatih/color"
	"github.com/sravindra99/Go-Stand-Up/utils"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	infoText := color.New(color.FgYellow, color.Bold)
	infoText.Println("Enter participants (space separated): ")
	text, _ := reader.ReadString('\n')
	utils.Standup(text)
}

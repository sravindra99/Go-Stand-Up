package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func Standup(input string) {
	var standup Session
	var choice string

	infoText := color.New(color.FgYellow, color.Bold)
	infoText.Println("Starting team standup!!!")
	personText := color.New(color.FgGreen, color.Bold)

	errorText := color.New(color.FgRed)

	err := standup.BuildUserMap(input)
	if err != nil {
		errorText.Println(err)
		return
	}
	var next bool
	for standup.IsDone(){
		infoText.Printf("Choose?(Y/n)")
		fmt.Scanf("%s", choice)
		next = choice != "n"
		if !next {
			break
		}
		person, err := standup.PickRandomUser()
		if err != nil {
			errorText.Println(err)
			return
		}
		fmt.Printf("The winner of the next turn is %s !!!\n", personText.Sprintf(person))

	}
	infoText.Println("That's it folks! Have a great day :)")
}

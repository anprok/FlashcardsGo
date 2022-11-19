package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var importFile, exportFile string
	flag.StringVar(&importFile, "import_from", "", "Enter your filename to import cards from")
	flag.StringVar(&exportFile, "export_to", "", "Enter your filename to export card to")
	flag.Parse()
	if importFile != "" {
		importCards(importFile)
	}
	for {
		printAndLog(fmt.Sprintln("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats)"))
		command := readString()

		switch {
		case command == "add":
			addTerm()
		case command == "remove":
			removeTerm()
		case command == "import":
			var filename string
			if importFile == "" {
				file, err := askFilename()
				filename = file
				if err != nil {
					continue
				}
			} else {
				filename = importFile
			}
			importCards(filename)
		case command == "export":
			var filename string
			if exportFile == "" {
				file, err := askFilename()
				filename = file
				if err != nil {
					continue
				}
			} else {
				filename = exportFile
			}
			exportCards(filename)
		case command == "ask":
			askTerms()
		case command == "exit":
			if exportFile != "" {
				exportCards(exportFile)
			}
			printAndLog(fmt.Sprintln("Bye bye!"))
			return
		case command == "log":
			logFile, err := askFilename()
			if err != nil {
				continue
			}
			saveLogs(logFile)
		case command == "hardest card":
			getHardestCards()
		case command == "reset stats":
			resetStats()
		default:
			printAndLog(fmt.Sprintln("Incorrect command!"))
			break
		}
	}
}

type CardsJSONObject struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Definition string `json:"definition"`
	Term       string `json:"term"`
	Errors     int    `json:"errors"`
}

var cards CardsJSONObject
var logStrings []string

func printAndLog(inp string) {
	fmt.Print(inp)
	logStrings = append(logStrings, inp)
}

func addTerm() {
	var term string
	var definition string
	printAndLog(fmt.Sprintln("The card:"))
	for {
		term = readString()
		if checkForTerm(term) {
			printAndLog(fmt.Sprintf("The term \"%v\" already exists. Try again:\n", term))
		} else {
			printAndLog(fmt.Sprintln("The definition of the card:"))
			for {
				definition = readString()
				if i := getIndexForDef(definition); i >= 0 {
					printAndLog(fmt.Sprintf("The definition \"%s\" already exists. Try again:\n", definition))
				} else {
					cards.Cards = append(cards.Cards, Card{
						Term:       term,
						Definition: definition,
						Errors:     0,
					})
					printAndLog(fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.\n", term, definition))
					break
				}
			}
			break
		}
	}
}

func askFilename() (string, error) {
	var filename string
	printAndLog(fmt.Sprintln("File name:"))
	_, err := fmt.Scanln(&filename)
	return filename, err
}

func importCards(filename string) {
	var imported CardsJSONObject
	data, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		printAndLog(fmt.Sprintln("File not found."))
		return
	}
	err = json.Unmarshal(data, &imported)
	if err != nil {
		log.Fatal(err)
	}
	printAndLog(fmt.Sprintf("%d cards have been loaded.\n", len(imported.Cards)))
	for _, element := range imported.Cards {
		if index := getIndexForTerm(element.Term); index >= 0 {
			cards.Cards[index].Definition = element.Definition
			cards.Cards[index].Errors = element.Errors
		} else {
			cards.Cards = append(cards.Cards, Card{
				Term:       element.Term,
				Definition: element.Definition,
				Errors:     element.Errors,
			})
		}

	}
}

func exportCards(filename string) {
	cardsJson, err := json.Marshal(cards)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = fmt.Fprint(file, string(cardsJson))
	if err != nil {
		log.Fatal(err)
	}
	printAndLog(fmt.Sprintf("%d cards have been saved.\n", len(cards.Cards)))
}

func removeTerm() {
	var card string
	printAndLog(fmt.Sprintln("Which card?"))
	_, err := fmt.Scanln(&card)
	if err != nil {
		return
	}

	if index := getIndexForTerm(card); index >= 0 {
		cards.Cards = append(cards.Cards[:index], cards.Cards[index+1:]...)
		printAndLog(fmt.Sprintln("The card has been removed."))
	} else {
		printAndLog(fmt.Sprintf("Can't remove \"%s\": there is no such card.\n", card))
	}
}

func askTerms() {
	var length int
	if len(cards.Cards) == 0 {
		printAndLog(fmt.Sprintln("No cards to ask."))
		return
	}
	printAndLog(fmt.Sprintln("How many times to ask?"))
	_, err := fmt.Scanln(&length)
	if err != nil {
		return
	}
	var i = length
	for {
		for index, card := range cards.Cards {
			if i == 0 {
				return
			}
			printAndLog(fmt.Sprintf("Print the definition of \"%s\"\n", card.Term))
			proposedAnswer := readString()
			if card.Definition == proposedAnswer {
				printAndLog(fmt.Sprintln("Correct!"))
			} else {
				correctTerm, ok := getTermForDefinition(proposedAnswer)
				if ok {
					printAndLog(fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", card.Definition, correctTerm))
				} else {
					printAndLog(fmt.Sprintf("Wrong. The right answer is \"%s\"\n", card.Definition))
				}
				cards.Cards[index].Errors++
			}
			i--
		}
	}
}

func getIndexForTerm(userValue string) int {
	for i, card := range cards.Cards {
		if card.Term == userValue {
			return i
		}
	}
	return -1
}

func getIndexForDef(userValue string) int {
	for i, card := range cards.Cards {
		if card.Definition == userValue {
			return i
		}
	}
	return -1
}

func checkForTerm(key string) bool {
	for _, card := range cards.Cards {
		if card.Term == key {
			return true
		}
	}
	return false
}

func getTermForDefinition(userValue string) (response string, ok bool) {
	for _, card := range cards.Cards {
		if card.Definition == userValue {
			return card.Term, true
		}
	}
	return "", false
}

func readString() string {
	var userInput string
	reader := bufio.NewReader(os.Stdin)
	userInput, _ = reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func getHardestCards() {
	if len(cards.Cards) == 0 {
		printAndLog(fmt.Sprintln("There are no cards with errors."))
		return
	}
	sortedCards := make([]Card, len(cards.Cards))
	copy(sortedCards, cards.Cards)

	sort.Slice(sortedCards, func(i, j int) bool {
		return sortedCards[i].Errors > sortedCards[j].Errors
	})
	strs := make([]string, 0, len(sortedCards))
	var max = sortedCards[0].Errors
	if max > 0 {
		for _, card := range sortedCards {
			if card.Errors == max {
				strs = append(strs, card.Term)
			}
		}
	}
	if len(strs) > 1 {
		printAndLog(fmt.Sprintf("The hardest cards are %s.\n", strings.Join(strs, ", ")))
	} else if len(strs) == 1 {
		printAndLog(fmt.Sprintf("The hardest card is \"%s\". You have %d errors answering it.\n", strs[0], max))
	} else {
		printAndLog("There are no cards with errors.\n")
	}
}

func resetStats() {
	for i, _ := range cards.Cards {
		cards.Cards[i].Errors = 0
	}
	printAndLog("Card statistics have been reset.\n")
}

func saveLogs(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range logStrings {
		fmt.Fprintln(file, line)
		if err != nil {
			log.Fatal(err)
		}
	}
	printAndLog(fmt.Sprintln("The log has been saved."))
}

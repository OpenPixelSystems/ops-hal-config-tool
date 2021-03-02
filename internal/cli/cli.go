package cli

import (
	"os"
	"log"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

type SuggestionSelection struct {
	Name string
	Suggestions []prompt.Suggest
}

var cmdSelection = SuggestionSelection{
	Name : "cmd",
	Suggestions : []prompt.Suggest{
		{"clocks", "New clock configuration"},
		{"exit", "Exit the program"},
	},
}

var clkSelection = SuggestionSelection{
	Name : "clk",
	Suggestions : []prompt.Suggest{
		{"inclk", "New input clock"},
		{"pll", "New PLL"},
		{"clk", "New peripheral or output clk"},
		{"save", "Save the new clk config"},
		{"exit", "Exit the program"},
	},
}

var activeSelection = clkSelection

func swapSelection(selection SuggestionSelection) {
	activeSelection = selection
}

func cmdExecutor(in string) {
	switch in {
	case "clocks":
		log.Println("Swapping to clock suggestion")
		swapSelection(clkSelection)
		return
	case "exit":
		os.Exit(0)
	}
}

func clkExecutor(in string) {
	switch in {
	case "exit":
		swapSelection(cmdSelection)
		return
	}
}

func executor(in string) {

	in = strings.TrimSpace(in)
	switch activeSelection.Name {
	case "cmd":
		cmdExecutor(in)
		return
	case "clk":
		clkExecutor(in)
		return
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(activeSelection.Suggestions, in.GetWordBeforeCursor(), true)
}

func RunCLI() {
	p := prompt.New(executor, completer,
		prompt.OptionTitle("config helper"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))

	p.Run()
}

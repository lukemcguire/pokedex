package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// supportedCommands is a registry of all available CLI commands.
// It maps command names to their implementations.
var supportedCommands = map[string]cliCommand{}

// cleanInput sanitizes and tokenizes user input by converting
// to lowercase and splitting into words.
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// startRepl begins the Read-Eval-Print Loop for the Pokedex CLI.
// It continuously reads user input, processes commands, and displays results.
func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, ok := supportedCommands[commandName]; ok {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

// init initializes the supported commands map with all available commands.
func init() {
	supportedCommands["catch"] = cliCommand{
		name:        "catch <pokemon_name>",
		description: "Attempts to catch a pokemon",
		callback:    commandCatch,
	}
	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	supportedCommands["explore"] = cliCommand{
		name:        "explore <area_name>",
		description: "Displays all Pokemon in the specified area <area_name>",
		callback:    commandExplore,
	}
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	supportedCommands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Gives details about a specific pokemon",
		callback:    commandInspect,
	}
	supportedCommands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 location areas",
		callback:    commandMap,
	}
	supportedCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapB,
	}
	supportedCommands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Print a list of all pokemon you have caught",
		callback:    commandPokedex,
	}
}

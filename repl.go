package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func replStart(){
    reader := bufio.NewScanner(os.Stdin)
    for{
        fmt.Print("Pokedex > ")
        reader.Scan()

        words:= cleanInput(reader.Text())
        if len(words) ==0{
            continue
        }

        commandName := words[0]

        command,exists := getCommands()[commandName]
        if exists{
            err := command.callback()
            if err != nil{
                fmt.Println(err)
            }
            continue
        }else{
            fmt.Println("unknown command")
            continue
        }
    }
}

func cleanInput(text string)[]string{
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

type cliCommand struct{
    name string
    description string
    callback func()error
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
        "help":{
            name: "help",
            description: "Displays a help command",
            callback: commandHelp,
        },
        "exit":{
            name:"exit",
            description: "Exists the program",
            callback: commandExit,
        },
        "map":{
            name: "map",
            description: "Consecutively shows 20 locations of the Pokemon world",
            callback: commandMap,
        },
        "mapb":{
            name: "mapb",
            description:"shows the previous 20 location | reverse of map",
            callback: commandMapb,
        },
    }
}


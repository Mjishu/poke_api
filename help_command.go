package main

import (
    "fmt"
    "os"
    "github.com/mjishu/ApiCalls"
)

func commandHelp()error{

    fmt.Println()
    fmt.Println("Welcome to the pokedex")
    fmt.Println("Usage: ")
    fmt.Println()
    for _,cmd := range getCommands(){
        fmt.Printf("%s: %s\n", cmd.name,cmd.description)
    }
    fmt.Println()
    return nil
}

func commandExit()error{
    os.Exit(0)
    return nil
} 

func commandMap()error{
    locations,_ := apicalls.CallPoke()
    for _,location := range locations{
        fmt.Printf("%v \n", location)
    }
    return nil
}

func commandMapb() error{
    fmt.Print("showing 20 locations")
    return nil
}

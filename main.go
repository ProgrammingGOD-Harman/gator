package main

import (
	"log"
	"os"

	"www.github.com/ProgrammingGOD-Harman/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

	// cfg, err := config.Read()
	//
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	//
	// fmt.Printf("Read config: %+v\n", cfg)
	//
	// err = cfg.SetUser("Dunk")
	//
	// if err != nil {
	// 	log.Fatalf("couldn't set current user: %v", err)
	// }
	//
	// cfg, err = config.Read()
	//
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	//
	// fmt.Printf("Read config again: %+v\n", cfg)
}

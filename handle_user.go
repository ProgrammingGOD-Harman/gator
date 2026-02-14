package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"www.github.com/ProgrammingGOD-Harman/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("New user %s Registered\n", name)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)

	if err != nil {
		return fmt.Errorf("couldn't login because user is not registered: %w", err)
	}

	err = s.cfg.SetUser(name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("There are no users in the database")
		return nil
	}

	fmt.Println("Users are:")
	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}

	return nil
}

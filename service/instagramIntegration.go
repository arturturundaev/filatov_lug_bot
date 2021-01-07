package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ahmdrz/goinsta/v2"
	"log"
	"os"
	"strings"
)

type MyInstabot struct {
	Insta *goinsta.Instagram
}

var Instabot MyInstabot

// login will try to reload a previous session, and will create a new one if it can't
func Login(login string, password string) bool {
	err := reloadSession()
	if err != nil {
		return createAndSaveSession(login, password)
	}

	return true
}

// reloadSession will attempt to recover a previous session
func reloadSession() error {

	insta, err := goinsta.Import("./goinsta-session")
	if err != nil {
		return errors.New("Couldn't recover the session")
	}

	if insta != nil {
		Instabot.Insta = insta
	}

	log.Println("Successfully logged in")
	return nil

}

func getInput(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(text)
	input, err := reader.ReadString('\n')
	check(err)
	return strings.TrimSpace(input)
}

// Logins and saves the session
func createAndSaveSession(login string, password string) bool {
	insta := goinsta.New(login, password)
	Instabot.Insta = insta
	err := Instabot.Insta.Login()
	if !check(err) {
		return false
	}

	err = Instabot.Insta.Export("./goinsta-session")

	if !check(err) {
		return false
	}
	log.Println("Created and saved the session")
	return true
}

// check will log.Fatal if err is an error
func check(err error) bool {
	if err != nil {
		return false
	}

	return true
}
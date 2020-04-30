package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrPageMoved  = errors.New("Page Moved")
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := webCall(); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request")
			break
		case ErrPageMoved:
			fmt.Println("Page Moved")
			break
		default:
			break
		}
	}
}

func webCall() error {
	guess := rand.Intn(10)
	fmt.Println(guess)
	switch guess {
	case 1, 9:
		return ErrBadRequest
	case 5:
		return ErrPageMoved
	default:
		return nil
	}
}

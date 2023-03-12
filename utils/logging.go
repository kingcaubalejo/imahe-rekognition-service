package utils

import "log"

func ErrorChecker(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

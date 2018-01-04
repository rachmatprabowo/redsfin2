package core

import (
	"log"
)

// CheckErr is function to check the error
func CheckErr(err error, msg string) bool {
	if err != nil {
		log.Fatal(err, ":", msg)
		return false
	}

	return true
}

package helpers

import (
	"log"
)

func HandleError(err error, message string) {
	if message == "" {
		message = "Error message:"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

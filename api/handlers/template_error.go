package api

import "fmt"

type templateErrorHandler func(error, string)

func handleErr(err error, context string) {
	if err != nil {
		// TODO (daniel): Replace this with a call to the logger and return an http error
		fmt.Println("Template rendering error with ", context+": ", err)
	}
}

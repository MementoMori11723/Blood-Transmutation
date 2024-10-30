package models

import "log"

func level_4() string {
	var (
		question = ""
		options  = ""
	)
	result, err := getDataFromGpt(question, options)
	if err != nil {
		log.Fatal(
			"Error occured at level 4, Error : ",
			err,
		)
	}
	return result
}

package models

import "log"

func level_1() string {
	var (
		question = ""
		options  = ""
	)
	result, err := getDataFromGpt(question, options)
	if err != nil {
		log.Fatal(
			"Error occured at level 1, Error : ",
			err,
		)
	}
	return result
}

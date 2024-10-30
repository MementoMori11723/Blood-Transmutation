package models

import "log"

func level_2() string {
	var (
		question = "" + level_1()
		options  = ""
	)
	result, err := getDataFromGpt(question, options)
	if err != nil {
		log.Fatal(
			"Error occured at level 2, Error : ",
			err,
		)
	}
	return result
}

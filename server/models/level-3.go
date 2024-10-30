package models

import "log"

func level_3() string {
	var (
		question = "" + level_2()
		options  = ""
	)
	result, err := getDataFromGpt(question, options)
	if err != nil {
		log.Fatal(
			"Error occured at level 3, Error : ",
			err,
		)
	}
	return result
}

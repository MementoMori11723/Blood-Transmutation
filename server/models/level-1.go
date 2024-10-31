package models

import "log"

func level_1(input, output string) string {
	var (
    question = "here are the two inputs : "
		options  = "You are a scientist who is working on blood conversion, and you are tasked to give all the blood combinations for the given input and output blood groups that can be possible for those two blood groups and also give their chemical name and present it in json format only, also you will only give one best combination for the given input and output blood groups."
	)
	result, err := getDataFromGpt(
    question + " blood group 1 (input): " + input + " blood group 2 (output): " + output,
    options,
  )
	if err != nil {
		log.Fatal(
			"Error occured at level 1, Error : ",
			err,
		)
	}
	return result
}

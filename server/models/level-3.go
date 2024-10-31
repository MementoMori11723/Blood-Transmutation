package models

import "log"

func level_3(input, output string) string {
	var (
    question = "here is your data : "
		options  = "you are a scientist who is working on blood conversion, and you are tasked to give an explanation of what is going on with that reaction and to validate the reaction to see if it works or not, if not then provide a better reaction, the result should be in json format."
	)
  preRes := level_2(input, output)
	result, err := getDataFromGpt(
    question + preRes,
    options,
  )
	if err != nil {
		log.Fatal(
			"Error occured at level 3, Error : ",
			err,
		)
	}
	return result
}

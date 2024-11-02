package models

import "log"

func level_1(input, output string) string {
	var (
    question = "here are the two inputs : "
		options  = "As a scientist working on blood conversion, your task is to identify a biochemical reaction that can convert a specified input blood group into a desired output blood group. Determine the type of reaction involved and list all the materials required to perform this reaction. Present your findings in JSON format only, using a structured format that includes the input and output blood groups, a description of the reaction, the reaction type (e.g., enzymatic, chemical modification), the materials required, and an explanation of why this is the best reaction for the conversion."
	)
  log.Print("Sending level 1 request to GPT")
	result, err := getDataFromGpt(
    question + " blood group 1 (input): " + input + " blood group 2 (output): " + output,
    options,
  )
  log.Print("Received level 1 response from GPT")
	if err != nil {
		log.Fatal(
			"Error occured at level 1, Error : ",
			err,
		)
	}
	return result
}

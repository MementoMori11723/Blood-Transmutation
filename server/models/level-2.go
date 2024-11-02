package models

import "log"

func level_2(input, output string) string {
	var (
    question = "Here is your data : " 
		options  = "Your task is to verify the validity of a proposed biochemical reaction for converting one blood group to another. If the reaction is not valid, provide a valid alternative reaction that achieves the desired blood conversion, including its reaction type and the materials required, presented in JSON format as in Level 1. If the reaction is valid, offer a detailed, step-by-step process on how to perform this reaction in a laboratory setting. Include safety precautions, special conditions (such as temperature and pH), and any equipment needed. Present this procedural information in JSON format, detailing each step, safety measures, special conditions, and equipment required."
	)
  log.Print("Sending level 2 request to GPT")
  preRes := level_1(input, output)
	result, err := getDataFromGpt(
    question + preRes,
    options,
  )
  log.Print("Received level 2 response from GPT")
	if err != nil {
		log.Fatal(
			"Error occured at level 2, Error : ",
			err,
		)
	}
	return result
}

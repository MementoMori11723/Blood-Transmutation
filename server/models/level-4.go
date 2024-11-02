package models

import "log"

func Level_4(input, output string) string {
	var (
    question = "Here is your data : "
    options  = "Your task is to format all the provided information about the blood conversion reaction into HTML code suitable for display on a website (inside a <div> tag). Ensure the content is user-friendly and written in simple terms so that a person with no medical knowledge can easily read and understand it. Include detailed explanations on how to conduct the reaction, the materials needed, safety precautions, instructions on how to store the blood after the reaction is completed, and information on how long the blood can be stored. Mention whether this reaction is the best method for blood conversion and, if not, explain its limitations. Present the entire content in HTML format, not in markdown, ensuring that all necessary information is clearly organized and easy to read, don't use markdown, also make sure to explain how blood works in depth using simple words for a better understanding, make each section clear and easy to read with atleast 3 paragraphs of explanation, and add atleast 10 sections."
	)
  log.Print("Sending level 4 request to GPT")
  preRes := level_3(input, output)
  result, err := getDataFromGpt(
    question + preRes, 
    options,
  )
  log.Print("Received level 4 response from GPT")
	if err != nil {
		log.Fatal(
			"Error occured at level 4, Error : ",
			err,
		)
	}
	return result
}

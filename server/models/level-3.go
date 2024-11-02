package models

import "log"

func level_3(input, output string) string {
	var (
    question = "here is your data : "
		options  = "Your task is to investigate whether anyone has previously worked on or is currently working on the biochemical reaction for the specified blood conversion. If no prior research exists, mention that as the status. If research has been conducted, provide details about the researchers involved, the time period of their studies, and summarize their findings. Include additional information to help people better understand blood and the reaction process, using simple terms suitable for someone with no medical knowledge. Present all this information in JSON format, including the research status, details of any researchers, and educational explanations about the blood conversion process."
	)
  log.Print("Sending level 3 request to GPT")
  preRes := level_2(input, output)
	result, err := getDataFromGpt(
    question + preRes,
    options,
  )
  log.Print("Received level 3 response from GPT")
	if err != nil {
		log.Fatal(
			"Error occured at level 3, Error : ",
			err,
		)
	}
	return result
}

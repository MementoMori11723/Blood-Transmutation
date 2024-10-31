package models

import "log"

func level_2(input, output string) string {
	var (
    question = "Here is your data : " 
		options  = "you are a scientist who is working on blood conversion, and you are tasked to give all the reaction that are possible for the given data, which you will use to balance the reaction and provide a result in json format."
	)
  preRes := level_1(input, output)
	result, err := getDataFromGpt(
    question + preRes,
    options,
  )
	if err != nil {
		log.Fatal(
			"Error occured at level 2, Error : ",
			err,
		)
	}
	return result
}

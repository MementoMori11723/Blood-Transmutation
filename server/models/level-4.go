package models

import "log"

func Level_4(input, output string) string {
	var (
    question = "Here is your data : "
    options  = "you are a scientist who is working on blood conversion, and you are tasked to check the status of given the reaction, if a reaction is successfully achived then give details about who has worked on that reaction and when the reaction occured, also add more information about the reaction so that people can understand better about blood, and use simple terms for better understanding and explain each topic in depth with some paragraph, also keep in mind that you are explaining this to a person with no knowledge of blood and how it works, and also menshion whather the reaction is the best way to convert blood or not and if it is not prefered they tell the limitation of that reaction as status, and show the result in html tag formate only, not in markdown formate so that it can be displayed on the website (inside of a div tag), and can you specify the details on how to conduct the reaction and what are the requirements for the reaction are and what materials do they need also give a step by step guide on how to conduct the reaction and what are the safety precautions that need to be taken while conducting the reaction, and also give the details on how to store the blood after the reaction is completed and how long can the blood be stored"
	)
  preRes := level_3(input, output)
  result, err := getDataFromGpt(
    question + preRes, 
    options,
  )
	if err != nil {
		log.Fatal(
			"Error occured at level 4, Error : ",
			err,
		)
	}
	return result
}

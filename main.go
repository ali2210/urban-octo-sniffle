package main

import (
  ."github.com/ali2210/urban-octo-sniffle/selector"
  "cloud.google.com/go/bigquery"
  "log"
  "fmt"
)


var(
	historyObject BigQuery = NewObject()
)
func main() {
	client := historyObject.NewRequestObject()
	
	fmt.Println("Client:", client)
	
	historyObject.StopClient(client)
	var schema bigquery.Schema
	history := []*bigquery.StructSaver{
		{Struct: History{Biography : "Great Akbar & Jodha", Location: "India"}, Schema: schema, InsertID:"1",},
		{Struct: History{Biography: "Alexender & Genius", Location : "Rome",}, Schema : schema, InsertID:"2",},
		{Struct: History{Biography: "Shah Jahan & First child", Location : "India",}, Schema: schema, InsertID:"3",},
	}
	
	err := historyObject.AddGoldernPapers(history, client); if err != nil {
		log.Fatalln("Something went wrong", err.Error())
		return 
	}
}
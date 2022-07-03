package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Meteroids struct{
	Name 			string			`json:"name"`
	ID      		string          `json:"id"`
	Name_Type 		string        	`json:"nametype"`
	Reclass   		string         	`json:"reclass`
	Mass      		string         	`json:"mass"`
	Fall      		string         	`json:"fall"`
	Year      		string         	`json:"year"`
	Reclat    		string         	`json:"reclat"`
	Reclong   		string         	`json:"reclong"`
}
func main (){
	content,err :=ioutil.ReadFile("meteroids.json")
	if err !=nil {
		fmt.Println(err.Error())
	}
	var meteroids []Meteroids
	err2 :=json.Unmarshal(content,&meteroids)
	if err2 !=nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	for  _,x :=range meteroids{
		fmt.Printf("Name: %s\n",x.Name)
		fmt.Printf("Mass: %s\n",x.Mass)
		fmt.Printf("Year: %s\n",x.Year)
	}
}
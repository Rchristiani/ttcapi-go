package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

//Create a structure for the data to live in
type TTCData struct {
	Time int
	Uri string
	Name string
}



func main() {
	// Probably doing too much in the main func, should have own function!

	//Make the call an get the response and the error
	apiCallRes, apiCallErr := http.Get("https://myttc.ca/finch_station.json")

	//Check if error exists
	if apiCallErr != nil {
		//Probably a better way to handle error
		fmt.Println(apiCallErr)
	} 
	// Make sure to defer the closing of the body from the response till after this main function returns
	defer apiCallRes.Body.Close()
	// Make new interface(proper term?)
	myData := new(TTCData)
	// Use the json package to decode the body and store it in the myData var
	// & indicates that this is a pointer.
	json.NewDecoder(apiCallRes.Body).Decode(&myData)


	// How can I create a general container in go? Like simple object
	// I don't want to have to define the structure, or is that just a thing you have to do?
	var generalContainer interface{}
	// Take apiCallRes.Body and return the []byte info for it
	ttcByte, ttcByteErr := ioutil.ReadAll(apiCallRes.Body)

	// Check for error
	if ttcByteErr != nil { 
		fmt.Println(ttcByteErr)
	}
	//Then try and unmarshal json into it
	//What does that even mean?
	json.Unmarshal(ttcByte,&generalContainer)
	//myData prints some data, generalContainer <nil> ?~!?!?
	fmt.Println(myData,generalContainer)

}
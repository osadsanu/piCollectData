package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//device contains the structure of the requested json
type device struct {
	Id       string `json:"id"`
	Status   string	`json:"status"`
	ActionID string  `json:"actionID"`
}

func main() {
	fmt.Println("hello")
	// list of ip addres ( we assume all the ips are static on the network)
	ipAddress := []string{
		//"192.168.0.15",
		//"192.168.0.16",
		//"192.168.0.17",
		//"192.168.0.18",
		"http://127.0.0.1:8081",
		"http://127.0.0.1:8082",
		"http://127.0.0.1:8083",
		"http://127.0.0.1:8084",
	}
	devices := make(map[string]device)
	//Infinite loop with a sleep
	for { 
		for _, ip := range ipAddress { //for each address is going to create a get request
			url := ip
			client := http.Client{
				Timeout: time.Millisecond * 1000, // Maximum of 2 secs
			}
			req, err := http.NewRequest(http.MethodGet, url, nil)
			//if err != nil {
			//	log.Fatal(err)
				//log.(err)
			//}
			res, getErr := client.Do(req)
			if getErr != nil {
				log.Fatal(getErr)
			}
			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
			//create the container object in order to store the data
			obj := &device{}
			jsonErr := json.Unmarshal(body, &obj)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			//checks if the action is different from the last requested action
			if devices[ip].ActionID != obj.ActionID{
				fmt.Println("Device: ",ip," just change the status to: ", obj.Status)
				devices[ip] = *obj
				//TODO
				//Generates a HTTP request to the server in order to Tell what is the new status.
			}else {
				
				//fmt.Println("Device: ",ip," has the same status", obj.Status)
			}

			//fmt.Println(obj)
		}
		//fmt.Println(devices)
		
		time.Sleep(500 * time.Millisecond)
	}
}

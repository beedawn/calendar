package main

import (
	"backend/calendar"

	"fmt"
	"time"
	"strings"
	"strconv"
)

func main() {

	//var cal calendar.Calendar
	var username, password string
	/*newR := cal.NewResource()
	startTime := time.Now()
	endTime := time.Now()
	newE, _ := newR.NewEvent(startTime, endTime)
	*/
	var calList = make([]calendar.Calendar,0)
	fmt.Println("Enter username:")
	fmt.Scan(&username)
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	if username == "test" && password == "test" {
		var input string
		for strings.ToLower(input) != "q" {
			fmt.Printf("1. Calendar\n")
			fmt.Println("Enter a selection, q to quit:")
			fmt.Scan(&input)
			if input == "1" {

				var calInput string

				for calInput != "q" {
					fmt.Println("1 Create new Calendar")
					fmt.Println("2 View a Calendar")
					fmt.Println("3 Update a resource")
					fmt.Println("4 Delete a resource")

					fmt.Scan(&calInput)
					switch calInput {
					case "1":
						cal := CreateResource(calList)
						calList = append(calList, cal)
					case "2":
						ViewResource(calList)
					case "3":
						UpdateResource()
					case "4":
						calList =	DeleteResource(calList)	
					}
				}
			}
		}
	}

	//	newR.DeleteEvent(*newE)

}


func CreateResource(c []calendar.Calendar) calendar.Calendar{
	fmt.Println("Create Resource!")
	var cal calendar.Calendar

	if len(c) == 0 {


	cal.Id = 1
	}

	if len(c) != 0 {
	cal.Id = c[len(c)-1].Id+1	
	}
	newR := cal.NewResource()
	startTime := time.Now()
	endTime := time.Now()
	newR.NewEvent(startTime, endTime)
	return cal
}

func ViewResource(cal []calendar.Calendar){
	fmt.Println("View Resource!")

	for _, item := range cal{
		fmt.Println(item)
		for _, resource := range item.Resource{
			fmt.Println(resource.Id)
		}
	}
}

func UpdateResource(){
	fmt.Println("Update Resource!")
}

func DeleteResource(c []calendar.Calendar) []calendar.Calendar{
	fmt.Println("Delete Resource!")
	var delInput, confirm string

		fmt.Println("Which Calendar should be deleted?")

		fmt.Scan(&delInput)
		fmt.Println(delInput+" will be deleted, are you sure?")
		fmt.Scan(&confirm)
		if confirm == "y" {
			fmt.Println("Delete")
		}
		num ,err := strconv.Atoi(delInput)
		if err != nil {
		return c
	}
	c = append(c[:num-1],c[num:]...)
	
	return c

}

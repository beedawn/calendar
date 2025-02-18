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
						fmt.Println("***************")
						fmt.Println("Creating calendar...")
						cal := CreateResource(calList)
						calList = append(calList, cal)
						fmt.Printf("Calendar %d, created!\n", cal.Id)
					case "2":
						ViewResource(calList)
					case "3":
						UpdateResource(&calList)
					case "4":
						calList =	DeleteResource(calList)	
					}
				}
			}
		}
	}

	//	newR.DeleteEvent(*newE)

}

//this also creates a calendar, i think a "CalendarManager" or similar abstraction is needed to hold all the calendars and move this functionality into something more useful than th emain file
func CreateResource(c []calendar.Calendar) calendar.Calendar{
	var cal calendar.Calendar

	if len(c) == 0 {


	cal.Id = 1
	}

	if len(c) != 0 {
	cal.Id = c[len(c)-1].Id+1	
	}
	newR, _ := cal.NewResource()
	startTime := time.Now()
	endTime := time.Now()
	newR.NewEvent(startTime, endTime)
	return cal
}

func ViewResource(cal []calendar.Calendar){

	if len(cal) ==0 {
		fmt.Println("No calendars!")
	} else{
	fmt.Println("Here is a list of all calendars:")
	for _, c := range cal{
			fmt.Printf("Calendar ID: %d\n",c.Id)
			for _, r := range c.Resource{
			fmt.Printf("Resource ID: %d\n", r.Id)
				for _, e := range r.Event {
				fmt.Printf("Event ID: %d\n", e.Id)
				}
			}
		}

	}
}

func UpdateResource(cal *[]calendar.Calendar){
	fmt.Println("Update Resource!")

		var openCal string
		fmt.Println("Which calendar would you like to open?")

		for i, c := range *cal{
		if i == len(*cal)-1{
			fmt.Printf("%d\n", c.Id)
			break
		}

		fmt.Printf("%d, ", c.Id)
	}
		fmt.Scan(&openCal)

	for i, c := range *cal{
	strId := strconv.Itoa(c.Id)
	if strId == openCal{
			fmt.Println(c)

		for _, resource := range c.Resource {
		fmt.Println(resource.Id)

		}
				var calR string
				fmt.Println("Which resource would you like to edit?")
				fmt.Scan(&calR)

		for j, resource := range c.Resource {
					rId:= strconv.Itoa(resource.Id)
					if calR == rId {
				var calEdit string
				fmt.Println("What would you like to do next?\n1: Add Event")
				fmt.Scan(&calEdit)
				
				if calEdit == "1" {
					fmt.Println("hello")
					newE, err := resource.NewEvent(time.Now(), time.Now())
					if err != nil {
						break
							}
					(*cal)[i].Resource[j].Event= append(resource.Event, *newE)
					fmt.Println(resource.Event)
				}

					}

		}
				fmt.Println(cal)

			}

		}
}

func DeleteCalendar(c []calendar.Calendar) []calendar.Calendar{
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
//	c = append(c[:num-1],c[num:]...)
	
	return c

}

package main

import (

	"backend/calendarmanager"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	var username, password string
	var calManager calendarmanager.CalendarManager


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
					fmt.Println("3 Update a Calendar")
					fmt.Println("4 Delete a Calendar")

					fmt.Scan(&calInput)
					switch calInput {
					case "1":
						fmt.Println("***************")
						fmt.Println("Creating calendar...")
						cal := calManager.NewCalendar()
						fmt.Printf("Calendar %d, created!\n", cal.Id)
					case "2":
						ViewResource(calManager)
					case "3":
						UpdateResource(&calManager)
						fmt.Println("Update Resource")
					case "4":
						DeleteCalendar(&calManager)

					}
				}
			}
		}
	}



}



func ViewResource(cm calendarmanager.CalendarManager) {

	if len(cm.Calendars) == 0 {
		fmt.Println("No calendars!")
	} else {
		fmt.Println("Here is a list of all calendars:")
		for _, c := range cm.Calendars {
			fmt.Printf("Calendar ID: %d\n", c.Id)
			for _, r := range c.Resource {
				fmt.Printf("Resource ID: %d\n", r.Id)
				for _, e := range r.Event {
					fmt.Printf("Event ID: %d\n", e.Id)
				}
			}
		}

	}
}

func UpdateResource(cm *calendarmanager.CalendarManager) {
	fmt.Println("Update Resource!")

	var openCal string
	fmt.Println("Which calendar would you like to open?")

	for i, c := range cm.Calendars {
		if i == len(cm.Calendars)-1 {
			fmt.Printf("%d\n", c.Id)
			break
		}

		fmt.Printf("%d, ", c.Id)
	}
	fmt.Scan(&openCal)

	for i, c := range cm.Calendars {
		strId := strconv.Itoa(c.Id)
		if strId == openCal {
			fmt.Println(c)

			for _, resource := range c.Resource {
				fmt.Println(resource.Id)

			}
			var calR string
			fmt.Println("Which resource would you like to edit?")
			fmt.Scan(&calR)
			//doesn't edit anything because theres no way to add resources...yet

			if len(c.Resource)==0{
			fmt.Println("nothin!")
			}
			for j, resource := range c.Resource {
				rId := strconv.Itoa(resource.Id)
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
						(cm.Calendars)[i].Resource[j].Event = append(resource.Event, *newE)
						fmt.Println(resource.Event)
					}

				}

			}
		
		}

	}
}

func DeleteCalendar(cm *calendarmanager.CalendarManager) {
	fmt.Println("Delete Resource!")
	var delInput, confirm string

	fmt.Println("Which Calendar should be deleted?")

	fmt.Scan(&delInput)
	fmt.Println(delInput + " will be deleted, are you sure?")
	fmt.Scan(&confirm)
	if confirm == "y" {
		fmt.Println("Delete")
	}
	num, err := strconv.Atoi(delInput)
	if err != nil {
		return
	}
	for _, cal := range cm.Calendars {
		if cal.Id == num {
			cm.DeleteCalendar(cal)
		}

	}

}

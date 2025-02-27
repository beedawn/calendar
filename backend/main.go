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
					fmt.Println("2 View All Calendars")
					fmt.Println("3 Open a Calendar")
					fmt.Println("4 Delete a Calendar")
					fmt.Println("5 Create new Resource")

					fmt.Scan(&calInput)
					switch calInput {
					case "1":
						CreateCalendar(&calManager)	
					case "2":
						ViewResource(calManager)
					case "3":
						OpenCalendar(&calManager)
					case "4":
						DeleteCalendar(&calManager)
					case "5":
						CreateResource(&calManager)

					}
				}
			}
		}
	}

}

func CreateCalendar(cm *calendarmanager.CalendarManager){
	cal := cm.NewCalendar()
	fmt.Printf("Calendar %d, created!\n", cal.Id)
}

func CreateResource(cm *calendarmanager.CalendarManager) {

	fmt.Println("Create a resource")
	cm.NewResource()
	for _, r := range cm.Resources {
		fmt.Printf("Resource Id: %d \n", r.Id)
	}
}

func ViewResource(cm calendarmanager.CalendarManager) {

	if len(cm.Calendars) == 0 {
		fmt.Println("No calendars!")
	} else {
		fmt.Println("Here is a list of all resources and calendars and events:")

		for _, r := range cm.Resources {
			fmt.Printf("Resource ID: %d\n", r.Id)

		}
		for i, c := range cm.Calendars {
			fmt.Printf("Calendar ID: %d\n", c.Id)
			for _, e := range cm.Calendars[i].Events {
				fmt.Printf("Event ID: %d\n", e.Id)
			}

		}

	}
}

func OpenCalendar(cm *calendarmanager.CalendarManager) {
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

			var calR, modStr string

			for modStr != "q" {

				fmt.Println("What would you like to do?")
				fmt.Println("1. Create Event")
				fmt.Println("2. Edit Resource")
				fmt.Println("3. Edit Event")
				fmt.Println("4. Delete Event")
				fmt.Scan(&modStr)

				switch modStr {
					case "1":
					fmt.Println("Create an event")
					var resourceId string
					fmt.Println("Which resource should it be added to?")
					fmt.Scan(&resourceId)

					for _, r := range cm.Resources {
						rId := strconv.Itoa(r.Id)

						if rId == resourceId {
							fmt.Println("match found")
							cm.NewEvent(time.Now(), time.Now(), &cm.Calendars[i], r)

						}
					}

					for _, e := range cm.Calendars[i].Events {
						fmt.Printf("Event ID: %d \n", e.Id)
					}

				case "2":
					fmt.Println("Which resource would you like to edit?")
					fmt.Scan(&calR)
					//doesn't edit anything because theres no way to add resources...yet

					if len(cm.Resources) == 0 {
						fmt.Println("nothin!")
					}
					for j, resource := range cm.Resources {
						rId := strconv.Itoa(resource.Id)
						if calR == rId {
							var calEdit string
							fmt.Println("What would you like to do next?\n1: Add Event")
							fmt.Scan(&calEdit)

							if calEdit == "1" {
								fmt.Println("hello")
								newE, err := cm.NewEvent(time.Now(), time.Now(), &cm.Calendars[j], resource)
								if err != nil {
									break
								}
								cm.Calendars[j].Events = append(cm.Calendars[j].Events, *newE)
								fmt.Println(cm.Calendars[j].Events)
							}

						}

					}
				case "3":
					fmt.Println("Edit Event")
				case "4":
					fmt.Println("Delete Event")
					fmt.Println("Enter the ID of the event you'd like to delete:")
					var delEvent string
					fmt.Scan(&delEvent)
					

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

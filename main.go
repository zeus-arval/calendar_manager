package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	DATE_PATTERN = "^20\\d{2}-([2-9]|1[0-2]?)-([0][1-9]|[1-2][0-9]|3[0-1])$"
	TIME_PATTERN = "^([0-1][0-9]|[2][0-3]):[0-5][0-9]$"
	PRIMARY = "primary"
	OPTIONS_COUNT = 5
)

type Option int
const (
	CREATE_EVENT = 1
	SHOW_EVENTS  = 2
	UPDATE_EVENT = 3
	REMOVE_EVENT = 4
	Exit         = 5
)

func fillOptions(options *[OPTIONS_COUNT]Option) {
	options[0] = CREATE_EVENT
	options[1] = SHOW_EVENTS
	options[2] = UPDATE_EVENT
	options[3] = REMOVE_EVENT
	options[4] = Exit
}

func fillOptionsArray(optionsArray *[OPTIONS_COUNT]string) {
	optionsArray[0] = "Create Event"
	optionsArray[1] = "Show Event"
	optionsArray[2] = "Update Events"
	optionsArray[3] = "Remove Event"
	optionsArray[4] = "Exit"
}

func clearScreen() {
	fmt.Print("\n\n\n\n\n\n\n\n")
}

func checkByRegex(value *string, pattern string) bool{
	match, _ := regexp.MatchString(pattern, *value)
	return match
}

func printOptionsList(options *[OPTIONS_COUNT]Option, optionsArray *[OPTIONS_COUNT]string) {
	for index, _ := range options {
		fmt.Printf("%v -> %v\n", index + 1, optionsArray[index])
	}
}

func scan() (string, error){
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func scanEventInput(value *string, description string) {
	fmt.Printf("Enter a %v for event\n", description)

	input, err := scan()
	if err != nil {
		log.Fatalf("Unable to read %v: %v", description, err)
	}
	*value = strings.Replace(input,  "\n", "", -1)
}

func scanFormattedEventInput(value *string, description string, format *string, pattern string){
	fmt.Printf("Enter a %v for event in format -> %v:\n", description, *format)
	isValueCorrect := false
	for !isValueCorrect {
		input, err := scan()
		if err != nil {
			log.Fatalf("Unable to read %v: %v", &description, err)
		}
		*value = strings.Replace(input,  "\n", "", -1)
		if checkByRegex(value, pattern){
			isValueCorrect = true
		} else {
			fmt.Println("Invalid format, try again!")
		}
	}
}

func createEvent(srv *calendar.Service) {
	var summary string
	var location string
	var description string
	var date string
	var startTime string
	var endTime string

	seconds := ":00"
	dateFormat := "YYYY-MM-DD"
	timeFormat := "HH:MM"

	clearScreen()

	scanEventInput(&summary, "summary")
	scanEventInput(&location, "location")
	scanEventInput(&description, "description")
	scanFormattedEventInput(&date, "start date", &dateFormat, DATE_PATTERN)
	scanFormattedEventInput(&startTime, "start time", &timeFormat, TIME_PATTERN)
	scanFormattedEventInput(&endTime, "end time", &timeFormat, TIME_PATTERN)

	event := &calendar.Event{
		Summary:     summary,
		Location:    location,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: date + "T" + startTime + seconds,
			TimeZone: "Europe/Tallinn",
		},
		End: &calendar.EventDateTime{
			DateTime: date + "T" + endTime + seconds,
			TimeZone: "Europe/Tallinn",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
	}

	event, err := srv.Events.Insert(PRIMARY, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
}

func showEvents(srv *calendar.Service) {
	t := time.Now().Format(time.RFC3339)
	clearScreen()

	events, err := srv.Events.List(PRIMARY).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("ID: [%v] -> %v (%v)\n", item.Id, item.Summary, date)
		}
	}
}

func checkAndModifyData(value *string, dataName string){
	fmt.Printf("Do you want to change %v -> [%v]? (yes/no)\n", dataName, *value)
	var input string
	var err error

	isInputCorrect := false
	for isInputCorrect {
		input, err = scan()
		input = strings.Replace(input, "\n", "", -1)
		if err != nil{
			log.Fatal(err)
		}
		if input == "yes" && input == "no"{
			isInputCorrect = true
		}
	}
	isInputCorrect = false
	if input == "yes"{
		for isInputCorrect {
			input, err := scan()
			if err != nil {
				log.Fatal(err)
			}
			if strings.Contains(dataName, "time") {
				if checkByRegex(&input, TIME_PATTERN) {
					isInputCorrect = true
					*value = input
				}
			} else if strings.Contains(dataName, "date") {
				if checkByRegex(&input, DATE_PATTERN) {
					isInputCorrect = true
					*value = input
				}
			} else {
				*value = input
			}
		}
	}else if input == "no"{
		return
	}
}

func modifyEventDatas(event *calendar.Event) calendar.Event{
	checkAndModifyData(&event.Summary, "string")
	checkAndModifyData(&event.Location, "string")
	checkAndModifyData(&event.Description, "string")
	checkAndModifyData(&event.Start.Date, "date")
	checkAndModifyData(&event.Start.DateTime, "start time")
	checkAndModifyData(&event.End.DateTime, "end time")
	return *event
}

func updateEvent(srv *calendar.Service){
	showEvents(srv)
	fmt.Println("\nChoose an event' ID for update")
	input, err := scan()
	if err != nil{
		fmt.Println("Input is incorrect")
		return
	}
	eventId := strings.Replace(input, "\n", "", -1)
	event, err := srv.Events.Get(PRIMARY, eventId).Do()
	if err != nil{
		log.Fatal(err)
	}
	*event = modifyEventDatas(event)

	_, err = srv.Events.Update(PRIMARY, eventId, event).Do()
	if err != nil{
		log.Fatal(err)
	}
}

func removeEvent(srv *calendar.Service){
	showEvents(srv)
	fmt.Println("Enter ID of an event you want to delete")
	input, err := scan()
	if err != nil{
		fmt.Println("Input is incorrect")
		return
	}

	eventId := strings.Replace(input, "\n", "", -1)
	err = srv.Events.Delete(PRIMARY, eventId).Do()
	if err != nil{
		fmt.Printf("Unable to delete event with ID: %v", eventId)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	err := browser.OpenURL(authURL)
	if err != nil {
		return nil
	}

	fmt.Println("Enter a token given in browser:")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	var options [OPTIONS_COUNT]Option
	var optionsArray [OPTIONS_COUNT]string

	fillOptions(&options)
	fillOptionsArray(&optionsArray)
	clearScreen()

	ctx := context.Background()

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	var userInput string
	for userInput != strconv.Itoa(Exit) {
		printOptionsList(&options, &optionsArray)

		if _, err := fmt.Scan(&userInput); err != nil {
			log.Fatalf("There is no option you have choosen: %v", err)
		}

		switch userInput {
		case strconv.Itoa(CREATE_EVENT):
			createEvent(srv)
		case strconv.Itoa(SHOW_EVENTS):
			showEvents(srv)
		case strconv.Itoa(UPDATE_EVENT):
			updateEvent(srv)
		case strconv.Itoa(REMOVE_EVENT):
			removeEvent(srv)
		}
	}
}

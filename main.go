package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	screen "github.com/inancgumus/screen"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const OPTIONS_COUNT = 5
const PRIMARY = "primary"

type Option int

const (
	CREATE_EVENT = 0
	SHOW_EVENTS  = 1
	UPDATE_EVENT = 2
	REMOVE_EVENT = 3
	Exit         = 4
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

func clearConsole() {
	screen.Clear()
}

func printOptionsList(options *[OPTIONS_COUNT]Option, optionsArray *[OPTIONS_COUNT]string) {
	for index, _ := range options {
		fmt.Printf("%v -> %v", index, optionsArray[index])
	}
}

func scanInput(value *string, description string) {
	fmt.Printf("Enter a %v for event\n", description)
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatalf("Unable to read %v: %v", description, err)
	}
}

func scanFormattedInput(value *string, description string, format string){
	fmt.Printf("Enter a %v for event in format -> %v:\n", description, format)
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatalf("Unable to read %v: %v", description, err)
	}
}

func createEvent(srv *calendar.Service) {
	var summary string
	var location string
	var description string
	var startDate string
	var startTime string
	var endDate string
	var endTime string

	var dateFormat string = "YYYY-MM-DD"
	var timeFormat string = "HH:MM:SS"

	scanInput(&summary, "summary")
	scanInput(&location, "location")
	scanInput(&description, "description")
	scanFormattedInput(&startDate, "start date", dateFormat)
	scanFormattedInput(&startTime, "start date", timeFormat)
	scanFormattedInput(&endDate, "start date", dateFormat)
	scanFormattedInput(&endTime, "start date", timeFormat)

	event := &calendar.Event{
		Summary:     summary,
		Location:    location,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startDate + "T" + startTime,
			TimeZone: "Europe/Tallinn",
		},
		End: &calendar.EventDateTime{
			DateTime: endDate + "T" + endTime,
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
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}

func updateEvent(srv *calendar.Service){

}

func removeEvent(srv *calendar.Service){
	showEvents(srv)

	//event, err := srv.Events.Delete(PRIMARY, )
	//if err != nil {
	//	log.Fatalf("Unable to create event. %v\n", err)
	//}
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
	browser.OpenURL(authURL)

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
	clearConsole()
	printOptionsList(&options, &optionsArray)

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

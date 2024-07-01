package main

import (
	"context" //Used to manage the context of the application, particularly for cancellation and timeouts.
	"fmt"     //Provides formatted I/O functions.
	"log"     //Provides logging functions.
	"os"      //Provides functions to access environment variables and OS functionality.
	"strings" //Provides functions to manipulate strings.
	"sync"    //Provides functions to manage concurrenc and Provides synchronization primitives such as mutexes.

	"github.com/shomali11/slacker" // An external package used to create Slack bots.
)

// StandupManager struct to manage standup meeting
type StandupManager struct {
	sync.Mutex
	employees   []string
	attendance  map[string]bool
	responses   map[string][]string
	questions   []string
	currentStep map[string]int
	state       string
}

// constructor for StandupManager
func NewStandupManager() *StandupManager {
	return &StandupManager{
		attendance:  make(map[string]bool),
		responses:   make(map[string][]string),
		questions:   []string{"What did you do last working day?", "What will you do today?", "Are there any impediments in your way?", "What is your availability for a new project/task?"},
		currentStep: make(map[string]int),
		state:       "idle",
	}
}

// printCommandEvents function to print command events
// func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
// 	for event := range analyticsChannel {
// 		fmt.Println("Command Events:")
// 		fmt.Println(event.Command)
// 		fmt.Println(event.Timestamp)
// 		fmt.Println(event.Parameters)
// 		fmt.Println(event.Event)
// 		fmt.Println()
// 	}
// }

// Starts a new standup meeting by initializing employee lists, setting attendance to false, and resetting responses and current steps. Changes the state to "attendance".
func (sm *StandupManager) StartStandup(botCtx slacker.BotContext, response slacker.ResponseWriter, employees []string) {
	sm.Lock()
	defer sm.Unlock()

	if sm.state != "idle" {
		response.Reply("A standup meeting is already in progress.")
		return
	}

	sm.employees = employees
	for _, employee := range employees {
		sm.attendance[employee] = false
		sm.responses[employee] = make([]string, 0)
		sm.currentStep[employee] = 0
	}

	sm.state = "attendance"
	response.Reply("Standup meeting started. Employees, please mark your attendance.")
}

// Marks an employee as present and asks the next question if the standup is in the attendance state.
func (sm *StandupManager) MarkAttendance(botCtx slacker.BotContext, response slacker.ResponseWriter, employee string) {
	sm.Lock()         //Locks the mutex, sm. If the lock is already in use, call goroutine blocks until the mutex is available.
	defer sm.Unlock() //Unlocks the mutex, m. A run-time error is thrown if m is not already locked.

	if sm.state != "attendance" {
		response.Reply("Attendance is not being taken at the moment.")
		return
	}

	if _, exists := sm.attendance[employee]; !exists {
		response.Reply(fmt.Sprintf("Employee %s is not on the list.", employee))
		return
	}

	sm.attendance[employee] = true
	response.Reply(fmt.Sprintf("Marked %s as present.", employee))
	sm.askNextQuestion(botCtx, response, employee)
}

// Asks the next question to the employee based on their current step.
func (sm *StandupManager) askNextQuestion(botCtx slacker.BotContext, response slacker.ResponseWriter, employee string) {
	step := sm.currentStep[employee]
	if step < len(sm.questions) {
		response.Reply(fmt.Sprintf("%s, %s", employee, sm.questions[step]))
		sm.currentStep[employee]++
	} else {
		response.Reply(fmt.Sprintf("%s, you have completed the standup questions.", employee))
	}
}

// Handles responses from employees and asks the next question.
func (sm *StandupManager) handleResponse(botCtx slacker.BotContext, response slacker.ResponseWriter, employee string, answer string) {
	sm.Lock()
	defer sm.Unlock()

	if _, exists := sm.responses[employee]; !exists || sm.currentStep[employee] == 0 {
		response.Reply(fmt.Sprintf("%s, you need to mark your attendance first.", employee))
		return
	}

	sm.responses[employee] = append(sm.responses[employee], answer)
	sm.askNextQuestion(botCtx, response, employee)
}

// Resets the standup manager and sets the state to "idle".
func (sm *StandupManager) ExitStandup(botCtx slacker.BotContext, response slacker.ResponseWriter) {
	sm.Lock()
	defer sm.Unlock()

	sm.employees = []string{}
	sm.attendance = make(map[string]bool)
	sm.responses = make(map[string][]string)
	sm.currentStep = make(map[string]int)
	sm.state = "idle"
	response.Reply("Standup meeting exited. You can start a new one now.")
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	// Create a new StandupManager instance
	manager := NewStandupManager()

	//go printCommandEvents(bot.CommandEvents())

	// Define the commands for the bot

	// Start standup command
	bot.Command("start standup <employees>", &slacker.CommandDefinition{
		Description: "Start the standup meeting and list employees (comma-separated)",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			employees := request.Param("employees")
			manager.StartStandup(botCtx, response, splitEmployees(employees))
		},
	})

	// Mark attendance command
	bot.Command("mark attendance <employee>", &slacker.CommandDefinition{
		Description: "Mark an employee as present",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			employee := request.Param("employee")
			manager.MarkAttendance(botCtx, response, employee)
		},
	})

	// Response command
	bot.Command("response <employee> <answer>", &slacker.CommandDefinition{
		Description: "Respond to the current question",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			employee := request.Param("employee")
			answer := request.Param("answer")
			manager.handleResponse(botCtx, response, employee, answer)
		},
	})

	// Exit standup command
	bot.Command("exit standup", &slacker.CommandDefinition{
		Description: "Exit the current standup meeting",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			manager.ExitStandup(botCtx, response)
		},
	})

	// Start the bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// splitEmployees function to split employees
func splitEmployees(employees string) []string {
	return strings.Split(employees, ",")
}

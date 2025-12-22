package main

// Key requirements
// OnMount - fetch existing todos from a file if it exists - and load them into ram through a struct
// allow continues modification of the list - the full crud
// exit the program on a special keystroke

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DataFileName = "db.txt"
const DelimiterString = "===========Task=========="

var reader = bufio.NewScanner(os.Stdin)
var tasks []Task = []Task{}

func main() {

	for {
		displayMenu()
		fmt.Print("Enter a command ")
		reader.Scan()
		input := reader.Text()
		app(input)
	}
}

func kill() {
	fmt.Print("Goodbye!")
	os.Exit(0)
}
func app(input string) {
	switch input {
	case "kill", "5":
		kill()
		break
	case "menu":
		displayMenu()
		break
	case "1":
		view()
		break
	case "2":
		createForm()

	default:
		displayMenu()

	}

}

func displayMenu() {
	fmt.Println("============================================================")
	fmt.Println("============================================================")
	fmt.Println("1. View tasks")
	fmt.Println("2. Create a task")
	fmt.Println("3. Edit a task")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Kill")
	fmt.Println("============================================================")
	fmt.Println("============================================================")
}
func view() {
	tasks := readFile()
	printTasks(tasks)

}

func readFile() []Task {
	b, err := os.ReadFile(DataFileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found - creating a new one")
			_, err := os.Create(DataFileName)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}

	}

	fileContent := string(b)
	tasks := parseFileContent(fileContent)

	return tasks
}

func parseFileContent(c string) []Task {
	t := strings.Split(c, DelimiterString)
	for i, singleTask := range t {
		if len(singleTask) == 0 {
			continue
		}

		singleTask = strings.TrimSpace(singleTask)
		//fmt.Println(singleTask)
		taskTitleAndContent := strings.Split(singleTask, "\n")
		task := Task{
			id:      i,
			title:   strings.TrimSpace(strings.ReplaceAll(taskTitleAndContent[0], "title:", "")),
			content: strings.TrimSpace(strings.ReplaceAll(taskTitleAndContent[1], "content:", "")),
		}
		tasks = append(tasks, task)
	}

	return ts
}

func printTasks(t []Task) {
	fmt.Println("============TasksList===============")

	for _, singleTask := range t {
		fmt.Println("=============================")
		fmt.Println(singleTask.id)
		fmt.Println(singleTask.title)
		fmt.Println(singleTask.content)
		fmt.Println("=============================")
	}

	displayMenu()
}

func createForm() {
	task := Task{}
	fmt.Println("==========Create Form=========")
	fmt.Println("1. Name the task")
	reader.Scan()
	title := reader.Text()
	task.title = title
	fmt.Println("2. Content the task")
	reader.Scan()
	content := reader.Text()
	task.content = content

}

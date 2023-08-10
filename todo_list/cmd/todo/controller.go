package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/todo_list/internal/core/domain/enums"
	"github.com/todo_list/internal/handlers"
)

func PrintMenu() {
	fmt.Println()
	fmt.Println("Please choose your option")
	fmt.Println("Press 1: To create a new task")
	fmt.Println("Press 2: To modify a task")
	fmt.Println("Press 3: To complete a task")
	fmt.Println("Press 4: To print all the task")
	fmt.Println("Press 5: To print a Task")
	fmt.Println("Press 6: To get activity logs between time ranges ")
	fmt.Println("Press 7: To get task stats between time ranges")

}

func CreateTaskHanlder(scanner *bufio.Scanner) (desc string, endDate string, createdBy int64, tag enums.Tags) {
	fmt.Println(" Please enter description of the task")
	scanner.Scan()
	desc = scanner.Text()
	fmt.Println(" please enter completion date of the task in format dd-mm-yyyy")
	scanner.Scan()
	endDate = scanner.Text()
	fmt.Println(" Please enter date of created by ")
	scanner.Scan()
	createdByStr := scanner.Text()
	createdBy, _ = strconv.ParseInt(createdByStr, 10, 64)
	fmt.Println(" do you want to add tags in the task(y/n)")
	scanner.Scan()
	addTag := scanner.Text()
	if strings.Compare(addTag, "y") == 0 {
		fmt.Printf(" Press 1 to mark as %s tag \n", enums.Work.String())
		fmt.Printf(" Press 2 to mark as %s tag \n", enums.Casual.String())
		fmt.Printf(" Press 3 to mark as %s tag \n", enums.HighPriority.String())
		scanner.Scan()
		tagStr := scanner.Text()
		tagInt, _ := strconv.ParseInt(tagStr, 10, 64)
		tag = enums.Tags(tagInt)
	}
	return desc, endDate, createdBy, tag
}

func UpdateOptions(scanner *bufio.Scanner) (int, string) {
	var (
		option    int
		updateStr string
	)
	fmt.Println("***********************************************")
	fmt.Println(" Press 1 to update description")
	fmt.Println(" Press 2 to update dueDate")
	fmt.Println(" Press 3 to update tag ")
	scanner.Scan()
	optionTxt := scanner.Text()
	option, _ = strconv.Atoi(optionTxt)
	fmt.Println(" Enter updated value")
	scanner.Scan()
	updateStr = scanner.Text()
	return option, updateStr

}
func PrintTask(response handlers.TaskRespBody) {
	fmt.Println("***********************************************")
	fmt.Println("Task id = ", response.Id)
	fmt.Println("Description = ", response.Desc)
	fmt.Println("Due Date =  ", response.DueDate)
	fmt.Println("Status =  ", response.Status)
	fmt.Println("Tag = ", response.Tag)
	fmt.Println("***********************************************")

}

func GetTaskId(scanner *bufio.Scanner) (id int64) {
	fmt.Println("***********************************************")
	fmt.Println("Please enter task id")
	scanner.Scan()
	taskIdStr := scanner.Text()
	id, _ = strconv.ParseInt(taskIdStr, 10, 64)
	return id
}

func GetTimeRange(scanner *bufio.Scanner) (string, string) {

	fmt.Println("Enter first time range in format dd-mm-yyyy hh:mm:ss")
	scanner.Scan()
	rangeOne := scanner.Text()
	fmt.Println("Enter second time range in format dd-mm-yyyy hh:mm:ss")
	scanner.Scan()
	rangeTwo := scanner.Text()
	return rangeOne, rangeTwo
}

func PrintActivityLog(logs []handlers.ActivityLog) {
	fmt.Println("************************* Activity Log ***************************")
	if len(logs) == 0 {
		fmt.Println(" No activity logs found in this range")
		return
	}
	for _, curLog := range logs {
		fmt.Printf(curLog.Log)
		for index := 0; index < 5; index++ {
			fmt.Println("                      |")
		}
	}
	fmt.Println(" **************************That's all for the logs******************")
}

func PrintTaskStatsReponse(input *handlers.TaskStatsResponse) string {
	return fmt.Sprintf(`\n Total Number of Active Task = %d\n Total number of complered Task = %d\n 
	Total number of spilled Task = %d\n`, input.ActiveCount, input.CompletedCount, input.SpilledCount)
}

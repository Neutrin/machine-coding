package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/todo_list/internal/core/domain/enums"
	"github.com/todo_list/internal/core/domain/validations"
	"github.com/todo_list/internal/core/services"
	"github.com/todo_list/internal/handlers"
	"github.com/todo_list/internal/repositiories"
)

func main() {
	validations.InitValidator()
	scanner := bufio.NewScanner(os.Stdin)
	//Todo : Write comparable function over here
	repo := repositiories.NewMapRepo(repositiories.NewBinarySearchIndex(repositiories.CompareFunction))
	service := services.NewService(repo)
	handler := handlers.NewConsole(service)

	for {
		PrintMenu()
		scanner.Scan()
		option := scanner.Text()
		if inputOption, err := strconv.Atoi(option); err == nil {
			switch inputOption {
			case 1:
				desc, endDate, createdBy, tags := CreateTaskHanlder(scanner)

				resp, err := handler.CreateTask(desc, endDate, createdBy, tags)
				if err != nil {
					fmt.Printf("Failed = %+v\n", err)
				} else {
					fmt.Printf("%s\n", resp.Resp)
				}
			case 2:
				id := GetTaskId(scanner)
				option, updateStr := UpdateOptions(scanner)
				resp := handlers.TaskCreatedResp{}
				switch option {
				case 1:
					resp = handler.UpdateDesc(id, updateStr)
				case 2:
					resp = handler.UpdateDueDate(id, updateStr)
				case 3:
					tagId, _ := strconv.ParseInt(updateStr, 10, 64)
					resp = handler.UpdateTag(id, enums.Tags(tagId))
				}
				if len(resp.ErrorStr) > 0 {
					fmt.Println(" failed to update with error = ", resp.ErrorStr)
				} else {
					fmt.Println(resp.Resp)
				}
			case 3:
				id := GetTaskId(scanner)
				resp := handler.MarkCompleted(id)
				if len(resp.ErrorStr) > 0 {
					fmt.Println(" failed to update with error = ", resp.ErrorStr)
				} else {
					fmt.Println(resp.Resp)
				}
			case 4:
				taskResp := handler.GetAllTask()
				if len(taskResp.Error) > 0 {
					fmt.Println(" error = ", taskResp.Error)
				} else {
					for _, curResp := range taskResp.TaskRespBody {
						PrintTask(curResp)
					}
				}
			case 5:

				id := GetTaskId(scanner)
				resp := handler.Task(id)
				if len(resp.Error) > 0 {
					fmt.Println(" error = ", resp.Error)
				} else {
					PrintTask(resp.TaskRespBody[0])
				}
			case 6:
				startTime, endTime := GetTimeRange(scanner)
				PrintActivityLog(handler.ActivityLogs(startTime, endTime))
			case 7:
				startTime, endTime := GetTimeRange(scanner)
				fmt.Println(PrintTaskStatsReponse(handler.TaskStats(startTime, endTime)))
			}
		} else {
			fmt.Println(" Invalid option try again")
			time.Sleep(4 * time.Second)
		}
		fmt.Println(" Want to conitnue (y/n)")
		scanner.Scan()
		option = scanner.Text()
		if strings.Compare(option, "y") != 0 {
			fmt.Println(" Thank you for using this application")
			repo.StopRoutine()
			fmt.Println("stopping the go routine")
			time.Sleep(10 * time.Second)
			break
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/online_coding_platform/internals/core/domains"
	"github.com/online_coding_platform/internals/core/domains/enums"
	"github.com/online_coding_platform/internals/core/services"
	"github.com/online_coding_platform/internals/repositiories"
)

func main() {
	var (
	// users []domains.User
	// err   error
	)
	repo := repositiories.NewUserRepoMap()
	userService := services.NewUserService(repo)
	questionRepo := repositiories.NewQuestionsRepo()
	questionService := services.NewQuestionService(questionRepo)
	contestRepo := repositiories.NewContestRepo()
	constestService := services.NewContestService(contestRepo, repo, questionRepo)
	leaderBoardService := services.NewLeaderBoardService(repo)

	//Command One Create User <user_name>
	userService.CreateUser("nitin")
	userService.CreateUser("kritika")
	userService.CreateUser("shubham")
	//List down all the users
	// if users, err = userService.AllUser(); err != nil {
	// 	fmt.Printf(" error = %+v\n", err)
	// 	os.Exit(1)
	// }

	// for _, curUser := range users {
	// 	PrintUser(curUser)
	// }
	//Command Twp create Quesion
	questionService.CreateQuestion(1, 80)
	questionService.CreateQuestion(1, 80)
	questionService.CreateQuestion(2, 130)
	questionService.CreateQuestion(3, 130)
	questionService.CreateQuestion(3, 200)

	//Command three list all quesions
	questions, err := questionService.ListQuestions()
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	for _, curQues := range questions {
		PrintQuestion(curQues)
	}
	fmt.Println(" question level 1")
	questions, err = questionService.ListQuestions(1)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	for _, curQues := range questions {
		PrintQuestion(curQues)
	}
	fmt.Println(" question level 2")
	questions, err = questionService.ListQuestions(2)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	for _, curQues := range questions {
		PrintQuestion(curQues)
	}
	fmt.Println(" question level 3")
	questions, err = questionService.ListQuestions(3)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	for _, curQues := range questions {
		PrintQuestion(curQues)
	}

	err = constestService.CreateContest("leetcode one", enums.Low, 1)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	err = constestService.CreateContest("leetcode two", enums.Medium, 1)

	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	err = constestService.CreateContest("leetcode three", enums.High, 1)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	contests, err := constestService.ListContest()
	_ = contests
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	// for _, curContest := range contests {
	// 	PrintContest(curContest)
	// }
	contests, err = constestService.ListContest(1)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	// for _, curContest := range contests {
	// 	PrintContest(curContest)
	// }
	contests, err = constestService.ListContest(2)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	// for _, curContest := range contests {
	// 	PrintContest(curContest)
	// }
	contests, err = constestService.ListContest(3)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	// for _, curContest := range contests {
	// 	PrintContest(curContest)
	// }
	err = constestService.AttendContest(1, 2)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	constestService.AttendContest(1, 3)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	err = constestService.RunContest(1, 1)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	users, err := userService.AllUser()
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}

	for _, curUser := range users {
		PrintUser(curUser)
	}

	leaders := leaderBoardService.GenerateLeaderboard(1)
	fmt.Println(" ***************** Leaderboard asceding***********************")
	for _, curLeader := range leaders {
		PrintLeaderBoard(curLeader)
	}
	fmt.Println(" ****************************************")
	leaders = leaderBoardService.GenerateLeaderboard(2)
	fmt.Println(" ***************** Leaderboard desc***********************")
	for _, curLeader := range leaders {
		PrintLeaderBoard(curLeader)
	}
	fmt.Println(" ****************************************")

}

func PrintUser(user domains.User) {
	fmt.Println(" ")
	fmt.Printf(" id = %d", user.ID())
	fmt.Printf(" name = %s", user.Name())
	fmt.Printf(" score = %d", user.Score())
	fmt.Println(" ")
}

func PrintQuestion(question domains.Question) {
	fmt.Println(" ")
	// fmt.Printf(" id = %d", question.Id())
	// fmt.Printf(" points = %d", question.Points())
	// fmt.Printf(" level = %d", question.Level())
	fmt.Println(" ")
}

func PrintContest(contest domains.Contest) {
	fmt.Println(" ***************** Contest***********************")
	fmt.Printf(" %+v\n", contest)
	fmt.Println(" ****************************************")
}

func PrintLeaderBoard(leaderboard domains.LeaderBoard) {
	fmt.Printf(" %+v\n", leaderboard)

}

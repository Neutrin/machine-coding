package services

import (
	"fmt"

	"github.com/online_coding_platform/internals/core/domains"
	custom_enums "github.com/online_coding_platform/internals/core/domains/enums"
	custom_errors "github.com/online_coding_platform/internals/core/errors"
	"github.com/online_coding_platform/internals/core/ports"
	"github.com/online_coding_platform/internals/core/utils"
)

type ContestService struct {
	repo         ports.ContestRepositiories
	userRepo     ports.UserRepo
	questionRepo ports.QuestionReposteries
}

func NewContestService(repo ports.ContestRepositiories, userRepo ports.UserRepo, quesRepo ports.QuestionReposteries) *ContestService {
	utils.InitRandom()
	return &ContestService{
		repo:         repo,
		userRepo:     userRepo,
		questionRepo: quesRepo,
	}
}

func (service *ContestService) CreateContest(name string, level int64, userId int64) error {
	if !custom_enums.ValidLevel(level) {
		return custom_errors.InvalidContestLevel
	}
	user, err := service.userRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	contest, err := domains.NewContest(level, name, user)
	if err != nil {
		return err
	}
	contest = contest.Register(user)
	service.repo.Save(contest)
	return nil
}

func (service *ContestService) ListContest(level ...int64) ([]domains.Contest, error) {
	if len(level) == 1 {
		return service.repo.Level(level[0])
	}
	return service.repo.All(), nil

}

func (service *ContestService) AttendContest(contestId int64, userId int64) error {
	contest, err := service.repo.ById(contestId)
	if err != nil {
		return err
	}
	user, err := service.userRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	contest = contest.Register(user)
	service.repo.Save(contest)
	return nil
}

func (service *ContestService) RunContest(contestId int64, startedBy int64) error {

	contest, err := service.repo.ById(contestId)
	if err != nil {
		return err
	}

	if contest.CreatedBy().ID() != startedBy {
		return custom_errors.ContestCreationAuthentication
	}
	if contest.Status() != custom_enums.Created {
		return custom_errors.ContestInvalidStatus
	}
	questions, err := service.questionRepo.GetByLevel(contest.Level())
	if err != nil {
		return err
	}
	scoreCal, err := domains.ScoreCalculatorFactory(contest.Level())
	if err != nil {
		return err
	}
	fmt.Println(" *******************starting contest name = ", contest.Name(), "*******************")
	for _, curUser := range contest.RegisterBy() {
		questionSolved := make([]domains.Question, 0)
		for _, curQuestion := range questions {
			if utils.Number()%2 == 0 {
				questionSolved = append(questionSolved, curQuestion)
			}
		}
		curScore := scoreCal.CalculateScore(questionSolved)
		fmt.Println(" for user = ", curUser.ID(), " score is ", curScore, " questions solved ", len(questionSolved))
		curUser = curUser.SetScore(curUser.Score() + curScore)
		service.userRepo.SaveUser(curUser)
	}
	fmt.Println("************************ Ended contest *********************************")
	return nil
}

package repositiories

import (
	"time"

	"github.com/online_coding_platform/internals/core/domains"
	custom_errors "github.com/online_coding_platform/internals/core/errors"
)

type ContestRepoModel struct {
	Id     int64
	Level  int64
	Name   string
	Status int64

	CreatedBy   userRepoModel
	RegistredBy []userRepoModel

	CreatedAt time.Time
}

type ContestRepo struct {
	mp map[int64]interface{}
}

func NewContestRepo() *ContestRepo {
	return &ContestRepo{
		mp: make(map[int64]interface{}),
	}
}

/*

type ContestRepositiories interface {
	Save(contest domains.Contest)
	All() []domains.Contest
	Level(level int64) ([]domains.Contest, error)
}

*/

func (repo *ContestRepo) Save(contest domains.Contest) {
	model := contestRepoModel(contest)
	repo.mp[model.Id] = model
}

func (repo *ContestRepo) All() []domains.Contest {
	contests := make([]domains.Contest, 0)
	for _, curContestRepo := range repo.mp {
		contests = append(contests, contestDomainToRepoModel(curContestRepo.(ContestRepoModel)))
	}
	return contests
}

func (repo *ContestRepo) Level(level int64) ([]domains.Contest, error) {
	contests := make([]domains.Contest, 0)
	for _, curContests := range repo.All() {
		if curContests.Level() == level {
			contests = append(contests, curContests)
		}
	}
	return contests, nil
}

func (repo *ContestRepo) ById(id int64) (domains.Contest, error) {
	repoModel, found := repo.mp[id]
	if !found {
		return domains.Contest{}, custom_errors.ContestNotFound
	}
	return contestDomainToRepoModel(repoModel.(ContestRepoModel)), nil
}

func contestRepoModel(contest domains.Contest) ContestRepoModel {
	model := ContestRepoModel{
		Id:        contest.Id(),
		Level:     contest.Level(),
		CreatedBy: repoUserModel(contest.CreatedBy()),
		CreatedAt: contest.CreatedAt(),
		Name:      contest.Name(),
		Status:    contest.Status(),
	}
	model.RegistredBy = make([]userRepoModel, 0)
	for _, curUser := range contest.RegisterBy() {
		model.RegistredBy = append(model.RegistredBy, repoUserModel(curUser))
	}
	return model
}

func contestDomainToRepoModel(contest ContestRepoModel) domains.Contest {
	users := make([]domains.User, 0)
	for _, curUser := range contest.RegistredBy {
		users = append(users, domainsUserModel(curUser))
	}
	return domains.NewContestWithId(contest.Id, domainsUserModel(contest.CreatedBy), contest.Name, contest.Level,
		users, contest.Status, contest.CreatedAt)
}

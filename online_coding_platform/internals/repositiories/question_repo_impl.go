package repositiories

import "github.com/online_coding_platform/internals/core/domains"

type QuestionsRepo struct {
	mp map[int64]interface{}
}

type QuestionDBModel struct {
	Id     int64
	Points int64
	Level  int64
}

/*
type QuestionReposteries interface {
	Save(question domains.Question) error
	GetAll() ([]domains.Question, error)
	GetByLevel(level int64) ([]domains.Question, error)
}
*/

func NewQuestionsRepo() *QuestionsRepo {
	return &QuestionsRepo{
		mp: make(map[int64]interface{}),
	}
}

func (repo *QuestionsRepo) Save(question domains.Question) error {
	repo.mp[question.Id()] = questionRepoModel(question)
	return nil
}

func (repo *QuestionsRepo) GetAll() ([]domains.Question, error) {

	questions := make([]domains.Question, 0, len(repo.mp))
	for _, quesDet := range repo.mp {
		questionRepo := quesDet.(QuestionDBModel)
		if question, err := domains.NewQuestionWithId(
			questionRepo.Id, questionRepo.Points, questionRepo.Level); err != nil {

			return questions, err
		} else {
			questions = append(questions, question)
		}

	}
	return questions, nil
}

func (repo *QuestionsRepo) GetByLevel(level int64) ([]domains.Question, error) {
	lvlQues := make([]domains.Question, 0)
	questions, err := repo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, curQues := range questions {
		if curQues.Level() == level {
			lvlQues = append(lvlQues, curQues)
		}
	}
	return lvlQues, nil
}

func questionRepoModel(question domains.Question) QuestionDBModel {
	return QuestionDBModel{
		Id:     question.Id(),
		Level:  question.Level(),
		Points: question.Points(),
	}
}

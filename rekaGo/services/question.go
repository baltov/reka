package services

import (
	"gorm.io/gorm"
	"math/rand"
	"rekaGo/models"
)

type QuestionService interface {
	GetQuestions(limit int) ([]models.Question, error)
}

type questionServiceImpl struct {
	db *gorm.DB
}

func NewQuestionService(db *gorm.DB) QuestionService {
	return &questionServiceImpl{
		db: db,
	}
}

func (s *questionServiceImpl) GetQuestions(limit int) ([]models.Question, error) {
	halfLimit := limit / 2

	// Fetch questions in both categories
	questions1, err := s.fetchQuestions("category_id IN (SELECT id FROM categories WHERE name LIKE ?)", "%ОППД%", halfLimit)
	if err != nil {
		return nil, err
	}

	questions2, err := s.fetchQuestions("category_id NOT IN (SELECT id FROM categories WHERE name LIKE ?)", "%ОППД%", halfLimit)
	if err != nil {
		return nil, err
	}

	// Combine all questions
	allQuestions := append(questions1, questions2...)

	// Increment the 'questioned' count for all retrieved questions
	if err := s.incrementQuestioned(allQuestions); err != nil {
		return nil, err
	}

	// Shuffle answers for each question
	shuffleAnswers(allQuestions)

	return allQuestions, nil
}

// fetchQuestions retrieves questions based on the provided condition and limit
func (s *questionServiceImpl) fetchQuestions(condition string, param string, limit int) ([]models.Question, error) {
	var questions []models.Question
	err := s.db.
		Preload("Answers").
		Preload("Images").
		Distinct().
		Order("RANDOM()").
		Where(condition, param).
		Limit(limit).
		Find(&questions).Error
	return questions, err
}

// incrementQuestioned updates the 'questioned' field for each question
func (s *questionServiceImpl) incrementQuestioned(questions []models.Question) error {
	for _, question := range questions {
		if err := s.db.Model(&question).Update("questioned", question.Questioned+1).Error; err != nil {
			return err
		}
	}
	return nil
}

// shuffleAnswers randomizes the order of answers for each question
func shuffleAnswers(questions []models.Question) {
	for i := range questions {
		rand.Shuffle(len(questions[i].Answers), func(j, k int) {
			questions[i].Answers[j], questions[i].Answers[k] = questions[i].Answers[k], questions[i].Answers[j]
		})
	}
}

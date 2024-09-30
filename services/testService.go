package services

import (
	"context"
	"fmt"
	"strings"
	"test/domain"
	"test/dto"
	"test/errors"
	"time"
)

type VariantsRepo interface {
	FindAll(context.Context) ([]domain.Variant, error)
}

type TestsUserRepo interface {
	Save(context.Context, domain.TestUser) error
	Update(context.Context, domain.TestUser) error
	CountByVariantID(context.Context, domain.UUID, domain.UUID) (int, error)
	GetByVariantID(context.Context, domain.UUID) (*domain.TestUser, error)
	FindByUserID(context.Context, domain.UUID, domain.UUID) (*domain.TestUser, error)
}

type AnswersRepo interface {
	Save(context.Context, domain.Answer) error
	CountByAnswerID(context.Context, domain.UUID) (int, error)
	GetByAnswerID(context.Context, domain.UUID) (*domain.Answer, error)
	Update(context.Context, domain.Answer) error
}

type ResultsRepo interface {
	Save(context.Context, domain.Result) error
}

type TasksRepo interface {
	FindByVariantID(context.Context, domain.UUID) ([]domain.Task, error)
	GetByVariantID(context.Context, domain.UUID) (*domain.Task, error)
	CountByVariantID(context.Context, domain.UUID) (int, error)
	FindTasksByTestUserID(context.Context, domain.UUID) ([]domain.Task, error)
}

type TestService struct {
	variantsRepo  VariantsRepo
	testsUserRepo TestsUserRepo
	answersRepo   AnswersRepo
	tasksRepo     TasksRepo
	resultsRepo   ResultsRepo
	authService   AuthService
}

func NewTestService(variantsRepo VariantsRepo, testsUserRepo TestsUserRepo, answersRepo AnswersRepo, tasksRepo TasksRepo, resultsRepo ResultsRepo, authService AuthService) *TestService {

	return &TestService{
		variantsRepo:  variantsRepo,
		testsUserRepo: testsUserRepo,
		answersRepo:   answersRepo,
		tasksRepo:     tasksRepo,
		resultsRepo:   resultsRepo,
		authService:   authService,
	}
}

func (s *TestService) ListVariants(ctx context.Context, request dto.ListVariantsReq) (*dto.ListVariantsRes, error) {
	if err := request.Validate(); err != nil {
		return nil, errors.NewValidationError("validate error")
	}
	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	variants, err := s.variantsRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.NewNotFoundError("not found variants")
	}

	return &dto.ListVariantsRes{Variants: variants}, nil
}

func (s *TestService) GetTask(ctx context.Context, request dto.GetTaskReq) (*dto.GetTaskRes, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request %w", err)
	}

	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {
		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	count, err := s.tasksRepo.CountByVariantID(ctx, request.VariantID)
	if err != nil {

		return nil, err
	}

	if count == 0 {

		return nil, errors.NewNotFoundError("not found task")
	}
	if request.TaskID == 1 {
		testUser := domain.NewTestUser(request.UserID, request.VariantID)
		if err := s.testsUserRepo.Save(ctx, *testUser); err != nil {
			return nil, fmt.Errorf("can`t save user")
		}
	}
	if request.TaskID == count {
		currentTimeNow := time.Now()

		testUser, err := s.testsUserRepo.GetByVariantID(ctx, request.VariantID)
		if err != nil {
			return nil, fmt.Errorf("can't get test User")
		}

		testUser.EndAt = &currentTimeNow

		if err := s.testsUserRepo.Update(ctx, *testUser); err != nil {
			return nil, fmt.Errorf("can`t save user")
		}
	}
	if request.TaskID > count {

		return &dto.GetTaskRes{}, nil
	}

	tasks, err := s.tasksRepo.FindByVariantID(ctx, request.VariantID)
	if err != nil {
		return nil, fmt.Errorf("not found task %w", err)
	}

	return &dto.GetTaskRes{
		TaskID:      request.TaskID,
		Description: tasks[request.TaskID].Description,
		Options:     tasks[request.TaskID].Options,
	}, nil

}

func (s *TestService) AnswerTask(ctx context.Context, request dto.AnswerTaskReq) (*dto.AnswerTaskRes, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request %w", err)
	}

	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {
		return nil, errors.NewAuthenticationError("user invalid authentication")
	}
	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	if request.TaskID == 0 {
		return nil, fmt.Errorf("not found task")
	}

	if request.Answer == "" {
		testUser, err := s.testsUserRepo.FindByUserID(ctx, request.UserID, request.VariantID)
		if err != nil {
			return nil, fmt.Errorf("not found tests")
		}

		answer := domain.NewAnswer(testUser.ID, request.Answer)
		if err := s.answersRepo.Save(ctx, *answer); err != nil {
			return nil, fmt.Errorf("can't save %w", err)
		}
		return &dto.AnswerTaskRes{AnswerID: answer.ID}, nil
	}

	count, err := s.tasksRepo.CountByVariantID(ctx, request.VariantID)
	if err != nil {
		return nil, errors.NewNotFoundError("not found task")
	}

	answer, err := s.answersRepo.GetByAnswerID(ctx, request.AnswerID)
	if err != nil {
		return nil, errors.NewNotFoundError("not found Answer")
	}

	if request.TaskID < count {
		answerUpdate := answer.UpdateAnswer(request.Answer)

		if err := s.answersRepo.Update(ctx, *answerUpdate); err != nil {
			return nil, fmt.Errorf("can't update answer")
		}
	}

	return &dto.AnswerTaskRes{AnswerID: answer.ID}, nil

}

func (s *TestService) Result(ctx context.Context, request dto.ResultVariantReq) (*dto.ResultVariantRes, error) {
	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {
		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	countAnswer, err := s.answersRepo.CountByAnswerID(ctx, request.AnswerID)
	if err != nil {

		return nil, fmt.Errorf("invalid count answer:%w", err)
	}

	if countAnswer == 0 {

		return nil, fmt.Errorf("not found answer:%w", err)
	}

	answer, err := s.answersRepo.GetByAnswerID(ctx, request.AnswerID)
	if err != nil {

		return nil, fmt.Errorf("not found answer:%w", err)
	}

	if answer.Answer == "" {
		return nil, fmt.Errorf("not found answer:%w", err)
	}
	var answerArray []string
	answerArray = strings.Split(answer.Answer, " ")

	tasks, err := s.tasksRepo.FindTasksByTestUserID(ctx, answer.TestUserID)
	if err != nil {
		return nil, err
	}
	var errorTest int

	for index, answer := range answerArray {
		if index >= len(tasks) {
			break
		}

		if tasks[index].CorrectAnswer != answer {
			errorTest++
		}
	}

	per := (len(tasks) - errorTest) / len(tasks)
	percent := fmt.Sprintf("%.2f", per)

	result := domain.NewResult(answer.TestUserID, percent)

	if err := s.resultsRepo.Save(ctx, *result); err != nil {

		return nil, fmt.Errorf("can't save %w", err)
	}

	return &dto.ResultVariantRes{Percent: percent}, nil
}

func (s *TestService) isAuthentication(ctx context.Context, userID domain.UUID) (bool, error) {
	user, err := s.authService.usersRepo.FindByID(ctx, userID)
	if err != nil {
		return false, errors.NewAuthenticationError("user not authentication")
	}

	return user.Auth == true, nil
}

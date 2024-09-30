package services

import (
	"context"
	"fmt"
	"test/domain"
	"test/dto"
	"test/errors"
)

type VariantsRepo interface {
	FindAll(context.Context) ([]domain.Variant, error)
}

type TestsUserRepo interface {
	Save(context.Context, domain.TestUser) error
	Update(context.Context, domain.TestUser) error
	CountByVariantID(context.Context, domain.UUID, domain.UUID) (int, error)
	GetByVariantID(context.Context, domain.UUID) (*domain.TestUser, error)
}

type AnswersRepo interface {
	Save(context.Context, domain.Answer) error
}

type ResultsRepo interface {
	Save(context.Context, domain.Result) error
}

type TasksRepo interface {
	FindByVariantID(context.Context, domain.UUID) ([]domain.Task, error)
	GetByVariantID(context.Context, domain.UUID) (*domain.Task, error)
	CountByVariantID(context.Context, domain.UUID) (int, error)
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
		return nil, fmt.Errorf("don't valid request %w", err)
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

	if request.TaskID > count {

		return &dto.GetTaskRes{}, nil
	}

	tasks, err := s.tasksRepo.FindByVariantID(ctx, request.VariantID)
	if err != nil {
		return nil, fmt.Errorf("not found task %w", err)
	}

	return &dto.GetTaskRes{TaskID: request.TaskID, Description: tasks[request.TaskID].Description,
		Options: tasks[request.TaskID].Options}, nil

	/*testUser, err := s.testsUserRepo.GetByVariantID(ctx, request.VariantID)
	if err != nil {
		return nil, fmt.Errorf("not found variant")
	}

	testUser.StartAt = time.Now()
	testUser.EndAt = nil

	if err := s.testsUserRepo.Update(ctx, *testUser); err != nil {
		return nil, fmt.Errorf("don`t update test")
	}

	return &dto.GetTasksRes{Task: *task}, nil*/
}

func (s *TestService) AnswerTask(ctx context.Context, request dto.AnswerTaskReq) (*dto.AnswerTaskRes, error) {
	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {
		return nil, errors.NewAuthenticationError("user invalid authentication")
	}
	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	return &dto.AnswerTaskRes{}, nil
}

func (s *TestService) Result(ctx context.Context, request dto.ResultVariantReq) (*dto.ResultVariantRes, error) {

	isAuth, err := s.isAuthentication(ctx, request.UserID)
	if err != nil {
		return nil, errors.NewAuthenticationError("user invalid authentication")
	}
	if isAuth == false {

		return nil, errors.NewAuthenticationError("user invalid authentication")
	}

	return &dto.ResultVariantRes{}, nil
}

func (s *TestService) isAuthentication(ctx context.Context, userID domain.UUID) (bool, error) {
	user, err := s.authService.usersRepo.FindByID(ctx, userID)
	if err != nil {
		return false, errors.NewAuthenticationError("user not authentication")
	}

	return user.Auth == true, nil
}

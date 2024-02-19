package service

import (
	"context"
	"saved/http/rest/internal/auth"
	mAccount "saved/http/rest/internal/model/account"
	mRepo "saved/http/rest/internal/model/repository"
	mRes "saved/http/rest/internal/model/response"
	mService "saved/http/rest/internal/model/service"
	mUser "saved/http/rest/internal/model/user"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type accountService struct {
	accountRepo mRepo.AccountRepository
	logger      *logrus.Logger
}

func InitAccountService(accountRepo mRepo.AccountRepository, logger *logrus.Logger) mService.AccountService {
	return &accountService{
		accountRepo: accountRepo,
		logger:      logger,
	}
}

func (s accountService) Count(ctx context.Context) (mRes.CountDto, error) {
	count, err := s.accountRepo.Count(ctx)
	if err != nil {
		return mRes.CountDto{Count: 0}, err
	}
	return mRes.CountDto{Count: count}, nil
}

func (s accountService) FindAll(ctx context.Context, query map[string][]string) ([]mAccount.Account, error) {
	res, err := s.accountRepo.FindAll(ctx, query)
	if err != nil {
		return []mAccount.Account{}, err
	}
	return res, nil
}

func (s accountService) FindById(ctx context.Context, id string) (mAccount.Account, error) {
	res, err := s.accountRepo.FindById(ctx, id)
	if err != nil {
		return mAccount.Account{}, err
	}
	return res, nil
}

func (s accountService) Create(ctx context.Context, payload mAccount.Account) (mAccount.Account, error) {
	user, err := auth.UserLogin(ctx, s.logger)
	if err != nil {
		user = mUser.User{}
	}
	payload.Id = primitive.NewObjectID()
	payload.CreatedBy = user.Id
	payload.CreatedAt = time.Now()
	payload.UpdatedBy = user.Id
	payload.UpdatedAt = time.Now()
	res, err := s.accountRepo.Create(ctx, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s accountService) Update(ctx context.Context, id string, payload mAccount.Account) (mAccount.Account, error) {
	user, err := auth.UserLogin(ctx, s.logger)
	if err != nil {
		user = mUser.User{}
	}
	doc, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mAccount.Account{}, err
	}
	payload.Id = doc
	payload.UpdatedBy = user.Id
	payload.UpdatedAt = time.Now()
	res, err := s.accountRepo.Update(ctx, payload)
	if err != nil {
		return mAccount.Account{}, err
	}
	return res, nil
}
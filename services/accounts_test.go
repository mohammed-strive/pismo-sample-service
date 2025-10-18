package services_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pismo-service/mocks"
	"pismo-service/models"
	"pismo-service/services"
)

var _ = Describe("Accounts", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockAccountRepository
		service  services.AccountService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockAccountRepository(mockCtrl)
		service = services.NewAccountService(mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("GetAccountById", func() {
		It("should return a valid user", func() {
			expectedAccount := models.Account{AccountId: uint64(1), DocumentNumber: "abcdef123456"}
			mockRepo.EXPECT().GetOneAccountById(context.TODO(), uint64(1)).Return(expectedAccount, nil)

			account, err := service.GetAccountById(context.TODO(), uint64(1))
			Expect(err).To(BeNil())
			Expect(account).To(Equal(expectedAccount))
		})

		It("should return an error if repo returns an error", func() {
			mockRepo.EXPECT().GetOneAccountById(context.TODO(), uint64(2)).Return(models.Account{}, fmt.Errorf("Repo error"))

			_, err := service.GetAccountById(context.TODO(), uint64(2))
			Expect(err).To(MatchError(fmt.Errorf("Account Service: GetAccountById: Error getting Account with ID-2: Repo error")))
		})
	})

	Context("CreateAccount", func() {
		It("create new account", func() {
			expectedAccount := models.Account{DocumentNumber: "abcdef123456"}
			mockRepo.EXPECT().CreateNewAccount(context.TODO(), expectedAccount).
				Return(models.Account{DocumentNumber: expectedAccount.DocumentNumber, AccountId: 1}, nil)

			result, err := service.CreateAccount(context.TODO(), expectedAccount)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(models.Account{
				DocumentNumber: "abcdef123456",
				AccountId:      uint64(1),
			}))
		})

		It("return an error if repo returns an error", func() {
			expectedAccount := models.Account{DocumentNumber: "abcdef123456"}
			mockRepo.EXPECT().CreateNewAccount(context.TODO(), expectedAccount).
				Return(models.Account{}, fmt.Errorf("repo error"))

			_, err := service.CreateAccount(context.TODO(), expectedAccount)
			Expect(err).To(MatchError(fmt.Errorf("Account Service: CreateAccount: Error creating new account: repo error")))
		})
	})
})

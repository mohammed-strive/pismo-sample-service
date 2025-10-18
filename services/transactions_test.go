package services_test

import (
	"context"
	"fmt"
	"pismo-service/mocks"
	"pismo-service/models"
	"pismo-service/services"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transactions", func() {
	var (
		mockCtrl *gomock.Controller
		repo     *mocks.MockTransactionRepository
		service  services.TransactionService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockTransactionRepository(mockCtrl)
		service = services.NewTransactionService(repo)
	})

	Context("CreateTransaction", func() {
		It("should create a new Transaction", func() {
			expectedTransaction := models.Transaction{
				AccountId:       1,
				TransactionId:   10,
				OperationTypeId: 4,
				Amount:          400.0,
			}
			repo.EXPECT().AddNewTransaction(context.TODO(), expectedTransaction).
				Return(expectedTransaction, nil)
			result, err := service.CreateTransaction(context.TODO(), expectedTransaction)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedTransaction))
		})

		It("should return error when repo returns an error", func() {
			repo.EXPECT().AddNewTransaction(context.TODO(), models.Transaction{}).
				Return(models.Transaction{}, fmt.Errorf("repo error"))
			_, err := service.CreateTransaction(context.TODO(), models.Transaction{})
			Expect(err).To(MatchError(fmt.Errorf("TransactionService: CreateTransaction: Error creating new Transaction: repo error")))
		})
	})
})

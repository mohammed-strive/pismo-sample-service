package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"pismo-service/controllers"
	"pismo-service/mocks"
	"pismo-service/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TransactionController", func() {
	var (
		router      *gin.Engine
		controller  controllers.TransactionController
		mockService *mocks.MockTransactionService
		mockCtrl    *gomock.Controller
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		mockCtrl = gomock.NewController(GinkgoT())
		mockService = mocks.NewMockTransactionService(mockCtrl)
		controller = controllers.NewTransactionController(mockService)

		router.POST("/transactions/", controller.CreateTransaction)
	})

	Context("CreateAccount", func() {
		It("should return 201 when transaction is successfully created", func() {
			input := `{"account_id":1, "operation_type_id":2, "amount":100.0}`
			req := httptest.NewRequest("POST", "/transactions/", bytes.NewBuffer([]byte(input)))
			resp := httptest.NewRecorder()

			expectedAccount := models.Transaction{
				AccountId:       1,
				OperationTypeId: 2,
				Amount:          100.0,
				TransactionId:   4,
			}
			mockService.EXPECT().
				CreateTransaction(gomock.AssignableToTypeOf(context.Background()), gomock.Any()).
				Return(expectedAccount, nil)

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusCreated))
			body, _ := io.ReadAll(resp.Body)

			var result models.Transaction
			err := json.Unmarshal(body, &result)
			Expect(err).To(BeNil())
			Expect(result.TransactionId).To(Equal(uint64(4)))
		})

		It("should return error when service returns error", func() {
			input := `{"account_id":"abc"}`
			req := httptest.NewRequest("POST", "/transactions/", bytes.NewBuffer([]byte(input)))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusBadRequest))
		})
	})
})

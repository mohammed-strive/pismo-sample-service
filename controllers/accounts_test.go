package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

var _ = Describe("AccountsController", func() {
	var (
		router      *gin.Engine
		controller  controllers.AccountController
		mockService *mocks.MockAccountService
		mockCtrl    *gomock.Controller
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		mockCtrl = gomock.NewController(GinkgoT())
		mockService = mocks.NewMockAccountService(mockCtrl)
		controller = controllers.NewAccountController(mockService)

		router.GET("/accounts/:accountId", controller.GetAccountByAccountId)
		router.POST("/accounts/", controller.CreateAccount)
	})

	Context("GetAccountByAccountId", func() {
		It("should return 200 when response is successful", func() {
			req := httptest.NewRequest("GET", "/accounts/123", nil)
			resp := httptest.NewRecorder()

			expectedAccount := models.Account{
				AccountId:      123,
				DocumentNumber: "abcdef123456",
			}
			mockService.EXPECT().
				GetAccountById(gomock.AssignableToTypeOf(context.Background()), uint64(123)).
				Return(expectedAccount, nil)

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			body, _ := io.ReadAll(resp.Body)

			var account models.Account
			err := json.Unmarshal(body, &account)
			Expect(err).To(BeNil())
			Expect(account.DocumentNumber).To(Equal("abcdef123456"))
		})

		It("should return error when service returns error", func() {
			req := httptest.NewRequest("GET", "/accounts/123", nil)
			resp := httptest.NewRecorder()

			expectedAccount := models.Account{
				AccountId:      123,
				DocumentNumber: "abcdef123456",
			}
			mockService.EXPECT().
				GetAccountById(gomock.AssignableToTypeOf(context.Background()), uint64(123)).
				Return(expectedAccount, fmt.Errorf("service error"))

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusNotFound))
		})
	})

	Context("CreateAccount", func() {
		It("should return 201 when transaction is successfully created", func() {
			input := `{"document_number":"123"}`
			req := httptest.NewRequest("POST", "/accounts/", bytes.NewBuffer([]byte(input)))
			resp := httptest.NewRecorder()

			expectedAccount := models.Account{
				AccountId:      123,
				DocumentNumber: "abcdef123456",
			}
			mockService.EXPECT().
				CreateAccount(gomock.AssignableToTypeOf(context.Background()), gomock.Any()).
				Return(expectedAccount, nil)

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusCreated))
			body, _ := io.ReadAll(resp.Body)

			var result models.Account
			err := json.Unmarshal(body, &result)
			Expect(err).To(BeNil())
			Expect(result.DocumentNumber).To(Equal("abcdef123456"))
		})

		It("should return error when service returns error", func() {
			input := `{"document_number":"abc"}`
			req := httptest.NewRequest("POST", "/accounts/", bytes.NewBuffer([]byte(input)))
			resp := httptest.NewRecorder()

			mockService.EXPECT().
				CreateAccount(gomock.AssignableToTypeOf(context.Background()), models.Account{DocumentNumber: "abc"}).
				Return(models.Account{}, fmt.Errorf("service error"))

			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})

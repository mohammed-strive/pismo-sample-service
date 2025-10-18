package controllers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"pismo-service/controllers"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("404", func() {
	var (
		router     *gin.Engine
		controller controllers.NoRouteController
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.Default()
		controller = controllers.NewNoRouteController()
		router.NoRoute(controller.NoRouteHandler)
	})

	Context("NoRouteController", func() {
		It("should return invalid path error when invalid path is requested", func() {
			req := httptest.NewRequest("GET", "/v1/invalid", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			Expect(resp.Result().StatusCode).To(Equal(404))
			var result any
			body, _ := io.ReadAll(resp.Body)
			err := json.Unmarshal(body, &result)
			fmt.Printf("%v", result)
			Expect(err).To(BeNil())
		})
	})
})

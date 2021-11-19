package customer

import (
	"TechnicalTest-Golang/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController struct {
	Usecase ICustomerUsecase
}

func (c CustomerController) Customer(r *gin.RouterGroup) {
	r.POST("/login", c.Login)
	r.POST("/logout", c.Logout)
}

func (c CustomerController) Login(g *gin.Context) {
	var body model.Customer
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: "Error Bind JSON, error = " + err.Error()})
		return
	}
	err = c.Usecase.Login(body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}
	g.JSON(http.StatusOK, model.Response{
		Message: "Success Login",
	})
}

func (c CustomerController) Logout(g *gin.Context) {
	c.Usecase.Logout()
	g.JSON(http.StatusOK, model.Response{
		Message: "Success Logout",
	})
}

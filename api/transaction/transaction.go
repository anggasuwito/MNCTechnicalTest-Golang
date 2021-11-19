package transaction

import (
	"TechnicalTest-Golang/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	Usecase ITransactionUsecase
}

func (t TransactionController) Transaction(r *gin.RouterGroup) {
	r.POST("/payment",t.Payment)
}

func (t TransactionController) Payment(g *gin.Context) {
	var body model.Payment
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: "Error Bind JSON, error = " + err.Error()})
		return
	}
	result, err := t.Usecase.Payment(body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}
	g.JSON(http.StatusOK, model.Response{
		Message: "Success Payment",
		Result:  result,
	})
}

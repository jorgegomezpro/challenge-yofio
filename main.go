package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"main.go/models"
)

var (
	creditTypes        []int32
	registeredServices Services
	ginLambda          *ginadapter.GinLambda
)

func init() {
	creditTypes = []int32{300, 500, 700}
	registeredServices = GetDependencies(creditTypes...)
}

func main() {
	lambda.Start(mainHandler)
}

func mainHandler(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if ginLambda == nil {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.POST("/credit-assignment", handlerAssignment)
		r.POST("/credit-assignments", handlerAssignments)
		r.POST("/statistics", handlerStatistics)
		ginLambda = ginadapter.New(r)
	}

	return ginLambda.Proxy(request)
}

func handlerAssignments(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()

	var request models.InvestmentRequest
	if err := ctx.ShouldBindWith(&request, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}

	assignments, err := registeredServices.creditAssigner.Assignments(int32(request.Investment))
	if err != nil {
		registeredServices.transactionDal.AddTransaction(false)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, assignments)
	registeredServices.transactionDal.AddTransaction(true)
}

func handlerAssignment(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()

	var request models.InvestmentRequest
	if err := ctx.ShouldBindWith(&request, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}

	a300, a500, a700, err := registeredServices.creditAssigner.Assign(int32(request.Investment))
	if err != nil {
		registeredServices.transactionDal.AddTransaction(false)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	assignments := models.ToInvestmentResponse(a300, a500, a700)
	if err != nil {
		registeredServices.transactionDal.AddTransaction(true)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, assignments)
	registeredServices.transactionDal.AddTransaction(true)
}

func handlerStatistics(ctx *gin.Context) {
	err, transactions := registeredServices.transactionDal.GetStatistics()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrbryside/assessment/internal/config"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/expense"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"log"
)

func main() {
	// init viper config
	config.Init()

	// get config
	port := config.NewProvider().Port
	dbUrl := config.NewProvider().DbUrl

	// init db to Store
	db.InitDB(db.NewPostgres(dbUrl))

	// init echo framework
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = util.Validator(validator.New())

	// init handler
	expenses := expense.NewExpense(db.DB)

	// register routes
	e.POST("/expenses", expenses.CreateExpenseHandler)
	e.GET("/expenses/:id", expenses.GetExpenseHandler)
	e.PUT("/expenses/:id", expenses.UpdateExpenseHandler)

	log.Printf("Server started at %v\n", port)
	log.Fatal(e.Start(port))
}

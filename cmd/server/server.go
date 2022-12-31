package main

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/mrbryside/assessment/internal/config"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/expense"
	"github.com/mrbryside/assessment/internal/pkg/middleware"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// init viper config
	config.Init()

	// get config
	port := config.NewProvider().Port
	dbUrl := config.NewProvider().DbUrl

	// init db to Store
	database, err := db.InitDB(db.NewPostgres(dbUrl))
	if err != nil {
		panic("can't connect database")
	}
	defer database.Close()

	// init echo framework
	e := echo.New()
	e.Use(eMiddleware.Logger())
	e.Use(eMiddleware.Recover())
	e.Validator = common.Validator(validator.New())

	// init handler
	expenses := expense.NewExpense(db.DB)

	// auth middleware
	e.Use(middleware.VerifyAuthorization)

	// register routes
	e.POST("/expenses", expenses.CreateExpenseHandler)
	e.GET("/expenses", expenses.GetExpensesHandler)
	e.GET("/expenses/:id", expenses.GetExpenseHandler)
	e.PUT("/expenses/:id", expenses.UpdateExpenseHandler)

	// start the server in a separate goroutine
	go func() {
		log.Printf("Server started at %v\n", port)
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down server")
		}
	}()

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}

package app

import (
	"fmt"
	"log"
	"net/http"

	"test/config"
	"test/domain/groups"

	db2 "test/db"
	"test/http/controllers"
	"test/http/middleware"
	"test/logger"
	"test/repository"
	"test/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CreateRouter(globalCfg config.Config) *mux.Router {
	db := db2.GetDB()

	// --- repos ---
	usersRepo := repository.NewUsersDBRepo(db)

	variantsRepo := repository.NewVariantsDBRepo(db)
	testsUserRepo := repository.NewTestsUserDBRepo(db)
	tasksRepo := repository.NewTasksDBRepo(db)
	answerRepo := repository.NewAnswersDBRepo(db)
	resultsRepo := repository.NewReultsDBRepo(db)

	//--- services ---

	tokenService := services.NewJWTTokenService([]byte(globalCfg.JWTSecret))
	authSvc := services.NewAuthService(usersRepo, tokenService)
	testSvc := services.NewTestService(variantsRepo, testsUserRepo, answerRepo, tasksRepo, resultsRepo, *authSvc)

	//controllers

	authCtrl := controllers.NewAuthController(authSvc)
	testCtrl := controllers.NewTestController(testSvc)

	router := mux.NewRouter()
	authMW := middleware.NewAuthMiddleware(tokenService)
	routerAuth := authMW.RouterWithGroup(groups.User)

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", authCtrl.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", authCtrl.Login).Methods(http.MethodPost)

	userRouter := router.PathPrefix("/{user_id}").Subrouter()
	userRouter.Use(routerAuth)
	router.HandleFunc("/loginout", authCtrl.LoginOut).Methods(http.MethodPut)
	router.HandleFunc("/variant", testCtrl.ListVariants).Methods(http.MethodGet)
	router.HandleFunc("/variants/{variant_id}/tasks/{task_id}", testCtrl.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/variants/{variant_id}/tasks/{task_id}", testCtrl.AnswerTask).Methods(http.MethodPost)
	router.HandleFunc("/variants/{variant_id}/results", testCtrl.Result)

	return router
}

func Run(globalCfg config.Config) {
	router := CreateRouter(globalCfg)
	handler := cors.Default().Handler(router)

	logger.Info(fmt.Sprintf("Starting server on %s:%d",
		globalCfg.ServerHost, globalCfg.ServerPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",
		globalCfg.ServerHost, globalCfg.ServerPort), handler))
}

package main

import (
	"DBForum/internal/app/database"
	forumHandlers "DBForum/internal/app/forum/handlers"
	forumRepo "DBForum/internal/app/forum/repository"
	forumUCase "DBForum/internal/app/forum/usecase"
	postHandlers "DBForum/internal/app/post/handlers"
	postRepo "DBForum/internal/app/post/repository"
	postUCase "DBForum/internal/app/post/usecase"
	serviceHandlers "DBForum/internal/app/service/handlers"
	serviceRepo "DBForum/internal/app/service/repository"
	serviceUCase "DBForum/internal/app/service/usecase"
	"fmt"
	router2 "github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	threadHandlers "DBForum/internal/app/thread/handlers"
	threadRepo "DBForum/internal/app/thread/repository"
	threadUCase "DBForum/internal/app/thread/usecase"

	userHandlers "DBForum/internal/app/user/handlers"
	userRepo "DBForum/internal/app/user/repository"
	userUCase "DBForum/internal/app/user/usecase"

	"log"
	"net/http"
)

func main() {

	postgres, err := database.NewPostgres()

	if err != nil {
		log.Fatal(err)
	}

	forumRepository := forumRepo.NewRepo(postgres.GetPostgres())
	if err := forumRepository.Prepare(); err != nil {
		log.Fatalln(err)
	}
	postRepository := postRepo.NewRepo(postgres.GetPostgres())
	if err := postRepository.Prepare(); err != nil {
		log.Fatalln(err)
	}
	serviceRepository := serviceRepo.NewRepo(postgres.GetPostgres())
	if err := serviceRepository.Prepare(); err != nil {
		log.Fatalln(err)
	}
	threadRepository := threadRepo.NewRepo(postgres.GetPostgres())
	if err := threadRepository.Prepare(); err != nil {
		log.Fatalln(err)
	}
	userRepository := userRepo.NewRepo(postgres.GetPostgres())
	if err := userRepository.Prepare(); err != nil {
		log.Fatalln(err)
	}

	forumUseCase := forumUCase.NewUseCase(*forumRepository, *userRepository, *threadRepository)
	postUseCase := postUCase.NewUseCase(*postRepository, *userRepository, *threadRepository, *forumRepository)
	serviceUseCase := serviceUCase.NewUseCase(*serviceRepository)
	threadUseCase := threadUCase.NewUseCase(*threadRepository, *postRepository)
	userUseCase := userUCase.NewUseCase(*userRepository)

	forumHandler := forumHandlers.NewHandler(*forumUseCase)
	postHandler := postHandlers.NewHandler(*postUseCase)
	serviceHandler := serviceHandlers.NewHandler(*serviceUseCase)
	threadHandler := threadHandlers.NewHandler(*threadUseCase)
	userHandler := userHandlers.NewHandler(*userUseCase)

	r := router2.New()

	r.POST("/api/forum/create", forumHandler.Create)
	r.GET("/api/forum/{slug}/details", forumHandler.Details)
	r.POST("/api/forum/{slug}/create", forumHandler.CreateThread)
	r.GET("/api/forum/{slug}/users", forumHandler.GetUsers)
	r.GET("/api/forum/{slug}/threads", forumHandler.GetThreads)

	r.GET("/api/post/{id}/details", postHandler.GetInfo)
	r.POST("/api/post/{id}/details", postHandler.ChangeMessage)

	r.POST("/api/service/clear", serviceHandler.ClearDB)
	r.GET("/api/service/status", serviceHandler.Status)

	r.POST("/api/thread/{slug_or_id}/create", threadHandler.CreatePost)
	r.GET("/api/thread/{slug_or_id}/details", threadHandler.ThreadInfo)
	r.POST("/api/thread/{slug_or_id}/details", threadHandler.ChangeThread)
	r.GET("/api/thread/{slug_or_id}/posts", threadHandler.GetPosts)
	r.POST("/api/thread/{slug_or_id}/vote", threadHandler.VoteThread)

	r.POST("/api/user/{nickname}/create", userHandler.CreateUser)
	r.GET("/api/user/{nickname}/profile", userHandler.GetUserInfo)
	r.POST("/api/user/{nickname}/profile", userHandler.ChangeUser)

	fmt.Printf("Starting server on port %s\n", ":5000")
	if err := fasthttp.ListenAndServe(":5000", r.Handler); err != nil {
		log.Fatal(err)
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func LoggingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"url":    r.URL,
			"method": r.Method,
			"body":   r.Body,
		}).Info()
		next.ServeHTTP(w, r)
	})
}

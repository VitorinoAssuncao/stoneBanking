package server

import (
	"log"
	"net/http"
	"stoneBanking/app/domain/entities/token"
	accounts "stoneBanking/app/gateway/web/account"
	transfers "stoneBanking/app/gateway/web/transfer"

	"github.com/gorilla/mux"
)

type Server struct {
	Router mux.Router
}

func New(usecase *UseCaseWrapper, tokenRepository token.Repository) *Server {
	router := mux.NewRouter().StrictSlash(true)
	controller_account := accounts.New(usecase.Accounts, tokenRepository)
	controller_transfer := transfers.New(usecase.Transfer, tokenRepository)
	router.HandleFunc("/account", controller_account.Create).Methods("POST")
	router.HandleFunc("/account/login", controller_account.LoginUser).Methods("POST")
	router.HandleFunc("/accounts", controller_account.GetAll).Methods("GET")
	router.HandleFunc("/account/balance", controller_account.GetBalance).Methods("GET")
	router.HandleFunc("/transfer", controller_transfer.Create).Methods("POST")
	router.HandleFunc("/transfer", controller_transfer.GetAllByAccountID).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
	server := Server{Router: *router}
	return &server
}

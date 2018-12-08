package main

import (
	"fmt"
	"net/http"

	"github.com/devdinu/slot_machine/config"
	"github.com/devdinu/slot_machine/game"
	"github.com/devdinu/slot_machine/handler"
	"github.com/devdinu/slot_machine/machine"
	"github.com/devdinu/slot_machine/score"
	"github.com/gorilla/mux"
)

func main() {
	err := config.Load("config.yaml")
	if err != nil {
		panic(fmt.Errorf("[Config] loading configuration failed: %v", err))
	}

	fmt.Println("Starting server at: ", config.Address())
	http.ListenAndServe(config.Address(), router())
}

func router() *mux.Router {
	mux := mux.NewRouter()
	gameServer, err := buildGameServer()
	if err != nil {
		panic(fmt.Errorf("Error building Game Server: %v", err))
	}
	mux.HandleFunc("/ping", Ping).Methods("GET", "HEAD")
	if config.AuthEnabled() {
		gameServer = handler.Authenticator(gameServer)
	}
	mux.Handle("/api/machines/atkins-diet/spins", gameServer).Methods("POST")
	return mux
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"response":"pong"}`))
}

func buildGameServer() (http.Handler, error) {
	gameCfg := config.Gaming()
	stopper, err := machine.NewRandomStopper(config.StopperLimit())
	if err != nil {
		return nil, err
	}
	scorer := score.NewScorer()
	machineCfg := machine.Config{
		ReelsOfSymbols: gameCfg.ReelsOfSymbols,
		Rows:           gameCfg.Rows,
	}
	mach := machine.NewMachine(stopper, machineCfg)
	service := game.NewService(mach, scorer)
	return game.NewServer(service, &handler.Refresher{}), nil
}

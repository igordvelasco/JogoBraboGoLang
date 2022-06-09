package main

import "net/http"

func (app *application) routes() http.Handler {
  
  mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/dificuldades", app.dificuldades)
	mux.HandleFunc("/leaderboard", app.leaderboard)
	mux.HandleFunc("/images", app.showImage)
  //mux.HandleFunc("/forms", app.Forms)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  return app.logRequest( secureHeaders( mux ) )
}
//https://cc5m.ricardomendes2.repl.co
package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/rmcs87/cc5m/pkg/models"
)

// func (app *application) Forms(rw http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/forms" {
// 		app.notFound(rw)
// 		return
// 	}
  


//   tmpl = template.Must(template.ParseFiles("./ui/html/dificuldades.page.tmpl.html"))
//   if r.Method != http.MethodPost {
//     tmpl.Execute(rw, nil)
//     return
//   }
// }

func (app *application) home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(rw)
		return
	}
  
  allCat,_ := app.images.AllCat()

	files := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
  
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, allCat)
	if err != nil {
		app.serverError(rw, err)
		return
	} 
}

func (app *application) dificuldades(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dificuldades" {
		app.notFound(rw)
		return
	}

  tmpl := template.Must(template.ParseFiles("./ui/html/forms.page.tmpl.html"))
  if r.Method != http.MethodPost {
    tmpl.Execute(rw, nil)
    return
  }

  details := r.FormValue("name")
  fmt.Printf("%s", details)
  app.points.Insert(details)
  tmpl.Execute(rw, details)
  
  allCat,_ := app.images.AllCat()

	files := []string{
		"./ui/html/dificuldades.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, allCat)
	if err != nil {
		app.serverError(rw, err)
		return
	} 
}

func (app *application) leaderboard(rw http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/leaderboard" {
  		app.notFound(rw)
  		return
  	}
  
    // images, err := app.images.LatestCat("ACAO")
    // if err != nil{
    //   app.serverError(rw, err)
    //   return
    // }
  
    ranking, _ := app.points.Latest()
  
  	files := []string{
  		"./ui/html/leaderboard.page.tmpl.html",
  		"./ui/html/base.layout.tmpl.html",
  		"./ui/html/footer.partial.tmpl.html",
  	}
  
  	ts, err := template.ParseFiles(files...)
  	if err != nil {
  		app.serverError(rw, err)
  		return
  	}
  	err = ts.Execute(rw, ranking)
  	if err != nil {
  		app.serverError(rw, err)
  		return
  	} 
}

//http://localhost:4000/snippet?id=1

func (app *application) showImage(rw http.ResponseWriter, r *http.Request) {
	
  categoria := r.URL.Query().Get("categoria")
  dificuldade := r.URL.Query().Get("dificuldade")
  
	if categoria == ""{
		app.notFound(rw)
		return
	}
  if dificuldade == ""{
		app.notFound(rw)
		return
	}

  s, err := app.images.LatestCatDif(categoria, dificuldade)
  if err == models.ErrNoRecord {
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
		"./ui/html/show.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
  s_random := rand.Intn(len(s))
	err = ts.Execute(rw, s[s_random])
	if err != nil {
		app.serverError(rw, err)
	}
}

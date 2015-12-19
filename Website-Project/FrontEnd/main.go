package main

import (
	"html/template"
	"net/http"
  )

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", cover)
	r.GET("/browse", browse)
	r.GET("/view", view)
	r.GET("/write", write)
	r.GET("/profile", profile)
	r.GET("/login", login)
	r.GET("/signup", signup)
	
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	tpl = template.New("roottemplate")
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
  }

func browse(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "browse.html", nil)
  }

func view(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "view.html", nil)
  }

func write(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "write.html", nil)
  }

func profile(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "profile.html", nil)
  }

func login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
  }

func signup(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
  }

func cover(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "cover.html", nil)
  }

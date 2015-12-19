package main

import (
	"html/template"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine"
  )

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", cover)
	r.GET("/browse", browse)
	r.GET("/view", view)
	r.GET("/write", write)
	r.GET("/user/:name", profile)
	r.GET("/login", login)
	r.GET("/signup", signup)
	r.GET("/logout", logout)
	r.POST("/api/login", loginProcess)
	r.POST("/api/signup", createUser)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	tpl = template.New("roottemplate")
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
  } 

func browse(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo (req)
	tpl.ExecuteTemplate(res, "browse.html", &sd)
  }

func view(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo (req)
	tpl.ExecuteTemplate(res, "view.html", &sd)
  }
  
func write(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo (req)
	tpl.ExecuteTemplate(res, "write.html", &sd)
  }


func cover(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo (req)
	tpl.ExecuteTemplate(res, "cover.html", &sd)
  }

func sessionInfo(req *http.Request) SessionData {

	memItem, err := getSession(req)
	var sd SessionData
	if err == nil {
	
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	  }
	return sd
  }

func getSession(req *http.Request) (*memcache.Item, error) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}, err
	  }

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	  }
	return item, nil
  }

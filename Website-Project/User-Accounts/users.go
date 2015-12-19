package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"net/http"
	"time"
  )

func profile(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo(req)
	var user User
	user.Username = ps.ByName("name")
	
	if user.Username != sd.Username {
	
		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Users", user.Username, 0, nil)
		err := datastore.Get(ctx, key, &user)
		
		if err != nil {
			panic(err)
		  }
	    } else {
	    	user = sd.User
	      }
    	sd.ViewingUser = user
    	tpl.ExecuteTemplate(res, "profile.html", &sd)
  }

func login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
  }

func signup(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
  }

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
	
		fmt.Fprint(res, "false")
		return
  	} else {
	  	fmt.Fprint(res, "true")
	  }
  }

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
	
	if err != nil {
		log.Errorf(ctx, "error creating password: %v", err)
		http.Error(res, err.Error(), 500)
		return
	  }
	
	t := time.Now()
	y, m, d := t.Date()
	s := fmt.Sprintf ("%v %v, %v", m, d, y)
	
	user := User{
		Email: req.FormValue("email"),
		Name: req.FormValue("name"),
		Username: req.FormValue("username"),
		About: req.FormValue("about"),
		Image: req.FormValue("image"),
		Password: string(hashedPass),
		JoinDate: s,
  	}
	
	key := datastore.NewKey(ctx, "Users", user.Username, 0, nil)
	key, err = datastore.Put(ctx, key, &user)

	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
  	}

	createSession(res, req, user)

	http.Redirect(res, req, "/", 302)
  }

func loginProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", req.FormValue("username"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("password"))) != nil {
	
		var sd SessionData
		sd.LoginFail = true
		tpl.ExecuteTemplate(res, "login.html", sd)
		return

	  } else {
	  	user.Username = req.FormValue("username")
	
	  	createSession(res, req, user)
	
	  	http.Redirect(res, req, "/", 302)
  	}
  }

func createSession(res http.ResponseWriter, req *http.Request, user User) {
	ctx := appengine.NewContext(req)

	id, _ := uuid.NewV4()
	
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",

  	}
	http.SetCookie(res, cookie)


	json, err := json.Marshal(user)

	if err != nil {
		log.Errorf(ctx, "error marshalling during user creation: %v", err)
		http.Error(res, err.Error(), 500)
		return
  	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
  	}
	memcache.Set(ctx, &sd)
  }

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")

	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	  }

	
	sd := memcache.Item{
		Key:        cookie.Value,
		Value:      []byte(""),
		Expiration: time.Duration(1 * time.Microsecond),
  	}
	memcache.Set(ctx, &sd)


	cookie.MaxAge = -1
	http.SetCookie(res, cookie)


	http.Redirect(res, req, "/login", 302)
  }

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

func follow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo(req)
	
	f := Follow {
		Follower: sd.UserName, 
		Following: ps.ByName("name"),
	  }
	
	
	if f.Follower != f.Following {
		ctx := appengine.NewContext(req)
		key := datastore.NewIncompleteKey(ctx, "Follow", nil)
		_, err := datastore.Put(ctx, key, &f)
		
		if err != nil {
			panic(err)
	  	}
	  }
	
	
	http.Redirect(res, req, "/", 302)
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
	
	user := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: string(hashedPass),
	  }
	
	key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
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
	key := datastore.NewKey(ctx, "Users", req.FormValue("userName"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("password"))) != nil {
	
		var sd SessionData
		sd.LoginFail = true
		tpl.ExecuteTemplate(res, "login.html", sd)
		return
	    } else {
		    user.UserName = req.FormValue("userName")
		
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

	
	http.Redirect(res, req, "/", 302)
  } 


func tweetProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	
	if err != nil {
		log.Infof(ctx, "Attempt to post tweet from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	  }
	
	
	var user User
	json.Unmarshal(memItem.Value, &user)
	
	
	log.Infof(ctx, user.UserName)
	
	tweet := Tweet{
		Msg: req.FormValue("tweet"),
		Time: time.Now(),
		UserName: user.UserName,
	  }
	
	err = putTweet(req, &user, &tweet)
	
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	  }
	
	time.Sleep(time.Millisecond * 500) 
	http.Redirect(res, req, "/", 302)
  }

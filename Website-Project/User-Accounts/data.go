package main

import "time"

type User struct {
	Email    string
	Name string
	Username string `datastore:"-"`
	Password string `json:"-"`
	About string
	Image string
	JoinDate string
  }

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	Stories []Story
	ViewingUser User
  }

type Scene struct {
	Text string
	Time time.Time
	Username string
  }

type Story struct {
	Name string
	Owner User
	Scenes []Scene
  }

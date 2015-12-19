package main

import "time"

type User struct {
	Email    string
	UserName string
	Password string
  }

type SessionData struct {
	User
	LoggedIn  bool
	Stories []Story
  }

type Scene struct {
	Text string
	Time time.Time
	UserName string
  }

type Story struct {
	Name string
	Owner User
	Scenes []Scene
  }

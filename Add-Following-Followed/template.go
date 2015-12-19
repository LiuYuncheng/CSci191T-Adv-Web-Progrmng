package main

import (
	"encoding/json"
	"net/http"
  )

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem, err := getSession(req)
	
	if err != nil {
	
		tpl.ExecuteTemplate(res, templateName, SessionData{})
		return
	  }
	
	var sd SessionData
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true
	tpl.ExecuteTemplate(res, templateName, &sd)
  }

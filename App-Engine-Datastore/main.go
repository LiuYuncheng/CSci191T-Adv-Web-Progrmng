package main

import (
  "fmt"
  "net/http"
  "strings"
  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
)


func init(){
  http.HandleFunc("/", handleImages)
}

type Image struct{
  Subject string
  URL string
}

func handleImages(res http.ResponseWriter, req *http.Request){

  if req.URL.Path != "/" {
    subject := strings.Split(req.URL.Path, "/")[1]
    showImage(res, req, subject)
    return
  }

  if req.Method == "POST" {
    saveImage(res, req)
    return
  }

  listImages(res, req)

}

func listImages(res http.ResponseWriter, req *http.Request){

  ctx := appengine.NewContext(req)

  q := datastore.NewQuery("Image").Order("Subject")

  html := "<h2>Database</h2>"

  iterator := q.Run(ctx)

  for{
  var entity Image
  _, err := iterator.Next(&entity)
  if err == datastore.Done{
    break
    } else if err != nil {
      http.Error(res, err.Error(), 500)
      return
    }
  
  html += `
    <h3>` + entity.Subject + `</h3>
    <br/>
    <img src = '` + entity.URL + `' />
  <br/> 
  `
  
  }
  

  res.Header().Set("Content-Type", "text/html")

  fmt.Fprintln(res, `
  
    <!DOCTYPE html>
  
    <html>
    
    <body>
      <h1>Image/Subject Database for Computer Recognition Training</h1>
      <h2>Submit New</h2>
    
      <form method="POST">
      <table>
      
        <tr>
          <td><label for = "subjec">Subject</label></td>
          <td><input type="text" name="url"></td>
        </tr>
      
        <tr>
          <td><label for="url">Image URL</label></td>
          <td><input type="text" name="url"></td>
        </tr>
      
        <tr>
          <td></td>
          <td><input type = "submit"></td>
        </tr>
      
      </table>
      
      </form>
    
      <dl>
        `+html+`
      </dl>
    
    </body>
    
    
    </html>
    
    
  `)

}

func showImage(res http.ResponseWriter, req *http.Request, subject string){

  ctx := appengine.NewContext(req)

  key := datastore.NewKey(ctx, "Image", subject, 0, nil)

  var entity Image

  err := datastore.Get(ctx, key, &entity)

  if err == datastore.ErrNoSuchEntity{
  
    http.NotFound(res, req)
  
    return
  
    } else if err != nil{
    
      http.Error(res, err.Error(), 500)
    
      return
    
    }

  res.Header().Set("Content-Type", "text/html")

  fmt.Fprintln(res, `
  
    <h2>` + entity.Subject + `</h2>
  
    <br/>
    
      <img src='` + entity.URL + `' />
    
    <br/>
  
  `)

}

func saveImage(res http.ResponseWriter, req *http.Request){

  subject := req.FormValue("subject")

  url := req.FormValue("url")

  ctx := appengine.NewContext(req)

  key := datastore.NewKey(ctx, "Image", subject, 0, nil)

  entity := Image{
    Subject: subject, 
    URL: url,
  }

  _, err := datastore.Put(ctx, key, &entity)

  if err != nil {
    http.Error(res, err.Error(), 500)
    return
  }

  http.Redirect(res, req, "/", 302)

}

package main

import (

	// "os"

	"html"
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/Ethanol48/medium-api-library/article"
	"github.com/Ethanol48/medium-api-library/user"
)

type ApiResponse struct {
    Message string `json:"message,omitempty"`
    Data    any    `json:"data,omitempty"` // Use `any` for flexible data types or replace with specific types
}


func main()  {
  mux := http.NewServeMux()

  /* user funcs */
  mux.HandleFunc("GET /user/metadata", func(w http.ResponseWriter, r *http.Request) {
    // url parameter
    usr := r.URL.Query().Get("usr")
    if (usr == "") {
      w.WriteHeader(422)
      fmt.Fprint(w, "we couldn't find the parameter 'usr' in your request")
      return
    }

    // create a user link
    userLink := fmt.Sprintf("https://medium.com/%s", usr)
    metadata := user.GetUserMetadata(userLink)


    fmt.Printf("metadata: %v\n", metadata)

    type MetadataResponse struct {
      Name string     `json:"name"`;
      Desc string    `json:"desc"`;
      About string  `json:"about"`;
      Followers string `json:"followers"`;
      Following string `json:"following"`;
    }

    resp := MetadataResponse{
      Name: metadata.Name,
      Desc: metadata.Desc,
      About: metadata.About,
      Followers: metadata.Followers, 
      Following: metadata.Following, 
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)


    json.NewEncoder(w).Encode(resp)
    
  })


  /* article funcs */
  mux.HandleFunc("GET /article/html", func(w http.ResponseWriter, r *http.Request) {
    // url parameter
    link := r.URL.Query().Get("link")
    if (link == "") {
      w.WriteHeader(422)
      fmt.Fprint(w, "we couldn't find the parameter 'link' in your request")
      return
    }

    // TODO: validation for link
    art := article.GetArticle(link)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)


    resp := ApiResponse{
      Message: "",
      Data: html.UnescapeString(art.ToHTML()),
    }

    json.NewEncoder(w).Encode(resp)
  })

  mux.HandleFunc("GET /article/markdown", func(w http.ResponseWriter, r *http.Request) {
    // url parameter
    link := r.URL.Query().Get("link")
    if (link == "") {
      w.WriteHeader(422)
      fmt.Fprint(w, "we couldn't find the parameter 'link' in your request")
      return
    }

    // TODO: validation for link
    art := article.GetArticle(link)

    resp := ApiResponse{
      Message: "",
      Data: html.UnescapeString(art.ToMarkdown()),
    }

    json.NewEncoder(w).Encode(resp)
  })

  mux.HandleFunc("GET /article/metadata", func(w http.ResponseWriter, r *http.Request) {

    // url parameter
    link := r.URL.Query().Get("link")
    if (link == "") {
      w.WriteHeader(422)
      fmt.Fprint(w, "we couldn't find the parameter 'link' in your request")
      return
    }

    // TODO: validation for link
    art := article.GetArticle(link)

    type MetadataResponse struct {
      Title string     `json:"title"`;
      Tags []string    `json:"tags"`;
      ReadTime string  `json:"readtime"`;
      Published string `json:"published"`;
    }

    resp := MetadataResponse{
      Title: art.Title,
      Tags: art.Tags,
      ReadTime: art.ReadTime,
      Published: art.Published,
    }

    json.NewEncoder(w).Encode(resp)
  })


  fmt.Println("Serving api @ http://0.0.0.0:8080")
  if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
    fmt.Println(err.Error())
  }
}

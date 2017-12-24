package main

import (
    "fmt"
    "net/http"
    "html/template"
    
    // "os"


    "github.com/go-martini/martini"


    // "github.com/Desten73/Go-site/models"
    // "models"
)


type Post struct {
    Id      string
    Title   string
    Content string
}

func NewPost(id, title, content string) *Post {
    return &Post{id, title, content}
}


var posts map[string]*Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
        return
    }

    t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
        return
    }

    t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
        return
    }

    id := r.FormValue("id")
    post, found := posts[id]
    if !found {
        http.NotFound(w, r)
        return
    }

    t.ExecuteTemplate(w, "write", post)

}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
    // m := martini.Classic()
    id := r.FormValue("id")
    title := r.FormValue("title")
    content := r.FormValue("content")

    var post *Post
    if id != "" {
        post = posts[id]
        post.Title = title
        post.Content = content
    } else {
        id = GenerateId()
        post := NewPost(id, title, content)
        posts[post.Id] = post
    }


    // m.Get("/", indexHandler)
    http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    if id == "" {
        http.NotFound(w, r)
        return
    }

    delete(posts, id)

    http.Redirect(w, r, "/", 302)
}

func main() {


    http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("./vendor/"))))

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))


    m := martini.Classic()
    m.Get("/", indexHandler)
    m.Get("/write", writeHandler)
    m.Get("/edit", editHandler)
    m.Get("/delete", deleteHandler)
    m.Get("/SavePost", savePostHandler)
    m.Run()

    // m := martini.Classic()
    // m.Get("/", func() string {
    //     return "Hello World marti"
    // })
    // m.Run()



}
package main

import (
    "fmt"
    "net/http"
    "html/template"

    "database/sql"

    _ "github.com/lib/pq"
    
    // "os"


    "github.com/go-martini/martini"


    // "github.com/Desten73/Go-site/models"
    // "models"
)


const (  
  host     = "ec2-23-23-245-89.compute-1.amazonaws.com"
  port     = 5432
  user     = "aiefinohpubdxt"
  password = "cc17ecb542a34940bd33694cfd066bd8af845f3ea0588d794e528b79551d983d"
  dbname   = "d67kuc4en8hccg"
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

type Table_view struct {
        id         string
        name        string
        description       string
        value   string
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


    // m.Get("/db_con", func() string {
    //     return "Hello World"
    // })


    m.Get("/db_con", func () string {

    fmt.Println("Подключение к БД!")
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
    panic(err)
     }

      defer db.Close()


     err = db.Ping()
     if err != nil {
    panic(err)
     }
     fmt.Println("Successfully connected!")
        

     var db_texts string

    rows, err := db.Query("SELECT id, name, description, value FROM Работа")
    bks := make([]*Table_view, 0)
    for rows.Next() {
        bk := new(Table_view)
        rows.Scan(&bk.id, &bk.name, &bk.description, &bk.value)
        bks = append(bks, bk)
    }
    for _, bk := range bks {

        m.Get("/db_con", func() string {
         return fmt.Sprintf ("Hello World")
     })
        db_texts+="[id]" + bk.id + "\n[name]" + bk.name + "[description]" + bk.description + "[value]" + bk.value+ "\n\n"
        fmt.Println(bk.id, bk.name, bk.description, bk.value)
    }


    return fmt.Sprintf (db_texts)

    })


    m.Get("/db_con/:name", func(params martini.Params) string {
    // (params["name"])



    fmt.Println("Подключение к БД!")
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
    panic(err)
     }

      defer db.Close()


     err = db.Ping()
     if err != nil {
    panic(err)
     }
     fmt.Println("Successfully connected!")
        

     var db_texts string

    rows, err := db.Query("SELECT id, name, description, value FROM Работа where id = " + params["name"])
    bks := make([]*Table_view, 0)
    for rows.Next() {
        bk := new(Table_view)
        rows.Scan(&bk.id, &bk.name, &bk.description, &bk.value)
        bks = append(bks, bk)
    }
    for _, bk := range bks {

        m.Get("/db_con", func() string {
         return fmt.Sprintf ("Hello World")
     })
        db_texts+="[id]" + bk.id + "\n[name]" + bk.name + "[description]" + bk.description + "[value]" + bk.value+ "\n\n"
        fmt.Println(bk.id, bk.name, bk.description, bk.value)
    }
    return fmt.Sprintf (db_texts)
    })

    m.Get("/name/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
    return fmt.Sprintf ("Hello %s", params["name"])
    })


    m.Run()




    

    // m := martini.Classic()
    // m.Get("/", func() string {
    //     return "Hello World marti"
    // })
    // m.Run()







}
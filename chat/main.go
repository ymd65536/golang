package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.temp1 =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	data := map[string]interface{}{
		"Host": req.Host,
	}
	if authCookie, err := req.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.temp1.Execute(w, data)
}

func main() {
	//ポート番号を自由に設定できるようにする。
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	var clientID = ""
	var key = ""

	gomniauth.SetSecurityKey(key)
	gomniauth.WithProviders(
		facebook.New(clientID, key, "http://localhost:8080/auth/callback/facebook/"),
		github.New(clientID, key, "http://localhost:8080/auth/callback/github/"),
		google.New(clientID, key, "http://localhost:8080/auth/callback/google/"),
	)

	r := newRoom()
	http.Handle("/chat", mustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()
	log.Println("Webサーバーを開始します。ポート： ", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

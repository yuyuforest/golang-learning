package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

type Comment struct {
	Nickname	string `json:"nickname"`
	Message 	string `json:"message"`
}

var comment Comment

func helloHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]
		formatter.Text(w, http.StatusOK, "Hello " + name + "!")
	}
}

func submitHandler(webRoot string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			req.ParseForm()
			comment.Nickname = req.FormValue("nickname")
			comment.Message = req.FormValue("message")
		}
		http.FileServer(http.Dir(webRoot + "/assets/")).ServeHTTP(w, req)
	}
}

func showHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, comment)
	}
}

func NotImplemented(w http.ResponseWriter, r *http.Request) { http.Error(w, "501 api not implemented", http.StatusNotImplemented) }

func NotImplementedHandler() http.Handler { return http.HandlerFunc(NotImplemented) }

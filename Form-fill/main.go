package main

import (
	"net/http"

	"github.com/joshua468/form-fill/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/submit_login", handlers.SubmitLoginHandler)
	http.HandleFunc("/view_submission", handlers.ViewSubmissionHandler)
	http.HandleFunc("/form_steps", handlers.FormStepsHandler)
	http.HandleFunc("/resume_form", handlers.ResumeFormHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

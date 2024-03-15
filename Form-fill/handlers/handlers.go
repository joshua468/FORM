package handlers

import (
	"html/template"
	"net/http"

	"github.com/joshua468/form-fill/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("handlers/templates/login.html"))
	tmpl.Execute(w, nil)
}

func SubmitLoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	submissionID, err := database.SaveLoginInfo(username, password)
	if err != nil {
		http.Error(w, "Failed to save login info", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Login info submitted successfully! SubmissionID: " + submissionID))
}

func ViewSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	submissionID := r.URL.Query().Get("id")
	submission, err := database.GetSubmission(submissionID)
	if err != nil {
		http.Error(w, "Failed to retrieve submission", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("handlers/templates/submission.html"))
	tmpl.Execute(w, submission)
}

func FormStepsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("handlers/templates/form_steps.html"))
	tmpl.Execute(w, nil)
}

func ResumeFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("handlers/templates/resume_form.html"))
	tmpl.Execute(w, nil)
}

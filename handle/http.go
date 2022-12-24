package handle

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	O "project_r/data"
)

type Message struct {
	Email    string
	Username string
	Errors   map[string]string
}
type User struct {
	Compte   bool
	Username string
	Email    string
	Password string
}

var oui User

var msg Message

func Index(w http.ResponseWriter, r *http.Request) {
	msg.Errors = make(map[string]string)
	tmpl := template.Must(template.ParseGlob("html/*"))
	var mail string
	var username string
	var password string
	if r.Method == http.MethodGet {
		fmt.Print("/   method = get Index  /")
	}

	if r.Method == http.MethodPost {

		M := r.FormValue("mail")
		A := r.FormValue("username")
		C := r.FormValue("password")
		mail += M
		username += A
		password += C
		ctx := context.Background()
		client := O.Createclient(ctx)
		fmt.Println("method : post " + " " + username + " " + password)
		cli := O.Check(client, w, r, ctx, username, mail, password)
		print(cli)
		if cli == 2 {
			http.Redirect(w, r, "/site", http.StatusSeeOther)
		}
		if cli == 0 {
			msg.Errors["Email"] = "Compte inexistant,réesayez"
			tmpl.ExecuteTemplate(w, "co.html", msg)
			return
		}
		if cli == 1 {
			msg.Errors["Email"] = "Mot de passe ou Email invalide"
			tmpl.ExecuteTemplate(w, "co.html", msg)
			return
		}
		if cli == 3 {
			msg.Errors["Email"] = "Veuillez remplir tout les champs "
			tmpl.ExecuteTemplate(w, "co.html", msg)
			return
		}
	}
	tmpl.ExecuteTemplate(w, "co.html", nil)
}

func Index1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("html/*"))
	tmpl.ExecuteTemplate(w, "registration.html", nil)
}

func Index2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("html/*"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func Index3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("html/*"))
	tmpl.ExecuteTemplate(w, "validation.html", nil)
}

func LogforRegsistration(w http.ResponseWriter, r *http.Request) {
	msg.Errors = make(map[string]string)
	tmpl := template.Must(template.ParseGlob("html/*"))
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur", http.StatusSeeOther)
	}
	user := r.FormValue("username")
	pswd := r.FormValue("password")
	pswd1 := r.FormValue("password1")
	email := r.FormValue("mail")
	t := r.Form["type"]
	pop := t[0]
	fmt.Print(user + "/" + email + "/" + pswd1 + "/" + pswd + "/" + pop)
	regs, user1, email1, pswd1, Compte := O.Hello(w, r, user, email, pswd, pop, pswd1)
	if regs == 0 {
		msg.Errors["Email"] = "Username et EMAIL DEJA PRIS"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 1 {
		msg.Errors["Email"] = "EMAIL DEJA PRIS"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 2 {
		msg.Errors["Email"] = "Username DEJA PRIS"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 4 {
		msg.Errors["Email"] = "Les mots de passes ne sont pas identiques"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 5 {
		msg.Errors["Email"] = "Veuillez remplir tous les champs"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 6 {
		msg.Errors["Email"] = "la taille du mot de passe doit etre superieur a 6"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 7 {
		msg.Errors["Email"] = "Veuillez choisir un type de compte"
		tmpl.ExecuteTemplate(w, "registration.html", msg)
		return
	}
	if regs == 3 {
		truuue := User{
			Compte:   Compte,
			Username: user1,
			Email:    email1,
			Password: pswd1}
		fmt.Println(truuue.Username)
		fmt.Println(truuue.Compte)
		fmt.Println(truuue.Email)
		fmt.Println(truuue.Password)
		tmpl.ExecuteTemplate(w, "validation.html", truuue)
	}
	tmpl.ExecuteTemplate(w, "z.html", nil)
}

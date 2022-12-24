package data

import (
	"context"
	"fmt"
	"log"
	"net/http"

	H "project_r/email"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Check(iencli *firestore.Client, w http.ResponseWriter, r *http.Request, ctx context.Context, a string, b string, c string) int {
	Comptesite := iencli.Collection("Compte")
	//tmpl := template.Must(template.ParseGlob("html/*"))
	//iter := iencli.Collection("Compte").Documents(ctx)
	doccc := Comptesite.Doc(a)
	docsnap, err := doccc.Get(ctx)
	dataMap := docsnap.Data()
	if a == "" || b == "" || c == "" {
		return 3
	}
	if err != nil {
		// TODO: Handle error.
		if status.Code(err) == codes.NotFound || dataMap == nil {
			fmt.Println("ERREUR/ le Compte n'existe pas")
			return 0

		}
	}
	fmt.Print(" le compte existe IL EXISTE")
	pswd, _ := docsnap.DataAt("Password")
	email, _ := docsnap.DataAt("Email")
	if pswd == c && email == b {
		fmt.Println("Tout es bon")
		fmt.Println("Connexion ...")
		return 2
	}
	return 1

	//	fmt.Print("IL EXISTE")
}

func Createclient(ctx context.Context) *firestore.Client {
	projectID := "aaaaa-fbebb"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func Hello(w http.ResponseWriter, r *http.Request, a string, b string, c string, d string, e string) (int, string, string, string, bool) {
	var Compte bool
	//tmpl := template.Must(template.ParseGlob("html/*"))
	// chek if username et email deja utilisé
	ctx := context.Background()
	client := Createclient(ctx)
	Comptesite := client.Collection("Compte")
	doccc := Comptesite.Doc(a)
	if d != "Etudiant" {
		Compte = true
	} else {
		Compte = false
	}
	if a == "" || b == "" || c == "" || d == "" || e == "" {
		return 5, a, b, c, Compte
	}
	if len(e) < 6 || len(c) < 6 {
		return 6, a, b, c, Compte
	}
	if d == "Type de Compte" {
		return 7, a, b, c, Compte
	}
	if checkifUsernameisUsed(ctx, a, client) && checkifEmailisUsed(ctx, b, client) {
		fmt.Println("Username et EMAIL DEJA PRIS")
		return 0, a, b, c, Compte

	}
	if checkifEmailisUsed(ctx, b, client) && !checkifUsernameisUsed(ctx, a, client) {
		fmt.Println("EMAIL DEJA PRIS")
		return 1, a, b, c, Compte
	}
	if checkifUsernameisUsed(ctx, a, client) && !checkifEmailisUsed(ctx, b, client) {
		fmt.Println("Username DEJA PRIS")

		return 2, a, b, c, Compte
	}

	if !checkifEmailisUsed(ctx, b, client) && !checkifUsernameisUsed(ctx, a, client) && e == c {
		createAccount(doccc, ctx, Compte, b, e, a)
		code, err := H.GenerateRandomCode()
		if err != nil {
			fmt.Println("Erreur")
		}
		fmt.Println(code, err)
		H.Sender("Welcome", "/html/email.html", []string{b}, a, code)
		fmt.Println("return 3")
		return 3, a, b, c, Compte
	}
	fmt.Println("return 4")
	return 4, a, b, c, Compte

}

func checkifEmailisUsed(ctx context.Context, a string, client *firestore.Client) bool {
	iter := client.Collection("Compte").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		cap, err := doc.DataAt("Email")
		if err != nil {
			// TODO: Handle error.
			break
		}
		result := cap.(string)
		if result == a {
			return true
		}
	}
	return false

}

func checkifUsernameisUsed(ctx context.Context, a string, client *firestore.Client) bool {
	iter := client.Collection("Compte").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		cap, err := doc.DataAt("Username")
		if err != nil {
			// TODO: Handle error.
			break
		}
		result := cap.(string)
		if result == a {
			return true
		}
	}
	return false

}

func createAccount(a *firestore.DocumentRef, ctx context.Context, account bool, mail string, Pswd string, Username string) {
	_, err := a.Create(ctx, map[string]interface{}{
		"Compte":   account,
		"Email":    mail,
		"ID":       10,
		"Password": Pswd,
		"Username": Username,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	fmt.Println("Compte crée")

}

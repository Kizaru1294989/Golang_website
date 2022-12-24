package verifemail

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/smtp"
	"text/template"
)

type Info struct {
	Username string
	code     string
}

func Sender(subject string, templatepath string, to []string, Username string, code string) {
	final_info := Info{
		Username: Username,
		code:     code,
	}
	var body bytes.Buffer
	fmt.Println("1")
	tmpl, err1 := template.ParseFiles(templatepath)
	fmt.Println("2")
	tmpl.Execute(&body, final_info)
	fmt.Println("3")
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "subject:" + subject + "\n" + headers + "\n\n" + body.String()
	authentification := smtp.PlainAuth(
		"",
		"#",
		"#",
		"smtp.gmail.com",
	)
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		authentification,
		"chimiflyimik@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println("erreur", err)
	}
	if err1 != nil {
		fmt.Println("erreur", err1)
	}
}

const otpChars = "1234567890"

func GenerateRandomCode() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < 6; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

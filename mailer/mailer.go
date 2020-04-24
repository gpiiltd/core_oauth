package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"time"
	"crypto/tls"
	gomail "gopkg.in/gomail.v2"
)

var auth = smtp.PlainAuth("", "info@my-gpi.io", "Kingstarr001", "mail.my-gpi.io")

// Data for composing mail data that will be made available in the mail template
type Data struct {
	//From      string
	Names      string
	CreatedAt time.Time
}

// Request a request object model
type Request struct {
	from       string
	to         string
	subject    string
	body       string
	attachment []string
}

// NewRequest for creating new Request object
func NewRequest(from string, to string, subject string, attachment []string) *Request {
	return &Request{
		from: 		from,
		to:         to,
		subject:    subject,
		attachment: attachment,
	}
}

// sendEmail for setting up email parameters
func (r *Request) sendEmail(data Data) bool {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", r.from, "GPI Mailer")
	m.SetHeader("To", r.to)
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)
	for _, v := range r.attachment {
		m.Attach(v)
	}

	d := gomail.NewDialer("mail.my-gpi.io", 25, "info@my-gpi.io", "Kingstarr001")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Send for sending out email
func (r *Request) Send(tempName string, data Data) {
	err := r.ParseTemplate(tempName, data)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	if ok := r.sendEmail(data); ok {

	} else {
		fmt.Printf("Failed to send the email to %s\n", r.to)
	}
}

// ParseTemplate for parsing email template
func (r *Request) ParseTemplate(templateFileName string, data Data) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"

    "fmt"
    "log"
    "os"
    "net/smtp"
    "net/mail"
    "strconv"
    "strings"
)

// Field names should start with an uppercase letter
type ContactFormDetails struct {
    Name string `json:"name" xml:"name" form:"name"`
    Email string `json:"email" xml:"email" form:"email"`
    Message string `json:"message" xml:"message" form:"message"`
}


func checkEmailServer(domain string) error {
	// Set up SMTP client
	smtpAddr := fmt.Sprintf("%s:25", domain)
	smtpClient, err := smtp.Dial(smtpAddr)
	if err != nil {
		return err
	}
	defer smtpClient.Close()
    return nil
}

func sendEmailHandler(c *fiber.Ctx) error {    
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    smtpUsername := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_APP_PASSWORD")

    contactForm := new(ContactFormDetails)
    if bodyErr := c.BodyParser(contactForm); bodyErr != nil {
        return bodyErr
    }

    // Get the fields from the request
    name := contactForm.Name
    email := contactForm.Email
    message := contactForm.Message

    // Create the email message
    from := mail.Address{Name: name, Address: email}
    to := mail.Address{Name: "Ran", Address: smtpUsername}
    subject := "New Message from Contact Form"
    body := "Name: " + name + "\nEmail: " + email + "\nMessage: " + message

    // Set up the SMTP server
    smtpServer := "smtp.gmail.com"
    smtpPort := 587
    
    domain := strings.Split(email, "@")[1]
    mail_err := checkEmailServer(domain) 
    if mail_err != nil {
        // email server failed, domain does not exist
        return c.Status(fiber.StatusInternalServerError).SendString(mail_err.Error())
    } else {
        // Connect to the SMTP server
        auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
        toAddr := []string{to.Address}
        msg := []byte("To: " + to.String() + "\r\n" +
            "From: " + from.String() + "\r\n" +
            "Subject: " + subject + "\r\n" +
            "\r\n" +
            body + "\r\n")

        smtpErr := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, from.Address, toAddr, msg)
        if smtpErr != nil {
            // Handle the error
            log.Println("Failed to send email:", smtpErr)
            return smtpErr        
	    }
        return c.SendStatus(fiber.StatusOK)
    }
}


func main() {
    app := fiber.New()

    app.Static("/static_test", "./public")

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	// Respond with "Hello, test!" on root path, "/test"
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Test!")
	})

    // Send email with 3 fields to my gmail account
    app.Post("/contact", sendEmailHandler)

    app.Listen(":3000")
}
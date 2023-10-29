package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"

    "log"
    "os"
    "net/smtp"
    "net/mail"
    "strconv"
)

func sendEmailHandler(c *fiber.Ctx) error {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    smtpUsername := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_APP_PASSWORD")

    // Get the fields from the request
    name := c.FormValue("name")
    email := c.FormValue("email")
    message := c.FormValue("message")
    log.Println(name, email, message)

    // Create the email message
    from := mail.Address{Name: name, Address: email}
    to := mail.Address{Name: "Ran", Address: smtpUsername}
    subject := "New Message from Contact Form"
    body := "Name: " + name + "\nEmail: " + email + "\nMessage: " + message

    // Set up the SMTP server
    smtpServer := "smtp.gmail.com"
    smtpPort := 587
    
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
package mailSetup

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"net/http"
	"net/smtp"
	//"gopkg.in/gomail.v2"
)



func SendMail(to string, subject string, body string) {
	
	// d := gomail.NewDialer("smtp.google.com", 587, "noreply.team71@gmail.com");
	
	// m := gomail.NewMessage();
	// m.SetHeader("From", "noreply.team71@gmail.com");
	// m.SetHeader("To", to);
	// m.SetHeader("Subject", subject);
	// m.SetBody("text/plain", body)

	// err := d.DialAndSend(m);

	// if err != nil {
	// 	fmt.Println("Email send error: ", err.Error())
	// }else {
	// 	fmt.Printf("Email send: %s\n", to)
	// }

	from := "noreply.team71@gmail.com"
    pass := os.Getenv("GMAIL_SECRET")
    
    msg := "From: Team 71<" + from + ">\n" +
        "To: " + to + "\n" +
        "Subject: "+subject+"\n\n" +
        body

    err := smtp.SendMail("smtp.gmail.com:587",
        smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
        from, []string{to}, []byte(msg))

    if err != nil {
        fmt.Printf("smtp error: %s", err)
        return
    }
    fmt.Println("Successfully sended to: " + to)
}




func encryptData(data string) (string) {
	password := os.Getenv("Email_Pass")
	// Convert data and password to byte arrays
	dataBytes := []byte(data)
	passwordBytes := []byte(password)

	// Base64 encode the data and password
	encodedData := base64.StdEncoding.EncodeToString(dataBytes)
	encodedPassword := base64.StdEncoding.EncodeToString(passwordBytes)

	// Extract a substring from the encoded data
	substring := encodedData[20:30]

	// Calculate the length-based value
	lengthValue := len(encodedData) * 5

	// Format the pass string
	pass := fmt.Sprintf("%s:%d", substring+encodedPassword, lengthValue)

	// Concatenate the encoded data with the formatted pass string
	result := encodedData + "{:+}" + pass

	return result
}

func SendMail2(usermail, subject, content string) error {
	// Prepare the data
	params := map[string]interface{}{
		"to":         usermail,
		"subject":    subject,
		"message":    content,
		"from":       "rmtomal@team71.link", // Replace with your MAIL_FROM
		"companyname": "Team 71",
		"timestamp":  time.Now().UnixMilli(),
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	fmt.Printf("Enc D: %s", encryptData(string(jsonData)) )

	// Prepare the payload
	payload := map[string]string{
		"encrypted": encryptData(string(jsonData)),
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Make the HTTP POST request
	url := "https://sender.team71.link/api/v1/mail-send" // Replace with your MAILER_URL
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response:", string(body))
	return nil
}

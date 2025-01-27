package mailSetup

import (
	"fmt"
	"os"

	"net/smtp"
	//"gopkg.in/gomail.v2"
)



func SendMail(to string, subject string, body string) {
	
	// d := gomail.NewDialer("smtp.google.com", 587, "noreply.team71@gmail.com", "ppsd eeoj qrdy ggsa");
	
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
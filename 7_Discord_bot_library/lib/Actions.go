package main;

import (
    "fmt";
    "net/smtp";
)

func SendEmailAlert() (error) {
    msg := []byte("To: nishan@vairav.net\r\n" +
    "Subject: discount Gophers!\r\n" +
    "\r\n" +
    "This is the email body.\r\n")
    auth := smtp.PlainAuth("","","","smtp.gmail.com");
    err := smtp.SendMail("smtp.gmail.com:587",auth,"@gmail.com",[]string{"nishan@vairav.net"},msg);
    if err != nil {
        fmt.Println("Error sending message",err);
        return err;
    }
    return nil;
}

func main() {
    SendEmailAlert()
}

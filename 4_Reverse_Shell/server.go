package main;

import (
  "fmt";
  "net";
  "log";
  "os";
  "bufio";
  "io";
)

// mai() function handles everything
func main() {

  // Creating a listener
  ln, err := net.Listen("tcp",":8080");
  if err != nil {
    log.Fatalln("Could not create a TCP Listener");
  }

  // Defering the close
  defer func(){
    ln.Close();
    }();
    fmt.Println("Ready to accept connection");

    // I will accept one connection for this time
    conn, err := ln.Accept();
    defer func(){
      conn.Close();
      }();
      fmt.Println("Established a Connection from bot ");

    // Getting the initial message from bot
    msg, err := bufio.NewReader(conn).ReadString('\n');
    if err != nil {
      fmt.Println("Error, ",err);
      return;
    }
    fmt.Println(msg);

    // Now the loop starts where we can actually send the command to the bot
    for { // Starts loop
      fmt.Print("Server: ");
      msg,err = bufio.NewReader(os.Stdin).ReadString('\n');  // Reading input from the terminal
      if err != nil {
        fmt.Println("Error recieveing input from u");
        return;
      }

      // Sending the data to bot
      io.WriteString(conn,msg);

      // Recieing the message from bot
      msg, err := bufio.NewReader(conn).ReadString('\n');
      if err != nil {
        fmt.Println("Error, ",err);
        return;
      }
      fmt.Println("Client: ",msg);


    } // Ends loop
}

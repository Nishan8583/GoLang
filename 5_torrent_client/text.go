package main;

import (
      "fmt"
      "net"
      "io"
)

func main() {
      conn, err := net.Dial("udp","62.138.0.158:6969")
      if err != nil {
            fmt.Println("Connection was not abe to establish",err);
      }
      defer conn.Close();
      fmt.Println("Connection Establishd");
      request := "GET /announce?compact=1&downloaded=0&info_hash=%0C%3C%E1%19g%29%C0%A8%AF%C7%90%5Ck%E9%B0%A5%92%7Co%BA&left=0&peer_id=%1Df%12%E1%B7%23%91Myu%40%29%0F%E0S%BC%DFJ%BFg&port=6009&uploaded=0 HTTP/1.1\r\nHost: 62.138.0.158:6969\r\n"

      io.WriteString(conn,request);

      fmt.Println("sent","Bytes successfully sent");
      response := []byte{};

      _,err = conn.Read(response);
      if err != nil {
            fmt.Println("Could not read from server",err);
      }
      fmt.Println("Response from server:",string(response));
}

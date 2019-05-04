package Tracker;

import (
      "net";
      "fmt";
      "io/ioutil";
)

func connect_tracker(tracker string) (error) {
      conn, err := net.Dial("udp",tracker);
      if err != nil {
            fmt.Println("Could not connect to server");
            return err;
      }
      fmt.Println("Connection with Server Estabished");
      defer conn.Close();
      //fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n");
      fmt.Fprintf(conn,"GET /announce?compact=0&downloaded=0&info_hash=%0C%3C%E1%19g%29%C0%A8%AF%C7%90%5Ck%E9%B0%A5%92%7Co%BA&left=0&peer_id=%93%0A%87%EA%BB%DE%C9%D5%3B%92%E3%98j%1Cg%B3%EA%97G-&port=6009&uploaded=0 HTTP/1.1\r\n\r\n")
//udp://163.172.180.68:80/announce?compact=1&downloaded=0&info_hash=%0C%3C%E1%19g%29%C0%A8%AF%C7%90%5Ck%E9%B0%A5%92%7Co%BA&left=0&peer_id=%93%0A%87%EA%BB%DE%C9%D5%3B%92%E3%98j%1Cg%B3%EA%97G-&port=6009&uploaded=0
      // conn.Read() reads a certain amount only, in the below case 1024 bytes
      //content := make([]byte,1024);
      //_,err = conn.Read(content);
      fmt.Println("Sending Data to server")
      content,err := ioutil.ReadAll(conn);
      if err != nil {
            fmt.Println("Could not read from server")
            return err;
      }
      fmt.Println(string(content));
      return nil;
}

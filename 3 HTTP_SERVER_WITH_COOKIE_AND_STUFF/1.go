package main;

import (
  "fmt";
  "net/http";
  "text/template";
  "github.com/satori/go.uuid";
  "io";
  "reflect";
)

type user struct {
  Username string
  Pass string
  Role string
}

var db_session = map[string]string{} //  sessionid to username
var user_session = map[string]user{} // username to userstruct


var tpl *template.Template;
//var userdb []string;
// init() function readies the template
func init() {
  tpl = template.Must(template.ParseGlob("./templates/*"));
  //db_session["root"] = "admin";
  user_session["admin"] = user{
        Username: "Nishan",
        Pass: "Password",
        Role: "Admin",
  }
  //userdb = append(userdb, "admin")
}

// index() function handles the default uri i.e "/"
func index(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w,"index.gohtml",nil);
}

func access_test(w http.ResponseWriter, req *http.Request) {
    c, err := req.Cookie("session-id");
    if err != nil {
        io.WriteString(w,"Please Login first");
        return;
    }
   // useranme := db_session[c.Value]
   if user_session[db_session[c.Value]].Role == "Admin" {
         io.WriteString(w,"You DO !!!!!!!!!!!!!! have access to this page")
      } else {
            io.WriteString(w,"Lol What are you trying to do, You do not have access lol");
      }
}

// login() function as the name suggests handles the login functionality
func login(w http.ResponseWriter, req *http.Request) {

  // Check if session id in cookie (User Already logged in?)
  c, err := req.Cookie("session-id");
  if req.Header.Get("Referer") == "http://localhost:8080/signup" {
        tpl.ExecuteTemplate(w,"login.gohtml",nil);
        return;
 }

  // If there was no session id, then user was not logged in
  if err != nil{
      fmt.Println("A user not logged in tried to access /login path, Now attempting to get form data ");
      username := req.FormValue("username");
      password := req. FormValue("password");

      fmt.Println("The submitted username and password: ",username,password);

      // If user suppllied no username and password
      if len(username) == 0 ||  len(password) == 0 {
          fmt.Println("No username or value supplied");
          tpl.ExecuteTemplate(w,"login.gohtml","ERROR, you have to put something in username and password");
          return;
      } else {
            _, ok := user_session[username];
            if !ok { // Checks if ther username is int database
                  fmt.Println("You have not signed up");
                  http.Redirect(w,req,"/signup",307);
                  return;
            }
            // Else if user did provide usrename and password generate session id
            if user_session[username].Pass != password {
                  io.WriteString(w,"You have supplied something wrong");
                  return;
                  }
            sid, err := uuid.NewV4();
            if err != nil {
              fmt.Println("Error while generating session id");
              return;
            }

            // If user info exists in the database
            //sid := "user-cookie"
            c = &http.Cookie{
              Name : "session-id",
              Value : sid.String(),
            }

            // Creating a user struct

            // Adding the user information in database
            //db_session[sid.String()] = username;
            db_session[sid.String()] = username;
            //user_session[username] = value

            // Actually setting the cookie
            http.SetCookie(w,c);



        }
        io.WriteString(w,c.Value);


  return;
  }
  fmt.Println("The user already had an session with cookie ",c);
}

// singup() handles the signup procedure
func signup(w http.ResponseWriter, req *http.Request) {

      // First check whether user already logged in
      c, err := req.Cookie("session-id");
      fmt.Println(c);
      fmt.Println(reflect.TypeOf(err));
      if err != nil {
            fmt.Println("There was no cookie");
            username := req.FormValue("firstname");
            password := req.FormValue("password");

            if len(username) == 0 || len(password) == 0 {
                  tpl.ExecuteTemplate(w,"signup.gohtml",nil)
                  return;
            }

                  //sid := "user-cookie"
                  //c = &http.Cookie{
                   // Name : "session-id",
                    //Value : sid.String(,
                    //Value: sid,
                  //}

                  // Creating a user struct
            value := user{
            Username: username,
            Pass: password,
            Role: "User",
            }

                  // Adding the user information in database
                  //db_session[sid.String()] = username;
                  //db_session[sid] = username;
            user_session[username] = value

                  // Actually setting the cookie
                  //http.SetCookie(w,c);


                  //tpl.ExecuteTemplate(w,"signup.gohtml",value);
            http.Redirect(w,req,"/login",307);
            fmt.Println("A new user was created : ",value);
            return;


      }

}

// signout() i got bored, so the user has to just user the whole link
func logout(w http.ResponseWriter, req *http.Request) {
      c, err := req.Cookie("session-id");
      if err != nil {
            http.Redirect(w,req,"/login",307);
            return;
      }
      c = &http.Cookie{
        Name : "session-id",
        Value : "-1",
        MaxAge : -1,
      }

      // Make sure to delete that first
      delete(db_session,c.Value);
      http.SetCookie(w,c);

}
func main() {
  http.HandleFunc("/",index);
  http.HandleFunc("/login",login);
  http.HandleFunc("/access",access_test);
  http.HandleFunc("/signup",signup);
  http.HandleFunc("/logout",logout);
  http.ListenAndServe(":8080",nil);
}

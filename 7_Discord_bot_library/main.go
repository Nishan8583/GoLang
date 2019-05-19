package main;

import (
    "fmt";
    "github.com/bwmarrin/discordgo";  // The Discord Go Library
    "strings";
    "io/ioutil";
    "os";
    "flag";

)

var threshold int;
var policyList []string;
var content []byte;
var policyBreakerList map[string]int;


// Prepration function that will get the policy
func init() {
      threshold = 1;
    //policyBreakerList["test"] = 1;
    // Declaring the flag so that the user can use it ti supply the file
    pFile := flag.String("file","policy.dat","Give the filename containing the policu Ex: -file=filename.dat")
    flag.Parse();

    // If no flag was supplied use the default policy.dat
    if (flag.NFlag() == 0) {
        fmt.Println("Using Default file since no file path was supplied");
        file, err := os.Open("policy.dat");
        if err != nil {
            fmt.Println("Error Opening Default policy file",err);
            return;
        }
        content,err = ioutil.ReadAll(file);
        if err != nil {
            fmt.Println(err);
            return
        }
    } else {  // In case the flag was supplied use that file path as policy.dat
        file, err := os.Open(*pFile);
        if err != nil {
            return;
        }
        content,err = ioutil.ReadAll(file);
        if err != nil {
            fmt.Println(err);
            return
        }
    }

    // Stribg Manipulation to ge the required policy manipulation
    policy_string := strings.Split(string(content),"\n");
    for _ ,value := range policy_string {
        if len(value) < 2 {
            break
        }
        temp := strings.Split(value,",")
        fmt.Println(temp[2])
        policyList = append(policyList,temp[0]);
    }
    fmt.Println(policyList);
}

// The Bot ID
var BotID string;

func main() {
    policyBreakerList = make(map[string]int);
    // Creating a bot type
    bot,err := discordgo.New("Bot <key here>");
    if err != nil {
        fmt.Println("An error has occured while trying to create the discord bot type",err);
        return;
    }

    // Getting Userinfo,  the bot ID will be necessary later
    user, err := bot.User("@me");  // We get ths information to hold bot idea
    if err != nil {
        fmt.Println("Error getting user information",err);
        return;
    }
    BotID = user.ID;

    // Whenever a message is recieved call this function
    bot.AddHandler(messageHandler);

    // First open a connection to the gateway, this was the problem I could not solve when writing my own library
    err = bot.Open()
    if err != nil {
        fmt.Println("An error while trying to open a connection",err);
        return
    }

    // Send the initial message
    Msg, err := bot.ChannelMessageSend("<Channel ID here>","Hello, The bot has activated itself");
    if err != nil {
        fmt.Println("Error sending Message",err);
    } else {
        fmt.Println(Msg);
    }

    // So that the program gets stuck, the handler function can continue doing its jov
    <-make(chan struct{});

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

    // If the message was sent by itself ignore
    if (m.Author.ID == BotID) {
        return;
    }
    fmt.Println(m.Content)
    if (checkVulgarList(m.Content) == true) {
        _,isPresent := policyBreakerList[m.Author.Username];
        if (isPresent){  // If username was present
            fmt.Println("This user has violated before");
            policyBreakerList[m.Author.Username] = policyBreakerList[m.Author.Username]+1;
        } else {
            fmt.Println("not Present");
            policyBreakerList[m.Author.Username] = 1;
        }

        msg := fmt.Sprintf("Why you gotta be vulgar %s Your count %d",m.Author.Username,policyBreakerList[m.Author.Username]);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)

        if (policyBreakerList[m.Author.Username] > threshold) {
             msg = fmt.Sprintf("You are in for some trouble %s",m.Author.Username);
             _,_ = s.ChannelMessageSend(m.ChannelID,msg)
       }
        return;
    }
    if (strings.Contains(m.Content,"bot")) {
        msg := fmt.Sprintf("You talking about me %s ?",m.Author.Username);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)
    }
}


func checkVulgarList(msg string) (bool){
    for _,value := range policyList {
        if (strings.Contains(strings.ToLower(msg),strings.ToLower(value))) {
            return true;
        }
    }

    return false;
}

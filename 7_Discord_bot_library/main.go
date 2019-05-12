package main;

import (
    "fmt";
    "github.com/bwmarrin/discordgo";  // The Discord Go Library
    "strings";
    "io/ioutil";
    "os";
    "flag";
    "strconv";
)

// Structure for holding policy info
type policy struct {
    word string;  // The word your policy wants to banned
    count int;  // The maximum thershold for the word count
    refresh int; // Refresh the user policy info in this time
}
var policyList []policy;
var content []byte;
var policyBreakerList map[string]int;


// Prepration function that will get the policy
func init() {

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
        v1,_ := strconv.Atoi(temp[1]);
        v2,_ := strconv.Atoi(temp[2]);
        policyList = append(policyList,policy{word : temp[0],count : v1,refresh : v2});
    }
    fmt.Println(policyList);
}

// The Bot ID
var BotID string;

func main() {
    policyBreakerList := make(map[string]int);
    // Creating a bot type
    bot,err := discordgo.New("Bot NTc2NjUwMzQ0ODc3MjYwODAw.XNZl4w.tUX0FBt87pXRt0Sow5M0k_SOcJ4");
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
    Msg, err := bot.ChannelMessageSend("576659713140457476","Hello, The bot has activated itself");
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
        msg := fmt.Sprintf("Why you gotta be vulgar %s Your count",m.Author.Username,policyBreakerList[m.Author.Username]);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)
        return;
    }
    if (strings.Contains(m.Content,"bot")) {
        msg := fmt.Sprintf("You talking about me %s ?",m.Author.Username);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)
    }
}


func checkVulgarList(msg string) (bool){
    for _,value := range policyList {
        if (strings.Contains(msg,value.word)) {
            return true;
        }
    }

    return false;
}

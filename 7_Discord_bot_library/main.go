package main;

import (
    "fmt";
    "github.com/bwmarrin/discordgo";  // The Discord Go Library
    "strings"
)

// The Bot ID
var BotID string;

func main() {
    bot,err := discordgo.New("Bot NTc2NjUwMzQ0ODc3MjYwODAw.XNZl4w.tUX0FBt87pXRt0Sow5M0k_SOcJ4");
    if err != nil {
        fmt.Println("An error has occured while trying to create the discord bot type");
        return;
    }

    user, err := bot.User("@me");  // We get ths information to hold bot idea
    if err != nil {
        fmt.Println("Error getting user information");
        return;
    }
    BotID = user.ID;
    bot.AddHandler(messageHandler);
    err = bot.Open()  // Open a connection it seems
    if err != nil {
        fmt.Println("An error while trying to open a connection");
        return
    }

    Msg, err := bot.ChannelMessageSend("576659713140457476","Hello, The bot has activated itself");
    if err != nil {
        fmt.Println("Error sending Message",err);
    } else {
        fmt.Println(Msg);
    }
    <-make(chan struct{});

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    if (m.Author.ID == BotID) {
        return;
    }
    fmt.Println(m.Content);
    if (checkVulgarList(m.Content) == true) {
        msg := fmt.Sprintf("Why you gotta be vulgar %s",m.Author.Username);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)
        return;
    }
    if (strings.Contains(m.Content,"bot")) {
        msg := fmt.Sprintf("You talking about me %s ?",m.Author.Username);
        _,_ = s.ChannelMessageSend(m.ChannelID,msg)
    }
}


func checkVulgarList(msg string) (bool){
    if (strings.Contains(msg,"fuck")) {
        return true;
    }
    return false;
}

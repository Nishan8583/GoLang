package Discord;

import (
    "fmt";
    "net/http";
    "io/ioutil";
    "bytes";
)

// discord structure holds API and channel information
type discord struct {
    Api string;  // The Api key for communication
    Channel string;
}

// Discord() function sets the Api and channel, will return the pointer to structure
func Discord(api,channel string) (*discord) {
    bot := discord{
        Api : api,
        Channel : channel,
    }
    return &bot;
}

// SetChannel() is used to change the channelID
func (d *discord) SetChannel(channel_id string) {
    d.Channel = channel_id;
}

// SendMessage() takes message string as parameter and sends message
func (d *discord) SendMessage(msg string) (error){
    url := "https://discordapp.com/api/channels/"+d.Channel+"/messages";
    body := "content :"+msg;
    req,err := http.NewRequest("POST",url,bytes.NewReader([]byte(body)));
    auth := "Bot "+d.Api;
    req.Header.Add("Authorization",auth);
    req.Header.Add("User-Agent","Discord Bot");
    req.Header.Add("Content-Type","application/json");

    client_d := &http.Client{};
    resp, err := client_d.Do(req);

    content,err := ioutil.ReadAll(resp.Body);
    if err != nil {
        fmt.Println("Error,",content);
    }
    fmt.Println(string(content));
    return nil;
}

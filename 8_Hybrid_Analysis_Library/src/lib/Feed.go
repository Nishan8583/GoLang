package Hybrid;

import (
    "fmt";
    "encoding/json";
    "io/ioutil";
)

// Feed() returns slice HybridMainType that will contain the feed
//
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Feed/get_feed_latest
func (h *GoHybrid) Feed() (FeedType, error){
    holder := FeedType{}
    h.req.Method = "GET";
    h.req.URL.Path = "/api/v2/feed/latest";

    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Could not get the feed");
        return holder,err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Could not read the feed body");
        return holder,err;
    }
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
        fmt.Println(err);
    }
    return holder,err;
}

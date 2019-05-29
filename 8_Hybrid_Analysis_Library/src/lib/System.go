package Hybrid;

import (
    "fmt";
    "io/ioutil";
    "encoding/json";
)

/*SystemVersion() gets the system element versions
Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/System/get_system_version*/
func (h *GoHybrid) SystemVersion() (SystemVersionType, error){
    holder := SystemVersionType{};
    h.req.Method = "GET";
    h.req.URL.Path = "/api/v2/system/version";

    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Error Getting System Version Info for server");
        return holder,err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Error Reading body");
        return holder, err;
    }
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
        fmt.Println(err);
    }
    return holder, err;
}

/*
SystemEnviroments() returns infomration about available execution enviroments
Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/System/get_system_environments
*/
func (h *GoHybrid) SystemEnviroments() ([]SystemEnviromentsType, error) {
    holder := []SystemEnviromentsType{};
    h.req.Method = "GET";
    h.req.URL.Path = "/api/v2/system/environments";

    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Error getting enviroment information");
        return holder,err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Error Reading Body");
        return holder, err;
    }
    fmt.Println(string(response));
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
        fmt.Println(err);
    }
    return holder,err;
}

/*
SystemStatus() contains a variety of webservice statistics, e.g. the total number of submissions, unique submissions, signature ID distribution, user comments, etc.
Reference Api:
*/
func (h *GoHybrid) SystemStats() (SystemStatsType, error) {
    holder := SystemStatsType{}
    h.req.Method = "GET";
    h.req.URL.Path = "/api/v2/system/stats"

    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Could not get system stats form server");
        return holder, err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Error Reading Body");
        return holder, err;
    }
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
        fmt.Println(err);
    }
    return holder,err;
}

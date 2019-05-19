/*Hybrid Analaysis package, a wrapper around the official hybrid analysis's rest api*/
package main;

import (
    "net/http";
    "fmt";
    "io/ioutil";
    "strings";
    "bytes";
)

var req *http.Request;
var err error;
var c http.Client;
// SetApiParams sets the api key to be used and http Request body, Will return error if found any
func SetParams(api string) (error) {
        // If i just declared error here, then i would have to put := since it was a new var
        // This would create a new req var, but assiging to the global req var
        req, err = http.NewRequest("POST","",nil)
        if err != nil {
            return err;
        }
        req.URL.Scheme = "https"
        req.URL.Host = "www.hybrid-analysis.com"
        req.Header.Add("api-key",api);
        req.Header.Add("user-agent","Falcon Sandbox");
        req.Header.Add("accept","application/json");
        return nil;
}

/* QueryDomain retrieves information about the domain
Ex: resp, err := QueryDomain("http.example.com")
    if err != nil {
    ...
    }
    fmt.Pritnln(resp);
*/
func Query(query_type,formdata string) (string, error){
    switch query_type {
    case "domain":
        req.URL.Path = "/api/v2/search/terms"
        body:= fmt.Sprintf(`{"domain":"%s"}`,formdata)
        //req.Header.Add("content-type","application/x-www-form-urlencoded")
        req.Body = ioutil.NopCloser(strings.NewReader(body));
    case "hash":
        req.URL.Path = "/api/v2/search/hash"
        req.Header.Add("content-type","application/x-www-form-urlencoded")
        body:= []byte(fmt.Sprintf(`{"hash":"%s"}`,formdata))
        req.Body = ioutil.NopCloser(bytes.NewBuffer(body));
    case "overview":
        req.Method = "GET";
        req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s`,formdata)
    }

    fmt.Println(req)
    resp, err := c.Do(req);
    if err != nil {
        fmt.Println("Could not get response")
        return "",err;
    }
    response,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "",err;
    }
    return string(response),nil;
}

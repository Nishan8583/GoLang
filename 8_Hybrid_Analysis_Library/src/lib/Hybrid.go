package Hybrid;

import (
    "net/http";
)




// Hybrid sets the api key to be used and http Request body, Will return error if found any
func HybridInit(api string) (GoHybrid,error) {
        // If i just declared error here, then i would have to put := since it was a new var
        // This would create a new req var, but assiging to the global req var
        hybridType := GoHybrid{};
        req, err := http.NewRequest("POST","",nil)
        if err != nil {
            return hybridType,err;
        }
        hybridType.req = req;
        hybridType.req.URL.Scheme = "https"
        hybridType.req.URL.Host = "www.hybrid-analysis.com"
        hybridType.req.Header.Add("api-key",api);
        hybridType.req.Header.Add("user-agent","Falcon Sandbox");
        hybridType.req.Header.Add("accept","application/json");
        hybridType.req.Header.Add("content-type","application/x-www-form-urlencoded")
        return hybridType,nil;
}






/*
func (h *GoHybrid) Query(query_type,formdata string) (string, error){
    switch query_type {
    case "domain":
        h.req.URL.Path = "/api/v2/search/terms"
        body:= fmt.Sprintf(`{"domain":"%s"}`,formdata)
        //req.Header.Add("content-type","application/x-www-form-urlencoded")
        h.req.Body = ioutil.NopCloser(strings.NewReader(body));
    case "hash":
        h.req.URL.Path = "/api/v2/search/hash"
        h.req.Header.Add("content-type","application/x-www-form-urlencoded")
        body:= []byte(fmt.Sprintf(`{"hash":"%s"}`,formdata))
        h.req.Body = ioutil.NopCloser(bytes.NewBuffer(body));
    case "overview":
        h.req.Method = "GET";
        h.req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s`,formdata)
    }

    fmt.Println(h.req)
    resp, err := h.client.Do(h.req);
    if err != nil {
        fmt.Println("Could not get response")
        return "",err;
    }
    response,err := ioutil.ReadAll(resp.Body);
    if err != nil {
        return "",err;
    }
    return string(response),nil;
}
*/

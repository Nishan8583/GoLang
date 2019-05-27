/*Hybrid Analaysis package, a wrapper around the official hybrid analysis's rest api*/
package Hybrid;

import (
    "net/http";
    "fmt";
    "io/ioutil";
    "strings";
    "net/url"
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



// Report API starts here
/*
ReportState(hash) gets the rport
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
err = h.ReportState(hash_in_string);
if err != nil {
  fmt.Println(err);
}
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__state
*/
func (h *GoHybrid) ReportState(hash string) (string, error) {
    sh, err := h.SearchHash(hash);
    if (err != nil) {
        fmt.Println("Error Calling SearchHash");
        return "",err;
    }
    fmt.Println(sh);
    h.req.Method = "GET";

    // Since the return is a slice of search hash result loop through it
    for _,value := range(sh) {
        h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/state",value.JobId);  // Using the jobid to get result
        fmt.Println(h);
        resp, err := h.client.Do(h.req);
        if (err != nil) {
          fmt.Println("Could not get Report");
          return "",nil;
        }
        response, err := ioutil.ReadAll(resp.Body);
        if (err != nil) {
          fmt.Println("Could not read body");
          return "",err;
        }
        fmt.Println(string(response))
    }





  return "",nil;

}

/*
ReportSummaryID(hash) gets the rport
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
err = h.ReportSummaryID(hash_in_string);
if err != nil {
  fmt.Println(err);
}
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__summary
*/
func (h *GoHybrid) ReportSummaryID(hash string) (string, error) {
    sh, _ := h.SearchHash(hash); // Get the JobId
    response := []byte{}
    h.req.Method = "GET";
    for _,value := range sh {
        h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/summary",value.JobId);
        fmt.Println(h);

        resp, err := h.client.Do(h.req);
        if (err != nil) {
          fmt.Println("Could not get Report");
          return "",nil;
        }
        response, err := ioutil.ReadAll(resp.Body);
        if (err != nil) {
          fmt.Println("Could not read body");
          return "",err;
        }
        fmt.Println(string(response))
    }






  return string(response),nil;

}

/*
ReportSummary([]string{hash1,hash2...) gets the report summary
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
resp, err = h.ReportSummary([]string{"603a72e1aad833b92a6ef7edac65849c3d899b4b7eaac399abf2f6d2cbb4b1e7","c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98"});
if err != nil {
  fmt.Println(err);
}
fmt.Println(string(resp))
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/post_report_summary
*/


func (h *GoHybrid) ReportSummary(hash []string) ([]byte, error) {

  h.req.Method = "POST";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/summary");
  h.req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  // Adding hashses in form data
  form := url.Values{}
  for _,value := range(hash) {
    form.Add("hashes[]",value);
  }

  h.req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()));

  resp, err := h.client.Do(h.req);
  if (err != nil) {
    return []byte("Could not get Report"),nil;
  }
  response, err := ioutil.ReadAll(resp.Body);
  if (err != nil) {
    return []byte("Could not read body"),err;
  }
  return response,nil;

}

/*
ReportTypeSample(hash,filetype) gets the report summary
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
resp, err = h.ReportSummary("603a72e1aad833b92a6ef7edac65849c3d899b4b7eaac399abf2f6d2cbb4b1e7","xml");
if err != nil {
  fmt.Println(err);
}
fmt.Println(string(resp))
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__file__type_
*/
func (h *GoHybrid) ReportTypeSample(hash,ftype string) ([]byte, error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/file/%s",hash,ftype);
  h.req.Header["accept"] = []string{`application/xml`};
  defer func(){
    h.req.Header["accept"] = []string{`application/json`}
    }();

    resp, err := h.client.Do(h.req);
    if err != nil {
      return []byte(`Could not get report`),err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
      return []byte(`Could not read the body of response`),err;
    }
    return response,nil;
}

/*
ReportScreenshots(hash string) ([]byte, error) gets screenshots
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
resp, err = h.ReportScreenshots("603a72e1aad833b92a6ef7edac65849c3d899b4b7eaac399abf2f6d2cbb4b1e7");
if err != nil {
  fmt.Println(err);
}
fmt.Println(string(resp))
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__screenshots
*/
func (h *GoHybrid) ReportScreenshots(hash string) ([]byte, error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/screenshots",hash);

    resp, err := h.client.Do(h.req);
    if err != nil {
      return []byte(`Could not get report`),err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
      return []byte(`Could not read the body of response`),err;
    }
    return response,nil;
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

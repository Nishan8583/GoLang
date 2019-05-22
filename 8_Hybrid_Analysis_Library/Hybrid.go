/*Hybrid Analaysis package, a wrapper around the official hybrid analysis's rest api*/
package main;

import (
    "net/http";
    "fmt";
    "io/ioutil";
    "strings";
    //"bytes";
    //"strconv";
    "time";
)

// The base type the package will be using
type GoHybrid struct {
  req *http.Request;
  err error;
  client http.Client;
}

// SetApiParams sets the api key to be used and http Request body, Will return error if found any
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
        return hybridType,nil;
}

// Overview part starts here
/*
Overview(hash_in _string) gets the overview of some hash info given
Original API Reference: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256_
*/
func (h *GoHybrid) Overview(formdata string) (string, error){
        h.req.Method = "GET";
        h.req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s`,formdata)

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

func (h * GoHybrid) OverviewSummary(hash string) (string, error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/overview/%s/summary",hash);
  fmt.Println(h);
  resp, err := h.client.Do(h.req); // Do the rest
  if err != nil {
    fmt.Println("Could not get overview summary");
    return "",err;
  }
  response, err := ioutil.ReadAll(resp.Body);
  if (err != nil) {
    fmt.Println("Could not read response body");
    return "",err;
  }
  return string(response),nil;

}
/*
OverviewReferesh(hash) refersehes overviews and downloads fresh data
Original API Reference: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256__refresh
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
resp,err = h.OverviewReferesh(hash_in_string);
if err != nil {
  fmt.Println(err);
}
fmt.Println(resp)
*/
func (h *GoHybrid) OverviewReferesh(formdata string) (string, error){
        h.req.Method = "GET";
        h.req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s/referesh`,formdata)

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

/*
DownloadSample(hash) downloads the file belonging to the hash, The file has to donwloadable though,
the filename will be sample-year-month-day-minute-nanosecond.gzip
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
err = h.DownloadSample(hash_in_string);
if err != nil {
  fmt.Println(err);
}
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256__sample
*/
func (h *GoHybrid) DownloadSample(hash string) (error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s/sample`,hash);
  h.req.Header["accept"] = []string{"application/gzip"};

  // Changing back to default so that other apis don't have problems
  defer func() {
    h.req.Header["accept"] = []string{"application/json"};
  }();

  fmt.Println("Attempting to Download file");
  resp,err := h.client.Do(h.req);
  if (err != nil) {
    fmt.Println("Could not Download file")
    return err;
  }

  response, err := ioutil.ReadAll(resp.Body);
  if (err != nil) {
    return err;
  }

  if (resp.StatusCode == 200) {
    fmt.Println("Downloading file successfull")
    t := time.Now()
    filename:= fmt.Sprintf("sample-%d-%d-%d-%d-%d.zip",t.Year(),t.Month(),t.Day(),t.Minute(),t.Nanosecond());
    fmt.Println("Filename:",filename);
    err := ioutil.WriteFile(filename,response,055);
    if (err != nil) {
      fmt.Println("Error Creating File");
      return err;
    }
    return nil;
  }
  return nil;
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
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/state",hash);

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
  return string(response),nil;

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
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/summary",hash);

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
  return string(response),nil;

}

/*
ReportSummary(hash) gets the rport
Ex:
h, err := HybridInit("<API-KEY>"); // The api key will be used
if err != nil {
  fmt.Println("Could not Create Hybrid Type",err);
  return;
}
fmt.Println(h);
err = h.ReportSummary(hash_in_string);
if err != nil {
  fmt.Println(err);
}
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__summary
*/
func (h *GoHybrid) ReportSummary(hash []string) (string, error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/summary",hash);
  body := ``
  for _,value := range(hash) {
    body = body + fmt.Sprintf(`hashes[]:%s&,`,value);
  }
  body = strings.TrimRight(body,"&");
  fmt.Println(body)
  h.req.Body = ioutil.NopCloser(strings.NewReader(body));
  fmt.Println(h.req);
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
  return string(response),nil;

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

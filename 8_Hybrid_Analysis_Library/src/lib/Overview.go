package Hybrid;

import (
    "fmt";
    "io/ioutil";
    "encoding/json"
    "time"
)

/*
Overview(hash_in _string) gets the general result of analysis/overview of some hash given.
It returns OverviewType, and error if any

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256_
*/
func (h *GoHybrid) Overview(hash string) (OverviewType, error){
  holder := OverviewType{};
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf(`/api/v2/overview/%s`,hash)

  resp, err := h.client.Do(h.req);
  if err != nil {
      fmt.Println("Could not get response")
      return holder,err;
  }
    response,err := ioutil.ReadAll(resp.Body);
    if err != nil {
        return holder,err;
    }
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
      fmt.Println(err);
    }
    return holder,nil;
}

/*OverviewSummary(hash) gets the summary result of scanned hash and
returns OverviewSummaryType and error if present

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256__summary*/
func (h * GoHybrid) OverviewSummary(hash string) (OverviewSummaryType, error) {
  holder := OverviewSummaryType{}
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/overview/%s/summary",hash);
  resp, err := h.client.Do(h.req); // Do the rest
  if err != nil {
    fmt.Println("Could not get overview summary");
    return holder,err;
  }
  response, err := ioutil.ReadAll(resp.Body);
  if (err != nil) {
    fmt.Println("Could not read response body");
    return holder,err;
  }
  err = json.Unmarshal(response,&holder);
  if (err != nil) {
      fmt.Println("Error parsing json ",err);
  }
  return holder,nil;

}

/*
DownloadSample(hash) downloads the file belonging to the hash, The file has to donwloadable though,
the filename will be sample-year-month-day-minute-nanosecond.gzip. It returns nil, or error if any


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
    err := ioutil.WriteFile(filename,response,055);
    if (err != nil) {
      fmt.Println("Error Creating File");
      return err;
    }
    fmt.Println("Sample Download under the Filename:",filename);
    return nil;
  }
  return nil;
}

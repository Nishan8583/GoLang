package Hybrid;
import (
    "fmt";
    "io/ioutil";
    "strings";
    "net/url"
    "encoding/json";
)
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
resp, err := h.SearchHash("c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98");
if err != nil {
    fmt.Println(err);
}
for _,value := range resp {
    res, err := ReportState(value.JobId);
    if err != nil {
      fmt.Println(err);
    }
    fmt.Println(res);
}
Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__state
*/
func (h *GoHybrid) ReportState(JobId string) (ReportStateType, error) {
    holder := ReportStateType{}
    h.req.Method = "GET";
    h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/state",JobId);  // Using the jobid to get result
    resp, err := h.client.Do(h.req);
    if (err != nil) {
      fmt.Println("Could not get Report");
      return holder,nil;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
      fmt.Println("Could not read body");
      return holder,err;
    }
    fmt.Println(string(response))
    err = json.Unmarshal(response,&holder);
    if (err != nil) {
        fmt.Println("Error parsing Json ",err);
    }
  return holder,nil;
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
func (h *GoHybrid) ReportSummaryID(JobId string) (ReportSummaryIDType, error) {
    holder := ReportSummaryIDType{}
    h.req.Method = "GET";
    h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/summary",JobId);  // Using the jobid to get result
    resp, err := h.client.Do(h.req);
    if (err != nil) {
      fmt.Println("Could not get Report");
      return holder,nil;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
      fmt.Println("Could not read body");
      return holder,err;
    }
    _ = json.Unmarshal(response,&holder);  // // An error may occur since if no value null is returned
  return holder,nil;
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

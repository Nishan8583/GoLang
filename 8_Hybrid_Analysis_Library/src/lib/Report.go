package Hybrid;
import (
    "fmt";
    "io/ioutil";
    "encoding/json";
    "encoding/base64"
    "time";
)

/*
ReportState(hash) gets the rport state

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
ReportSummaryID(hash) gets the report

Note: for /report/summary use this function in a loop

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/post_report_summary
*/
func (h *GoHybrid) ReportSummaryID(JobId string) (HybridMainType, error) {
    holder := HybridMainType{}
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
ReportTypeSample(hash,filetype) gets the report summary


Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__file__type_
*/
func (h *GoHybrid) ReportTypeSample(JobId,ftype string) ([]byte, error) {
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/file/%s",JobId,ftype);
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
ReportScreenshots(jobID string) ([]byte, error) gets screenshots of the jobID specified

Filename: sample-...

Reference API: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__screenshots
*/
func (h *GoHybrid) ReportScreenshots(JobID string) (error) {
    holder := []ReportScreenshotsType{}
  h.req.Method = "GET";
  h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/screenshots",JobID);

    resp, err := h.client.Do(h.req);
    if err != nil {
        fmt.Println("Could not Perform Request")
      return err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("could not read the response body")
      return err;
    }
    _= json.Unmarshal(response,&holder);
    i := 0;
    for _,value := range holder {
        data, err := base64.StdEncoding.DecodeString(value.Image);
        if (err != nil) {
            fmt.Println("Could not decode the image value");
            return err;
        }
        t := time.Now();
        filename:= fmt.Sprintf("sample-%s-%d-%d-%d-%d-%d-%d.png",JobID,i,t.Year(),t.Month(),t.Day(),t.Minute(),t.Nanosecond());
        err = ioutil.WriteFile(filename,data,055);
        if (err != nil) {
            fmt.Println("Sorry But could not create file ",filename);
            return err;
        }
        fmt.Println("Successfully Created file ",filename);
        i = i+1;
    }
    return nil;
}

// ReportDownloadExtractedFiles(JobId) will download all the extracted files by a malware
// The file will be download in the directory the path the program was running in
//
// filename: {jobid}-Fxtracted_Files-....
//
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__dropped_files
func (h *GoHybrid) ReportDownloadExtractedFiles(JobID string) (error) {
    h.req.Method = "GET";
    h.req.URL.Path = fmt.Sprintf("/api/v2/report/%s/dropped-files",JobID);

    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Could not query the server");
        return err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Error Reading Response Body");
        return err;
    }
    t := time.Now();
    filename:= fmt.Sprintf("%s-Extracted-Files-%d-%d-%d-%d-%d.zip",JobID,t.Year(),t.Month(),t.Day(),t.Minute(),t.Nanosecond());
    err = ioutil.WriteFile(filename,response,055);
    if (err != nil) {
        fmt.Println("Sorry But could not create file ",filename);
        return err;
    }
    fmt.Println("Successfully Created file ",filename);
    return nil;
}

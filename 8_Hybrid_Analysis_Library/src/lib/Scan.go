package Hybrid;

import (
  "mime/multipart"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "bytes"
  "os"
  "net/url"
  "strings"
)

/*ScanState() gets list of all the available scanner and their state
The return type is []ScannerType

Referece Api: https://www.hybrid-analysis.com/docs/api/v2#/Quick_Scan/get_quick_scan_state
*/
func (h *GoHybrid) ScanState() ([]ScannerType, error) {
  holder := []ScannerType{};
  h.req.Method = "GET";
  h.req.URL.Path = "/api/v2/quick-scan/state"

  resp, err := h.client.Do(h.req);
  if (err != nil) {
    fmt.Println("Could not get scanners state");
    return holder,err
  }
  response, err := ioutil.ReadAll(resp.Body);
  defer resp.Body.Close();
  if (err != nil) {
    fmt.Println("Could not read response body")
    return holder, err
  }

  err = json.Unmarshal(response,&holder);
  return holder,err;
}

/*
ScanFile(filename, keys mapp) scans the file and returns the structure ScannerResultType

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Quick_Scan/post_quick_scan_file
*/
func (h *GoHybrid) ScanFile(filename string, keys map[string]string) (ScannerResultType, error) {

  holder := ScannerResultType{}
  file, err := os.Open(filename);
  if (err != nil) {
    return holder,err;
  }
  defer file.Close();
  content, err := ioutil.ReadAll(file);
  if (err != nil) {
    fmt.Println("Content of file could not be read");
    return holder, err;
  }

  h.req.Method = "POST";
  h.req.URL.Path = "/api/v2/quick-scan/file"

  body := new(bytes.Buffer);
  writer := multipart.NewWriter(body);
  defer writer.Close();
  part, err := writer.CreateFormFile("file",file.Name());
  if (err != nil) {
    fmt.Println("Could not create multipart form for file");
    return holder,err;
  }
  part.Write(content);
  h.req.Header["Content-Type"] = []string{writer.FormDataContentType()};
  for key,value := range keys {
    writer.WriteField(key,value)
  }

  h.req.Body = ioutil.NopCloser(body);

  resp, err := h.client.Do(h.req);
  if (err != nil) {
    fmt.Println("Could not get respone form server");
    return holder, err
  }
  response, err := ioutil.ReadAll(resp.Body)
  if (err != nil) {
    fmt.Println("Could not read body")
    return holder, err;
  }
  err = json.Unmarshal(response,&holder);
  return holder, err
}

/*
ScanURLFile(url, keys...) the url is the link to file to scan

Refernce Api: https://www.hybrid-analysis.com/docs/api/v2#/Quick_Scan/post_quick_scan_url_to_file
*/
func (h *GoHybrid) ScanURLFile(url_query string, keys map[string]string) (ScannerResultType, error) {
  holder := ScannerResultType{};

  h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/quick-scan/url-to-file"
	h.req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()

	formValue := url.Values{};
	formValue.Add("url",url_query);
	for key, value := range keys {
		formValue.Add(key,value);
	}

	h.req.Body = ioutil.NopCloser(strings.NewReader(formValue.Encode()))



	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return holder, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response,&holder);
	if (err != nil) {
		return holder,err
	}
	return holder,err
}

/*
ScanURL(url,keys..) pass the url to scan, not file url

Refernce Api: https://www.hybrid-analysis.com/docs/api/v2#/Quick_Scan/post_quick_scan_url_for_analysis
*/
func (h *GoHybrid) ScanURL(url_query string, keys map[string]string) (ScannerResultType, error) {
  holder := ScannerResultType{};

  h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/quick-scan/url-for-analysis"
	h.req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()

	formValue := url.Values{};
	formValue.Add("url",url_query);
	for key, value := range keys {
		formValue.Add(key,value);
	}

	h.req.Body = ioutil.NopCloser(strings.NewReader(formValue.Encode()))



	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return holder, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response,&holder);
	if (err != nil) {
		return holder,err
	}
	return holder,err
}

/*
ScanResultId(id, keys), gets the report of the scan result, pass in the JobId
It returns ScannerResultType

Refernce Api: https://www.hybrid-analysis.com/docs/api/v2#/Quick_Scan/get_quick_scan__id_
*/
func (h *GoHybrid) ScanResultId(id string) (ScannerResultType, error) {
  holder := ScannerResultType{};

  h.req.Method = "GET"
	h.req.URL.Path = fmt.Sprintf("/api/v2/quick-scan/%s",id);
	h.req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()




	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return holder, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response,&holder);
	if (err != nil) {
		return holder,err
	}
	return holder,err
}

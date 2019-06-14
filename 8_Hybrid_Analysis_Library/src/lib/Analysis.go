package Hybrid

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"encoding/json"
	"net/url"
	"strings"
)

/* AnalyzeFile(filename,map[string]stirng{"environment_id","300"})
AnalyzeFile() submits the file specified, with the parameters to hybrid analysys for analysis
submits file for analysis to hybrid analysis, and returns AnalyzeFileType strucutre containing enviromentid,jobid and hash

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Submission/post_submit_file
Note: map[string]string here at least environment is cumposlary
*/
func (h *GoHybrid) AnalyzeFile(filename string, keys map[string]string) (AnalyzeFileType,error) {
	holder := AnalyzeFileType{}
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/file"

	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()

  // Getting the pointer to file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not read file", err)
		return holder, err
	}
  defer file.Close();


	fileContents, err := ioutil.ReadAll(file)
	if (err  != nil) {
		fmt.Println("Could not read the content of file")
		return holder, err;
	}

  body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)  // Now everything the writer writes as multipart strcuture it will be in the buffer
	part, err := writer.CreateFormFile("file", file.Name());  // Create a multipart part with key file
  part.Write(fileContents)  // Write the content of file to the key
	h.req.Header["Content-Type"] = []string{writer.FormDataContentType()}  // Get the contenttype
	if err != nil {
		fmt.Println("Could not create part")
		return holder, err
	}

	err = writer.Close()
	if (err != nil) {
		fmt.Println("Error could not close the writer object")
		return holder, err;
	}

	for key1, value := range keys {
		writer.WriteField(key1,value)
	}
	h.req.Body = ioutil.NopCloser(body)
	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return holder, err
	}
	defer resp.Body.Close();
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response,&holder);
	if (err != nil) {
		return holder,err
	}
	return holder,err
}

/* AnalyzeURLFile(url,map[string]stirng{"environment_id","300"})
AnalyzeURLFile submits file from the specified url for analysis to hybrid analysis,
and returns AnalyzeFile strucutre containing enviromentid,jobid and hash

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Submission/post_submit_url_to_file
Note: map[string]string here at least environment is cumposlary
Also while providing the url, please be sure that the url is a download link
*/
func (h *GoHybrid) AnalyzeURLFile(url_query string, keys map[string]string) (AnalyzeFileType,error) {
	holder := AnalyzeFileType{}
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/url-to-file"
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

/* AnalyzeURL(url,map[string]stirng{"environment_id","300"})
AnalyzeURL() submits url for analysis to hybrid analysis, and returns AnalyzeFileType
strucutre containing enviromentid,jobid and hash

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Submission/post_submit_url_for_analysis
Note: map[string]string here at least environment is cumposlary
Also please be sure that the url you provide is a url link and not a donwload link
*/
func (h *GoHybrid) AnalyzeURL(url_query string, keys map[string]string) (AnalyzeFileType,error) {
	holder := AnalyzeFileType{}
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/url-for-analysis"
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

/* AnalyzeURLHash(url)
AnalyzeURLHash(url)submits file for analysis to hybrid analysis, and returns a strucutre containing enviromentid,jobid and hash

Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Submission/post_submit_url_for_analysis
Note: map[string]string here at least environment is cumposlary
Also please be sure that the url you provide is a url link a donwload link
*/
func (h *GoHybrid) AnalyzeURLHash(url_query string) (AnalyzeFileType,error) {
	holder := AnalyzeFileType{}
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/hash-for-url"
	h.req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()

	formValue := url.Values{};
	formValue.Add("url",url_query);

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

/*AnalyzeDroppedFiles() submits dropped file for analysis with the sha256 hash specified rleated to the given jobid
analysis to hybrid analysis,and returns a strucutre containing enviromentid,jobid and hash
Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Submission/post_submit_dropped_file

Its basically submitting dropped files by the malware for further analysis
*/
func (h *GoHybrid) AnalyzeDroppedFiles(jobid,sha256 string) (string,error) {
	//holder := AnalyzeFileType{}
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/dropped-file"
	h.req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	// Defering to ensure change of content-type
	defer func() {
		h.req.Header["Content-Type"] = []string{"application/json"}
	}()

	formValue := url.Values{};
	formValue.Add("id",jobid);
	formValue.Add("file_hash",sha256)

	h.req.Body = ioutil.NopCloser(strings.NewReader(formValue.Encode()))

	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return "", err
	}
	response, err := ioutil.ReadAll(resp.Body)

	if (err != nil) {
		return string(response),err
	}
	return string(response),err
}

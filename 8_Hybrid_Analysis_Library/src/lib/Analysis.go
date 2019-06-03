package Hybrid

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	//"io"
	"os"
)

func (h *GoHybrid) AnalyzeFile(filename string) error {
	h.req.Method = "POST"
	h.req.URL.Path = "/api/v2/submit/file"

	defer func() {
		h.req.Header["content-type"] = []string{"application/json"}
	}()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not read file", err)
		return err
	}
  defer file.Close();

	fileContents, err := ioutil.ReadAll(file)

  body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("File", file.Name())
  part.Write(fileContents)

  h.req.Header["content-type"] = []string{"multipart/form-data"}

	if err != nil {
		fmt.Println("Could not create part")
		return err
	}

	//writer.WriteField("file",string(content));
	writer.WriteField("environment_id", "300")
	err = writer.Close()

	h.req.Body = ioutil.NopCloser(body)
	resp, err := h.client.Do(h.req)
	if err != nil {
		fmt.Println("Error sending file")
		return err
	}
	response, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))
	return err
}

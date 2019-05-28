// Implements /search/... parts of the Api, /search/hashes is not implemented, you can run /search/hash in loop
package Hybrid

import (
    "fmt";
    "io/ioutil";
    "strings";
    "net/url"
    "encoding/json"
)

/* SearchHash(hash) returns sliceof the SearchHashType, in case of error error value is thrown\
Referece Api: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
*/
func (h *GoHybrid) SearchHash(hash string) ([]SearchHashType, error) {
    // Reference: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
    holder := []SearchHashType{};  // Destination of json Unmarshalling

    form := url.Values{};
    form.Add("hash",hash);
    h.req.URL.Path = "/api/v2/search/hash";
    h.req.Method = "POST";
    h.req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()));
    resp, err := h.client.Do(h.req);
    if (err != nil) {
        fmt.Println("Error while requesting",err);
        return holder,err;
    }
    response, err := ioutil.ReadAll(resp.Body);
    if (err != nil) {
        fmt.Println("Error while Reading");
        return holder,err
    }
    err = json.Unmarshal(response,&holder);
    if err != nil {
        fmt.Println(err);
    }
    defer resp.Body.Close()
    return holder,nil;

}

/*SearchTerms(type,term) gets information about the related field
Ex:
resp,err := h.SearchQuery("domain","google.com");
if err != nil {
  fmt.Println(err);
}
fmt.Println(resp);
Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_terms*/
func (h *GoHybrid) SearchQuery(types,term string) (SearchTermsType, error) {
  holder := SearchTermsType{};
  h.req.Method = "POST";
  h.req.URL.Path = "/api/v2/search/terms";
  form := url.Values{};
  form.Add(types,term);
  h.req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()));

  resp, err := h.client.Do(h.req); // do the request
  if (err != nil) {
    fmt.Println("Error Sending request");
    return holder,err;
  }
  response, err := ioutil.ReadAll(resp.Body);
  if (err != nil) {
    fmt.Println("Error Could not read response body");
    return holder,err;
  }
  err = json.Unmarshal(response,&holder);
  if (err != nil) {
    fmt.Println(err);
  }
  return holder,nil;
}

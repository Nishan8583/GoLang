/*This is a package which uses the official hybrid Analysis API Details for the original Api can be found here https://www.hybrid-analysis.com/docs/api/v2#/
:

	h, err := Hybrid.HybridInit("<API-KEY Here>") // The api key will be used
	if err != nil {
		fmt.Println("Could not Create Hybrid Type", err)
		return
	}

	//resp,err := h.ReportSummary([]string{"603a72e1aad833b92a6ef7edac65849c3d899b4b7eaac399abf2f6d2cbb4b1e7","c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98"});
	//resp,err := h.OverviewSummary("c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98");
	//resp,err := h.SearchHash("c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98");
	//resp,err := h.SearchQuery("domain","google.com");
	//resp, err := h.AnalyzeFile("path_to_file.exe",map[string]string{"environment_id":"300");
	//resp, err := h.AnalyzeURLFile("https://www.blackhat.com/presentations/bh-usa-04/bh-us-04-chambet/bh-us-04-chambet-google-up.pdf",map[string]string{"environment_id":"300"});
	//resp, err := h.AnalyzeURL("https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7",map[string]string{"environment_id":"300"});
	//resp, err := h.AnalyzeURLHash("https://www.blackhat.com/presentations/bh-usa-04/bh-us-04-chambet/bh-us-04-chambet-google-up.pdf");
	//resp, err := h.AnalyzeDroppedFiles("e5b9ce395b80a7b55af07915923cd282589c6c4c9f079efc25827dcf11e6b9ec","7cff31263aaf801db3d229b44f77658736188f18578dce4594d71c514c8e412f");
	//resp, err := h.ScanState();
	resp, err := h.ScanResultId("5cf8c662028838c232ebc6bb")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
*/
package Hybrid;

import (
    "net/http";
)

/*HybridInit() sets the api key and other essentials in the main GoHybrid type
to be used later for calling functions. It Will return GoHybrid type and nil
if error, will return empty type and error
*/
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

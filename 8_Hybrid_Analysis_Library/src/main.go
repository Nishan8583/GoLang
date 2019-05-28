package main;

import (
	"fmt";
	"./lib";
)

func main() {
	h, err := Hybrid.HybridInit("<API-KEY>") // The api key will be used
	if err != nil {
		fmt.Println("Could not Create Hybrid Type", err)
		return
	}

  // Using SOme Sample Hahses
	//resp, err := h.ReportSummary([]string{"603a72e1aad833b92a6ef7edac65849c3d899b4b7eaac399abf2f6d2cbb4b1e7", "c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98"})
	//resp,err := h.OverviewSummary("c7acf3c1167ae28439a22bec62e35303fd34043c600a6ad333cfe115a2b12e98");
	resp, err := h.ReportSummaryID("5ce509b503883867419823dd");
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(resp.Processes);

}

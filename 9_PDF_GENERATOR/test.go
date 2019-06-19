package main

import (
        "fmt"
        "github.com/jung-kurt/gofpdf"
  "encoding/json"

  "Elasticsearch"

)

var pdf *gofpdf.Fpdf;

func main() {
        pdf = gofpdf.New("P", "mm", "A4", "")
        // Tilte Page
        pdf.AddPage() // Adding a page
        pdf.SetTitle("Report", true)

        pdf.SetFont("Arial", "B", 60)
        pdf.SetFillColor(255, 0, 0)

        pdf.WriteAligned(0, 35, "Report Sample\n", "c")
        pdf.Write(35, "\n")
        //pdf.WriteAligned(0, 35, "\n", "c")
        //pdf.Write(35,"REPORT")
        //x1 := pdf.GetLineWidth();
        _, y := pdf.GetXY()
        //fmt.Println(x);
        pdf.Image("./gopher.png", 55, y, 0, 35, false, "", 0, "")
        pdf.SetFont("Times", "", 20)
        pdf.Write(35, "\nName:Nishan MAharjan\n")
        pdf.Write(1, "\nID:Don\n")

        // Second Page
        pdf.AddPage()
        pdf.SetFont("Arial", "I", 50)
        pdf.WriteAligned(0, 35, "NetFlowSummary\n", "c")

        es := Elasticsearch.ES{
                User:     "admin",
                Password: "admin",
        }
        err := es.es_get.Build_es_type()
        if err != nil {
                fmt.Println("Threat Intel ERROR: Could Build elasticsearch type: ", err)
                return
        }

        url := "https://127.0.0.1:9200/netflow-2019.05.02/_search"

        // Building Query for Request Body
        queries := `
          {
              "size":1000
          }
          `

        // Actually Performing the query
        p, err := es_get.Query("GET", url, queries)
        if err != nil {
                fmt.Println("Threat Intel ERROR: Could not get response from Elasticsearch server ", err, "Trying to connect again")
                return
        }
  var response_struct Elasticsearch.Ctms_response;
  err = json.Unmarshal(p,&response_struct)
  if (err != nil) {
    fmt.Println(err)
  }
        // Writing to file
        err = pdf.OutputFileAndClose("test.pdf")
        if err != nil {
                fmt.Println("Could not write to file", err)
        }
}

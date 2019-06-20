package main

import (
        "encoding/json"
        "fmt"
        "github.com/jung-kurt/gofpdf"

        "./Elasticsearch"
)

type Bucket struct {
        Key   string `json:"key"`
        Count int    `json:"doc_count"`
}

type UniqType struct {
        DcountError int `json:"doc_count_error_upper_bound"`
        Sum int `json:"sum_other_doc_count"`
        Buck []Bucket `json:"buckets"`
}
type AggregationsType struct {
        Uniq UniqType `json:"uniq_gender"`
}
type Summary struct {
        Took int `json:"took"`
        TimedOut bool `json:"timed_out"`
        //Value map[string]map[string]map[string][]Bucket `json:"aggregations"`
        Shard map[string]int `json:"_shards"`
        Hits map[string]int `json:"hits"`
        Aggregations AggregationsType `json:"aggregations"`

}
/*
        "aggregations":
                {"uniq_gender":
                                {"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":
                                                [{"key":"10.10.80.69","doc_count":131746},{"key":"8.8.8.8","doc_count":57073},{"key":"10.10.4.99","doc_count":702},{"key":"fe80:0000:0000:0000:020c:29ff:fedd:caa9","doc_count":24},{"key":"10.10.4.17","doc_count":3}]}}}
'*/
var pdf *gofpdf.Fpdf    // gofpdf pointer type that we will use to create pdf
var es Elasticsearch.ES // Elasticsearch type for communication

/*init() initalizes the basic of pdf creation and elasticsearch query*/
func init() {
        pdf = gofpdf.New("P", "mm", "A4", "")

        // Tilte Page
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 60)
        pdf.WriteAligned(0, 35, "Report Sample\n", "c")
        pdf.Write(35, "\n")

        y := pdf.GetY() // Get the current Y position
        pdf.Image("./gopher.png", 55, y, 0, 35, false, "", 0, "")
        pdf.SetFont("Times", "", 20)
        pdf.Write(35, "\nName:Nishan MAharjan\n")
        pdf.Write(1, "\nID:Don\n")

        // Second Page
        pdf.AddPage()

        // Elasticsearch part starts here
        es = Elasticsearch.ES{ // This could be different
                User:     "admin",
                Password: "admin",
        }
        err := es.Build_es_type()
        if err != nil {
                fmt.Println("Threat Intel ERROR: Could Build elasticsearch type: ", err)
                return
        }



}

/* WriteSummary will write the basics of the report, Ex total counts
index_name is the index to query, fields is the slice of field name you want summary from*/
func WriteSummary(index_name string, fields []string) {
        fmt.Println("Looping through index")
        for index := range fields {
                url := fmt.Sprintf(`https://127.0.0.1:9200/%s/_search?`, index_name)
                queries := fmt.Sprintf(`
                        {
                        "size":"0",
                        "aggs" : {
                        "uniq_gender" : {
                        "terms" : { "field" : "%s.keyword" }
                        }
                        }
                        }`, fields[index])
                fmt.Println(queries)
                p, err := es.Query("GET", url, queries)
                if err != nil {
                        fmt.Println("Report Generation error ERROR: Could not get response from Elasticsearch server ", err, "Trying to connect again")
                        return
                }

                fmt.Println("Trying to print")
                fmt.Println(string(p))
                response_struct := Summary{};
                fmt.Println("tyring to Unmarshal")
                err = json.Unmarshal(p, &response_struct)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println("HEre")
                // Writing Total TotalHits
                buf := fmt.Sprintf("Total Hits: \n")
                pdf.Write(35, buf)
                for _,value := range response_struct.Aggregations.Uniq.Buck {
                        msg := fmt.Sprintf(`Source_IP: %s               Count: %d\n`,value.Key,value.Count)
                        pdf.Write(35, msg)
                        fmt.Println(value);
                }

        }
}

// main() function will call other necessary functions
func main() {
        // Writing to file
        fmt.Println("Calling writesummary")
        WriteSummary("netflow-2019.05.03", []string{"source_ip"})
        err := pdf.OutputFileAndClose("test.pdf")
        if err != nil {
                fmt.Println("Could not write to file", err)
        }
}

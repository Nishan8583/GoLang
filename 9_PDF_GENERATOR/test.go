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
        DcountError int      `json:"doc_count_error_upper_bound"`
        Sum         int      `json:"sum_other_doc_count"`
        Buck        []Bucket `json:"buckets"`
}
type AggregationsType struct {
        Uniq UniqType `json:"uniq_gender"`
}
type Summary struct {
        Took     int  `json:"took"`
        TimedOut bool `json:"timed_out"`
        Shard        map[string]int   `json:"_shards"`
        Hits         map[string]int   `json:"hits"`
        Aggregations AggregationsType `json:"aggregations"`
}

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

        // Setting font as needed
        pdf.SetFont("Helvetica","",10)

        // a slice of Summary{} that will hold Summary{} structure for each field
        response_struct := []Summary{}

        // Looping through each fields requestd
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

                p, err := es.Query("GET", url, queries)
                if err != nil {
                        fmt.Println("Report Generation error ERROR: Could not get response from Elasticsearch server ", err, "Trying to connect again")
                        return
                }

                temp := Summary{}

                err = json.Unmarshal(p, &temp)
                if (err != nil) {
                        fmt.Println("Error unmarshalling json",err);
                }

                response_struct = append(response_struct,temp);
        }
        for i :=0; i < len(response_struct); i++ {
                pdf.Write(10,fmt.Sprintf(`%s                    Count\n`,fields[i]))
                //DrawLine();
                for _, v := range(response_struct[i].Aggregations.Uniq.Buck){
                        pdf.Write(10,fmt.Sprintf(`%s                    %d\n`,v.Key,v.Count))
                }
        }
}

func DrawLine() {
        x,y := pdf.GetXY();
        pdf.SetDrawColor(1,0,25)
        pdf.SetLineWidth(1)
        pdf.Line(x,y,x,y+150);
}
// main() function will call other necessary functions
func main() {
        // Writing to file
        pdf.Cell(0,20,"xtet \n")
        pdf.Cell(0,20,"xtet \n")
        fmt.Println("Calling writesummary")
        pdf.SetFont("Arial","B",15)
        pdf.Write(35,"shit coding")

        WriteSummary("netflow-2019.05.03", []string{"source_ip","destination_ip","source_port","destination_port"})
        err := pdf.OutputFileAndClose("test.pdf")
        if err != nil {
                fmt.Println("Could not write to file", err)
        }
}
func drawGrid() {
        w, h := pdf.GetPageSize()
        //pdf.SetFont("courier", "", 12)
        //pdf.SetTextColor(80, 80, 80)
        pdf.SetDrawColor(200, 200, 200)
        for x := 0.0; x < w; x = x + (w / 20.0) {
                pdf.Line(x, 0, x, h)
                _, lineHt := pdf.GetFontSize()
                pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
        }
        for y := 0.0; y < h; y = y + (w / 20.0) {
                pdf.Line(0, y, w, y)
                pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
        }

}
/*func SummaryInput() {
        pdf.SetFont("times","",14);
        pdf.SetTextColor(180,180,180);

}
*/

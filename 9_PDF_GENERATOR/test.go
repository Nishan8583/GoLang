package main;

import (
  "fmt";
  "github.com/jung-kurt/gofpdf";
)

func main() {
  pdf := gofpdf.New("P","mm","A4","");
  pdf.AddPage(); // Adding a page
  pdf.SetTitle("Report",true);
  //pdf.SetFontSize(10)
  pdf.SetFont("Arial","B",80)
  pdf.SetFillColor(255, 0, 0);

  pdf.WriteAligned(0, 35, "CTMS Report\n", "")
  pdf.Write(35,"\n")
  //pdf.WriteAligned(0, 35, "\n", "c")
  //pdf.Write(35,"REPORT")
  //x1 := pdf.GetLineWidth();
  _,y := pdf.GetXY();
  //fmt.Println(x);
  pdf.Image("./1.png",55,y,0,35,false,"",0,"")
  pdf.SetFont("Times","",20);
  pdf.Write(35,"\nTodays Date")
  err := pdf.OutputFileAndClose("test.pdf");
  if (err != nil) {
    fmt.Println("Could not write to file",err)
  }
}

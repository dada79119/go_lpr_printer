package main

import (
	"bytes"
	"fmt"
	"github.com/signintech/gopdf"
	"log"
	"os/exec"
	"time"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("font", "./util/font/wt001.ttf") // add font-family names font.
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("font", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(30)
	pdf.SetY(40)
	pdf.Text("Receipt")
	pdf.Br(20)
	pdf.SetX(30)
	pdf.Text("print")
	pdf.Br(20)
	pdf.SetX(30)
	pdf.Text("Monday")
	pdf.Br(20)
	pdf.SetX(30)
	pdf.Text(TimeNow().Format("2006-01-02 15:04:05"))
	pdf.Br(20)
	pdf.SetX(30)
	pdf.Text("NT.5")
	//pdf.Cell(nil, "Receipt")
	//pdf.WritePdf("hello.pdf")
	pdfbuf, err := pdf.GetBytesPdfReturnErr()
	if err != nil {
		log.Print(err.Error())
		return
	}
	cmd := exec.Command("lpr")
	var out bytes.Buffer
	cmd.Stdin = bytes.NewBuffer(pdfbuf)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Print(err.Error())
		return
	}
	fmt.Printf("%s\n", out.Bytes())
}

func TimeNow() time.Time {
	now := time.Now()
	local, _ := time.LoadLocation("Asia/Taipei") //localization
	return now.In(local)
}
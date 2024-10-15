package scrape

import (
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/Rob9nn/scrap-n-go/pkg/log"
)

var (
  dir string
)

func init() {
  dir = "data/"
}

func Execute(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.ErrorLog(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

  getHtml(url, body)
}

func getHtml(url string, body []byte) {
  fileName, _ := strings.CutPrefix(url, "https://")
  fileName = strings.ReplaceAll(fileName, ".", "-")
  fileName = strings.ReplaceAll(fileName, "/", "_")

  log.InfoLog("Writing bytes into file : ", dir + fileName)
  err := os.WriteFile(dir + fileName, body, 0644)

	if err != nil {
		log.ErrorLog(err)
	}
}

func getCss() {
  log.ErrorLog("Not Yet Implemented")
}

func buildPdf() {
  log.ErrorLog("Not Yet Implemented")
}

package miaopai

import (
	"fmt"
	"github.com/cnych/gospider/utils"
	"strings"
)

var (
	fakeMobileHeaders = map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Charset":  "UTF-8,*;q=0.5",
		"Accept-Encoding": "gzip,deflate,sdch",
		"Accept-Language": "en-US,en;q=0.8",
		"User-Agent":      "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
	}
)

type Miaopai struct {
	url            string
	outputDir      string
	outputFilename string
}

func NewMiaopai(url, outputDir, outputFilename string) *Miaopai {
	return &Miaopai{
		url:            url,
		outputDir:      outputDir,
		outputFilename: outputFilename,
	}
}

func (m *Miaopai) Download() error {
	doc, err := utils.GetDocument(m.url, fakeMobileHeaders)
	if err != nil {
		return err
	}
	if src, _ := doc.Find("video").Attr("src"); src != "" {
		var destFile string
		if m.outputFilename == "" {
			title := doc.Find("div.perinfo").Find("div.title").Text()
			destFile = fmt.Sprintf("%s/%s.mp4", m.outputDir, strings.Trim(title, " "))
		} else {
			destFile = fmt.Sprintf("%s/%s", m.outputDir, m.outputFilename)
		}
		if err = utils.DownloadFile(src, destFile); err != nil {
			return err
		}
	}
	return nil
}

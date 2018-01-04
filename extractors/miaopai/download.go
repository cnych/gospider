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
	showInfo       bool
}

func NewMiaopai(url, outputDir, outputFilename string, showInfo bool) *Miaopai {
	return &Miaopai{
		url:            url,
		outputDir:      outputDir,
		outputFilename: outputFilename,
		showInfo:       showInfo,
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
		var downloadPre utils.DownloadPre = func(site, title, size string) {
			fmt.Printf("site:				%s\n", site)
			fmt.Printf("title:				%s\n", title)
			fmt.Printf("size:				%s\n", size)
			fmt.Printf("Downloading %s ...\n", title)
		}
		var downloadPost utils.DownloadPost = func(info string) {
			fmt.Println(info)
		}

		if err = utils.DownloadFile(src, destFile, downloadPre, downloadPost); err != nil {
			return err
		}
	}
	return nil
}

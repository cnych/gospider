package extractors

import (
	"github.com/cnych/gospider/extractors/miaopai"
)

type Extractor interface {
	Download() error
}

func NewExtractor(url, outputDir, outputFilename string, showInfo bool) Extractor {
	// TODO, 根据url来实例化解析器
	return miaopai.NewMiaopai(url, outputDir, outputFilename, showInfo)
}

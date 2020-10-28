package main

import (
	"singleTaskCrawfer/engine"
	"singleTaskCrawfer/laomassf/parser"
)

func main() {
	engine.Run(engine.Request{
		ParserFunc: parser.GetBookListPage,
	})
}

package config

const (
	// Parser Names
	GetbookListPage = "GetbookListPage"
	ParseBookList   = "ParseBookList"
	ParseBookDetail = "ParseBookDetail"
	NilParser       = "NilParser"

	// Elasticsearch
	IndexName = "laoma"
	TypeName  = "bookInfo"

	// RPC Endpoints
	ItemSaverRpcFunName = "ItemSaverService.Save"
	CrawlServiceFunName = "CrawlService.Process"
)

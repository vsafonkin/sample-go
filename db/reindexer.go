package db

import (
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/builtin"
)

type Item struct {
	ID       int64  `reindex:"id,,pk"`    // 'id' is primary key
	Name     string `reindex:"name"`      // add index by 'name' field
	Articles []int  `reindex:"articles"`  // add index by articles 'articles' array
	Year     int    `reindex:"year,tree"` // add sortable index by 'year' field
}

func NewReindexerConnect() *reindexer.Reindexer {
	db := reindexer.NewReindex("builtin:///tmp/reindex/sampledb")
	db.OpenNamespace("items", reindexer.DefaultNamespaceOptions(), Item{})
	return db
}

package goepub

import (
	"encoding/json"
	"log"
	"testing"
)

func TestEpubReade(t *testing.T) {
	epub, err := NewEpub("data/诡秘之主-爱潜水的乌贼.epub")
	if err != nil {
		t.Error(err.Error())
		return
	}

	buf, _ := json.Marshal(epub.OPF.TocNcx)
	log.Println(string(buf))
}

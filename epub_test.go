package goepub

import (
	"encoding/json"
	"log"
	"testing"
)

func TestEpubReade(t *testing.T) {
	epub, err := NewEpub("")
	if err != nil {
		t.Error(err.Error())
		return
	}

	buf, _ := json.Marshal(epub.OPF.Manifest)
	log.Println(string(buf))
}

package config

import (
	"encoding/json"
	"testing"
)

func TestGetConfig(t *testing.T) {
	jsonData, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", jsonData)
}

package models

import "encoding/json"

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"

type BGInfo struct {
	Proxies      []string        `json:"proxies"`
}

type RawBgInfo struct {
	Proxies      json.RawMessage `db:"proxies"`
}

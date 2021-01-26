package proxies

import (
	"log"

	"strings"
)

func (f *FClient) CreateProxyAndUser() ([]*ProxyAndUA, error) {
	var res []*ProxyAndUA
	if len(f.proxies) == 0 {
		log.Fatal("Proxies not found")
	}

	for _, v := range f.proxies {
		p := strings.Split(v, ":")
		if len(p) != 4 {
			continue
		}
		x := ProxyAndUA{
			Host: p[0],
			Port: p[1],
			User: p[2],
			Pass: p[3],
		}
		res = append(res, &x)
	}
	return res, nil
}

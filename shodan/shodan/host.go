package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	Latutude     float32 `json:"latitude"`
	CountryCode  string  `json:"country_code"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"county_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
}

type Host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(query string) (*HostSearch, error) {
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseUrl, s.apiKey, query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

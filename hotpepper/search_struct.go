package hotpepper

import "encoding/xml"

type Genre struct {
	Code string `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name string `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Large_area struct {
	Code string `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name string `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Large_service_area struct {
	XMLName xml.Name `xml:"large_service_area,omitempty" json:"large_service_area,omitempty"`
	Code    string   `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    string   `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Middle_area struct {
	Code string `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name string `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Gourmet struct {
	Attrxmlns         string `xml:"xmlns,attr"  json:",omitempty"`
	Results_available string `xml:"http://webservice.recruit.co.jp/HotPepper/ results_available,omitempty" json:"results_available,omitempty"`
	Results_returned  string `xml:"http://webservice.recruit.co.jp/HotPepper/ results_returned,omitempty" json:"results_returned,omitempty"`
	Results_start     string `xml:"http://webservice.recruit.co.jp/HotPepper/ results_start,omitempty" json:"results_start,omitempty"`
	Shop              []Shop `xml:"http://webservice.recruit.co.jp/HotPepper/ shop,omitempty" json:"shop,omitempty"`
}

type Service_area struct {
	Code string `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name string `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Shop struct {
	Address            string             `xml:"http://webservice.recruit.co.jp/HotPepper/ address,omitempty" json:"address,omitempty"`
	Close              string             `xml:"http://webservice.recruit.co.jp/HotPepper/ close,omitempty" json:"close,omitempty"`
	Genre              Genre              `xml:"http://webservice.recruit.co.jp/HotPepper/ genre,omitempty" json:"genre,omitempty"`
	Id                 string             `xml:"http://webservice.recruit.co.jp/HotPepper/ id,omitempty" json:"id,omitempty"`
	Large_area         Large_area         `xml:"http://webservice.recruit.co.jp/HotPepper/ large_area,omitempty" json:"large_area,omitempty"`
	Large_service_area Large_service_area `xml:"http://webservice.recruit.co.jp/HotPepper/ large_service_area,omitempty" json:"large_service_area,omitempty"`
	Lat                string             `xml:"http://webservice.recruit.co.jp/HotPepper/ lat,omitempty" json:"lat,omitempty"`
	Lng                string             `xml:"http://webservice.recruit.co.jp/HotPepper/ lng,omitempty" json:"lng,omitempty"`
	Logo_image         string             `xml:"http://webservice.recruit.co.jp/HotPepper/ logo_image,omitempty" json:"logo_image,omitempty"`
	Lunch              string             `xml:"http://webservice.recruit.co.jp/HotPepper/ lunch,omitempty" json:"lunch,omitempty"`
	Middle_area        Middle_area        `xml:"http://webservice.recruit.co.jp/HotPepper/ middle_area,omitempty" json:"middle_area,omitempty"`
	Name               string             `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
	Open               string             `xml:"http://webservice.recruit.co.jp/HotPepper/ open,omitempty" json:"open,omitempty"`
	Service_area       Service_area       `xml:"http://webservice.recruit.co.jp/HotPepper/ service_area,omitempty" json:"service_area,omitempty"`
	Small_area         Small_area         `xml:"http://webservice.recruit.co.jp/HotPepper/ small_area,omitempty" json:"small_area,omitempty"`
	Station_name       string             `xml:"http://webservice.recruit.co.jp/HotPepper/ station_name,omitempty" json:"station_name,omitempty"`
	Urls               Urls               `xml:"http://webservice.recruit.co.jp/HotPepper/ urls,omitempty" json:"urls,omitempty"`
}

type Small_area struct {
	Code string `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name string `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Urls struct {
	L  string `xml:"http://webservice.recruit.co.jp/HotPepper/ l,omitempty" json:"l,omitempty"`
	M  string `xml:"http://webservice.recruit.co.jp/HotPepper/ m,omitempty" json:"m,omitempty"`
	S  string `xml:"http://webservice.recruit.co.jp/HotPepper/ s,omitempty" json:"s,omitempty"`
	Pc string `xml:",chardata" json:",omitempty"`
}

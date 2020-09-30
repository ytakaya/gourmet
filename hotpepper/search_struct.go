package hotpepper

import "encoding/xml"

type Address struct {
	XMLName xml.Name `xml:"address,omitempty" json:"address,omitempty"`
	Address string   `xml:",chardata" json:",omitempty"`
}

type Close struct {
	XMLName xml.Name `xml:"close,omitempty" json:"close,omitempty"`
	Close   string   `xml:",chardata" json:",omitempty"`
}

type Code struct {
	XMLName xml.Name `xml:"code,omitempty" json:"code,omitempty"`
	Code    string   `xml:",chardata" json:",omitempty"`
}

type Genre struct {
	XMLName xml.Name `xml:"genre,omitempty" json:"genre,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Id struct {
	XMLName xml.Name `xml:"id,omitempty" json:"id,omitempty"`
	Id      string   `xml:",chardata" json:",omitempty"`
}

type L struct {
	XMLName xml.Name `xml:"l,omitempty" json:"l,omitempty"`
	L       string   `xml:",chardata" json:",omitempty"`
}

type Large_area struct {
	XMLName xml.Name `xml:"large_area,omitempty" json:"large_area,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Large_service_area struct {
	XMLName xml.Name `xml:"large_service_area,omitempty" json:"large_service_area,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Lat struct {
	XMLName xml.Name `xml:"lat,omitempty" json:"lat,omitempty"`
	Lat     string   `xml:",chardata" json:",omitempty"`
}

type Lng struct {
	XMLName xml.Name `xml:"lng,omitempty" json:"lng,omitempty"`
	Lng     string   `xml:",chardata" json:",omitempty"`
}

type Logo_image struct {
	XMLName    xml.Name `xml:"logo_image,omitempty" json:"logo_image,omitempty"`
	Logo_image string   `xml:",chardata" json:",omitempty"`
}

type Lunch struct {
	XMLName xml.Name `xml:"lunch,omitempty" json:"lunch,omitempty"`
	Lunch   string   `xml:",chardata" json:",omitempty"`
}

type M struct {
	XMLName xml.Name `xml:"m,omitempty" json:"m,omitempty"`
	M       string   `xml:",chardata" json:",omitempty"`
}

type Middle_area struct {
	XMLName xml.Name `xml:"middle_area,omitempty" json:"middle_area,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Name struct {
	XMLName xml.Name `xml:"name,omitempty" json:"name,omitempty"`
	Name    string   `xml:",chardata" json:",omitempty"`
}

type Open struct {
	XMLName xml.Name `xml:"open,omitempty" json:"open,omitempty"`
	Open    string   `xml:",chardata" json:",omitempty"`
}

type Gourmet struct {
	XMLName           xml.Name           `xml:"results,omitempty" json:"results,omitempty"`
	Attrxmlns         string             `xml:"xmlns,attr"  json:",omitempty"`
	Results_available *Results_available `xml:"http://webservice.recruit.co.jp/HotPepper/ results_available,omitempty" json:"results_available,omitempty"`
	Results_returned  *Results_returned  `xml:"http://webservice.recruit.co.jp/HotPepper/ results_returned,omitempty" json:"results_returned,omitempty"`
	Results_start     *Results_start     `xml:"http://webservice.recruit.co.jp/HotPepper/ results_start,omitempty" json:"results_start,omitempty"`
	Cshop             []*Cshop           `xml:"http://webservice.recruit.co.jp/HotPepper/ shop,omitempty" json:"shop,omitempty"`
}

type Results_available struct {
	XMLName           xml.Name `xml:"results_available,omitempty" json:"results_available,omitempty"`
	Results_available string   `xml:",chardata" json:",omitempty"`
}

type Results_returned struct {
	XMLName          xml.Name `xml:"results_returned,omitempty" json:"results_returned,omitempty"`
	Results_returned string   `xml:",chardata" json:",omitempty"`
}

type Results_start struct {
	XMLName       xml.Name `xml:"results_start,omitempty" json:"results_start,omitempty"`
	Results_start string   `xml:",chardata" json:",omitempty"`
}

type S struct {
	XMLName xml.Name `xml:"s,omitempty" json:"s,omitempty"`
	S       string   `xml:",chardata" json:",omitempty"`
}

type Service_area struct {
	XMLName xml.Name `xml:"service_area,omitempty" json:"service_area,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Cshop struct {
	XMLName            xml.Name            `xml:"shop,omitempty" json:"shop,omitempty"`
	Address            *Address            `xml:"http://webservice.recruit.co.jp/HotPepper/ address,omitempty" json:"address,omitempty"`
	Close              *Close              `xml:"http://webservice.recruit.co.jp/HotPepper/ close,omitempty" json:"close,omitempty"`
	Genre              *Genre              `xml:"http://webservice.recruit.co.jp/HotPepper/ genre,omitempty" json:"genre,omitempty"`
	Id                 *Id                 `xml:"http://webservice.recruit.co.jp/HotPepper/ id,omitempty" json:"id,omitempty"`
	Large_area         *Large_area         `xml:"http://webservice.recruit.co.jp/HotPepper/ large_area,omitempty" json:"large_area,omitempty"`
	Large_service_area *Large_service_area `xml:"http://webservice.recruit.co.jp/HotPepper/ large_service_area,omitempty" json:"large_service_area,omitempty"`
	Lat                *Lat                `xml:"http://webservice.recruit.co.jp/HotPepper/ lat,omitempty" json:"lat,omitempty"`
	Lng                *Lng                `xml:"http://webservice.recruit.co.jp/HotPepper/ lng,omitempty" json:"lng,omitempty"`
	Logo_image         *Logo_image         `xml:"http://webservice.recruit.co.jp/HotPepper/ logo_image,omitempty" json:"logo_image,omitempty"`
	Lunch              *Lunch              `xml:"http://webservice.recruit.co.jp/HotPepper/ lunch,omitempty" json:"lunch,omitempty"`
	Middle_area        *Middle_area        `xml:"http://webservice.recruit.co.jp/HotPepper/ middle_area,omitempty" json:"middle_area,omitempty"`
	Name               *Name               `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
	Open               *Open               `xml:"http://webservice.recruit.co.jp/HotPepper/ open,omitempty" json:"open,omitempty"`
	Service_area       *Service_area       `xml:"http://webservice.recruit.co.jp/HotPepper/ service_area,omitempty" json:"service_area,omitempty"`
	Small_area         *Small_area         `xml:"http://webservice.recruit.co.jp/HotPepper/ small_area,omitempty" json:"small_area,omitempty"`
	Station_name       *Station_name       `xml:"http://webservice.recruit.co.jp/HotPepper/ station_name,omitempty" json:"station_name,omitempty"`
	Urls               *Urls               `xml:"http://webservice.recruit.co.jp/HotPepper/ urls,omitempty" json:"urls,omitempty"`
}

type Small_area struct {
	XMLName xml.Name `xml:"small_area,omitempty" json:"small_area,omitempty"`
	Code    *Code    `xml:"http://webservice.recruit.co.jp/HotPepper/ code,omitempty" json:"code,omitempty"`
	Name    *Name    `xml:"http://webservice.recruit.co.jp/HotPepper/ name,omitempty" json:"name,omitempty"`
}

type Station_name struct {
	XMLName      xml.Name `xml:"station_name,omitempty" json:"station_name,omitempty"`
	Station_name string   `xml:",chardata" json:",omitempty"`
}

type Urls struct {
	XMLName xml.Name `xml:"urls,omitempty" json:"urls,omitempty"`
	Pc      *Pc      `xml:"http://webservice.recruit.co.jp/HotPepper/ pc,omitempty" json:"pc,omitempty"`
}

type Pc struct {
	XMLName xml.Name `xml:"pc,omitempty" json:"pc,omitempty"`
	L       *L       `xml:"http://webservice.recruit.co.jp/HotPepper/ l,omitempty" json:"l,omitempty"`
	M       *M       `xml:"http://webservice.recruit.co.jp/HotPepper/ m,omitempty" json:"m,omitempty"`
	S       *S       `xml:"http://webservice.recruit.co.jp/HotPepper/ s,omitempty" json:"s,omitempty"`
	Pc      string   `xml:",chardata" json:",omitempty"`
}

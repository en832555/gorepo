package pxml

import "encoding/xml"

type Plant struct {
	XMLName xml.Name "Xml:\"plant\""
	Id int "xml:\"id,attr\""
	Name string "xml:\"name\""
	Origin []string "xml:\"origin\""
}
package pxml

import "encoding/xml"

type (
	Nesting struct {
		XMLName xml.Name      "xml:\"nesting\""
		Child   []*Plant "xml:\"parent>child>plant\""
		Parent  []*Plant "xml:\"parent>plant\""
	}
)
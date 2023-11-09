package types

import (
	"encoding/xml"
)

type XMLLog struct {
	XMLName xml.Name     `xml:"log"`
	Commits []*XmlCommit `xml:"commit"`
}

type XmlCommit struct {
	XMLName xml.Name `xml:"commit"`
	Hash    string   `xml:"hash"`
	Names   string   `xml:"names"`
	Date    string   `xml:"date"`
	Author  string   `xml:"author"`
	Subject string   `xml:"subject"`
	Body    string   `xml:"body"`
}

func ParseLogXML(out string) (XMLLog, error) {
	var log XMLLog
	if err := xml.Unmarshal([]byte(out), &log); err != nil {
		return log, err
	}
	return log, nil
}

package types

import (
	"encoding/xml"
)

// XmlLog is used to read commit messages from the git log
// Via log --format the log command creates xml output that can be parsed
// into these structs. XmlLog being just the container for the list of XmlCommit structs
//
//	<commit>
//	  <hash>55d061</hash>
//	  <names><![CDATA[head]]></names>
//	  <date>2023-04-23</date>
//	  <author><![CDATA[Sebastian Feldmann <sf@sebastian-feldmann.info>]]></author>
//	  <subject><![CDATA[Fix example docs]]></subject>
//	  <body><![CDATA[]]></body>
//	</commit>
type XmlLog struct {
	XMLName xml.Name     `xml:"log"`
	Commits []*XmlCommit `xml:"commit"`
}

// XmlCommit is used to represent  commits in a commit log
type XmlCommit struct {
	XMLName xml.Name `xml:"commit"`
	Hash    string   `xml:"hash"`
	Names   string   `xml:"names"`
	Date    string   `xml:"date"`
	Author  string   `xml:"author"`
	Subject string   `xml:"subject"`
	Body    string   `xml:"body"`
}

// ParseLogXml is used to read infos from the git log
func ParseLogXml(out string) (XmlLog, error) {
	var log XmlLog
	if err := xml.Unmarshal([]byte(out), &log); err != nil {
		return log, err
	}
	return log, nil
}

package log

import (
	"github.com/captainhook-go/captainhook/git/types"
)

const (
	XmlFormat = "<commit>%n" +
		"<hash>%h</hash>%n" +
		"<names><![CDATA[%d]]></names>%n" +
		"<date>%ci</date>%n" +
		"<author><![CDATA[%an]]></author>%n" +
		"<subject><![CDATA[%s]]></subject>%n" +
		"<body><![CDATA[%n%b%n]]></body>%n" +
		"</commit>"
)

func AbbrevCommit(g *types.Cmd) {
	g.AddOption("--abbrev-commit")
}

func AuthoredBy(name string) func(g *types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--author=" + name)
	}
}

func Format(format string) func(g *types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--format='" + format + "'")
	}
}

func ParseXML(out string) ([]*types.Commit, error) {
	var log []*types.Commit
	xmlLog, err := types.ParseLogXML(out)
	if err != nil {
		return log, err
	}
	for _, c := range xmlLog.Commits {
		log = append(log, types.CreateCommitFromXML(c))
	}
	return log, nil
}

func FromTo(from, to string) func(g *types.Cmd) {
	if to == "" {
		to = "HEAD"
	}
	return func(g *types.Cmd) {
		g.AddOption(from + ".." + to)
	}
}

func InTimeFrame(after, before string) func(g *types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--after=" + after)
		g.AddOption("--before=" + before)
	}
}

func NameOnly(g *types.Cmd) {
	g.AddOption("--name-only")
}

func NameStatus(g *types.Cmd) {
	g.AddOption("--name-status")
}

func NoCommitID(g *types.Cmd) {
	g.AddOption("--no-commit-id")
}

func NoMerges(g *types.Cmd) {
	g.AddOption("--no-merges")
}

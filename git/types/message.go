package types

import (
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type CommitMessage struct {
	commentChar  string
	raw          string
	rawLines     []string
	contentLines []string
}

func (m *CommitMessage) CommentChar() string {
	return m.commentChar
}

func (m *CommitMessage) Raw() string {
	return m.raw
}

func (m *CommitMessage) Message() string {
	subject := m.Subject()
	body := m.Body()
	glue := ""

	if body != "" {
		glue = "\n\n"
	}
	return subject + glue + body
}

func (m *CommitMessage) Lines() []string {
	return m.contentLines
}

func (m *CommitMessage) Subject() string {
	if len(m.contentLines) > 0 {
		return m.contentLines[0]
	}
	return ""
}

func (m *CommitMessage) Body() string {
	return strings.Join(m.BodyLines(), "\n")
}

func (m *CommitMessage) BodyLines() []string {
	if len(m.contentLines) > 2 {
		return m.contentLines[2:]
	}
	return []string{}
}

func (m *CommitMessage) IsFixup() bool {
	return strings.HasPrefix(m.raw, "fixup!")
}

func (m *CommitMessage) IsSquash() bool {
	return strings.HasPrefix(m.raw, "squash!")
}

func (m *CommitMessage) IsEmpty() bool {
	return strings.TrimSpace(m.Message()) == ""
}

func NewCommitMessage(msg string, commentChar string) *CommitMessage {
	rawLines := io.SplitLines(msg)

	m := CommitMessage{
		commentChar:  commentChar,
		raw:          msg,
		rawLines:     rawLines,
		contentLines: extractContentLines(rawLines, commentChar),
	}
	return &m
}

func NewCommitMessageFromFile(file, commentChar string) (*CommitMessage, error) {
	data, err := io.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewCommitMessage(string(data), commentChar), nil
}

func extractContentLines(rawLines []string, commentChar string) []string {
	var lines []string
	for _, line := range rawLines {
		// if we handle a comment line
		if strings.HasPrefix(line, commentChar) {
			// check if we should ignore all following lines
			if strings.Contains(line, "------------------------ >8 ------------------------") {
				break
			}
			// or only the current one
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

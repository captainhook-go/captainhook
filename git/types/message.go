package types

import (
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

// CommitMessage represents a git commit message
// You can access the message's Subject and Body
type CommitMessage struct {
	commentChar  string
	raw          string
	rawLines     []string
	contentLines []string
}

// CommentChar returns the configured comment character, default is `#`
func (m *CommitMessage) CommentChar() string {
	return m.commentChar
}

// Raw returns the raw content of the commit message with comments and everything
func (m *CommitMessage) Raw() string {
	return m.raw
}

// Message returns the commit message without comments
func (m *CommitMessage) Message() string {
	subject := m.Subject()
	body := m.Body()
	glue := ""

	if body != "" {
		glue = "\n\n"
	}
	return subject + glue + body
}

// Lines returns all lines except comment lines
func (m *CommitMessage) Lines() []string {
	return m.contentLines
}

// Subject returns the first line of the commit message
func (m *CommitMessage) Subject() string {
	if len(m.contentLines) > 0 {
		return m.contentLines[0]
	}
	return ""
}

// Body returns a string starting from line 3
func (m *CommitMessage) Body() string {
	return strings.Join(m.BodyLines(), "\n")
}

// BodyLines returns all lines starting from line 3
func (m *CommitMessage) BodyLines() []string {
	if len(m.contentLines) > 2 {
		return m.contentLines[2:]
	}
	return []string{}
}

// IsFixup indicates if a commit is a `--fixup` commit
func (m *CommitMessage) IsFixup() bool {
	return strings.HasPrefix(m.raw, "fixup!")
}

// IsSquash indicates if a commit is a `--squash` commit
func (m *CommitMessage) IsSquash() bool {
	return strings.HasPrefix(m.raw, "squash!")
}

// IsEmpty returns true if the commit message does not have any content
func (m *CommitMessage) IsEmpty() bool {
	return strings.TrimSpace(m.Message()) == ""
}

// NewCommitMessage is the CommitMessage struct constructor
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

// NewCommitMessageFromFile creates a CommitMessage struct by reading the contents a file
func NewCommitMessageFromFile(file, commentChar string) (*CommitMessage, error) {
	data, err := io.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewCommitMessage(string(data), commentChar), nil
}

// extractContentLines finds all none comment lines in a commit message
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

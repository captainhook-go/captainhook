package rules

import (
	"fmt"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type CapitalizeSubject struct{}

func (r *CapitalizeSubject) AppliesTo(msg types.CommitMessage) (bool, string) {
	firstChar := io.SubString(msg.Subject(), 0, 1)
	if strings.ToUpper(firstChar) == firstChar {
		return true, ""
	}
	return false, "subject line has to start with an upper case letter"
}

type LimitBodyLineLength struct {
	length int
}

func (r *LimitBodyLineLength) AppliesTo(msg types.CommitMessage) (bool, string) {
	for nr, line := range msg.BodyLines() {
		if len(line) > r.length {
			return false, fmt.Sprintf("line %d of your body exceeds the line limit of %d", nr, r.length)
		}
	}
	return true, ""
}

type LimitSubjectLineLength struct {
	length int
}

func (r *LimitSubjectLineLength) AppliesTo(msg types.CommitMessage) (bool, string) {
	subjectLength := len(msg.Subject())
	if subjectLength > r.length {
		return false, fmt.Sprintf("subject length of %d exceeds the limit of %d", subjectLength, r.length)
	}
	return true, ""
}

type MsgNotEmpty struct{}

func (r *MsgNotEmpty) AppliesTo(msg types.CommitMessage) (bool, string) {
	if msg.Message() == "" {
		return false, "commit message can not be empty"
	}
	return true, ""
}

type NoPeriodOnSubjectEnd struct{}

func (r *NoPeriodOnSubjectEnd) AppliesTo(msg types.CommitMessage) (bool, string) {
	return true, ""
}

type SeparateSubjectFromBodyWithBlankLine struct{}

func (r *SeparateSubjectFromBodyWithBlankLine) AppliesTo(msg types.CommitMessage) (bool, string) {
	lines := msg.Lines()
	if len(lines) > 1 && lines[1] != "" {
		return false, "subject and body should be separated with a blank line"
	}
	return true, ""
}

type UseImperativeMood struct{}

func (r *UseImperativeMood) AppliesTo(msg types.CommitMessage) (bool, string) {
	hint := "a commit message subject should always complete the following sentence\n" +
		"this commit will [YOUR COMMIT MESSAGE].\n"

	blacklist := []string{
		"added",
		"changed",
		"created",
		"deleted",
		"fixed",
		"reformatted",
		"removed",
		"updated",
		"uploaded",
	}
	subject := msg.Subject()
	for _, word := range blacklist {
		if strings.Contains(subject, word) {
			return false, fmt.Sprintf("%ssubject should not contain '%s'", hint, word)
		}
	}
	return true, ""
}

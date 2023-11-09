package io

type CollectorIO struct {
	verbosity int
	arguments map[string]string
	messages  []*CollectedMessage
}

type CollectedMessage struct {
	Verbosity int
	Message   string
}

func NewCollectorIO(verbosity int, arguments map[string]string) *CollectorIO {
	io := CollectorIO{verbosity: verbosity, arguments: arguments}
	return &io
}

func (c *CollectorIO) Verbosity() int {
	return c.verbosity
}

func (c *CollectorIO) Arguments() map[string]string {
	return c.arguments
}
func (c *CollectorIO) Argument(name, defaultValue string) string {
	value, ok := c.arguments[name]
	if !ok {
		value = defaultValue
	}
	return value
}

func (c *CollectorIO) StandardInput() []string {
	return make([]string, 0)
}

func (c *CollectorIO) IsInteractive() bool {
	return false
}

func (c *CollectorIO) IsDebug() bool {
	return c.verbosity == DEBUG
}

func (c *CollectorIO) IsVerbose() bool {
	return c.verbosity == VERBOSE
}

func (c *CollectorIO) Write(message string, newline bool, verbosity int) {
	var linebreak = ""
	if newline {
		linebreak = "\n"
	}
	c.messages = append(c.messages, &CollectedMessage{verbosity, message + linebreak})
}

func (c *CollectorIO) Ask(message string, defaultValue string) string {
	value, err := getUserInput(message)
	if err != nil {
		c.Write("can't read from std input", true, NORMAL)
	}
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func (c *CollectorIO) HasCollectedMessages() bool {
	return len(c.messages) > 0
}

func (c *CollectorIO) HasCollectedMessagesForVerbosity(verbosity int) bool {
	for _, m := range c.messages {
		if verbosity >= m.Verbosity {
			return true
		}
	}
	return false
}

func (c *CollectorIO) Messages() []*CollectedMessage {
	return c.messages
}

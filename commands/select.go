package commands

import (
	"errors"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/utf7"
)

// A SELECT command.
// If ReadOnly is set to true, the EXAMINE command will be used instead.
// See RFC 3501 section 6.3.1
type Select struct {
	Mailbox  string
	ReadOnly bool
}

func (cmd *Select) Command() *imap.Command {
	name := imap.Select
	if cmd.ReadOnly {
		name = imap.Examine
	}

	mailbox, _ := utf7.Encoder.String(cmd.Mailbox)

	return &imap.Command{
		Name:      name,
		Arguments: []interface{}{mailbox},
	}
}

func (cmd *Select) Parse(fields []interface{}) error {
	if len(fields) < 1 {
		return errors.New("No enough arguments")
	}

	if mailbox, ok := fields[0].(string); !ok {
		return errors.New("Mailbox name must be a string")
	} else if mailbox, err := utf7.Decoder.String(mailbox); err != nil {
		return err
	} else {
		cmd.Mailbox = imap.CanonicalMailboxName(mailbox)
	}

	return nil
}

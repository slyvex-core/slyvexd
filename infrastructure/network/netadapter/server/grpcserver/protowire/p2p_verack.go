package protowire

import (
	"github.com/slyvex-core/slyvexd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SlyvexdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SlyvexdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *SlyvexdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}

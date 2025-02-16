package protowire

import (
	"github.com/slyvex-core/slyvexd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SlyvexdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SlyvexdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *SlyvexdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}

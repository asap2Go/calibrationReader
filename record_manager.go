package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
)

func getValue(m *a2l.Module, c *a2l.Characteristic) (error, interface{}) {
	var err error

	//get record layout for characteristic
	if !c.DepositSet {
		err = errors.New("no deposit set in characteristic " + c.Name)
		return err, nil
	}

	rl, exists := m.RecordLayouts[c.Deposit]
	if !exists {
		err = errors.New("no record layout found for deposit identifier" + c.Deposit + " of characteristic " + c.Name)
		return err, nil
	}

	return err, nil
}

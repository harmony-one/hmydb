package cmd

type oneAddress struct {
	address string
}

func (oneAddress oneAddress) String() string {
	return oneAddress.address
}

func (oneAddress *oneAddress) Set(s string) error {
	oneAddress.address = s
	return nil
}

func (oneAddress oneAddress) Type() string {
	return "one-address"
}

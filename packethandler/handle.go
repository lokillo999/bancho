package packethandler

import (
	"fmt"
	"git.zxq.co/ripple/go-bancho/packethandler/logindata"
	"time"
)

const deliverTimeout = time.Millisecond * 15

// Handle takes an input and writes data to an output. Not very hard.
func Handle(input []byte, output *[]byte, token string) (string, error) {
	sendBackToken := false

	// The user wants to login
	if token == "" {
		sendBackToken = true
		d, err := logindata.Unmarshal(input)
		if err != nil {
			return "", err
		}
		token, err = Login(d)
		if err != nil {
			return token, err
		}
	}

	// Make up response, putting together all the accumulated packets.
	for {
		packet := Tokens[token].Pop()
		if packet == nil {
			break
		}
		*output = append(*output, packet.Content...)
	}
	fmt.Printf("% x\n", *output)

	if sendBackToken {
		return token, nil
	}
	return "", nil
}
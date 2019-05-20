package main

import "time"

type ByteCOunter int

func (c *ByteCOunter) Write(p []byte) (int, error) {
	*c += ByteCOunter(len(p))
	return len(p), nil
}

func main() {
	time.Sleep(time.Seconds)

}

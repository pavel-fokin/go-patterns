package main

func FanIn(msgs ...<-chan string) chan string {
	out := make(chan string)

	for _, each := range msgs {

		go func(input <-chan string) {
			for {
				// read value from input
				val, ok := <-input
				// break if channel is closed
				if !ok {
					break
				}
				out <- val
			}
		}(each)

	}
	return out
}

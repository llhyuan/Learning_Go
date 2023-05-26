package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )
//
// func main() {
// 	s := "The quick brown fox jumped over the lazy dog"
// 	sr := strings.NewReader(s)
//
// 	out, err := countLetters(sr)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(out)
// }
//
// func countLetters(r io.Reader) (map[string]int, error) {
// 	buf := make([]byte, 100)
// 	out := map[string]int{}
//
// 	for {
// 		n, err := r.Read(buf)
// 		for _, b := range buf[:n] {
// 			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
// 				out[string(b)]++
// 			}
// 		}
//
// 		if err == io.EOF {
// 			return out, nil
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
//
// 	}
// }

func main() {
	tick := time.NewTicker(2 * time.Second)

	ch := make(chan bool)

	go func() {
		time.Sleep(15 * time.Second)

		close(ch)
		tick.Stop()
	}()

	for {
		select {
		case <-ch:
			fmt.Println("The program is timed out.")
			return
		case tickerTime := <-tick.C:
			fmt.Println("The current ticker is ", tickerTime)
		}
	}
}

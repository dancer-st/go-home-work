package main

import (
	"fmt"
	"time"
)

func main() {
	addRequestfun := addRequest()

	addRequestfun("string")

	time.Sleep(2 * time.Second)

	addRequestfun("string2")
	addRequestfun("string")
}

func addRequest() func(string) {
	var requests = make(map[int64][]string)

	return func(request string) {
		now := time.Now().Unix()
		tenSecondsAgo := now - 10
		duplicate := false

		for key, _ := range requests {
			if key < tenSecondsAgo {
				fmt.Println("delete key", key)
				delete(requests, key)
			} else {
				if contains(request, requests[key]) {
					duplicate = true
				}
			}
		}

		if duplicate {
			fmt.Println("Duplicate request " + request)
		} else {
			requests[now] = append(requests[now], request)
		}

		fmt.Println(requests)
	}
}

func contains(str1 string, slice []string) bool {
	for _, value := range slice {
		if value == str1 {
			return true
		}
	}

	return false
}

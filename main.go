package main

import (
	"fmt"
	"github.com/aThebigbot/base64/base64"
	"io/ioutil"
	"os"
)

func main() {
	switch os.Args[1] {
	case "--encode", "-e":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
			return
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			if len(os.Args) <= 4 {
				content, err := ioutil.ReadFile(os.Args[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(base64.Encode(string(content)))
			} else {
				content, err := ioutil.ReadFile(os.Args[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				ioutil.WriteFile(os.Args[4], []byte(base64.Encode(string(content))), 0644)
			}
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			if len(os.Args) <= 4 {
				fmt.Println(base64.Encode(os.Args[3]))
			} else {
				ioutil.WriteFile(os.Args[4], []byte(base64.Encode(os.Args[3])), 0644)
			}
		} else {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
		}
	case "--decode", "-d":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
			return
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			if len(os.Args) <= 4 {
				content, err := ioutil.ReadFile(os.Args[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(base64.Decode(string(content)))
			} else {
				content, err := ioutil.ReadFile(os.Args[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				ioutil.WriteFile(os.Args[4], []byte(base64.Decode(string(content))), 0644)
			}
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			if len(os.Args) <= 4 {
				fmt.Println(base64.Decode(os.Args[3]))
			} else {
				ioutil.WriteFile(os.Args[4], []byte(base64.Decode(os.Args[3])), 0644)
			}
		} else {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
		}
	default:
		fmt.Println("Usage: base64 [--encode | --decode] [-f <file> | -s string]")
	}
}

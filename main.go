package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/aThebigbot/base64/base64"
)

func main() {
	switch os.Args[1] {
	case "--encode":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
			return;
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			content, err := ioutil.ReadFile(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(base64.Encode(string(content)))
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			fmt.Println(base64.Encode(os.Args[3]))
		} else {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
		}
	case "--decode":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
			return;
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			content, err := ioutil.ReadFile(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(base64.Decode(string(content)))
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			fmt.Println(base64.Decode(os.Args[3]))
		} else {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
		}
	case "-e":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
			return;
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			content, err := ioutil.ReadFile(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(base64.Encode(string(content)))
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			fmt.Println(base64.Encode(os.Args[3]))
		} else {
			fmt.Println("Usage: base64 --encode [-f|-s] <data>")
		}
	case "-d":
		if len(os.Args) <= 3 {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
			return;
		}
		if os.Args[2] == "-f" || os.Args[2] == "--file" {
			content, err := ioutil.ReadFile(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(base64.Decode(string(content)))
		} else if os.Args[2] == "-s" || os.Args[2] == "--string" {
			fmt.Println(base64.Decode(os.Args[3]))
		} else {
			fmt.Println("Usage: base64 --decode [-f|-s] <data>")
		}
	default:
		fmt.Println("Usage: base64 [--encode | --decode] [-f <file> | -s string]")
	}
}
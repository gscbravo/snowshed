package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// get -t flag
	alltypes := []string{"twit", "disc", "mast"}
	tflagptr := flag.String("t", "", fmt.Sprintf("snowflake id type, valid: %v", alltypes))
	flag.Parse()
	tflag := strings.ToLower(strings.TrimSpace(*tflagptr))

	// leave if empty
	if tflag == "" {
		flag.Usage()
		os.Exit(1)
	}

	if tflag == "twit" {
		// get user input and turn to uint64
		fmt.Print("Twitter Snowflake ID: ")
		var snowtext string
		fmt.Scanln(&snowtext)
		snowflake, _ := strconv.ParseUint(snowtext, 10, 64)

		// unix timestamp in milliseconds
		timestamp := (snowflake >> 22) + 1288834974657
		// misc process info
		machineid := (snowflake & 0x3FF000) >> 12
		seqnum := snowflake & 0xFFF

		// output parts
		fmt.Printf("Timestamp: %v\n", time.Unix(int64(timestamp/1000), 0).UTC().Format("Mon 2006-01-02 3:04:05 PM MST"))
		fmt.Printf("Machine ID: %v\n", machineid)
		fmt.Printf("Sequence Number: %v\n", seqnum)
	} else if tflag == "disc" {
		// get user input and turn to uint64
		fmt.Print("Discord Snowflake ID: ")
		var snowtext string
		fmt.Scanln(&snowtext)
		snowflake, _ := strconv.ParseUint(snowtext, 10, 64)

		// unix timestamp in milliseconds
		timestamp := (snowflake >> 22) + 1420070400000
		// misc process info
		workerid := (snowflake & 0x3E0000) >> 17
		processid := (snowflake & 0x1F000) >> 12
		increment := snowflake & 0xFFF

		// output parts
		fmt.Printf("Timestamp: %v\n", time.Unix(int64(timestamp/1000), 0).UTC().Format("Mon 2006-01-02 3:04:05 PM MST"))
		fmt.Printf("Worker ID: %v\n", workerid)
		fmt.Printf("Process ID: %v\n", processid)
		fmt.Printf("Increment: %v\n", increment)
	} else if tflag == "mast" {
		// get user input and turn to uint64
		fmt.Print("Mastodon Snowflake ID: ")
		var snowtext string
		fmt.Scanln(&snowtext)
		snowflake, _ := strconv.ParseUint(snowtext, 10, 64)

		// unix timestamp in milliseconds
		timestamp := snowflake >> 16
		// misc data hashed using table name, secret salt, and timestamp
		// can't be reversed further without brute force
		seqdata := snowflake & 0xFFFF

		// output parts
		fmt.Printf("Timestamp: %v\n", time.Unix(int64(timestamp/1000), 0).UTC().Format("Mon 2006-01-02 3:04:05 PM MST"))
		fmt.Printf("Sequence Data: %v\n", seqdata)
	} else {
		fmt.Printf("Type not recognized, valid: %v\n", alltypes)
		os.Exit(1)
	}
}

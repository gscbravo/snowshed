# snowshed

## About

Simple reverse engineering of various Snowflake IDs.

### Snowflake ID Types

- twitter
	- timestamp (42 bits)
	- machine id (10 bits)
	- sequence number (12 bits)
- discord
	- timestamp (42 bits)
	- worker id (5 bits)
	- process id (5 bits)
	- increment (12 bits)
- mastodon
	- timestamp (48 bits)
	- sequence data (16 bits)

## Usage

```
go install github.com/gscbravo/snowshed@latest
showshed -t TYPE
```

Valid types are listed above in the section under "Snowflake ID Types".

## TODO

- [ ] More Snowflake IDs when I'm aware of them

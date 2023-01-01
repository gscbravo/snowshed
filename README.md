# snowshed

## About

Simple reverse engineering of various Snowflake IDs.

### Snowflake ID Types

- (twi)tter
	- timestamp (42 bits)
	- machine id (10 bits)
	- sequence number (12 bits)
- (disc)ord
	- timestamp (42 bits)
	- worker id (5 bits)
	- process id (5 bits)
	- increment (12 bits)
- (mast)odon
	- timestamp (48 bits)
	- sequence data (16 bits)

## Usage

```
go install github.com/gscbravo/snowshed
showshed -t TYPE
```

## TODO

- [ ] More Snowflake IDs when I'm aware of them

# librmonitor

Go library for handling RMonitor protocol messages.

RMonitor is a popular text protocol for communicating race information over serial or TCP connection.

Used by software such as MyLaps Orbits, Westhold Race Manager, Tag Heuer and more.

## Docs

Mostly self-explanatory.

Each message type is represented as a struct. For each message type, 
there is a corresponding `ToX` method to parse a line (and its string parts) into 
a message type.

There's a top-level `Parse` method as well that takes in any string message and
attempts to convert to corresponding message type. 

## Data

In `data/sebring.txt` there is test data to parse, newline delimited text.

## Simulator

In `cmd/simulator` is a simulator app that reads from a text file and serves RMonitor
clients on a TCP port and address of your choice. The default data in `data/sebring.txt` 
is a good starting point for this, but you can substitute any file of your own. 

## License

Apache License 2.0, see LICENSE.
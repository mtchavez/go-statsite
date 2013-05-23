go-statsite [![Build Status](https://travis-ci.org/mtchavez/go-statsite.png)](https://travis-ci.org/mtchavez/go-statsite)
===========

Go package to talk to Statsite - https://github.com/armon/statsite

## Requirements

### Statsite

Follow install instructions [here](https://github.com/armon/statsite#install) to get
statsite up and running

## Installation

```go
  go get github.com/mtchavez/go-statsite/statsite
```

## Usage

### Client

Setting up a new client takes a string of the ip and port statsite is running on.

```go
client, err := statsite.NewClient("0.0.0.0:8125")
if err != nil {
  log.Println("Unable to connect to statsite: ", err)
}
```
If any error occurs connecting to statsite it will be returned in ```err```.

### Sending Messages

The following messages are implemented as ```structs```.

* Message - Generic message (takes a Key, Value and MType)
* KVMsg - Key/Value message (takes a Key and a Value)
* GMsg - Key/Value message compatible with statsd gauges (takes a Key and a Value)
* TimeMsg - Timer message (takes a Key and a Value)
* CountMsg - Counter message (takes a Key and a Value)
* SetMsg - Unique Set message (takes a Key and a Value)

Example:

```
client, err := statsite.NewClient("0.0.0.0:8125")
if err != nil {
  log.Println("Unable to connect to statsite: ", err)
}
msg := &statsite.CountMsg{"rewards", "1"}
for i := 0; i < 20; i++ {
  client.Emit(msg)
}
```

Your statsite log should have similar output if successful:

```
May 23 14:28:06 el-chavo-2.local statsite[20400] <Debug>: Accepted client connection: 127.0.0.1 64538 [8]
May 23 14:28:06 el-chavo-2.local statsite[20400] <Debug>: Closed client connection. [8]
May 23 14:28:06 el-chavo-2.local statsite[20400] <Debug>: Closed connection. [8]
counts.rewards|20.000000|1369344489
```

## Test

Tests use Gocheck for assertions. You can see more [here](http://labix.org/gocheck)

Then run using the normal ```go test``` command

## TODO

* Add client integration tests
* Implement Statsite binary protocol

## License

Written by Chavez

Released under the MIT License: http://www.opensource.org/licenses/mit-license.php

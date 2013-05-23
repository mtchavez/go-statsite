package statsite

import (
	"fmt"
)

type Messenger interface {
	String() string
}

type Message struct {
	Key   string
	Value string
	Mtype string
}

type KVMsg struct {
	Key   string
	Value string
}

type GMsg struct {
	Key   string
	Value string
}

type TimeMsg struct {
	Key   string
	Value string
}

type CountMsg struct {
	Key   string
	Value string
}

type SetMsg struct {
	Key   string
	Value string
}

var MTypes []string = []string{
	"kv", // - Simple Key/Value.
	"g",  // - Same as kv, compatibility with statsd gauges
	"ms", // - Timer.
	"c",  // - Counter.
	"s",  // - Unique Set
}

func (m Message) String() string { return formatted(m.Key, m.Value, m.Mtype) }

func (m KVMsg) String() string { return formatted(m.Key, m.Value, "kv") }

func (m GMsg) String() string { return formatted(m.Key, m.Value, "g") }

func (m TimeMsg) String() string { return formatted(m.Key, m.Value, "ms") }

func (m CountMsg) String() string { return formatted(m.Key, m.Value, "c") }

func (m SetMsg) String() string { return formatted(m.Key, m.Value, "s") }

func formatted(k string, v string, t string) string {
	return fmt.Sprintf("%v:%v|%v\n", k, v, t)
}

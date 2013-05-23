package statsite

import (
	. "launchpad.net/gocheck"
)

func (s *MySuite) TestMtypes(c *C) {
	c.Assert(MTypes, HasLen, 5)
	expected := []string{"kv", "g", "ms", "c", "s"}
	c.Assert(MTypes, DeepEquals, expected)
}

func (s *MySuite) TestMessengerInterface(c *C) {
	var m Messenger

	msg := &Message{}
	c.Assert(msg, Implements, &m)

	msg1 := &KVMsg{}
	c.Assert(msg1, Implements, &m)

	msg2 := &GMsg{}
	c.Assert(msg2, Implements, &m)

	msg3 := &TimeMsg{}
	c.Assert(msg3, Implements, &m)

	msg4 := &CountMsg{}
	c.Assert(msg4, Implements, &m)

	msg5 := &SetMsg{}
	c.Assert(msg5, Implements, &m)
}

func (s *MySuite) TestMessageString(c *C) {
	msg := &Message{"users", "1", "asdf"}
	c.Assert(msg.String(), Equals, "users:1|asdf\n")
}

func (s *MySuite) TestKVMsgString(c *C) {
	msg := &KVMsg{"user_4", "john_doe"}
	c.Assert(msg.String(), Equals, "user_4:john_doe|kv\n")
}

func (s *MySuite) TestGmsgString(c *C) {
	msg := &GMsg{"user_4", "john_doe"}
	c.Assert(msg.String(), Equals, "user_4:john_doe|g\n")
}

func (s *MySuite) TestTimeMsgString(c *C) {
	msg := &TimeMsg{"user_create", "145"}
	c.Assert(msg.String(), Equals, "user_create:145|ms\n")
}

func (s *MySuite) TestCountMsgString(c *C) {
	msg := &CountMsg{"redeemed", "1"}
	c.Assert(msg.String(), Equals, "redeemed:1|c\n")
}

func (s *MySuite) TestSetMsgString(c *C) {
	msg := &SetMsg{"us_users", "john_doe"}
	c.Assert(msg.String(), Equals, "us_users:john_doe|s\n")
}

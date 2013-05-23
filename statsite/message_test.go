package statsite

import (
	. "launchpad.net/gocheck"
)

func (s *MySuite) TestMtypes(c *C) {
	c.Assert(MTypes, HasLen, 5)
	expected := []string{"kv", "g", "ms", "c", "s"}
	c.Assert(MTypes, DeepEquals, expected)
}

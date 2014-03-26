package arrays

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestContainsString(c *C) {
	strs := []string{"A", "b", "c"}
	c.Check(Contains("c", strs), Equals, true)
}

func (s *S) TestContainsStringNotFound(c *C) {
	strs := []string{"A", "b", "c"}
	c.Check(Contains("z", strs), Equals, false)
}

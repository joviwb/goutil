package cliutil

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	args := ParseLine(`./app top sub -a ddd --xx "msg"`)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = ParseLine(`./app top sub -a ddd --xx "abc
def"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "abc\ndef", args[6])

	args = ParseLine(`./app top sub -a ddd --xx "abc
def ghi"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "abc\ndef ghi", args[6])

	args = ParseLine(`./app top sub --msg "has multi words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has multi words", args[4])

	args = StringToOSArgs(`./app top sub --msg "has inner 'quote'"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote'", args[4])

	args = StringToOSArgs(`./app top sub --msg "'has' inner quote"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "'has' inner quote", args[4])

	args = StringToOSArgs(`./app top sub --msg "has inner 'quote' words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote' words", args[4])

	args = StringToOSArgs(`./app top sub --msg "has 'inner quote' words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has 'inner quote' words", args[4])

	args = StringToOSArgs(`./app top sub --msg "has 'inner quote words'"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has 'inner quote words'", args[4])

	args = StringToOSArgs(`./app top sub --msg "'has inner quote' words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "'has inner quote' words", args[4])

	args = StringToOSArgs(" ")
	assert.Len(t, args, 0)

	args = StringToOSArgs("./app")
	assert.Len(t, args, 1)
}

func TestParseLine_errLine(t *testing.T) {
	// exception line string.
	args := ParseLine(`./app top sub -a ddd --xx msg"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = ParseLine(`./app top sub -a ddd --xx "msg`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = StringToOSArgs(`./app top sub -a ddd --xx "msg text`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])

	args = StringToOSArgs(`./app top sub -a ddd --xx "msg "text"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])
}
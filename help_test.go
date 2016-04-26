package main_test

import (
	"bytes"
	"os"

	cli "github.com/heroku/cli"

	"github.com/lunixbochs/vtclean"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Help", func() {
	var stdout string
	var stderr string
	exit := 9999

	BeforeEach(func() {
		cli.Stdout = new(bytes.Buffer)
		cli.Stderr = new(bytes.Buffer)
		cli.ExitFn = func(code int) {
			if exit == 9999 {
				exit = code
			}
		}
	})

	AfterEach(func() {
		exit = 9999
		cli.Stdout = os.Stdout
		cli.Stderr = os.Stderr
	})

	JustBeforeEach(func() {
		stdout = vtclean.Clean(cli.Stdout.(*bytes.Buffer).String(), false)
		stderr = vtclean.Clean(cli.Stderr.(*bytes.Buffer).String(), false)
	})

	Context("heroku help", func() {
		BeforeEach(func() {
			cli.Start("heroku", "help")
		})

		It("exits with code 0", func() { Expect(exit).To(Equal(0)) })
		It("shows the help", func() {
			Expect(stdout).To(HavePrefix("Usage: heroku COMMAND [--app APP] [command-specific-options]"))
		})
	})

	Context("heroku hlp", func() {
		BeforeEach(func() {
			cli.Start("heroku", "hlp")
		})

		It("exits with code 2", func() { Expect(exit).To(Equal(2)) })
		It("has no stdout", func() { Expect(stdout).To(Equal("")) })
		It("shows invalid command message", func() {
			Expect(stderr).To(Equal(` !    hlp is not a heroku command.
 !    Perhaps you meant help.
 !    Run heroku help for a list of available commands.
`))
		})
	})

	Context("heroku help version", func() {
		BeforeEach(func() {
			cli.Start("heroku", "help", "version")
		})

		It("exits with code 0", func() { Expect(exit).To(Equal(0)) })
		It("shows help for version command", func() {
			Expect(stdout).To(HavePrefix("Usage: heroku version"))
		})
	})
})

package gexec_test

import (
	"os/exec"
	. "github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExitMatcher", func() {
	var command *exec.Cmd
	var session *Session

	BeforeEach(func() {
		var err error
		command = exec.Command(fireflyPath, "0")
		session, err = Start(command, nil, nil)
		Ω(err).ShouldNot(HaveOccurred())
	})

	Describe("when passed something that is not a session", func() {
		It("should error", func() {
			failures := interceptFailures(func() {
				Ω(command).Should(Exit())
			})

			Ω(failures[0]).Should(ContainSubstring("Exit must be passed a gexit session"))
		})
	})

	Context("with no exit code", func() {
		It("should say the right things when it fails", func() {
			Ω(session).ShouldNot(Exit())

			failures := interceptFailures(func() {
				Ω(session).Should(Exit())
			})

			Ω(failures[0]).Should(ContainSubstring("Expected process to exit.  It did not."))

			Eventually(session).Should(Exit())

			Ω(session).Should(Exit())

			failures = interceptFailures(func() {
				Ω(session).ShouldNot(Exit())
			})

			Ω(failures[0]).Should(ContainSubstring("Expected process not to exit.  It did."))
		})
	})

	Context("with an exit code", func() {
		It("should say the right things when it fails", func() {
			Ω(session).ShouldNot(Exit(0))
			Ω(session).ShouldNot(Exit(1))

			failures := interceptFailures(func() {
				Ω(session).Should(Exit(0))
			})

			Ω(failures[0]).Should(ContainSubstring("Expected process to exit.  It did not."))

			Eventually(session).Should(Exit(0))

			Ω(session).Should(Exit(0))

			failures = interceptFailures(func() {
				Ω(session).Should(Exit(1))
			})

			Ω(failures[0]).Should(ContainSubstring("to match exit code:"))

			Ω(session).ShouldNot(Exit(1))

			failures = interceptFailures(func() {
				Ω(session).ShouldNot(Exit(0))
			})

			Ω(failures[0]).Should(ContainSubstring("not to match exit code:"))
		})
	})
})

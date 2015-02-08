# stringer

This is a fork of the golang.org/x/tools
[stringer](https://github.com/golang/tools/tree/master/cmd/stringer)
utility.  It has been modified to generate code that more closely
matches the style a human would use.  The reason for these changes are
purely practical: they allow tools like
[golint](https://github.com/golang/lint) to run on an entire source
tree without reporting problems in generated code.

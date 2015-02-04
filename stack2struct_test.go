package stack2struct

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

type testElement struct {
	lineNumber  int
	packageName string
	fileName    string
	methodName  string
}

type testStack []*testElement

func (t *testStack) AddEntry(lineNumber int, packageName, fileName, methodName string) {
	*t = append(*t, &testElement{lineNumber, packageName, fileName, methodName})
}

func TestStack2Struct(t *testing.T) {
	Convey("#splitAtLastSlash", t, func() {
		testLine := "foo/bar/baz"
		left, right := splitAtLastSlash(testLine)
		So(left, ShouldEqual, "foo/bar")
		So(right, ShouldEqual, "baz")
	})

	Convey("#removeSpaceAndSuffix", t, func() {
		testLine := "foo:bar baz"
		result := removeSpaceAndSuffix(testLine)
		So(result, ShouldEqual, "foo:bar")
	})

	Convey("#Parse", t, func() {
		buf, _ := ioutil.ReadFile("_fixtures/stack_trace")

		stack := make(testStack, 0, 0)
		Parse(buf, &stack)

		expectedFirstEntry := testElement{13,
			"codevault.io/go_projects/raygun4go/stack2struct",
			"stack2struct_test.go",
			"func·001()"}

		So(len(stack), ShouldEqual, 5)
		firstEntry := stack[0]
		So(firstEntry.lineNumber, ShouldEqual, expectedFirstEntry.lineNumber)
		So(firstEntry.packageName, ShouldEqual, expectedFirstEntry.packageName)
		So(firstEntry.fileName, ShouldEqual, expectedFirstEntry.fileName)
		So(firstEntry.methodName, ShouldEqual, expectedFirstEntry.methodName)

		So(*stack[0], ShouldResemble, expectedFirstEntry)
	})

}

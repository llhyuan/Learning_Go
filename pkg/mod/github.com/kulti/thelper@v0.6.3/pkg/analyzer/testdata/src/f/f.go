// Code generated by generator. DO NOT EDIT.

package f

import (
	"context"
	"testing"
)

// -----------------------------
// Free functions
// -----------------------------

func nonTestHelper(f int) {}

func helperWithoutHelper(f *testing.F) {} // want "test helper function should start from f.Helper()"

func helperWithHelper(f *testing.F) {
	f.Helper()
}

func helperWithEmptyStringBeforeHelper(f *testing.F) {

	f.Helper()
}

func helperWithHelperAfterAssignment(f *testing.F) { // want "test helper function should start from f.Helper()"
	_ = 0
	f.Helper()
}

func helperWithHelperAfterOtherCall(f *testing.F) { // want "test helper function should start from f.Helper()"
	ff()
	f.Helper()
}

func helperWithHelperAfterOtherSelectionCall(f *testing.F) { // want "test helper function should start from f.Helper()"
	f.Fail()
	f.Helper()
}

func helperParamNotFirst(s string, i int, f *testing.F) { // want "parameter \\*testing.F should be the first or after context.Context"
	f.Helper()
}

func helperParamSecondWithoutContext(s string, f *testing.F, i int) { // want "parameter \\*testing.F should be the first or after context.Context"
	f.Helper()
}

func helperParamSecondWithContext(ctx context.Context, f *testing.F) {
	f.Helper()
}

func helperWithIncorrectName(o *testing.F) { // want "parameter \\*testing.F should have name f"
	o.Helper()
}

func helperWithAnonymousHelper(f *testing.F) {
	f.Helper()
	func(f *testing.F) {}(f) // want "test helper function should start from f.Helper()"
}

func helperWithNoName(_ *testing.F) {
}

// -----------------------------
// Methods of helper type
type helperType struct{}
// -----------------------------

func (h helperType) nonTestHelper(f int) {}

func (h helperType) helperWithoutHelper(f *testing.F) {} // want "test helper function should start from f.Helper()"

func (h helperType) helperWithHelper(f *testing.F) {
	f.Helper()
}

func (h helperType) helperWithEmptyStringBeforeHelper(f *testing.F) {

	f.Helper()
}

func (h helperType) helperWithHelperAfterAssignment(f *testing.F) { // want "test helper function should start from f.Helper()"
	_ = 0
	f.Helper()
}

func (h helperType) helperWithHelperAfterOtherCall(f *testing.F) { // want "test helper function should start from f.Helper()"
	ff()
	f.Helper()
}

func (h helperType) helperWithHelperAfterOtherSelectionCall(f *testing.F) { // want "test helper function should start from f.Helper()"
	f.Fail()
	f.Helper()
}

func (h helperType) helperParamNotFirst(s string, i int, f *testing.F) { // want "parameter \\*testing.F should be the first or after context.Context"
	f.Helper()
}

func (h helperType) helperParamSecondWithoutContext(s string, f *testing.F, i int) { // want "parameter \\*testing.F should be the first or after context.Context"
	f.Helper()
}

func (h helperType) helperParamSecondWithContext(ctx context.Context, f *testing.F) {
	f.Helper()
}

func (h helperType) helperWithIncorrectName(o *testing.F) { // want "parameter \\*testing.F should have name f"
	o.Helper()
}

func (h helperType) helperWithAnonymousHelper(f *testing.F) {
	f.Helper()
	func(f *testing.F) {}(f) // want "test helper function should start from f.Helper()"
}

func (h helperType) helperWithNoName(_ *testing.F) {
}

func ff() {}
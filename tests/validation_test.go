package tests_test

import (
	"encoding/json"
	"errors"
	"testing"

	testExclusiveMaximum "github.com/atombender/go-jsonschema/tests/data/validation/exclusiveMaximum"
	testExclusiveMinimum "github.com/atombender/go-jsonschema/tests/data/validation/exclusiveMinimum"
	testMaxLength "github.com/atombender/go-jsonschema/tests/data/validation/maxLength"
	testMaximum "github.com/atombender/go-jsonschema/tests/data/validation/maximum"
	testMinLength "github.com/atombender/go-jsonschema/tests/data/validation/minLength"
	testMinimum "github.com/atombender/go-jsonschema/tests/data/validation/minimum"
	testMultipleOf "github.com/atombender/go-jsonschema/tests/data/validation/multipleOf"
	testPattern "github.com/atombender/go-jsonschema/tests/data/validation/pattern"
	testPrimitiveDefs "github.com/atombender/go-jsonschema/tests/data/validation/primitive_defs"
	testRequiredFields "github.com/atombender/go-jsonschema/tests/data/validation/requiredFields"
	"github.com/atombender/go-jsonschema/tests/helpers"
)

func TestMaxStringLength(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   *testMaxLength.MaxLength
	}{
		{
			desc:    "no violations",
			data:    `{"myString": "hi"}`,
			wantErr: nil,
			asObj:   &testMaxLength.MaxLength{MyString: "hi"},
		},
		{
			desc:    "myString has the max allowed length",
			data:    `{"myString": "hello"}`,
			wantErr: nil,
			asObj:   &testMaxLength.MaxLength{MyString: "hello"},
		},
		{
			desc:    "myString too long",
			data:    `{"myString": "hello world"}`,
			wantErr: errors.New("field myString length: must be <= 5"),
			asObj:   &testMaxLength.MaxLength{MyString: "hello world"},
		},
		{
			desc:    "myString not present",
			data:    `{}`,
			wantErr: errors.New("field myString in MaxLength: required"),
		},
		{
			desc:    "myNullableString too long",
			data:    `{"myString": "hi","myNullableString": "hello world"}`,
			wantErr: errors.New("field myNullableString length: must be <= 10"),
			asObj:   &testMaxLength.MaxLength{MyString: "hi", MyNullableString: pointer("hello world")},
		},
		{
			desc:    "myString and myNullableString too long",
			data:    `{"myString": "hello","myNullableString": "hello world"}`,
			wantErr: errors.New("field myNullableString length: must be <= 10"),
			asObj:   &testMaxLength.MaxLength{MyString: "hello", MyNullableString: pointer("hello world")},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testMaxLength.MaxLength{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)

			if tC.asObj != nil {
				err = tC.asObj.Verify()

				helpers.CheckError(t, tC.wantErr, err)
			}
		})
	}
}

func TestMinStringLength(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   *testMinLength.MinLength
	}{
		{
			desc:    "no violations",
			data:    `{"myString": "hello"}`,
			wantErr: nil,
			asObj:   &testMinLength.MinLength{MyString: "hello"},
		},
		{
			desc:    "myString too short",
			data:    `{"myString": "hi"}`,
			wantErr: errors.New("field myString length: must be >= 5"),
			asObj:   &testMinLength.MinLength{MyString: "hi"},
		},
		{
			desc:    "myString not present",
			data:    `{}`,
			wantErr: errors.New("field myString in MinLength: required"),
		},
		{
			desc:    "myNullableString too short",
			data:    `{"myString": "hello","myNullableString": "hi"}`,
			wantErr: errors.New("field myNullableString length: must be >= 10"),
			asObj:   &testMinLength.MinLength{MyString: "hello", MyNullableString: pointer("hi")},
		},
		{
			desc:    "myString and myNullableString too short",
			data:    `{"myString": "hi","myNullableString": "hello"}`,
			wantErr: errors.New("field myNullableString length: must be >= 10"),
			asObj:   &testMinLength.MinLength{MyString: "hi", MyNullableString: pointer("hello")},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testMinLength.MinLength{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)

			if tC.asObj != nil {
				err = tC.asObj.Verify()

				helpers.CheckError(t, tC.wantErr, err)
			}
		})
	}
}

func TestRequiredFields(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
	}{
		{
			desc:    "object without required property fails validation",
			data:    `{}`,
			wantErr: errors.New("field myNullableObject in RequiredNullable: required"),
		},
		{
			desc:    "required properties may be null",
			data:    `{ "myNullableObject": null, "myNullableStringArray": null, "myNullableString": null }`,
			wantErr: nil,
		},
		{
			desc:    "required properties may have a non-null value",
			data:    `{ "myNullableObject": { "myNestedProp": "世界" }, "myNullableStringArray": ["hello"], "myNullableString": "world" }`,
			wantErr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testRequiredFields.RequiredNullable{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestPattern(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   *testPattern.Pattern
	}{
		{
			desc:  "no violations",
			data:  `{"myString": "0x12345abcde."}`,
			asObj: &testPattern.Pattern{MyString: "0x12345abcde."},
		},
		{
			desc:    "myString does not match pattern",
			data:    `{"myString": "0x123456"}`,
			wantErr: errors.New("field ^0x[0-9a-f]{10}\\.$ pattern match: must match MyString"),
			asObj:   &testPattern.Pattern{MyString: "0x123456"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testPattern.Pattern{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)

			if tC.asObj != nil {
				err = tC.asObj.Verify()

				helpers.CheckError(t, tC.wantErr, err)
			}
		})
	}
}

func TestPrimitiveDefs(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   testPrimitiveDefs.PrimitiveDefs
	}{
		{
			desc:  "no violations",
			data:  `{"myString": "hello"}`,
			asObj: testPrimitiveDefs.PrimitiveDefs{MyString: "hello"},
		},
		{
			desc:    "myString too short",
			data:    `{"myString": "hi"}`,
			wantErr: errors.New("field  length: must be >= 5"),
			asObj:   testPrimitiveDefs.PrimitiveDefs{MyString: "hi"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			prim := testPrimitiveDefs.PrimitiveDefs{}

			err := json.Unmarshal([]byte(tC.data), &prim)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestMultipleOf(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   testMultipleOf.MultipleOf
	}{
		{
			desc:  "no violations",
			data:  `{"myInteger": 10, "myNumber": 2.4}`,
			asObj: testMultipleOf.MultipleOf{MyInteger: 10, MyNumber: 2.4},
		},
		{
			desc:    "myInt not a multiple of 2",
			data:    `{"myInteger": 11, "myNumber": 2.4}`,
			wantErr: errors.New("field myInteger: must be a multiple of 2"),
			asObj:   testMultipleOf.MultipleOf{MyInteger: 11, MyNumber: 2.4},
		},
		{
			desc:    "myNumber not a multiple of 1.2",
			data:    `{"myInteger": 10, "myNumber": 2.5}`,
			wantErr: errors.New("field myNumber: must be a multiple of 1.2"),
			asObj:   testMultipleOf.MultipleOf{MyInteger: 10, MyNumber: 2.5},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			mo := testMultipleOf.MultipleOf{}

			err := json.Unmarshal([]byte(tC.data), &mo)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestMaximum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   testMaximum.Maximum
	}{
		{
			desc:  "no violations",
			data:  `{"myInteger": 1, "myNumber": 1.0}`,
			asObj: testMaximum.Maximum{MyInteger: 1, MyNumber: 1.0},
		},
		{
			desc:    "myInt exceeds maximum of 2",
			data:    `{"myInteger": 3, "myNumber": 1.0}`,
			wantErr: errors.New("field myInteger: must be <= 2"),
			asObj:   testMaximum.Maximum{MyInteger: 3, MyNumber: 1.0},
		},
		{
			desc:    "myNumber exceeds maximum of 1.2",
			data:    `{"myInteger": 1, "myNumber": 1.3}`,
			wantErr: errors.New("field myNumber: must be <= 1.2"),
			asObj:   testMaximum.Maximum{MyInteger: 1, MyNumber: 1.3},
		},
		{
			desc:  "boundary case - exactly at maximum",
			data:  `{"myInteger": 2, "myNumber": 1.2}`,
			asObj: testMaximum.Maximum{MyInteger: 2, MyNumber: 1.2},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			mo := testMaximum.Maximum{}

			err := json.Unmarshal([]byte(tC.data), &mo)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestMinimum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
		asObj   testMinimum.Minimum
	}{
		{
			desc:  "no violations",
			data:  `{"myInteger": 3, "myNumber": 1.5}`,
			asObj: testMinimum.Minimum{MyInteger: 3, MyNumber: 1.5},
		},
		{
			desc:    "myInt below minimum of 2",
			data:    `{"myInteger": 1, "myNumber": 1.5}`,
			wantErr: errors.New("field myInteger: must be >= 2"),
			asObj:   testMinimum.Minimum{MyInteger: 1, MyNumber: 1.5},
		},
		{
			desc:    "myNumber below minimum of 1.2",
			data:    `{"myInteger": 3, "myNumber": 1.1}`,
			wantErr: errors.New("field myNumber: must be >= 1.2"),
			asObj:   testMinimum.Minimum{MyInteger: 3, MyNumber: 1.1},
		},
		{
			desc:  "boundary case - exactly at minimum",
			data:  `{"myInteger": 2, "myNumber": 1.2}`,
			asObj: testMinimum.Minimum{MyInteger: 2, MyNumber: 1.2},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			mo := testMinimum.Minimum{}

			err := json.Unmarshal([]byte(tC.data), &mo)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestExclusiveMaximum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		data     string
		wantErr  error
		asObj    testExclusiveMaximum.ExclusiveMaximum
		asObjOld testExclusiveMaximum.ExclusiveMaximumOld
	}{
		{
			desc:     "no violations",
			data:     `{"myInteger": 1, "myNumber": 1.1}`,
			asObj:    testExclusiveMaximum.ExclusiveMaximum{MyInteger: 1, MyNumber: 1.1},
			asObjOld: testExclusiveMaximum.ExclusiveMaximumOld{MyInteger: 1, MyNumber: 1.1},
		},
		{
			desc:     "myInt exceeds exclusive maximum of 2",
			data:     `{"myInteger": 2, "myNumber": 1.1}`,
			wantErr:  errors.New("field myInteger: must be < 2"),
			asObj:    testExclusiveMaximum.ExclusiveMaximum{MyInteger: 2, MyNumber: 1.1},
			asObjOld: testExclusiveMaximum.ExclusiveMaximumOld{MyInteger: 2, MyNumber: 1.1},
		},
		{
			desc:     "myNumber exceeds exclusive maximum of 1.2",
			data:     `{"myInteger": 1, "myNumber": 1.2}`,
			wantErr:  errors.New("field myNumber: must be < 1.2"),
			asObj:    testExclusiveMaximum.ExclusiveMaximum{MyInteger: 1, MyNumber: 1.2},
			asObjOld: testExclusiveMaximum.ExclusiveMaximumOld{MyInteger: 1, MyNumber: 1.2},
		},
		{
			desc:     "boundary case - just below exclusive maximum",
			data:     `{"myInteger": 1, "myNumber": 1.19}`,
			asObj:    testExclusiveMaximum.ExclusiveMaximum{MyInteger: 1, MyNumber: 1.19},
			asObjOld: testExclusiveMaximum.ExclusiveMaximumOld{MyInteger: 1, MyNumber: 1.19},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			mo := testExclusiveMaximum.ExclusiveMaximum{}

			err := json.Unmarshal([]byte(tC.data), &mo)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)

			mo2 := testExclusiveMaximum.ExclusiveMaximumOld{}

			err = json.Unmarshal([]byte(tC.data), &mo2)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObjOld.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestExclusiveMinimum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		data     string
		wantErr  error
		asObj    testExclusiveMinimum.ExclusiveMinimum
		asObjOld testExclusiveMinimum.ExclusiveMinimumOld
	}{
		{
			desc:     "no violations",
			data:     `{"myInteger": 3, "myNumber": 1.3}`,
			asObj:    testExclusiveMinimum.ExclusiveMinimum{MyInteger: 3, MyNumber: 1.3},
			asObjOld: testExclusiveMinimum.ExclusiveMinimumOld{MyInteger: 3, MyNumber: 1.3},
		},
		{
			desc:     "myInt below exclusive minimum of 2",
			data:     `{"myInteger": 2, "myNumber": 1.3}`,
			wantErr:  errors.New("field myInteger: must be > 2"),
			asObj:    testExclusiveMinimum.ExclusiveMinimum{MyInteger: 2, MyNumber: 1.3},
			asObjOld: testExclusiveMinimum.ExclusiveMinimumOld{MyInteger: 2, MyNumber: 1.3},
		},
		{
			desc:     "myNumber below exclusive minimum of 1.2",
			data:     `{"myInteger": 3, "myNumber": 1.2}`,
			wantErr:  errors.New("field myNumber: must be > 1.2"),
			asObj:    testExclusiveMinimum.ExclusiveMinimum{MyInteger: 3, MyNumber: 1.2},
			asObjOld: testExclusiveMinimum.ExclusiveMinimumOld{MyInteger: 3, MyNumber: 1.2},
		},
		{
			desc:     "boundary case - just above exclusive minimum",
			data:     `{"myInteger": 3, "myNumber": 1.21}`,
			asObj:    testExclusiveMinimum.ExclusiveMinimum{MyInteger: 3, MyNumber: 1.21},
			asObjOld: testExclusiveMinimum.ExclusiveMinimumOld{MyInteger: 3, MyNumber: 1.21},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			mo := testExclusiveMinimum.ExclusiveMinimum{}

			err := json.Unmarshal([]byte(tC.data), &mo)
			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObj.Verify()

			helpers.CheckError(t, tC.wantErr, err)

			mo2 := testExclusiveMinimum.ExclusiveMinimumOld{}

			err = json.Unmarshal([]byte(tC.data), &mo2)

			helpers.CheckError(t, tC.wantErr, err)

			err = tC.asObjOld.Verify()

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func pointer[T any](v T) *T {
	return &v
}

package rapidjson

import (
	"testing"

	"github.com/stretchr/testify/assert" // Assertion package
)

var (
	testJSON1 = `{
        "member1" : 12345,
        "member2" : [1, 2, 3, 4, 5],
        "member3" : {
            "sub1" : 1.234,
            "sub2" : true,
            "sub3" : null
        },
        "member4" : "rapidjson is awesome!"
    }`
)

func TestParse(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()
}

func TestOutput(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()

	expected := `{"member1":12345,"member2":[1,2,3,4,5],"member3":{"sub1":1.234,"sub2":true,"sub3":null},"member4":"rapidjson is awesome!"}`
	assert.Equal(t, expected, json.String())
}

func TestGetters(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()

	ct := json.GetContainer()

	member1, err := ct.GetMemberOrNil("member1").GetInt()
	assert.Nil(t, err, "should not error on member1")
	assert.Equal(t, 12345, member1)

	member2, err := ct.GetMemberOrNil("member2").GetIntArray()
	assert.Nil(t, err, "should not error on member2")
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, member2)
	arraySize, err := ct.GetMemberOrNil("member2").GetArraySize()
	assert.Nil(t, err, "should not error on array size")
	assert.Equal(t, 5, arraySize)

	member3, err := ct.GetMember("member3")
	assert.Nil(t, err, "should not error on member3")
	sub1, err := member3.GetMemberOrNil("sub1").GetFloat()
	assert.Nil(t, err, "should not error on sub1")
	assert.Equal(t, 1.234, sub1)
	sub2, err := member3.GetMemberOrNil("sub2").GetBool()
	assert.Nil(t, err, "should not error on sub2")
	assert.Equal(t, true, sub2)
	sub3, err := member3.GetMember("sub3")
	assert.Nil(t, err, "should not error on sub3")
	assert.Equal(t, TypeNull, sub3.GetType())

	member4, err := ct.GetMemberOrNil("member4").GetString()
	assert.Nil(t, err, "should not error on member4")
	assert.Equal(t, "rapidjson is awesome!", member4)
}

func TestGetValue(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()

	ct := json.GetContainer()

	member1, err := ct.GetMemberOrNil("member1").GetValue()
	assert.Nil(t, err, "should not error on member1")
	assert.Equal(t, int64(12345), member1)

	member3, err := ct.GetMember("member3")
	assert.Nil(t, err, "should not error on member3")
	sub1, err := member3.GetMemberOrNil("sub1").GetValue()
	assert.Nil(t, err, "should not error on sub1")
	assert.Equal(t, 1.234, sub1)
	sub2, err := member3.GetMemberOrNil("sub2").GetValue()
	assert.Nil(t, err, "should not error on sub2")
	assert.Equal(t, true, sub2)

	member4, err := ct.GetMemberOrNil("member4").GetValue()
	assert.Nil(t, err, "should not error on member4")
	assert.Equal(t, "rapidjson is awesome!", member4)
}

func TestSetters(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()

	ct := json.GetContainer()

	err = ct.GetMemberOrNil("member1").SetValue(54321)
	assert.Nil(t, err, "should not error on set member1")

	err = ct.GetMemberOrNil("member2").ArrayAppend(6)
	assert.Nil(t, err, "should not error on set member2")

	err = ct.GetMemberOrNil("member3").AddValue("sub4", false)
	assert.Nil(t, err, "should not error on add member3.sub4")

	expected := `{"member1":54321,"member2":[1,2,3,4,5,6],"member3":{"sub1":1.234,"sub2":true,"sub3":null,"sub4":false},"member4":"rapidjson is awesome!"}`
	assert.Equal(t, expected, json.String())
}

func TestRemoves(t *testing.T) {
	json, err := NewParsedStringJson(testJSON1)
	assert.Nil(t, err, "should not error on parsing")
	defer json.Free()

	ct := json.GetContainer()

	err = ct.RemoveMember("member1")
	assert.Nil(t, err, "should not error on remove member1")

	err = ct.GetMemberOrNil("member2").ArrayRemove(2)
	assert.Nil(t, err, "should not error on remove array member2 element 2")

	err = ct.RemoveMemberAtPath("member3.sub3")
	assert.Nil(t, err, "should not error on remove member3.sub3")

	expected := `{"member4":"rapidjson is awesome!","member2":[1,2,4,5],"member3":{"sub1":1.234,"sub2":true}}`
	assert.Equal(t, expected, json.String())
}
package main

//
//import (
//	"github.com/bitly/go-simplejson"
//	"github.com/gookit/color"
//	"github.com/stretchr/testify/assert"
//	"reflect"
//	"testing"
//)
//
//func main() {
//
//	js, _ := simplejson.NewJson([]byte(`{
//		"test": {
//			"string_array": ["asdf", "ghjk", "zxcv"],
//			"string_array_null": ["abc", null, "efg"],
//			"array": [1, "2", 3],
//			"arraywithsubs": [{"subkeyone": 1},
//			{"subkeytwo": 2, "subkeythree": 3}],
//			"int": 10,
//			"float": 5.150,
//			"string": "simplejson",
//			"bool": true,
//			"sub_obj": {"a": 1}
//		}
//	}`))
//
//	str := js.Get("test").Get("sub_obj")
//
//	color.Red.Println(reflect.TypeOf(str))
//
//}
//func TestSimplejson(t *testing.T) {
//	var ok bool
//	var err error
//
//	js, err := simplejson.NewJson([]byte(`{
//		"test": {
//			"string_array": ["asdf", "ghjk", "zxcv"],
//			"string_array_null": ["abc", null, "efg"],
//			"array": [1, "2", 3],
//			"arraywithsubs": [{"subkeyone": 1},
//			{"subkeytwo": 2, "subkeythree": 3}],
//			"int": 10,
//			"float": 5.150,
//			"string": "simplejson",
//			"bool": true,
//			"sub_obj": {"a": 1}
//		}
//	}`))
//
//	assert.NotEqual(t, nil, js)
//	assert.Equal(t, nil, err)
//
//	_, ok = js.CheckGet("test")
//	assert.Equal(t, true, ok)
//
//	_, ok = js.CheckGet("missing_key")
//	assert.Equal(t, false, ok)
//
//	aws := js.Get("test").Get("arraywithsubs")
//	assert.NotEqual(t, nil, aws)
//	var awsval int
//	awsval, _ = aws.GetIndex(0).Get("subkeyone").Int()
//	assert.Equal(t, 1, awsval)
//	awsval, _ = aws.GetIndex(1).Get("subkeytwo").Int()
//	assert.Equal(t, 2, awsval)
//	awsval, _ = aws.GetIndex(1).Get("subkeythree").Int()
//	assert.Equal(t, 3, awsval)
//
//	i, _ := js.Get("test").Get("int").Int()
//	assert.Equal(t, 10, i)
//
//	f, _ := js.Get("test").Get("float").Float64()
//	assert.Equal(t, 5.150, f)
//
//	s, _ := js.Get("test").Get("string").String()
//	assert.Equal(t, "simplejson", s)
//
//	b, _ := js.Get("test").Get("bool").Bool()
//	assert.Equal(t, true, b)
//
//	mi := js.Get("test").Get("int").MustInt()
//	assert.Equal(t, 10, mi)
//
//	mi2 := js.Get("test").Get("missing_int").MustInt(5150)
//	assert.Equal(t, 5150, mi2)
//
//	ms := js.Get("test").Get("string").MustString()
//	assert.Equal(t, "simplejson", ms)
//
//	ms2 := js.Get("test").Get("missing_string").MustString("fyea")
//	assert.Equal(t, "fyea", ms2)
//
//	ma2 := js.Get("test").Get("missing_array").MustArray([]interface{}{"1", 2, "3"})
//	assert.Equal(t, ma2, []interface{}{"1", 2, "3"})
//
//	msa := js.Get("test").Get("string_array").MustStringArray()
//	assert.Equal(t, msa[0], "asdf")
//	assert.Equal(t, msa[1], "ghjk")
//	assert.Equal(t, msa[2], "zxcv")
//
//	msa2 := js.Get("test").Get("string_array").MustStringArray([]string{"1", "2", "3"})
//	assert.Equal(t, msa2[0], "asdf")
//	assert.Equal(t, msa2[1], "ghjk")
//	assert.Equal(t, msa2[2], "zxcv")
//
//	msa3 := js.Get("test").Get("missing_array").MustStringArray([]string{"1", "2", "3"})
//	assert.Equal(t, msa3, []string{"1", "2", "3"})
//
//	mm2 := js.Get("test").Get("missing_map").MustMap(map[string]interface{}{"found": false})
//	assert.Equal(t, mm2, map[string]interface{}{"found": false})
//
//	strs, err := js.Get("test").Get("string_array").StringArray()
//	assert.Equal(t, err, nil)
//	assert.Equal(t, strs[0], "asdf")
//	assert.Equal(t, strs[1], "ghjk")
//	assert.Equal(t, strs[2], "zxcv")
//
//	strs2, err := js.Get("test").Get("string_array_null").StringArray()
//	assert.Equal(t, err, nil)
//	assert.Equal(t, strs2[0], "abc")
//	assert.Equal(t, strs2[1], "")
//	assert.Equal(t, strs2[2], "efg")
//
//	gp, _ := js.GetPath("test", "string").String()
//	assert.Equal(t, "simplejson", gp)
//
//	gp2, _ := js.GetPath("test", "int").Int()
//	assert.Equal(t, 10, gp2)
//
//	assert.Equal(t, js.Get("test").Get("bool").MustBool(), true)
//
//	js.Set("float2", 300.0)
//	assert.Equal(t, js.Get("float2").MustFloat64(), 300.0)
//
//	js.Set("test2", "setTest")
//	assert.Equal(t, "setTest", js.Get("test2").MustString())
//
//	js.Del("test2")
//	assert.NotEqual(t, "setTest", js.Get("test2").MustString())
//
//	js.Get("test").Get("sub_obj").Set("a", 2)
//	assert.Equal(t, 2, js.Get("test").Get("sub_obj").Get("a").MustInt())
//
//	js.GetPath("test", "sub_obj").Set("a", 3)
//	assert.Equal(t, 3, js.GetPath("test", "sub_obj", "a").MustInt())
//}

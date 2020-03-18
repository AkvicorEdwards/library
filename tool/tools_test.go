package tools

import (
	"fmt"
	"testing"
)

var tests = struct {
	StringSplitByParagraph []struct {
		a string
		b []string
	}
	StringDeleteVoidAndFormat []struct{ a, b string }
	StringSplitBySpace        []struct {
		a string
		b []string
	}
	JsonWithFalseEscapeHTML []struct {
		a []string
		b string
	}
	Max []struct{ a, b, c int }
	Min []struct{ a, b, c int }
}{
	[]struct {
		a string
		b []string
	}{
		{
			"aaa",
			[]string{
				"aaa",
			},
		},
		{
			"aaa bbb ccc",
			[]string{
				"aaa bbb ccc",
			},
		},
		{
			"aaa\nbbb\nvcc",
			[]string{
				"aaa", "bbb", "vcc",
			},
		},
		{
			`
aksjdf aklsjdf lkw alks kvalkw kl
  lkasjfd
	aksdfj klasjfl ;w,

阿科技
	蓝色空间阿飞
联赛离开东京 了思考的解放  阿拉科技
我就哦i‘】`,
			[]string{
				"aksjdf aklsjdf lkw alks kvalkw kl",
				"  lkasjfd",
				"	aksdfj klasjfl ;w,",
				"阿科技",
				"	蓝色空间阿飞",
				"联赛离开东京 了思考的解放  阿拉科技",
				"我就哦i‘】",
			},
		},
	},
	[]struct{ a, b string }{
		{"abcdef", "abcdef"},
		{"akl sjdfliqj", "aklsjdfliqj"},
		{"akl sjd\nfl\n\riqj", "aklsjdfliqj"},
		{`
alskdjf

aksfj aklsjdf lkwj kalj fk

	af
lks
skjlwia jkasdf卡江苏大丰 卢卡角色的反抗军
	卢卡斯就地方

    阿里克斯郡地方
  拉科技斯蒂芬
ljasdf 
s
`, "alskdjfaksfjaklsjdflkwjkaljfkaflksskjlwiajkasdf卡江苏大" +
			"丰卢卡角色的反抗军卢卡斯就地方阿里克斯郡地方拉科技斯蒂芬ljasdfs"},
	},
	[]struct {
		a string
		b []string
	}{
		{
			"aaa",
			[]string{
				"aaa",
			},
		},
		{
			"aaa bbb ccc",
			[]string{
				"aaa", "bbb", "ccc",
			},
		},
		{
			"aaa\nbbb\nvcc",
			[]string{
				"aaa", "bbb", "vcc",
			},
		},
		{
			`
aksjdf aklsjdf lkw alks kvalkw kl
  lkasjfd
	aksdfj klasjfl ;w,

阿科技
	蓝色空间阿飞
联赛离开东京 了思考的解放  阿拉科技
我就哦i‘】`,
			[]string{
				"aksjdf", "aklsjdf", "lkw", "alks", "kvalkw", "kl", "lkasjfd", "aksdfj", "klasjfl", ";w,", "阿科技",
				"蓝色空间阿飞", "联赛离开东京", "了思考的解放", "阿拉科技", "我就哦i‘】",
			},
		},
	},
	[]struct {
		a []string
		b string
	}{
		{
			[]string{
				"a",
				"b",
				"c",
			},
		`["a","b","c"]`,
		},
		{
			[]string{"abc"},
			`["abc"]`,
		},
		{
			[]string{""},
			`[""]`,
		},
		{
			[]string{},
			`[]`,
		},
	},
	[]struct{ a, b, c int }{
		{1, 2, 2},
		{1, 1, 1},
		{4, 2, 4},
	},
	[]struct{ a, b, c int }{
		{1, 2, 1},
		{1, 1, 1},
		{4, 2, 2},
	},
}

func TestStringSplitBySpace(t *testing.T) {
	for _, test := range tests.StringSplitBySpace {
		tt := StringSplitBySpace(test.a)
		cp := func() bool {
			if len(tt) != len(test.b) {
				fmt.Println(len(tt), len(test.b))
				return false
			}
			for i := range tt {
				if tt[i] != test.b[i] {
					fmt.Println(tt[i])
					fmt.Println(test.b[i])
					return false
				}
			}
			return true
		}
		if !cp() {
			t.Errorf("STD: %s \nGot: %s", test.b, StringSplitBySpace(test.a))
		}

	}
}

func TestStringSplitByParagraph(t *testing.T) {
	for _, test := range tests.StringSplitByParagraph {
		tt := StringSplitByParagraph(test.a)
		cp := func() bool {
			if len(tt) != len(test.b) {
				fmt.Println(len(tt), len(test.b))
				return false
			}
			for i := range tt {
				if tt[i] != test.b[i] {
					fmt.Println(tt[i])
					fmt.Println(test.b[i])
					return false
				}
			}
			return true
		}
		if !cp() {
			t.Errorf("STD: %s \nGot: %s", test.b, StringSplitByParagraph(test.a))
		}

	}
}

func TestStringDeleteVoidAndFormat(t *testing.T) {
	for _, test := range tests.StringDeleteVoidAndFormat {
		if StringDeleteVoidAndFormat(test.a) != test.b {
			t.Errorf("STD: %s \nGot: %s", test.b, StringDeleteVoidAndFormat(test.a))
		}
	}
}

func TestJsonWithFalseEscapeHTML(t *testing.T) {
	for _, test := range tests.JsonWithFalseEscapeHTML {
		if JsonWithFalseEscapeHTML(test.a) != test.b {
			t.Errorf("STD: %v \nGot: %v", test.b, JsonWithFalseEscapeHTML(test.a))
		}
	}
}

func TestMax(t *testing.T) {
	for _, test := range tests.Max {
		if Max(test.a, test.b) != test.c {
			t.Errorf("STD: %d Got: %d", test.c, Max(test.a, test.b))
		}
	}
}

func TestMin(t *testing.T) {
	for _, test := range tests.Min {
		if Min(test.a, test.b) != test.c {
			t.Errorf("STD: %d Got: %d", test.c, Min(test.a, test.b))
		}
	}
}

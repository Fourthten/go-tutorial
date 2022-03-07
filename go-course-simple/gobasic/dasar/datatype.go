package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// str to int
	var str = "123"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	// int to str
	var _num = 123
	var _str = strconv.Itoa(_num)
	fmt.Println(_str)

	// str to num
	var str_ = "123"
	var num_, err_ = strconv.ParseInt(str_, 10, 64)
	if err_ == nil {
		fmt.Println(num_)
	}
	var strs = "101"
	var nums, errs = strconv.ParseInt(strs, 2, 8)
	if errs == nil {
		fmt.Println(nums)
	}

	// int64 to str
	var _nums = int64(24)
	var _strs = strconv.FormatInt(_nums, 8)
	fmt.Println(_strs)

	// str to dec
	var strs_ = "25.15"
	var nums_, errs_ = strconv.ParseFloat(strs_, 32)
	if errs_ == nil {
		fmt.Println(nums_)
	}

	// float64 to str
	var num2 = float64(25.15)
	var str2 = strconv.FormatFloat(num2, 'f', 6, 64)
	fmt.Println(str2)

	// str to bool
	var str3 = "true"
	var bul, err3 = strconv.ParseBool(str3)
	if err3 == nil {
		fmt.Println(bul)
	}

	// bool to str
	var bul2 = true
	var str4 = strconv.FormatBool(bul2)
	fmt.Println(str4)

	// casting
	var a float64 = float64(24)
	fmt.Println(a)
	var b int32 = int32(24.00)
	fmt.Println(b)

	var text1 = "hello"
	var _b = []byte(text1)
	fmt.Printf("%d %d %d \n", _b[0], _b[1], _b[2])

	var byte1 = []byte{100, 70, 80}
	var s = string(byte1)
	fmt.Printf("%s \n", s)

	var c int64 = int64('h')
	fmt.Println(c)
	var d string = string(100)
	fmt.Println(d)

	var data = map[string]interface{}{
		"name":    "setia",
		"grade":   4,
		"height":  160.5,
		"isMale":  true,
		"hobbies": []string{"swim", "play"},
	}
	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}

	var isExists = strings.Contains("setia", "set")
	fmt.Println(isExists)

	var isPrefix1 = strings.HasPrefix("setia", "se")
	fmt.Println(isPrefix1)
	var isPrefix2 = strings.HasPrefix("setia", "ia")
	fmt.Println(isPrefix2)

	var isSuffix1 = strings.HasSuffix("setia", "se")
	fmt.Println(isSuffix1)
	var isSuffix2 = strings.HasSuffix("setia", "ia")
	fmt.Println(isSuffix2)

	var howMany = strings.Count("setia", "se")
	fmt.Println(howMany)

	var index1 = strings.Index("setia", "t")
	fmt.Println(index1)
	var index2 = strings.Index("setias", "s")
	fmt.Println(index2)

	var text = "banana"
	var find = "a"
	var replaceWith = "u"
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)
	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)
	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3)

	var str5 = strings.Repeat("na", 4)
	fmt.Println(str5)

	var string1 = strings.Split("ajung setia", " ")
	fmt.Println(string1)
	var string2 = strings.Split("ajung", "")
	fmt.Println(string2)

	var data1 = []string{"banana", "papaya", "tomato"}
	var str6 = strings.Join(data1, "-")
	fmt.Println(str6)

	var str7 = strings.ToLower("SeTia")
	fmt.Println(str7)

	var str8 = strings.ToUpper("ajUnG")
	fmt.Println(str8)

	var _text = "grape orange mango"
	var regex, err4 = regexp.Compile(`[a-z]+`)
	if err4 != nil {
		fmt.Println(err4.Error())
	}
	var res1 = regex.FindAllString(_text, 2)
	fmt.Printf("%#v \n", res1)
	var res2 = regex.FindAllString(_text, -1)
	fmt.Printf("%#v \n", res2)

	var isMatch = regex.MatchString(_text)
	fmt.Println(isMatch)
	var str1_ = regex.FindString(_text)
	fmt.Println(str1_)
	var idx = regex.FindStringIndex(_text)
	fmt.Println(idx)
	var str2_ = _text[0:4]
	fmt.Println(str2_)
	var str3_ = regex.FindAllString(_text, -1)
	fmt.Println(str3_)
	var str4_ = regex.FindAllString(_text, 1)
	fmt.Println(str4_)
	var str5_ = regex.ReplaceAllString(_text, "potato")
	fmt.Println(str5_)
	var str6_ = regex.ReplaceAllStringFunc(_text, func(each string) string {
		if each == "orange" {
			return "lemon"
		}
		return each
	})
	fmt.Println(str6_)
	var regex_, _ = regexp.Compile(`[a-m]+`)
	var str7_ = regex_.Split(_text, -1)
	fmt.Printf("%#v \n", str7_)
}

// Format Eksponen		Penjelasan
// b					-ddddp±ddd, a, eksponen biner (basis 2)
// e					-d.dddde±dd, a, eksponen desimal (basis 10)
// E					-d.ddddE±dd, a, eksponen desimal (basis 10)
// f					-ddd.dddd, tanpa eksponen
// g					Akan menggunakan format eksponen e untuk eksponen besar dan f untuk selainnya
// G					Akan menggunakan format eksponen E untuk eksponen besar dan f untuk selainnya

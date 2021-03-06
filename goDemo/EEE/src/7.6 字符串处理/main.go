package main

import (
	"fmt"
)

//字符串在我们平常的web开发中经常用到 包括用户的输入 数据库读取的数据等
//我们经常需要对字符串 进行分割 转换等操作 本小节 将通过Go标准库中的strings和strconv两个包中的函数来讲解如何进行有效快速的操作

//字符串操作
//下面这些函数来自于strings包  这里介绍一些我平常经常用到的函数 更详细的请参考官方的文档
func Contatins(s, substr string) bool //字符串s中是否包含substr 返回bool值
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true
	
func Join(a[]string,sep string)string//字符串链接  把slice a通过sep链接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, "-"))//foo-bar-baz
	
func Index(s,sep string)int//在字符串s中查找sep所在的位置 返回位置值 找不到返回-1
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1	

func Repeat(s string,count int) string//重复s字符串 count次 最后返回重复的字符串
	fmt.Println("ba"+strings.Repeat("na",2))//output: banana
	
	
func Replace(s,old,new string,n int)string
//在字符串 中把old字符串替换为new字符串 n表示替换的次数 小于0表示全部替换
	fmt.Println(strings.Replace("olik olik olik", "k", "ky", 2))      //output : oliky oliky olik
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo


func Split(s,sep string)[]string
//把字符串按照sep分割 返回slice
fmt.Printf("%q\n", strings.Split("a,b,c", ","))
fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
fmt.Printf("%q\n", strings.Split(" xyz ", ""))
fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
//Output:["a" "b" "c"]
//["" "man " "plan " "canal panama"]
//[" " "x" "y" "z" " "]
//[""]


func Trim(s string,cutset string)string//在s字符串中去除cutset指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	Output:["Achtung"]
	
func Field(s string)[]string
//去除s字符串的空格符 并且按照空格分割返回slice
fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
//Output:Fields are: ["foo" "bar" "baz"]	


//字符串转换
//字符串转化的函数在strconv中 如下也只是列表一些常用的
//append系列函数将整数 等 转换为字符串后 添加到现有的字节数组中
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}
//	4567false"abcdefg"'单'


//Format系列函数把其他类型的转换为字符串
	package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}
//false 123.23 1234 12345 1023

//parse系列函数把字符串转换为其他类型
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}
	c, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	d, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a, b, c, d)
}

	//false 123.23 1234 12345
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
		
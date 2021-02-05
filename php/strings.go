package php

import (
	"bytes"
	"html"
	"strings"
	"unicode/utf8"
)

// Explode - Split a string by string
func Explode(s, sep string) []string {
	return strings.Split(s, sep)
}

// HtmlSpecialChars - Convert special characters to HTML entities
func HtmlSpecialChars(s string) string {
	return html.EscapeString(s)
}

// HtmlSpecialCharsDecode - Convert special HTML entities back to characters
func HtmlSpecialCharsDecode(s string) string {
	return html.UnescapeString(s)
}

// Implode - Join array elements with a string
func Implode(a []string, sep string) string {
	return strings.Join(a, sep)
}

// Join - Alias of implode()
func Join(a []string, sep string) string {
	return Implode(a, sep)
}

// Trim - Strip whitespace (or other characters) from the beginning and end of a string
func Trim(s, cutSet string) string {
	if cutSet == "" {
		return strings.TrimSpace(s)
	}
	return strings.Trim(s, cutSet)
}

// Ltrim - Strip whitespace (or other characters) from the beginning of a string
func Ltrim(s, cutSet string) string {
	return strings.TrimLeft(s, cutSet)
}

// Rtrim - 从字符串末尾去除空格（或其他字符）
func Rtrim(s, cutSet string) string {
	return strings.TrimRight(s, cutSet)
}

// StrRepeat - 重复一个字符串
func StrRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StrReplace - 用替换字符串替换所有出现的搜索字符串
func StrReplace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// StrToLower - 使字符串小写
func StrToLower(s string) string {
	return strings.ToLower(s)
}

// StrToUpper - 使字符串大写
func StrToUpper(s string) string {
	return strings.ToUpper(s)
}

// StrFind - 查找字符串的第一个匹配项
func StrFind(s, substr string) int {
	return strings.Index(s, substr)
}

// StrPos - 查找字符串中子字符串首次出现的位置
func StrPos(s, substr string) int {
	return strings.Index(s, substr)
}

// Stripos - 查找不区分大小写的子字符串在字符串中首次出现的位置
func StrIpos(s, substr string) int {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Index(s, substr)
}

// StrLastPos - 查找字符串中最后一次出现子字符串的位置
func StrLastPos(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// Strripos - Find the position of the last occurrence of a case-insensitive substring in a string
func Strripos(s, substr string) int {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.LastIndex(s, substr)
}

// StrLastChr - 查找字符串中字符的最后一次出现
func StrLastChr(s, substr string) string {
	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}
	return s[i:]
}

// StrLen - 获取字符串长度
func StrLen(s string) int {
	return len(s)
}

// MbStrLen - 获取字符串长度
func MbStrLen(s string) int {
	return utf8.RuneCountInString(s)
}

// StrRev - 反转字符串
func StrReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SubStr - 返回字符串的一部分
func SubStr(s string, start int, length ...int) string {
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(s) + l
			if end < 0 {
				end = 0
			}
			return s[start:end]
		}
		end := start + l
		if end > len(s) {
			end = len(s)
		}
		return s[start:end]
	}
	return s[start:]
}

// MbSubStr - 获取字符串的一部分
func MbSubStr(s string, start int, length ...int) string {
	runes := []rune(s)
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(runes) + l
			if end < 0 {
				end = 0
			}
			return string(runes[start:end])
		}
		end := start + l
		if end > len(runes) {
			end = len(runes)
		}
		return string(runes[start:end])
	}
	return string(runes[start:])
}

// SubStrCount - 计算子字符串出现的次数
func SubStrCount(s, substr string) int {
	return strings.Count(s, substr)
}

// UppercaseFirst - 使字符串的第一个字符大写
func UppercaseFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// UppercaseWords — 将字符串中每个单词的第一个字符大写
func UppercaseWords(s string) string {
	return strings.Title(s)
}

// ChunkSplit - 将字符串分成较小的块
func ChunkSplit(str string, chunkLen int, end string) string {
	if chunkLen <= 0 {
		return str
	}
	var buf bytes.Buffer
	l := len(str)
	for i := 0; i < l; i += chunkLen {
		if chunkLen > l-i {
			chunkLen = l - i
		}
		buf.WriteString(str[i : i+chunkLen])
		buf.WriteString(end)
	}
	return buf.String()
}
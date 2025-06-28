package value

import (
	"fmt"
	"regexp"
	"strings"
)

// Title はタイトル
type Title string

func NewTitle(str string) Title {
	return Title(str)
}

func (t Title) string() string {
	return string(t)
}

func (t Title) Contains(str string) bool {
	return strings.Contains(t.string(), str)
}

func (t Title) AddSuffix(prefix string) Title {
	return NewTitle(t.string() + prefix)
}

func (t Title) TrimAfter(separator string) Title {
	idx := strings.Index(t.string(), separator)
	if idx != -1 {
		trimmed := t.string()[:idx]
		return NewTitle(trimmed)
	}

	return t
}

func (t Title) ReplaceAllString(re string, repl string) Title {
	replaced := regexp.MustCompile(re).ReplaceAllString(t.string(), repl)
	return NewTitle(replaced)
}

// DisassembleIntoParts はタイトルを部分に分解する
func (t Title) DisassembleIntoParts(re string) (TitleParts, error) {
	return newTitleParts(t, re)
}

// TitlePart はタイトル部分
type TitlePart string

func newTitlePart(str string) TitlePart {
	return TitlePart(str)
}

// TitleParts はタイトル部分の集合
type TitleParts []TitlePart

func newTitleParts(title Title, re string) (TitleParts, error) {
	matches := regexp.MustCompile(re).FindStringSubmatch(title.string())
	if len(matches) == 0 {
		return nil, fmt.Errorf("regex is not matched. title = %s, re = %s", title, re)
	}
	if len(matches) == 1 {
		return nil, fmt.Errorf("regex group is not matched. title = %s, re = %s", title, re)
	}
	matchGroups := matches[1:] // 最初の要素は正規表現の対象の文字列なので、スキップする

	tps := make(TitleParts, 0, len(matches))
	for _, match := range matchGroups {
		tps = append(tps, newTitlePart(match))
	}

	return tps, nil
}

func (tp TitlePart) string() string {
	return string(tp)
}

// Assemble はフォーマット文字列と部分のインデックスから、タイトルを組み立てる
//
// example:
// title := []TitleParts{"first", "second"}
// title.Assemble("%s and %s", 0, 1)
// //=> "first and second"
func (tps TitleParts) Assemble(format string, partIndex ...int) (Title, error) {
	if len(tps) <= maxValue(partIndex) {
		return "", fmt.Errorf("partIndex is out of range. tps = %v, len(tps) = %d, partIndex = %v", tps, len(tps), partIndex)
	}

	parts := make([]any, len(partIndex))
	for i, index := range partIndex {
		parts[i] = tps[index].string()
	}

	formatted := fmt.Sprintf(format, parts...)
	return NewTitle(formatted), nil
}

func maxValue(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

package DFA

import (
	"github.com/hakutyou/goapi/core/utils"
	"io/ioutil"
	"strings"
)

func init() {
	words := strings.Split(invalidWords, ",")
	for _, v := range words {
		InvalidWord[v] = nil
	}
	zip, _ := ioutil.ReadFile(sensitiveFile)
	outData, err := utils.UnCompress(zip)
	if err != nil {
		panic(err)
	}
	outSensitive := strings.Split(string(outData), ",")
	for _, v := range outSensitive {
		Set[v] = nil
	}
	addSensitiveToMap(Set)
}

// 生成词集合
func addSensitiveToMap(set map[string]interface{}) {
	for key := range set {
		str := []rune(key)
		nowMap := sensitiveWord
		for i := 0; i < len(str); i++ {
			if _, ok := nowMap[string(str[i])]; !ok { // 如果该key不存在
				thisMap := make(map[string]interface{})
				thisMap["isEnd"] = false
				nowMap[string(str[i])] = thisMap
				nowMap = thisMap
			} else {
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}
			if i == len(str)-1 {
				nowMap["isEnd"] = true
			}
		}
	}
}

// 敏感词汇转换为 *
func changeSensitiveWords(txt string) (word string) {
	str := []rune(txt)
	nowMap := sensitiveWord
	start := -1
	tag := -1
	for i := 0; i < len(str); i++ {
		if _, ok := InvalidWord[string(str[i])]; ok {
			continue // 如果是无效词汇直接跳过
		}
		if thisMap, ok := nowMap[string(str[i])].(map[string]interface{}); ok {
			// 记录敏感词第一个文字的位置
			tag++
			if tag == 0 {
				start = i
			}
			// 判断是否为敏感词的最后一个文字
			if isEnd, _ := thisMap["isEnd"].(bool); isEnd {
				//将敏感词的第一个文字到最后一个文字全部替换为 "*"
				for y := start; y < i+1; y++ {
					str[y] = '*'
				}
				// 重置标志数据
				nowMap = sensitiveWord
				start = -1
				tag = -1

			} else { // 不是最后一个，则将其包含的 map 赋值给 nowMap
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}

		} else { // 如果敏感词不是全匹配，则终止此敏感词查找。从开始位置的第二个文字继续判断
			if start != -1 {
				i = start
			}
			//重置标志参数
			nowMap = sensitiveWord
			start = -1
			tag = -1
		}
	}
	return string(str)
}

// 添加敏感词汇
func sensitiveAppend(appendContext []byte) bool {
	outData, err := ioutil.ReadFile(sensitiveFile)
	if err != nil {
		return false
	}
	outData, err = utils.UnCompress(outData)
	if err != nil {
		return false
	}
	outData = append(outData, appendContext...)
	outData, err = utils.Compress(outData)
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(sensitiveFile, outData, 0644)
	if err != nil {
		return false
	}
	return true
}

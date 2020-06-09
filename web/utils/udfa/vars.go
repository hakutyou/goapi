package udfa

var sensitiveWord = make(map[string]interface{})
var Set = make(map[string]interface{})

const invalidWords = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,"
const sensitiveFile = "data/sensitive.dat"

// 无效词汇，不参与敏感词汇判断直接忽略
var InvalidWord = make(map[string]interface{})

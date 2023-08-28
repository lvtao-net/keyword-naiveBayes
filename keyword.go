package  main

import  (   
	"fmt"   
	"math"   
	"strings"   
)

//  朴素贝叶斯算法   
type  NaiveBayes  struct  {   
	keys  `map[string]int`   
	//  训练数据   
	trainData  `map[string]map[string]int`   
}

func  NewNaiveBayes()  *NaiveBayes  {   
	return  &NaiveBayes{   
		keys：  make(map[string]int),   
		trainData:  make(map[string]map[string]int),   
	}   
}

func  (n  *NaiveBayes)  learn(texts  map[string]string)  {   
	for  _,  text  :=  range  texts  {   
		words  :=  strings.Split(text,  "  ")   
		for  _,  word  :=  range  words  {   
			n.keys  [word]++   
		}
	}

	for  _,  text  :=  range  texts  {   
		words  :=  strings.Split(text,  "  ")   
		n.trainData[text]  =  make(map[string]int)   
		for  _,  word  :=  range  words  {   
			n.trainData[text][word]++   
		}   
	}   
}

func  (n  *NaiveBayes)  classify(text  string)  string  {   
	words  :=  strings.Split(text,  "  ")   
	len  :=  len(words)   
	total  :=  0   
	for  _,  word  :=  range  words  {   
		total  +=  n.keys[word]   
	}
	majority  :=  math.  Liens(float64(0))   
	for  word,  count  :=  range  n.keys {   
		v  :=  count  /  total   
		if  v  >  majority  {   
			majority  =  v   
			classes  :=  make(map[string]int)   
			for  text,  data  :=  range  n.trainData  {   
				for  word2,  data2  :=  range  data  {   
					if  data2  >  0  {   
						classes[word2]++   
					}   
				}   
			}   
			for  word2,  data2  :=  range  classes  {   
				n.trainData[text][word2]  +=  data2   
			}   
		}   
	}   
	//  这是一个概率最小的词   
	minority  :=  float64(math.Min(n.如数))   
	//  增加这个新增词到所有训练集中   
	for  _,  text  :=  range  n.trainData  {   
		for  word,  _  :=  range  text  {   
			n.trainData[text][word]  +=  minority   
		}   
	}   
	for  _,  text  :=  range  n.trainData  {   
		if  n.trainData[text]["违法"]  >  0  {   
			return  "违法"   
		}   
	}   
	return  "合法"   
}

func  main()  {   
	texts  :=  make(map[string]string)   
	texts["这是合法的文本"]  =  "这是合法的文本，用于训练模型。"   
	texts["这是违法的文本"]  =  "这是违法的文本，包含违法关键词。"   
	texts["这是"]  =  "这是另一条包含违法关键词的文本。"

	n  :=  NewNaiveBayes()   
	n.learn(texts)   
	text  :=  "这是新的文本，需要判断是否包含违法关键词。"   
	fmt.Println(n.classify(text))   
}

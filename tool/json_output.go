package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func Dependencies_To_Json() {
	now := time.Now()
	dirName := fmt.Sprintf("outputs/%s", now.Format("2006-01-02_15-04"))
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		panic(err)
	}

	InstantiationToJson(dirName)
	StructEmbeddingToJSON(dirName)
	StructaggregationToJson(dirName)
	ConstReferToJSON(dirName)
	GlobalReferToJSON(dirName)
	FunctionCallToJSON(dirName)
	os.Exit(0)
}

func InstantiationToJson(dirName string) {
	// 遍历 DependencyList，找到所有的 Instantiation 数据
	var instantiationsJSON []InstantiationJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(Instantiation); ok {
					instantiationsJSON = append(instantiationsJSON, InstantiationJSON{
						Src:        dep.Src,
						Des:        dep.Des,
						StructName: ins.StructName,
						ObjName:    ins.ObjName,
						Token:      ins.Token,
						Pos:        ins.pos.String(),
					})
				}
			}
		}
	}

	fileName := "Instantiation.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(instantiationsJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Instantiation_JSON data has been written to file %s\n", filePath)
}

func StructaggregationToJson(dirName string) {
	var structAggregationsJSON []StructAggregationJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(structAggregation); ok {
					structAggregationsJSON = append(structAggregationsJSON, StructAggregationJSON{
						Src:     dep.Src,
						Des:     dep.Des,
						Whole:   ins.whole,
						Part:    ins.part,
						ObjName: ins.ObjName,
						Pos:     ins.pos.String(),
					})
				}
			}
		}
	}

	// 输出 JSON 格式数据到文件
	fileName := "StructAggregation.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(structAggregationsJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("StructAggregation_JSON data has been written to file %s\n", filePath)
}

func StructEmbeddingToJSON(dirName string) {
	var structEmbeddingsJSON []StructEmbeddingJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(structEmbedding); ok {
					structEmbeddingsJSON = append(structEmbeddingsJSON, StructEmbeddingJSON{
						Src:       dep.Src,
						Des:       dep.Des,
						Container: ins.container,
						Member:    ins.member,
						Pos:       ins.pos.String(),
					})
				}
			}
		}
	}

	fileName := "StructEmbedding.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(structEmbeddingsJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("StructEmbedding_JSON data has been written to file %s\n", filePath)
}

func ConstReferToJSON(dirName string) {
	var constRefersJSON []ConstReferJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(constRefer); ok {
					constRefersJSON = append(constRefersJSON, ConstReferJSON{
						Src:  dep.Src,
						Des:  dep.Des,
						Name: ins.name,
						Pos:  ins.pos.String(),
					})
				}
			}
		}
	}

	fileName := "ConstRefer.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(constRefersJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ConstRefer_JSON data has been written to file %s\n", filePath)
}

func GlobalReferToJSON(dirName string) {
	var globalRefersJSON []GlobalReferJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(globalRefer); ok {
					globalRefersJSON = append(globalRefersJSON, GlobalReferJSON{
						Src:  dep.Src,
						Des:  dep.Des,
						Name: ins.name,
						Type: ins.Type,
						Pos:  ins.pos.String(),
					})
				}
			}
		}
	}

	fileName := "GlobalRefer.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(globalRefersJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GlobalRefer_JSON data has been written to file %s\n", filePath)
}

func FunctionCallToJSON(dirName string) {
	var functionCallsJSON []FunctionCallJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(functionCall); ok {
					functionCallsJSON = append(functionCallsJSON, FunctionCallJSON{
						Src:     dep.Src,
						Des:     dep.Des,
						FunName: ins.funName,
						Pos:     ins.pos.String(),
					})
				}
			}
		}
	}

	fileName := "FunctionCall.txt"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(functionCallsJSON, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("FunctionCall_JSON data has been written to file %s\n", filePath)
}

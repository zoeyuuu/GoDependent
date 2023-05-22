package tool

import (
	"GoDependent/file"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func DependenciesToJson() {
	now := time.Now()
	dirName := fmt.Sprintf("outputs/"+file.Name+
		"/ByDependType/%s", now.Format("2006-01-02_15-04"))
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
	//os.Exit(0)
}

func InstantiationToJson(dirName string) {
	// 遍历 DependencyList，找到所有的 Instantiation 数据
	var instantiationsJSON []InstantiationJSON
	for _, dep := range DependencyList {
		for _, rels := range dep.Relations {
			for _, rel := range rels {
				if ins, ok := rel.(Instantiation); ok {
					instantiationsJSON = append(instantiationsJSON, InstantiationJSON{
						Src:        dep.RelativeSrc,
						Des:        dep.RelativeTar,
						StructName: ins.StructName,
						ObjName:    ins.ObjName,
						Token:      ins.Token,
						//Pos:        ins.Pos.String(),
					})
				}
			}
		}
	}

	fileName := "Instantiation.json"
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
						Src:     dep.RelativeSrc,
						Des:     dep.RelativeTar,
						Whole:   ins.Whole,
						Part:    ins.Part,
						ObjName: ins.ObjName,
						//Pos:     ins.Pos.String(),
					})
				}
			}
		}
	}

	// 输出 JSON 格式数据到文件
	fileName := "StructAggregation.json"
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
						Src:       dep.RelativeSrc,
						Des:       dep.RelativeTar,
						Container: ins.Container,
						Member:    ins.Member,
						//Pos:       ins.Pos.String(),
					})
				}
			}
		}
	}

	fileName := "StructEmbedding.json"
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
						Src:  dep.RelativeSrc,
						Des:  dep.RelativeTar,
						Name: ins.Name,
						//Pos:  ins.Pos.String(),
					})
				}
			}
		}
	}

	fileName := "ConstRefer.json"
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
						Src:  dep.RelativeSrc,
						Des:  dep.RelativeTar,
						Name: ins.Name,
						Type: ins.Type,
						//Pos:  ins.Pos.String(),
					})
				}
			}
		}
	}

	fileName := "GlobalRefer.json"
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
						Src:     dep.RelativeSrc,
						Des:     dep.RelativeTar,
						FunName: ins.FunName,
						//Pos:     ins.Pos.String(),
					})
				}
			}
		}
	}

	fileName := "FunctionCall.json"
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

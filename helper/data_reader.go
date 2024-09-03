package helper

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func ReadData[T any](model string, result *T) error{
	workingDirectory:="/home/zhalok/Desktop/practice/data"
	dataPath := filepath.Join(workingDirectory,model+".json")

	file,err := os.Open(dataPath)

	if err != nil{
		return err
		
		
	}

	defer file.Close()
	
	byteData,err := io.ReadAll(file)

	if err != nil{
		return err
		
	}

	
	err = json.Unmarshal(byteData,result)

	return err


}

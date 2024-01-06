// Package file
package file

import (
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/j23063519/clean_architecture/pkg/util"
)

// saving data to file
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// if file exists return true
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// remove file extension from file name
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// random name
func RandomNameFromFile(file *multipart.FileHeader) string {
	return util.RandomString(16) + filepath.Ext(file.Filename)
}

// saving file with gin
func SavingFile(c *gin.Context, file *multipart.FileHeader, dst string) (path string, err error) {
	// create directory
	os.Mkdir(dst, 0755)

	// create a random file name
	fileName := RandomNameFromFile(file)

	// if file exists then remove it
	path = dst + fileName
	os.Remove(path)

	// saving by gin
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.ErrorJSON("SavingFile", "save file", err)
		return
	}

	return
}

// delete file
func DeleteFile(path string) (err error) {
	err = os.Remove(path)
	if err != nil {
		log.ErrorJSON("DeleteFile", "delete file", err)
		return
	}

	// if folder does not have any file then remove it
	dirPath := strings.Split(path, "/")
	dirPath = dirPath[:len(dirPath)-1]
	subPath := ""
	for _, v := range dirPath {
		subPath += v + "/"
	}
	if len(ListFile(subPath)) < 1 {
		os.RemoveAll(subPath)
	}
	return
}

// list file by path
func ListFile(dst string) (paths []string) {
	var listFunc = func(path string, info fs.FileInfo, err error) error {
		var strRet string

		// if file not found then return error
		if info == nil {
			return err
		}

		// recursive loop
		if info.IsDir() {
			strRet += "/" + dst + info.Name() + "/"
			ListFile(strRet)
			return nil
		}

		// add path
		strRet += path

		// if file exist then put path into the paths
		if strings.HasSuffix(strRet, ".jpg") || strings.HasSuffix(strRet, ".jpeg") || strings.HasSuffix(strRet, ".png") {
			paths = append(paths, strRet)
		}

		return nil
	}

	// list func
	err := filepath.Walk(dst, listFunc)
	if err != nil {
		log.ErrorJSON("ListFile", "list file", err)
		return
	}

	return
}

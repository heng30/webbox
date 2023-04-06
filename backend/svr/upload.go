package svr

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	_ "log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type muploadMap struct {
	mmap  map[string]int
	mutex sync.Mutex
}

func (m *muploadMap) getPartIndex(uuid string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.mmap[uuid]; ok {
		return m.mmap[uuid]
	}
	return -1
}

func (m *muploadMap) setPartIndex(uuid string, index int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.mmap[uuid]; ok {
		if m.mmap[uuid]+1 != index {
			return errors.New(fmt.Sprintf("Invalid partIndex %d", index))
		}
		m.mmap[uuid] = index
		return nil
	}

	if index != 1 {
		return errors.New(fmt.Sprintf("Invalid partIndex %d", index))
	}

	m.mmap[uuid] = index
	return nil
}

func (m *muploadMap) delPartIndex(uuid string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.mmap, uuid)
}

var mupMap muploadMap = muploadMap{
	mmap: make(map[string]int),
}

func upload(r gin.IRouter) {
	r.POST("/upload", func(c *gin.Context) {
		filename, _ := c.GetQuery("filename")
		apath, err := getAbsPath(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		_, err = os.Stat(apath)
		if !os.IsNotExist(err) {
			c.JSON(http.StatusBadRequest, errorBody(errors.New(fmt.Sprintf("%s is exist!", apath))))
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		err = c.SaveUploadedFile(file, apath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}

func mupload(r gin.IRouter) {
	r.POST("/mupload", func(c *gin.Context) {
		filename, _ := c.GetQuery("filename")
		uuid, _ := c.GetQuery("uuid")
		partIndexStr, _ := c.GetQuery("partindex")
		totalPartStr, _ := c.GetQuery("totalpart")

		if filename == "" || uuid == "" || partIndexStr == "" || totalPartStr == "" {
			c.JSON(http.StatusBadRequest, errorBody(errors.New("Invalid arguments")))
			return
		}

		partIndex, err := strconv.Atoi(partIndexStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(errors.New(fmt.Sprintf("Invalid partindex: %s", partIndexStr))))
			return
		}

		totalPart, err := strconv.Atoi(totalPartStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(errors.New(fmt.Sprintf("Invalid totalpart: %s", totalPartStr))))
			return
		}

		apath, err := getAbsPath(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		if partIndex == 1 {
			_, err = os.Stat(apath)
			if !os.IsNotExist(err) {
				c.JSON(http.StatusBadRequest, errorBody(errors.New(fmt.Sprintf("%s is exist!", apath))))
				return
			}
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		err = saveUploadedFilePart(file, apath)
		if err != nil {
			mupMap.delPartIndex(uuid)
			os.Remove(apath)
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		if totalPart != 1 {
			if partIndex == totalPart {
				mupMap.delPartIndex(uuid)
			} else {
				if err := mupMap.setPartIndex(uuid, partIndex); err != nil {
					os.Remove(apath)
					c.JSON(http.StatusBadRequest, errorBody(err))
					return

				}
			}
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}

func saveUploadedFilePart(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

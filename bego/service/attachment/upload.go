package attachment

import (
	"bego/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func SaveImage(m *models.FileInfo, tmp multipart.File, mime string, filename string, created time.Time, data map[string]string, size int64) (updata *models.FileInfo, err error) {

	m.Gid = data["gid"]
	m.Type = data["type"]
	m.CreateTime = created
	m.UpdateTime = created
	m.Filename = filename

	if err, m.Ext = CheckFileType(filename, mime); err != nil {
		return nil, errors.New("非法类型!")
	}

	if err := CheckSize(size); err != nil {
		return nil, errors.New("文件过大!")
	}

	// 文件大小
	f := float32(size) / 1024
	m.Size = fmt.Sprintf("%f", f) + "kb"

	// 文件保存路径
	path := GenImagePath(m)
	os.MkdirAll(strings.Trim(path, "/"), 0755)
	savename := GetSaveName("date") + "." + m.Ext
	m.Path = path + "/" + savename

	if err := m.Insert(); err != nil || m.Id <= 0 {
		return nil, err
	}

	return m, nil
}

func GenImagePath(img *models.FileInfo) string {
	updir, _ := web.AppConfig.String("upload::upImg")
	return updir + beego.Date(img.CreateTime, "20060102")
}

func GetSaveName(t string) (s string) {
	switch t {
	case "date":
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 19)))
		s = hex.EncodeToString(md5Ctx.Sum(nil))
	case "hash":
		randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
		hashname := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
		s = string(hashname[:])
	default:
		s = beego.Date(time.Now(), "20060102") + string(time.Now().UnixNano())
	}
	return s
}

func CheckSize(size int64) error {
	maxsize, _ := web.AppConfig.String("MaxMemory")
	ms, _ := strconv.Atoi(maxsize)
	if size > int64(ms) {
		return errors.New("invialid files")
	}
	return nil
}

func CheckFileType(filename, mime string) (err error, s string) {
	var ext string
	// test image mime type
	switch mime {
	case "image/jpeg":
		ext = "jpg"

	case "image/png":
		ext = "png"

	case "image/gif":
		ext = "gif"

	default:
		ext = filepath.Ext(filename)
		switch ext {
		case ".jpg", ".png", ".gif", ".jpeg":
		default:
			return fmt.Errorf("unsupport image format `%s`", filename), ""
		}
	}
	return nil, ext
}

// DelOldFile 删除该游戏对应的图片信息
func DelOldFile(num int64, filelist []models.FileInfo, types string) {
	if num != 0 {
		for _, v := range filelist {
			if v.Type == types {
				models.DelFile(v.Id)
				os.Remove(strings.Trim(v.Path, "/"))
			}
		}
	}
}

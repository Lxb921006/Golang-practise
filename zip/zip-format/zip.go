package zip_format

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Zip struct {
	Src     string `json:"src"`
	Dst     string `json:"dst"`
	ZipName string `json:"zip-name"`
}

func (z *Zip) Check() (err error) {
	_, err = os.Stat(z.Src)
	if err != nil {
		return
	}

	_, err = os.Stat(z.Dst)
	if err != nil {
		return
	}

	z.ZipName = filepath.Join(z.Src, filepath.Base(z.Src)+".zip")

	return
}

func (z *Zip) ZipFile() (err error) {
	fc, err := os.Create(z.ZipName)
	if err != nil {
		return
	}

	defer fc.Close()

	zipWriter := zip.NewWriter(fc)

	defer zipWriter.Close()

	err = filepath.Walk(z.Src, func(path string, info fs.FileInfo, err error) error {
		head, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		file, err := filepath.Rel(z.Src, path)
		if err != nil {
			return err
		}

		head.Name = file

		ch, err := zipWriter.CreateHeader(head)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fo, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(ch, fo)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return
	}

	return
}

func (z *Zip) UnZipFile() (err error) {
	return
}

func NewZip(src, dst, zipName string) *Zip {
	z := &Zip{src, dst, zipName}

	err := z.Check()
	if err != nil {
		log.Fatalln(err)
	}

	return z
}

package archive_learn_01

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func Unpack(srcFile string) {
	fr, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	tr := tar.NewReader(fr)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			log.Println(err)
			continue
		}
		// 读取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, err := os.Create(fi.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err := io.Copy(fw, tr); err != nil {
			log.Println(err)
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		fw.Close()
	}

}

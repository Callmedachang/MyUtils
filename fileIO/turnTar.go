package fileIO
import (
	"archive/tar"
	"io"
	"fmt"
	"os"
	"compress/gzip"
	"path"
)
// Create a buffer to write our archive to.
// main functions shows how to TarGz a directory and
// UnTarGz a fileIO
//func main() {
//  targetFilePath := "testdata.tar.gz"
//  srcDirPath := "testdata"
//  TarGz(srcDirPath, targetFilePath)
//  UnTarGz(targetFilePath, srcDirPath+"_temp")
//}

// Gzip and tar from source directory or fileIO to destination fileIO
// you need check fileIO exist before you call this function
func TarGz(srcDirPath string, destFilePath string) {
	//创建目标目录（文件）
	fw, err := os.Create(destFilePath)
	dealErr(err)
	defer fw.Close()

	// Gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Check if it's a fileIO or a directory
	f, err := os.Open(srcDirPath)
	dealErr(err)
	fi, err := f.Stat()//得到fileinfo
	dealErr(err)
	if fi.IsDir() {
		// handle source directory
		fmt.Println("Cerating tar.gz from directory...")
		tarGzDir(srcDirPath, path.Base(srcDirPath), tw)
	} else {
		// handle fileIO directly
		fmt.Println("Cerating tar.gz from " + fi.Name() + "...")
		tarGzFile(srcDirPath, fi.Name(), tw, fi)
	}
	fmt.Println("Well done!")
}

// Deal with directories
// if find files, handle them with tarGzFile
// Every recurrence append the base path to the recPath
// recPath is the path inside of tar.gz
func tarGzDir(srcDirPath string, recPath string, tw *tar.Writer) {
	// Open source diretory
	dir, err := os.Open(srcDirPath)
	dealErr(err)
	defer dir.Close()

	// Get fileIO info slice
	fis, err := dir.Readdir(0)
	dealErr(err)
	for _, fi := range fis {
		// Append path
		curPath := srcDirPath + "/" + fi.Name()
		// Check it is directory or fileIO
		if fi.IsDir() {
			// Directory
			// (Directory won't add unitl all subfiles are added)
			fmt.Printf("Adding path...%s\\n", curPath)
			tarGzDir(curPath, recPath+"/"+fi.Name(), tw)
		} else {
			// File
			fmt.Printf("Adding fileIO...%s\\n", curPath)
		}

		tarGzFile(curPath, recPath+"/"+fi.Name(), tw, fi)
	}
}

// Deal with files
func tarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) {
	if fi.IsDir() {
		// Create tar header
		hdr := new(tar.Header)
		// if last character of header name is '/' it also can be directory
		// but if you don't set Typeflag, error will occur when you untargz
		hdr.Name = recPath + "/"
		hdr.Typeflag = tar.TypeDir
		hdr.Size = 0
		//hdr.Mode = 0755 | c_ISDIR
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err := tw.WriteHeader(hdr)
		dealErr(err)
	} else {
		// File reader
		fr, err := os.Open(srcFile)
		dealErr(err)
		defer fr.Close()

		// Create tar header
		hdr := new(tar.Header)
		hdr.Name = recPath
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err = tw.WriteHeader(hdr)
		dealErr(err)

		// Write fileIO data
		_, err = io.Copy(tw, fr)
		dealErr(err)
	}
}

// Ungzip and untar from source fileIO to destination directory
// you need check fileIO exist before you call this function
func UnTarGz(srcFilePath string, destDirPath string) {
	fmt.Println("UnTarGzing " + srcFilePath + "...")
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)

	fr, err := os.Open(srcFilePath)
	dealErr(err)
	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)

	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		//handleError(err)
		fmt.Println("UnTarGzing fileIO..." + hdr.Name)
		// Check if it is diretory or fileIO
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create fileIO
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to fileIO
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			dealErr(err)
			_, err = io.Copy(fw, tr)
			dealErr(err)
		}
	}
	fmt.Println("Well done!")
}

func dealErr(err error){
	if err!=nil {
		fmt.Println(err)
	}
}
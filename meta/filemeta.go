package meta

import "FILESTORE-SERVER/db"

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var flieMetas map[string]FileMeta

func init()  {
	flieMetas = make(map[string]FileMeta)
}


//UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta)  {
	flieMetas[fmeta.FileSha1] = fmeta
}

func UpdateFileMetaDb(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1,fmeta.FileName,fmeta.FileSize,fmeta.Location)
}
//GetFileMeta:通sha1值获取文件的原信息对象
func GetFileMeta(fileSha1 string ) FileMeta {

	return flieMetas[fileSha1]
}

func GetFileMetaDb(filehash string) (*FileMeta,error)  {
	fileT,err := db.GetFileMeta(filehash)

	if err !=nil {
		return nil,err
	}

	fileMe := FileMeta{FileSha1:fileT.FileHash,FileName:fileT.FileName,FileSize:fileT.FileSize,Location:fileT.FileAddr}

	return &fileMe,nil
}

//删除
func RemoveFileMetaHandler(fileSha1 string)  {
	delete(flieMetas,fileSha1)
}
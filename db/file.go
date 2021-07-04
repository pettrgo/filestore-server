package db

import (
	"fmt"

	mydb "filestore-server/db/mysql"
)

//文件上传完成，保存meta
func OnFileUploadFinished(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`, `file_addr`)" +
			"values(?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(fileHash, fileName, fileSize, fileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been upload before", fileHash)
			return false
		}
		return true
	}
	return false
}

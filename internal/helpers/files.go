package helpers

import "fmt"

func FileDir(filesDir, userID string) string {
	return fmt.Sprintf("%s/%s", filesDir, userID)
}
func FilePath(filesDir, userID, fileID string) string {
	return fmt.Sprintf("%s/%s", FileDir(filesDir, userID), fileID)
}

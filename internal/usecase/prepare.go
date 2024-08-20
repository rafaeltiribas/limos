package usecase

func PrepareApplication() {
	if CheckFile() {
		DeleteFile()
	}
}

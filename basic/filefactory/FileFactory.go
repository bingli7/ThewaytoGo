package filefactory

type fileFactory struct {
	fd		int     // 文件描述符
	name	string  // 文件名
}

func NewFile(fd int, name string) *fileFactory {
	return &fileFactory{fd, name}
}
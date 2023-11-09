package types

type FileList struct {
	list []string
}

func (f *FileList) All() []string {
	return []string{}
}

func (f *FileList) OfType(suffix string) []string {
	return []string{}
}

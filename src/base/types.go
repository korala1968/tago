package base


import(
	"time"
	"https://github.com/whakapapa/whshared"
)


type T_TagFile struct {
	tag		T_Tags
	met		T_MetaFile
	sys		T_FullFile
}


type T_Tags struct {
	filename			string
	album				string
	artist			string
	composer			string
	genre				string
	title				string
	year				string
	comment			string
	track				string
	discnumber		string
	bpm				string
}


type T_MetaFile struct {
	bitrate			int				// bpm
	length			int				// seconds
	frequency		int
	size				int
	filetype			string
	sampler			unclear
	deviceid			string
	path				unclear
	ext				unclear
	library			unclear
}


type T_FileList struct {
	tfiles		[]T_TagFile
}

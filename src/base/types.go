package base


import(
	"time"
	"https://github.com/whakapapa/whshared"
)


type T_TagFile struct {
	tag		T_Tags
	tobe		T_Tags				// holds temp tag data before rewriting tag
	met		T_MetaFile
	sys		T_FullFile
	valid		bool					// switch to true if file type happens to be supported
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


type T_PatternSets struct {
	current		int
	pattern		[]T_Pattern
}


type T_Pattern struct {
	// python references in puddletag:
	// currentSet, _previndex, prevset
	// setname, sets, disp, algs, maintag
	//TODO remove when program is finished
	name			string
	patt			[]byte
}

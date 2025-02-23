package ngwords

// NGWords は全てのNGワードのカテゴリーをまとめた構造体
type NGWords struct {
	HateWords             []string
	ViolenceWords         []string
	SexualHarassmentWords []string
	SpamWords             []string
	FakeInfoWords         []string
	SelfHarmWords         []string
}

// NewNGWords は全てのNGワードをまとめたNGWordsオブジェクトを返すコンストラクタ
func NewNGWords() NGWords {
	return NGWords{
		HateWords: []string{
			"死ね", "クズ", "バカ", "ゴミ", "嫌い", "最低", "殺す",
		},
		ViolenceWords: []string{
			"暴力", "殴る", "ぶっ潰す", "刺す", "暴行",
		},
		SexualHarassmentWords: []string{
			"セクハラ", "キモい", "エッチ", "変態", "セックス",
		},
		SpamWords: []string{
			"http://", "https://", "広告", "宣伝", "リンク", "買うなら",
		},
		FakeInfoWords: []string{
			"詐欺", "偽物", "嘘", "サクラ", "最悪",
		},
		SelfHarmWords: []string{
			"自殺", "死にたい", "傷つけたい", "命を絶つ",
		},
	}
}

// GetAllNGWords は全てのNGワードを1つの構造体で取得するメソッド
func (ng *NGWords) GetAllNGWords() map[string][]string {
	return map[string][]string{
		"HateWords":             ng.HateWords,
		"ViolenceWords":         ng.ViolenceWords,
		"SexualHarassmentWords": ng.SexualHarassmentWords,
		"SpamWords":             ng.SpamWords,
		"FakeInfoWords":         ng.FakeInfoWords,
		"SelfHarmWords":         ng.SelfHarmWords,
	}
}

// GetAllNGWordsCombined は全てのNGワードを1つのスライスにまとめて取得するメソッド
func (ng *NGWords) GetAllNGWordsCombined() []string {
	allWords := append(ng.HateWords, ng.ViolenceWords...)
	allWords = append(allWords, ng.SexualHarassmentWords...)
	allWords = append(allWords, ng.SpamWords...)
	allWords = append(allWords, ng.FakeInfoWords...)
	allWords = append(allWords, ng.SelfHarmWords...)
	return allWords
}

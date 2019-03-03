package whatlanggo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"unicode"
)

func TestDetect(t *testing.T) {
	tests := map[string]Info{
		"Además de todo lo anteriormente dicho, también encontramos...": {Spa, unicode.Latin, 1},
		"बहुत बहुत (धन्यवाद / शुक्रिया)!":                               {Hin, unicode.Devanagari, 1},
		"अनुच्छेद १: सबहि लोकानि आजादे जम्मेला आओर ओखिनियो के बराबर सम्मान आओर अघ्कार प्राप्त हवे। ओखिनियो के पास समझ-बूझ आओर अंत:करण के आवाज होखता आओर हुनको के दोसरा के साथ भाईचारे के बेवहार करे के होखला": {Bho, unicode.Devanagari, 1},
		"ኢትዮጵያ አፍሪቃ ውስጥ ናት":         {Amh, unicode.Ethiopic, 1},
		"لغتي العربية ليست كما يجب": {Arb, unicode.Arabic, 1},
		"我爱你": {Cmn, unicode.Han, 1},
		"আমি তোমাকে ভালোবাস ": {Ben, unicode.Bengali, 1},
		"울란바토르": {Kor, unicode.Hangul, 1},
		"ყველა ადამიანი იბადება თავისუფალი და თანასწორი თავისი ღირსებითა და უფლებებით":        {Kat, unicode.Georgian, 1},
		"Όλοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και τα δικαιώματα.": {Ell, unicode.Greek, 1},
		"ಎಲ್ಲಾ ಮಾನವರ ಉಚಿತ ಮತ್ತು ಘನತೆ ಮತ್ತು ಹಕ್ಕುಗಳಲ್ಲಿ ಸಮಾನ ಹುಟ್ಟಿದ.":                         {Kan, unicode.Kannada, 1},
		"நீங்கள் ஆங்கிலம் பேசுவீர்களா?":                                                       {Tam, unicode.Tamil, 1},
		"มนุษย์ทุกคนเกิดมามีอิสระและเสมอภาคกันในศักดิ์ศรีและสิทธิ":                            {Tha, unicode.Thai, 1},
		"નાણાં મારા લોહીમાં છે":    {Guj, unicode.Gujarati, 1},
		" ਗੁਰੂ ਗ੍ਰੰਥ ਸਾਹਿਬ ਜੀ":     {Pan, unicode.Gurmukhi, 1},
		"నన్ను ఒంటరిగా వదిలేయ్":    {Tel, unicode.Telugu, 1},
		"എന്താണ് നിങ്ങളുടെ പേര് ?": {Mal, unicode.Malayalam, 1},
		"ମୁ ତୁମକୁ ଭଲ ପାଏ |":        {Ori, unicode.Oriya, 1},
		"အားလုံးလူသားတွေအခမဲ့နှင့်ဂုဏ်သိက္ခာနှင့်လူ့အခွင့်အရေးအတွက်တန်းတူဖွားမြင်ကြသည်။": {Mya, unicode.Myanmar, 1},
		"වෙලාව කියද?":                        {Sin, unicode.Sinhala, 1},
		"ពួកម៉ាកខ្ញុំពីរនាក់នេះ":             {Khm, unicode.Khmer, 1},
		"其疾如風、其徐如林、侵掠如火、不動如山、難知如陰、動如雷震。":     {Cmn, unicode.Han, 1},
		"知彼知己、百戰不殆。不知彼而知己、一勝一負。不知彼不知己、毎戰必殆。": {Cmn, unicode.Han, 1},
		"支那の上海の或町です。":                        {Jpn, _HiraganaKatakana, 1},
		"或日の暮方の事である。":                        {Jpn, _HiraganaKatakana, 1},
		"今日は":                                {Jpn, _HiraganaKatakana, 1},
		"コンニチハ":                              {Jpn, _HiraganaKatakana, 1},
		"ﾀﾅｶ ﾀﾛｳ":                            {Jpn, _HiraganaKatakana, 1},
		"どうもありがとう":                           {Jpn, _HiraganaKatakana, 1},
	}

	for key, value := range tests {
		got := Detect(key)

		if value.Lang != got.Lang || value.Script != got.Script {
			t.Fatalf("%s want %v %v got %v %v", key, LangToString(value.Lang), Scripts[value.Script], LangToString(got.Lang), Scripts[got.Script])
		}
	}
}

func TestDetectLang(t *testing.T) {
	tests := map[string]Lang{
		"Та нічого, все нормально. А в тебе як?": Ukr,
		"Vouloir, c'est pouvoir":                                Fra,
		"Where there is a will there is a way":                  Eng,
		"Mi ŝategas la japanan kaj studas ĝin kelkajn jarojn 😊": Epo,
		"Te echo de menos":                                      Spa,
		"Buona notte e sogni d'oro!":                            Ita,
	}

	for text, want := range tests {
		got := DetectLang(text)
		if got != want {
			t.Fatalf("%s want %v got %v", text, LangToString(want), LangToString(got))
		}
	}
}

// Test detect with empty options and supported language and script
func TestDetectWithOptionsEmptySupportedLang(t *testing.T) {
	want := Info{Epo, unicode.Latin, 1}
	got := DetectWithOptions("La viro amas hundojn. Hundo estas la plej bona amiko de viro", Options{})
	if want.Lang != got.Lang && want.Script != got.Script {
		t.Fatalf("want %v %v got %v %v", want.Lang, want.Script, got.Lang, got.Script)
	}
}

// Test detect with empty options and nonsupported script(Balinese)
func TestDetectWithOptionsEmptyNonSupportedLang(t *testing.T) {
	want := Info{-1, nil, 0}
	got := DetectWithOptions("ᬅᬓ᭄ᬱᬭᬯ᭄ᬬᬜ᭄ᬚᬦ", Options{})
	if want.Lang != got.Lang && want.Script != got.Script {
		t.Fatalf("want %v %v got %v %v", want.Lang, want.Script, got.Lang, got.Script)
	}
}

func TestDetectWithOptionsWithBlacklist(t *testing.T) {
	text := "האקדמיה ללשון העברית"
	//All languages with Hebrew text blacklisted ... returns correct script but invalid language
	options1 := Options{
		Blacklist: map[Lang]bool{
			Heb: true,
			Ydd: true,
		},
	}
	want := Info{-1, unicode.Hebrew, 1}
	got := DetectWithOptions(text, options1)
	if got.Lang != want.Lang && want.Script != got.Script {
		t.Fatalf("Want %s %s got %s %s", LangToString(want.Lang), Scripts[want.Script], LangToString(got.Lang), Scripts[got.Script])
	}

	text = "Tu me manques"
	want = Info{Fra, unicode.Latin, 1}
	options3 := Options{
		Blacklist: map[Lang]bool{
			Kur: true,
		},
	}
	got = DetectWithOptions(text, options3)
	if got.Lang != want.Lang && want.Script != got.Script {
		t.Fatalf("Want %s %s got %s %s", LangToString(want.Lang), Scripts[want.Script], LangToString(got.Lang), Scripts[got.Script])
	}
}

func TestWithOptionsWithWhitelist(t *testing.T) {
	text := "Mi ne scias!"
	want := Info{Epo, unicode.Latin, 1}
	options2 := Options{
		Whitelist: map[Lang]bool{
			Epo: true,
			Ukr: true,
		},
	}
	got := DetectWithOptions(text, options2)
	if got.Lang != want.Lang && want.Script != got.Script {
		t.Fatalf("Want %s %s got %s %s", LangToString(want.Lang), Scripts[want.Script], LangToString(got.Lang), Scripts[got.Script])
	}
}

func TestDetectLangWithOptions(t *testing.T) {
	text := "All evil come from a single cause ... man's inability to sit still in a room"
	want := Eng
	//without blacklist
	got := DetectLangWithOptions(text, Options{})
	if want != got {
		t.Fatalf("want %s got %s", LangToString(want), LangToString(got))
	}

	//with blacklist
	options := Options{
		Blacklist: map[Lang]bool{
			Jav: true,
			Tgl: true,
			Nld: true,
			Uzb: true,
			Swe: true,
			Nob: true,
			Ceb: true,
			Ilo: true,
		},
	}
	got = DetectLangWithOptions(text, options)
	if want != got {
		t.Fatalf("want %s got %s", LangToString(want), LangToString(got))
	}
}

func Test_detectLangBaseOnScriptUnsupportedScript(t *testing.T) {
	want := Info{-1, nil, 0}
	gotLang, gotConfidence := detectLangBaseOnScript("ᬅᬓ᭄ᬱᬭᬯ᭄ᬬᬜ᭄ᬚᬦ", Options{}, unicode.Balinese)
	if want.Lang != gotLang && want.Confidence != gotConfidence {
		t.Fatalf("want %v %v got %v %v", want.Lang, want.Script, gotLang, gotConfidence)
	}
}

func TestWithMultipleExamples(t *testing.T) {
	examplesFile, err := os.Open("testdata/examples.json")
	if err != nil {
		t.Fatal("Error opening testdata/examples.json")
	}

	defer examplesFile.Close()

	byteValue, err := ioutil.ReadAll(examplesFile)
	if err != nil {
		t.Fatal("Error reading testdata/examples.json")
	}

	var examples map[string]string
	err = json.Unmarshal(byteValue, &examples)
	if err != nil {
		t.Fatal("Error Unmarshalling json")
	}

	for lang, text := range examples {
		want := CodeToLang(lang)
		info := Detect(text)
		if info.Lang != want && !info.IsReliable() {
			t.Fatalf("want %v, got %v", Langs[want], Langs[info.Lang])
		}
	}
}

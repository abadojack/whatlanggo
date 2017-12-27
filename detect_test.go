package whatlanggo

import (
	"testing"
	"unicode"
)

func Test_Detect(t *testing.T) {
	tests := map[string]Info{
		"Además de todo lo anteriormente dicho, también encontramos...": {Spa, unicode.Latin},
		"बहुत बहुत (धन्यवाद / शुक्रिया)!":                               {Hin, unicode.Devanagari},
		"अनुच्छेद १: सबहि लोकानि आजादे जम्मेला आओर ओखिनियो के बराबर सम्मान आओर अघ्कार प्राप्त हवे। ओखिनियो के पास समझ-बूझ आओर अंत:करण के आवाज होखता आओर हुनको के दोसरा के साथ भाईचारे के बेवहार करे के होखला": {Bho, unicode.Devanagari},
		"ኢትዮጵያ አፍሪቃ ውስጥ ናት":         {Amh, unicode.Ethiopic},
		"لغتي العربية ليست كما يجب": {Arb, unicode.Arabic},
		"我爱你": {Cmn, unicode.Han},
		"আমি তোমাকে ভালোবাস ": {Ben, unicode.Bengali},
		"울란바토르":               {Kor, unicode.Hangul},
		"ყველა ადამიანი იბადება თავისუფალი და თანასწორი თავისი ღირსებითა და უფლებებით":        {Kat, unicode.Georgian},
		"Όλοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και τα δικαιώματα.": {Ell, unicode.Greek},
		"ಎಲ್ಲಾ ಮಾನವರ ಉಚಿತ ಮತ್ತು ಘನತೆ ಮತ್ತು ಹಕ್ಕುಗಳಲ್ಲಿ ಸಮಾನ ಹುಟ್ಟಿದ.":                         {Kan, unicode.Kannada},
		"நீங்கள் ஆங்கிலம் பேசுவீர்களா?":                                                       {Tam, unicode.Tamil},
		"มนุษย์ทุกคนเกิดมามีอิสระและเสมอภาคกันในศักดิ์ศรีและสิทธิ":                            {Tha, unicode.Thai},
		"નાણાં મારા લોહીમાં છે":                                                               {Guj, unicode.Gujarati},
		" ਗੁਰੂ ਗ੍ਰੰਥ ਸਾਹਿਬ ਜੀ":                                                                {Pan, unicode.Gurmukhi},
		"నన్ను ఒంటరిగా వదిలేయ్":                                                               {Tel, unicode.Telugu},
		"എന്താണ് നിങ്ങളുടെ പേര് ?":                                                            {Mal, unicode.Malayalam},
		"ମୁ ତୁମକୁ ଭଲ ପାଏ |":                                                                   {Ori, unicode.Oriya},
		"အားလုံးလူသားတွေအခမဲ့နှင့်ဂုဏ်သိက္ခာနှင့်လူ့အခွင့်အရေးအတွက်တန်းတူဖွားမြင်ကြသည်။": {Mya, unicode.Myanmar},
		"වෙලාව කියද?":                        {Sin, unicode.Sinhala},
		"ពួកម៉ាកខ្ញុំពីរនាក់នេះ":             {Khm, unicode.Khmer},
		"其疾如風、其徐如林、侵掠如火、不動如山、難知如陰、動如雷震。":     {Cmn, unicode.Han},
		"知彼知己、百戰不殆。不知彼而知己、一勝一負。不知彼不知己、毎戰必殆。": {Cmn, unicode.Han},
		"支那の上海の或町です。":                        {Jpn, unicode.Hiragana},
		"或日の暮方の事である。":                        {Jpn, unicode.Hiragana},
		"今日は":                                {Jpn, unicode.Hiragana},
		"コンニチハ":                              {Jpn, unicode.Katakana},
		"ﾀﾅｶ ﾀﾛｳ":                            {Jpn, unicode.Katakana},
		"どうもありがとう":                           {Jpn, unicode.Hiragana},
	}

	for key, value := range tests {
		got := Detect(key)

		if value.Lang != got.Lang || value.Script != got.Script {
			t.Fatalf("%s want %v %v got %v %v", key, LangToString(value.Lang), Scripts[value.Script], LangToString(got.Lang), Scripts[got.Script])
		}
	}
}

func Test_DetectLang(t *testing.T) {
	tests := map[string]Lang{
		"Та нічого, все нормально. А в тебе як?":                Ukr,
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

func Test_DetectWithOptions(t *testing.T) {
	//without blacklist
	want := Info{Epo, unicode.Latin}
	got := DetectWithOptions("La viro amas hundojn. Hundo estas la plej bona amiko de viro", Options{})
	if want.Lang != got.Lang && want.Script != got.Script {
		t.Fatalf("want %v %v got %v %v", want.Lang, want.Script, got.Lang, got.Script)
	}

	text := "האקדמיה ללשון העברית"
	//All languages with Hebrew text blacklisted ... returns correct script but invalid language
	options1 := Options{
		Blacklist: map[Lang]bool{
			Heb: true,
			Ydd: true,
		},
	}
	want = Info{-1, unicode.Hebrew}
	got = DetectWithOptions(text, options1)
	if got.Lang != want.Lang && want.Script != got.Script {
		t.Fatalf("Want %s %s got %s %s", LangToString(want.Lang), Scripts[want.Script], LangToString(got.Lang), Scripts[got.Script])
	}

	text = "Mi ne scias!"
	want = Info{Epo, unicode.Latin}
	options2 := Options{
		Whitelist: map[Lang]bool{
			Epo: true,
			Ukr: true,
		},
	}
	got = DetectWithOptions(text, options2)
	if got.Lang != want.Lang && want.Script != got.Script {
		t.Fatalf("Want %s %s got %s %s", LangToString(want.Lang), Scripts[want.Script], LangToString(got.Lang), Scripts[got.Script])
	}

	text = "Tu me manques"
	want = Info{Fra, unicode.Latin}
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

func Test_DetectLangWithOptions(t *testing.T) {
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

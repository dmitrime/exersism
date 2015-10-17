package foodchain

const TestVersion = 1

var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow"}
var words = []string{
	" that wriggled and jiggled and tickled inside her",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
}

func Song() string {
	sing := ""
	for i := 1; i <= 8; i++ {
		sing += Verse(i)
		if i != 8 {
			sing += "\n\n"
		}
	}
	return sing
}

func Verse(v int) string {
	if v < 1 || v > 8 {
		return "La-la-la..."
	}
	if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}
	sing := "I know an old lady who swallowed a " + animals[v-1] + ".\n"
	if v > 1 {
		sing += words[v-1] + "\n"
		for i := v - 1; i > 0; i-- {
			sing += "She swallowed the " + animals[i] + " to catch the " + animals[i-1]
			if i == 2 {
				sing += words[0]
			}
			sing += ".\n"
		}
	}
	sing += "I don't know why she swallowed the fly. Perhaps she'll die."
	return sing
}

func Verses(verses ...int) string {
	sing := ""
	for n, a := range verses {
		sing += Verse(a)
		if n != len(verses)-1 {
			sing += "\n\n"
		}
	}
	return sing
}

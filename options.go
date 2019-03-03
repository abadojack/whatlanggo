package whatlanggo

// Options represents options that can be set when detecting a language or/and script such
// blacklisting languages to skip checking.
type Options struct {
	Whitelist map[Lang]bool
	Blacklist map[Lang]bool
}

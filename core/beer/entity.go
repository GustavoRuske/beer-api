package beer

type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle    = 1
	TypeLagger = 2
	TypeMalt   = 3
	TypeStout  = 4
)

func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLagger:
		return "Lagger"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknown"
}

type BeerStyle int

const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (t BeerStyle) String() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StyleFruit:
		return "Fruit"
	case StyleGolden:
		return "Golden"
	case StyleHoney:
		return "Honey"
	case StyleIPA:
		return "IPA"
	case StyleLime:
		return "Lime"
	case StylePale:
		return "Pale"
	case StylePilsner:
		return "Pilsner"
	case StyleRed:
		return "Red"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	}

	return "Unknown"
}

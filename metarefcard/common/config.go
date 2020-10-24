package common

// Config contains all the configuration data for the app
type Config struct {
	AppName       string `yaml:"AppName"`
	Version       string `yaml:"Version"`
	DebugOutput   bool   `yaml:"DebugOutput"`
	VerboseOutput bool   `yaml:"VerboseOutput"`

	DevicesFile string `yaml:"DevicesFile"`
	Devices     Devices

	DefaultImage    Dimensions2d `yaml:"DefaultImage"`
	PixelMultiplier float64      `yaml:"PixelMultiplier"`
	HotasImagesDir  string       `yaml:"HotasImagesDir"`
	LogoImagesDir   string       `yaml:"LogoImagesDir"`
	JpgQuality      int          `yaml:"JpgQuality"`

	FontsDir          string  `yaml:"FontsDir"`
	InputFont         string  `yaml:"InputFont"`
	InputFontSize     float64 `yaml:"InputFontSize"`
	InputMinFontSize  int     `yaml:"InputMinFontSize"`
	DefaultLineHeight int     `yaml:"DefaultLineHeight"`
	InputPixelXInset  int     `yaml:"InputPixelXInset"`
	InputPixelYInset  int     `yaml:"InputPixelYInset"`

	ImageHeading struct {
		Font             string  `yaml:"Font"`
		FontSize         float64 `yaml:"FontSize"`
		Inset            Point2d `yaml:"Inset"`
		TextHeight       float64 `yaml:"TextHeight"`
		TextColour       string  `yaml:"TextColour"`
		BackgroundHeight float64 `yaml:"BackgroundHeight"`
		BackgroundColour string  `yaml:"BackgroundColour"`
	} `yaml:"ImageHeading"`

	Watermark struct {
		Text       string  `yaml:"Text"`
		TextColour string  `yaml:"TextColour"`
		Font       string  `yaml:"Font"`
		FontSize   float64 `yaml:"FontSize"`
		Location   Point2d `yaml:"Location"`
	} `yaml:"Watermark"`

	BackgroundColour string   `yaml:"BackgroundColour"`
	LightColour      string   `yaml:"LightColour"`
	DarkColour       string   `yaml:"DarkColour"`
	AlternateColours []string `yaml:"AlternateColours"`
}

// Point2d contains x and y
type Point2d struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
}

// Dimensions2d contains width and height
type Dimensions2d struct {
	W int `yaml:"w"` // Width
	H int `yaml:"h"` // Height
}

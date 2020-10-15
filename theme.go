package pterm

var (
	ThemeDefault = Theme{
		PrimaryStyle:            NewStyle(FgCyan),
		SecondaryStyle:          NewStyle(FgLightMagenta),
		HighlightStyle:          NewStyle(Bold, FgYellow),
		InfoMessageStyle:        NewStyle(FgLightCyan),
		InfoPrefixStyle:         NewStyle(FgBlack, BgCyan),
		SuccessMessageStyle:     NewStyle(FgGreen),
		SuccessPrefixStyle:      NewStyle(FgBlack, BgGreen),
		WarningMessageStyle:     NewStyle(FgYellow),
		WarningPrefixStyle:      NewStyle(FgBlack, BgYellow),
		ErrorMessageStyle:       NewStyle(FgLightRed),
		ErrorPrefixStyle:        NewStyle(FgBlack, BgLightRed),
		FatalMessageStyle:       NewStyle(FgLightRed),
		FatalPrefixStyle:        NewStyle(FgBlack, BgLightRed),
		DescriptionMessageStyle: NewStyle(FgWhite),
		DescriptionPrefixStyle:  NewStyle(FgLightWhite, BgDarkGray),
		ScopeStyle:              NewStyle(FgGray),
	}
)

// TODO: Add documentation

type Theme struct {
	PrimaryStyle            Style
	SecondaryStyle          Style
	HighlightStyle          Style
	InfoMessageStyle        Style
	InfoPrefixStyle         Style
	SuccessMessageStyle     Style
	SuccessPrefixStyle      Style
	WarningMessageStyle     Style
	WarningPrefixStyle      Style
	ErrorMessageStyle       Style
	ErrorPrefixStyle        Style
	FatalMessageStyle       Style
	FatalPrefixStyle        Style
	DescriptionMessageStyle Style
	DescriptionPrefixStyle  Style
	ScopeStyle              Style
}

func (t Theme) WithPrimaryStyle(style Style) Theme {
	t.PrimaryStyle = style
	return t
}

func (t Theme) WithSecondaryStyle(style Style) Theme {
	t.SecondaryStyle = style
	return t
}

func (t Theme) WithHighlightStyle(style Style) Theme {
	t.HighlightStyle = style
	return t
}

func (t Theme) WithInfoMessageStyle(style Style) Theme {
	t.InfoMessageStyle = style
	return t
}

func (t Theme) WithInfoPrefixStyle(style Style) Theme {
	t.InfoPrefixStyle = style
	return t
}

func (t Theme) WithSuccessMessageStyle(style Style) Theme {
	t.SuccessMessageStyle = style
	return t
}

func (t Theme) WithSuccessPrefixStyle(style Style) Theme {
	t.SuccessPrefixStyle = style
	return t
}

func (t Theme) WithWarningMessageStyle(style Style) Theme {
	t.WarningMessageStyle = style
	return t
}

func (t Theme) WithWarningPrefixStyle(style Style) Theme {
	t.WarningPrefixStyle = style
	return t
}

func (t Theme) WithErrorMessageStyle(style Style) Theme {
	t.ErrorMessageStyle = style
	return t
}

func (t Theme) WithErrorPrefixStyle(style Style) Theme {
	t.ErrorPrefixStyle = style
	return t
}

func (t Theme) WithFatalMessageStyle(style Style) Theme {
	t.FatalMessageStyle = style
	return t
}

func (t Theme) WithFatalPrefixStyle(style Style) Theme {
	t.FatalPrefixStyle = style
	return t
}

func (t Theme) WithDescriptionMessageStyle(style Style) Theme {
	t.DescriptionMessageStyle = style
	return t
}

func (t Theme) WithDescriptionPrefixStyle(style Style) Theme {
	t.DescriptionPrefixStyle = style
	return t
}

package types

type TOML struct {
	Clouds    []Cloud    `toml:"cloud"`
	Projects  []Project  `toml:"project"`
	Redirects []Redirect `toml:"redirect"`
	Files     []File     `toml:"file"`
}

type Redirect struct {
	ID          string `toml:"id" json:"id"`
	Title       string `toml:"title" json:"title"`
	Description string `toml:"description" json:"description"`
	Image       string `toml:"image" json:"image"`
	URL         string `toml:"url" json:"url"`
}

type Project struct {
	Title   string `toml:"title" json:"title"`
	Preview string `toml:"preview" json:"preview"`
	URL     string `toml:"url" json:"url"`
}

type File struct {
	ID   string `toml:"id" json:"id"`
	File string `toml:"file" json:"file"`
}

type Cloud struct {
	Image string `toml:"image" json:"image"`
}

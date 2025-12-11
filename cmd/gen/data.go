package main

func newTemplateData() map[string]any {
	return map[string]any{
		"SiteTitle": "Nicholas Cheng",
		"Links": []NavLink{
			{"/", "Home"},
		},
		"Posts": blogEntries,
	}
}

var blogEntries = []BlogEntries{
	{
		Title:         "On Leading a Team",
		Slug:          "on-leading-a-team",
		TemplatePath:  "pages/blog/on-leading-a-team.tmpl",
		PublishedDate: "2025-12-11",
	},
}

type NavLink struct {
	URL  string
	Name string
}

type BlogEntries struct {
	Title         string
	Slug          string
	TemplatePath  string
	PublishedDate string
}

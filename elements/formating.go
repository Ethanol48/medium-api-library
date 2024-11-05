package elements

import (
	"fmt"
	"os"
	"strings"
)

/* Markdown */

func (p P) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString(p.Content)
	return sb.String()

}

func (cb CodeBlock) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString("```\n")
	sb.WriteString(cb.Content)
	sb.WriteString("\n```")
	return sb.String()

}

func (i Image) ToMarkdown() string {
	var sb strings.Builder

	lin := i.Base + "/" + i.Id

	sb.WriteString(fmt.Sprintf("![%s](%s)", i.Alt, lin))
	return sb.String()

}

func (t Title) ToMarkdown() string {
	var sb strings.Builder

	switch t.Name {
	case "h1":
		sb.WriteString("# ")

	case "h2":
		sb.WriteString("## ")

	case "h3":
		sb.WriteString("### ")

	case "h4":
		sb.WriteString("#### ")
	}

	sb.WriteString(t.Content)

	return sb.String()

}

func (l List) ToMarkdown() string {
	var sb strings.Builder
	var prefix string

	switch l.Name {
	case "ol":
		prefix = "numeric"

	case "ul":
		prefix = ""

	}

	for i, v := range l.Compts {
		if prefix == "numeric" {
			sb.WriteString(fmt.Sprintf("%v. %s", i, v))
		} else {
			sb.WriteString(fmt.Sprintf("- %s", v))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (p PlaceHolder) ToMarkdown() string {
	return "null"
}

func (a Article) ToMarkdown() string {
	var sb strings.Builder

	for _, v := range a.Content {
		sb.WriteString(v.ToMarkdown())
		sb.WriteString("\n\n")

	}

	return sb.String()
}

func (a Article) ToMarkdownFile(path string) {

	err := os.WriteFile(path, []byte(a.ToMarkdown()), 0644) // 0644 is a common permission setting allowing reading for everyone and full write access to the owner of the file.
	if err != nil {
		// Handle the error here
		fmt.Println("Error writing to file: ", err)
	}

	fmt.Println("File created at: ", path)
}

func (b Blockquote) ToMarkdown() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("> %s", b.Content))
	return sb.String()
}

func (a Article) ToHtmlFile(path string) {

	err := os.WriteFile(path, []byte(a.ToHTML()), 0644) // 0644 is a common permission setting allowing reading for everyone and full write access to the owner of the file.
	if err != nil {
		// Handle the error here
		fmt.Println("Error writing to file: ", err)
	}

	fmt.Println("File created at: ", path)
}

/* HTML */

func (p P) ToHTML() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<p>%s</p>", p.Content))

	return sb.String()
}

func (cb CodeBlock) ToHTML() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<pre>%s</pre>", cb.Content))

	return sb.String()
}

func (b Blockquote) ToHTML() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<blockquote>%s</blockquote>", b.Content))

	return sb.String()
}

func (i Image) ToHTML() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<img src='%s/%s' alt='%s'>", i.Base, i.Id, i.Alt))
	return sb.String()
}

func (t Title) ToHTML() string {
	var sb strings.Builder

	switch t.Name {
	case "h1":
		sb.WriteString(fmt.Sprintf("<h1>%s</h1>", t.Content))
	case "h2":
		sb.WriteString(fmt.Sprintf("<h2>%s</h2>", t.Content))
	case "h3":
		sb.WriteString(fmt.Sprintf("<h3>%s</h3>", t.Content))
	case "h4":
		sb.WriteString(fmt.Sprintf("<h4>%s</h4>", t.Content))
	}

	return sb.String()
}

func (l List) ToHTML() string {
	var sb strings.Builder
	var tmp string
	var wrappingTags string

	switch l.Name {
	case "ol":
		wrappingTags = "<ol>%s</ol>"
	case "ul":
		wrappingTags = "<ul>%s</ul>"
	}

	for _, v := range l.Compts {
		tmp = tmp + fmt.Sprintf("<li>%s</li>", v)
	}

	sb.WriteString(fmt.Sprintf(wrappingTags, tmp))
	return sb.String()
}

func (p PlaceHolder) ToHTML() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<div>%s</div>", p.ToMarkdown()))
	return sb.String()
}

func (a Article) ToHTML() string {
	var sb strings.Builder

	for i := 0; i < len(a.Content); i++ {
		sb.WriteString(a.Content[i].ToHTML())
	}

	return sb.String()
}

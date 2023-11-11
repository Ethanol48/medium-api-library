package elements

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Ethanol48/medium-api/utilities"

	"github.com/gocolly/colly"
)

type Markdown interface {
	ToMarkdown() string
}

type Element interface {
	GetName() string
	ToMarkdown() string
	addData(map[string][]string)
}

type Article struct {
	Title     string
	Publisehd string
	ReadTime  string
	Content   []Element
}

type P struct {
	Name    string
	Content string
}

type Title struct {
	Name    string
	Content string
}

type Image struct {
	Name string
	Id   string
	Base string
	Alt  string
}

type List struct {
	Name   string
	Compts []string
}

type PlaceHolder struct {
	Elem colly.HTMLElement
}

// implementations GetName()

func (elem *PlaceHolder) GetName() string {
	return elem.Elem.Name
}

func (elem *List) GetName() string {
	return elem.Name
}

func (elem *Image) GetName() string {
	return elem.Name
}

func (elem *Title) GetName() string {
	return elem.Name
}

func (elem *P) GetName() string {
	return elem.Name
}

// addData

func (l *PlaceHolder) addData(m map[string][]string) {
	// elem.Content = m["content"]
}

func (l *List) addData(m map[string][]string) {
	for i := 0; i < len(m["content"]); i++ {
		l.Compts = append(l.Compts, m["content"][i])
	}
}

func (elem *Image) addData(m map[string][]string) {
	// elem.Name = m["name"]
	// elem.Src = m["src"]
	// elem.Alt = m["alt"]

}

func (elem *Title) addData(m map[string][]string) {
	// elem.Name = m["name"]
	// elem.Content = m["content"]
}

func (elem *P) addData(m map[string][]string) {
	// elem.Name = m["name"]
	// elem.Content = m["content"]
}

func (p P) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString(p.Content)
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

func createElement(elem *colly.HTMLElement) (Element, error) {

	switch elem.Name {

	case "img":
		return &Image{
			Name: elem.Name,
		}, nil

	case "ol":
		return &List{
			Name: elem.Name,
		}, nil

	case "ul":
		return &List{
			Name: elem.Name,
		}, nil

	case "h1":
		return &Title{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	case "h2":
		return &Title{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	case "h3":
		return &Title{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	case "h4":
		return &Title{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	case "p":
		return &P{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	default:

		return &PlaceHolder{
			Elem: *elem,
		}, errors.New("this element is not permited")

	}
}

func ExtractDataArticle(elem *colly.HTMLElement, idx int) Element {

	var e Element

	if idx == 0 {

		e, err := createElement(elem)
		if err != nil {
			fmt.Printf("There was an error with this element: %+v\n\n", e)
			panic(err)
		}

		return e
	}

	if elem.Name == "ul" || elem.Name == "ol" {
		// crear una lista y empezar appending
		// aqui deberia haber lista

		l, err := createElement(elem)
		if err != nil {
			fmt.Printf("There was an error with this element: %+v\n\n", e)
			panic(err)
		}

		tmp := make([]string, 0)

		elem.ForEach("li", func(_ int, h *colly.HTMLElement) {
			tmp = append(tmp, utilities.TrimMoreThanOneSpace(h.Text))
		})

		m := make(map[string][]string, 0)
		m["content"] = tmp

		l.addData(m)

		e = l

	} else if elem.Name == "li" {

		e = &PlaceHolder{}

	} else if elem.Name == "img" {
		// create image object

		siblings := elem.DOM.Siblings().Nodes

		// obtain link from first sibling
		base, id := utilities.ExtractLinkImg(siblings[0].Attr[0].Val)

		i := Image{
			Name: "img",
			Id:   id,
			Base: base,
			Alt:  elem.Attr("alt"),
		}

		e = &i

	} else {
		// tiene que ser o un header o un paragraph

		// fmt.Println("The element we are analizing: ")
		// fmt.Printf("%+v\n\n", elem)

		if elem.Name == "p" {

			p := P{
				Name:    elem.Name,
				Content: utilities.TrimMoreThanOneSpace(elem.Text),
			}

			e = &p

		} else {

			h := Title{
				Name:    elem.Name,
				Content: utilities.TrimMoreThanOneSpace(elem.Text),
			}

			e = &h
		}
	}

	return e

}

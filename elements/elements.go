package elements

import (
	"errors"
	"fmt"

	"github.com/Ethanol48/medium-api/utilities"

	"github.com/gocolly/colly"
)

type Element interface {
	GetName() string
	ToMarkdown() string
	ToHTML() string
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

/* GetName() */

func (elem *PlaceHolder) GetName() string { return elem.Elem.Name }
func (elem *List) GetName() string        { return elem.Name }
func (elem *Image) GetName() string       { return elem.Name }
func (elem *Title) GetName() string       { return elem.Name }
func (elem *P) GetName() string           { return elem.Name }

/* addData */

func (l *List) addData(m map[string][]string) {
	for i := 0; i < len(m["content"]); i++ {
		l.Compts = append(l.Compts, m["content"][i])
	}
}

func (l *PlaceHolder) addData(m map[string][]string) {}
func (elem *Image) addData(m map[string][]string)    {}
func (elem *Title) addData(m map[string][]string)    {}
func (elem *P) addData(m map[string][]string)        {}

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
		// it's another permitted element (header or paragragh)
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

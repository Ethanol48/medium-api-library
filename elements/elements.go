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
}

type Article struct {
	Title     string
	Published string
	ReadTime  string
	Content   []Element
}

// TODO: elements: 'code blocks', gists, parse through text for code, strong, italic text or links

type P struct {
	Name    string
	Content string
}

type Blockquote struct {
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
	Name string
	Elem colly.HTMLElement
}

/* GetName() */

func (elem *PlaceHolder) GetName() string { return elem.Name }
func (elem *List) GetName() string        { return elem.Name }
func (elem *Image) GetName() string       { return elem.Name }
func (elem *Title) GetName() string       { return elem.Name }
func (elem *P) GetName() string           { return elem.Name }
func (elem *Blockquote) GetName() string  { return elem.Name }

/* addData */

// func (l *List) addData(m map[string][]string) {
// 	for i := 0; i < len(m["content"]); i++ {
// 		l.Compts = append(l.Compts, m["content"][i])
// 	}
// }

// func (l *PlaceHolder) addData(m map[string][]string) {}
// func (elem *Image) addData(m map[string][]string)    {}
// func (elem *Title) addData(m map[string][]string)    {}
// func (elem *P) addData(m map[string][]string)        {}

func createElement(elem *colly.HTMLElement) (Element, error) {

	switch elem.Name {

	case "img":

		siblings := elem.DOM.Siblings().Nodes

		// obtain link from first sibling
		base, id := utilities.ExtractLinkImg(siblings[0].Attr[0].Val)

		i := Image{
			Name: "img",
			Id:   id,
			Base: base,
			Alt:  elem.Attr("alt"),
		}

		return &i, nil

	case "ol":

		l := List{
			Name:   elem.Name,
			Compts: make([]string, 0),
		}

		elem.ForEach("li", func(_ int, h *colly.HTMLElement) {
			l.Compts = append(l.Compts, utilities.TrimMoreThanOneSpace(h.Text))
		})

		return &l, nil

	case "ul":

		l := List{
			Name:   elem.Name,
			Compts: make([]string, 0),
		}

		elem.ForEach("li", func(_ int, h *colly.HTMLElement) {
			l.Compts = append(l.Compts, utilities.TrimMoreThanOneSpace(h.Text))
		})

		return &l, nil

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

	case "li":
		return &PlaceHolder{
			Name: "",
			Elem: *elem,
		}, nil

	case "blockquote":
		return &Blockquote{
			Name:    elem.Name,
			Content: utilities.TrimMoreThanOneSpace(elem.Text),
		}, nil

	default:

		return &PlaceHolder{
			Name: "",
			Elem: *elem,
		}, errors.New("this element is not permited")

	}
}

func ExtractDataArticle(elem *colly.HTMLElement) Element {

	e, err := createElement(elem)
	if err != nil {
		fmt.Printf("There was an error with this element: %+v\n\n", e)
		panic(err)
	}

	return e

}

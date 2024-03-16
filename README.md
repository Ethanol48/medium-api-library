## Unofficial Medium library

This project intends to create a simple library in Golang for 
reading data from medium.com website, from lists, users and tags.

The end goal is to create a simple REST API to obtain data programmatically

### How it came to be

This is a personal project created because I wanted to extract information about my
articles in medium.com to import it into my website.

So I was searching for an alternative from the official medium API as it is deprecated until I encountered an unofficial Medium API that had a lot of 
features, but I had some problems with it, first not open-sourced, you need to register on a website 
called RapidAPI, in the free plan you have 100 free calls, and after that it becomes <strong>4 cents</strong> per call!!, 
after some calls the bill could be astronomical. So I decided to create a module to easily extract data programmatically and use this to create an API that anyone can use for free.


## How it was built?

The way it works this "library" uses [colly]([google.com](https://pkg.go.dev/github.com/gocolly/colly@v1.2.0#section-readme)) as a way to get the HTML and parse through it with CSS selectors. In the case of medium.com the names of classes in the elements are inconsistent, so we need to rely on the structure of the HTML to identify the objects. This poses some problems for selecting certain structures in the articles. 

### Modules

The logic is distributed in different modules, dividing the responsibility of the code by the different webpages in medium.com `article`, `lists`, `user` and `elements`, the last one is used like a factory, we can use the method `elements.createElement(*colly.HTMLElement)` using the 'node' that we get with colly and outputs an object with the `Element` interface that we can use to convert the data into HTML or Markdown. 


## This is a Project under development 🔧

There are some features to be implemented in the library, like obtaining tags of the topics of the user and the articles.

## Collaboration

Are you a fanatic of open-source? And want to build something awesome?

PRs are welcome, docs, implementation, every help is welcomed.


</br>
</br>

## Functionality TODO:

  - [ ] Link Validation
  - [x] Articles
    - [x] Get content of Article
      - [x] Markdown
      - [x] HTML
      - [x] Support for nested elements (strong, italic)
    - [x] Metadata of Article
      - [x] Title
      - [x] Tags / Topics
    - [x] Images
      - [x] Download Images of Markdown file to local fs

  - [ ] Users
    - [ ] Top 10 Articles
    - [ ] Topic of User
    - [x] Metadata
  
  - [x] Lists
    - [x] Top 10 Articles

### API
  - [ ] API it-self 
    - [ ] Create routes
// Package markdown implements markdown parser and HTML renderer.
//
// It parses markdown into AST format which can be serialized to HTML
// (using Renderer) or possibly other formats (using alternate renderers).
//
//
// Convert markdown to HTML
//
// The simplest way to convert markdown text to HTML
//
//  md := []byte("## markdown document")
//  html := markdown.ToHTML(md, nil, nil)
//
// Customizing parsing and HTML rendering
//
// Youc can customize parser and HTML renderer:
//
//  md := []byte("markdown document")
//  extensions := markdown.CommonExtensions | markdown.AutoHeadingIDs
//  parser := markdown.NewParserWithExensions(extensions)
//  htmlParams := markdown.CommonHTMLFlags | markdown.HrefTargetBlank
//  renderer := markdown.NewRenderer(htmlParams)
//  html := markdown.ToHTML(md, parser, renderer)
//
// For a cmd-line tool see https://github.com/gomarkdown/markdown/tree/master/cmd/mdtohtml
package markdown

package main

import (
	"github.com/aquilax/tripcode"
	"github.com/gosimple/slug"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"time"
)

func hfLang(t string) string {
	return t
}

func hfTime(t time.Time) string {
	return t.Format("01.02.2006 15:04")
}

func hfSlug(s string) string {
	return slug.Make(s) + ".html"
}

func getTripcode(s string) string {
	return tripcode.Tripcode(s)
}

func getVote(t string) int {
	if t == "y" {
		return 1
	}
	if t == "n" {
		return -1
	}
	return 0
}

func inHoneypot(t string) bool {
	if len(t) > 0 {
		return true
	}
	return false
}

func renderText(t string) template.HTML {
	unsafe := blackfriday.MarkdownCommon([]byte(t))
	return template.HTML(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
}

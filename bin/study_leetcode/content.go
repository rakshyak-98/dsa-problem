package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

var (
	rePreBlock   = regexp.MustCompile(`(?s)<pre[^>]*>(.*?)</pre>`)
	reHTMLTag    = regexp.MustCompile(`<[^>]+>`)
	reBlankLine  = regexp.MustCompile(`\n{3,}`)
	reStrongOpen = regexp.MustCompile(`<strong[^>]*>`)
	reStrongClose = regexp.MustCompile(`</strong>`)
	reEmOpen     = regexp.MustCompile(`<em[^>]*>`)
	reEmClose    = regexp.MustCompile(`</em>`)
	reCodeOpen   = regexp.MustCompile(`<code[^>]*>`)
	reCodeClose  = regexp.MustCompile(`</code>`)
	reParaOpen   = regexp.MustCompile(`<p[^>]*>`)
	reParaClose  = regexp.MustCompile(`</p>`)
	reLiOpen     = regexp.MustCompile(`<li[^>]*>`)
	reLiClose    = regexp.MustCompile(`</li>`)
	reBr         = regexp.MustCompile(`<br\s*/?>`)
)

func fetchQuestionContent(slug string) (string, error) {
	var resp struct {
		Data struct {
			Question struct {
				Content string `json:"content"`
			} `json:"question"`
		} `json:"data"`
	}
	err := postGraphQL(gqlRequest{
		Query: `query getQuestionContent($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				content
			}
		}`,
		Variables:     map[string]any{"titleSlug": slug},
		OperationName: "getQuestionContent",
	}, &resp)
	if err != nil {
		return "", err
	}
	content := strings.TrimSpace(resp.Data.Question.Content)
	if content == "" {
		return "", fmt.Errorf("empty content for %s", slug)
	}
	return content, nil
}

func htmlToText(raw string) string {
	s := raw
	s = rePreBlock.ReplaceAllStringFunc(s, func(block string) string {
		inner := rePreBlock.FindStringSubmatch(block)
		if len(inner) < 2 {
			return block
		}
		text := stripTags(inner[1])
		text = html.UnescapeString(text)
		text = strings.TrimSpace(text)
		return "\n\n```\n" + text + "\n```\n\n"
	})

	s = reStrongOpen.ReplaceAllString(s, "**")
	s = reStrongClose.ReplaceAllString(s, "**")
	s = reEmOpen.ReplaceAllString(s, "*")
	s = reEmClose.ReplaceAllString(s, "*")
	s = reCodeOpen.ReplaceAllString(s, "`")
	s = reCodeClose.ReplaceAllString(s, "`")
	s = reLiOpen.ReplaceAllString(s, "\n- ")
	s = reLiClose.ReplaceAllString(s, "")
	s = reParaOpen.ReplaceAllString(s, "\n")
	s = reParaClose.ReplaceAllString(s, "\n")
	s = reBr.ReplaceAllString(s, "\n")
	s = strings.ReplaceAll(s, "&nbsp;", " ")

	s = stripTags(s)
	s = html.UnescapeString(s)
	s = reBlankLine.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s)
}

func stripTags(s string) string {
	return strings.TrimSpace(reHTMLTag.ReplaceAllString(s, ""))
}

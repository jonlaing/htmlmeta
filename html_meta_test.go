package htmlmeta

import (
	"io"
	"testing"
)

type MockPage struct {
	Content string
	done    bool
}

func NewMockPage(s string) MockPage {
	return MockPage{Content: s}
}

func (m *MockPage) Read(p []byte) (n int, err error) {
	if m.done {
		return 0, io.EOF
	}
	for i, b := range []byte(m.Content) {
		p[i] = b
	}
	m.done = true
	return len(m.Content), nil
}

func TestTitle(t *testing.T) {
	title := "foo bar"
	mp := NewMockPage("<html><head><title>" + title + "</title></head></html>")

	hm := Extract(&mp)

	if hm.Title != title {
		t.Error("Expected title to be", title, ", but was:", hm.Title)
	}
}

func TestDescription(t *testing.T) {
	desc := "foo bar"
	mp := NewMockPage("<html><head><meta property=\"description\" content=\"" + desc + "\"></head></html>")

	hm := Extract(&mp)

	if hm.Description != desc {
		t.Error("Expected desc to be", desc, ", but was:", hm.Description)
	}
}

func TestOGTitle(t *testing.T) {
	title := "foo bar"
	mp := NewMockPage("<html><head><meta property=\"og:title\" content=\"" + title + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGTitle != title {
		t.Error("Expected og:title to be", title, ", but was:", hm.OGTitle)
	}
}

func TestOGDescription(t *testing.T) {
	desc := "foo bar"
	mp := NewMockPage("<html><head><meta property=\"og:description\" content=\"" + desc + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGDescription != desc {
		t.Error("Expected og:description to be", desc, ", but was:", hm.OGDescription)
	}
}

func TestOGImage(t *testing.T) {
	image := "http://google.com/images/blah.jpg"
	mp := NewMockPage("<html><head><meta property=\"og:image\" content=\"" + image + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGImage != image {
		t.Error("Expected og:image to be", image, ", but was:", hm.OGImage)
	}
}

func TestOGAuthor(t *testing.T) {
	author := "jonlaing"
	mp := NewMockPage("<html><head><meta property=\"og:author\" content=\"" + author + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGAuthor != author {
		t.Error("Expected og:author to be", author, ", but was:", hm.OGAuthor)
	}
}

func TestOGPublisher(t *testing.T) {
	publisher := "jonlaing"
	mp := NewMockPage("<html><head><meta property=\"og:publisher\" content=\"" + publisher + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGPublisher != publisher {
		t.Error("Expected og:publisher to be", publisher, ", but was:", hm.OGPublisher)
	}
}

func TestOGSiteName(t *testing.T) {
	sitename := "Google"
	mp := NewMockPage("<html><head><meta property=\"og:site_name\" content=\"" + sitename + "\"></head></html>")

	hm := Extract(&mp)

	if hm.OGSiteName != sitename {
		t.Error("Expected og:site_name to be", sitename, ", but was:", hm.OGSiteName)
	}
}

func TestFullExtraction(t *testing.T) {
	title := "foobar"
	description := "boo far"
	ogTitle := "Foo Bar"
	ogDesc := "Boo Far"
	ogImage := "http://google.com/images/blah.jpg"
	ogAuthor := "Jon Laing"
	ogPublisher := "jonlaing"
	ogSiteName := "Google"

	mp := NewMockPage(`
	<html>
		<head>
			<title>` + title + `</title>
			<meta property="description" content="` + description + `">
			<meta property="og:title" content="` + ogTitle + `">
			<meta property="og:description" content="` + ogDesc + `">
			<meta property="og:image" content="` + ogImage + `">
			<meta property="og:author" content="` + ogAuthor + `">
			<meta property="og:publisher" content="` + ogPublisher + `">
			<meta property="og:site_name" content="` + ogSiteName + `">
		</head>
	</html>`)

	hm := Extract(&mp)

	if hm.Description != description {
		t.Error("Expected description to be", description, ", but was:", hm.Description)
	}
	if hm.OGTitle != ogTitle {
		t.Error("Expected og:title to be", ogTitle, ", but was:", hm.OGTitle)
	}
	if hm.OGDescription != ogDesc {
		t.Error("Expected og:description to be", ogDesc, ", but was:", hm.OGDescription)
	}
	if hm.OGImage != ogImage {
		t.Error("Expected og:image to be", ogImage, ", but was:", hm.OGImage)
	}
	if hm.OGAuthor != ogAuthor {
		t.Error("Expected og:author to be", ogAuthor, ", but was:", hm.OGAuthor)
	}
	if hm.OGPublisher != ogPublisher {
		t.Error("Expected og:publisher to be", ogPublisher, ", but was:", hm.OGSiteName)
	}
	if hm.OGSiteName != ogSiteName {
		t.Error("Expected og:site_name to be", ogSiteName, ", but was:", hm.OGSiteName)
	}
}

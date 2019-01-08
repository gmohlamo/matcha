// Generated by ego.
// DO NOT EDIT

//line views/home.ego:1

package views

import "fmt"
import "html"
import "io"
import "context"

import "github.com/blavkboy/matcha/views/components/headers"
import "github.com/blavkboy/matcha/views/components"
import "github.com/blavkboy/matcha/models"

func RenderHome(w io.Writer, u *models.User) {

//line views/home.ego:10
	_, _ = io.WriteString(w, "\n<head>\n")
//line views/home.ego:11

	produceHead(w, headers.HomeStyles, headers.HomeScripts)

//line views/home.ego:14
	_, _ = io.WriteString(w, "\n</head>\n<body>\n  <div class=\"columns\">\n    <div class=\"column is-one-quarter\">\n      ")
//line views/home.ego:18

	components.HomeSidePanel(w, u)

//line views/home.ego:21
	_, _ = io.WriteString(w, "\n    </div>\n    <div class=\"column is-three-quarter\">\n      <figure class=\"image is-128x128\">\n        <img src=\"https://bulma.io/images/placeholders/128x128.png\">\n      </figure>\n    </div>\n  </div>\n</body>\n")
//line views/home.ego:29
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

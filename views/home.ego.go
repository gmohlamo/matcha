// Generated by ego.
// DO NOT EDIT

//line views/home.ego:1
package views

import "fmt"
import "html"
import "io"
import "context"

import "github.com/gmohlamo/matcha/views/components/headers"
import "github.com/gmohlamo/matcha/views/components"
import "github.com/gmohlamo/matcha/models"

func RenderHome(w io.Writer, u *models.User) {

//line views/home.ego:10
	_, _ = io.WriteString(w, "\n<head>\n")
//line views/home.ego:11
	produceHead(w, headers.HomeStyles, headers.HomeScripts)

//line views/home.ego:14
	_, _ = io.WriteString(w, "\n</head>\n<body style=\"height:100%; width:100%;\">\n  ")
//line views/home.ego:16
	components.HomeNavbar(w, u)

//line views/home.ego:19
	_, _ = io.WriteString(w, "\n  <div class=\"columns\"> \n    <div class=\"column is-one-quarter\">\n      ")
//line views/home.ego:21
	components.HomeSidePanel(w, u)

//line views/home.ego:24
	_, _ = io.WriteString(w, "\n    </div>\n    <div class=\"column is-three-quarters\" id=\"rightColumn\">\n      ")
//line views/home.ego:26
	components.RenderProfileColumn(w, u)

//line views/home.ego:29
	_, _ = io.WriteString(w, "\n    </div>\n  </div>\n</body>\n")
//line views/home.ego:32
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

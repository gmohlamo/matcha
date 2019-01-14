// Generated by ego.
// DO NOT EDIT

//line views/index.ego:1
package views

import "fmt"
import "html"
import "io"
import "context"

import (
	"github.com/blavkboy/matcha/views/components"
	"github.com/blavkboy/matcha/views/components/headers"
)

func RenderIndex(w io.Writer) {

//line views/index.ego:11
	_, _ = io.WriteString(w, "\n<head>\n")
//line views/index.ego:12
	produceHead(w, headers.IndexStyles, headers.IndexScripts)

//line views/index.ego:15
	_, _ = io.WriteString(w, "\n</head>\n<body>\n    <section class=\"hero is-primary is-bold is-fullheight\" style=\"align-items: center; justify-content: center;\">\n        <div class=\"hero-body\">\n        <div class=\"container\" style=\"text-align: center;\">\n            <h1 class=\"title is-large\">\n            Lovr\n            </h1>\n            <h2 class=\"subtitle is-large\">\n            Dating App\n            </h2>\n            <div class=\"container\" id=\"landing_modals\">\n              ")
//line views/index.ego:27
	components.LoginModal(w)
	components.RegisterModal(w)

//line views/index.ego:31
	_, _ = io.WriteString(w, "\n                <div class=\"buttons is-centered\">\n                    <center><span class=\"button is-small is-primary is-inverted is-small\" style=\"margin-right: 5px;\" id=\"login\">Login</span></center>\n                    <center><span class=\"button is-small is-primary is-inverted is-small\" style=\"margin-left: 5px;\" id=\"register\">Register</span></center>\n                </div>\n            </div>\n        </div>\n        </div>\n    </section>\n</body>\n")
//line views/index.ego:40
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

// Generated by ego.
// DO NOT EDIT

//line views/index.ego:1

package views

import "fmt"
import "html"
import "io"
import

//line views/index.ego:6
"context"

func RenderIndex(w io.Writer) {

	_, _ = io.WriteString(w, "\n<!DOCTYPE html>\n<html>\n<head>\n  ")
//line views/index.ego:9

	ProduceHead(w)

//line views/index.ego:12
	_, _ = io.WriteString(w, "\n</head>\n  <body>\n    <section class=\"hero is-primary is-bold is-fullheight\" style=\"align-items: center; justify-content: center;\">\n      <div class=\"hero-body\">\n        <div class=\"container\" style=\"text-align: center;\">\n          <h1 class=\"title is-large\">\n            Lovr\n          </h1>\n          <h2 class=\"subtitle is-large\">\n            Dating App\n          </h2>\n          <div class=\"container\">\n            <div class=\"modal\">\n              <div class=\"modal-background\"></div>\n              <div class=\"modal-card animated zoomIn\" style=\"width: 80%;\">\n              <header class=\"modal-card-head\">\n                <p class=\"modal-card-title\">Lovr Login</p>\n                <button class=\"delete\" aria-label=\"close\"></button>\n              </header>\n              <section class=\"modal-card-body\">\n                <input class=\"input\" id=\"log_username\" type=\"text\" placeholder=\"Username/Email\">\n                <span></span>\n                <input class=\"input\" id=\"log_password\" type=\"password\" placeholder=\"Password\">\n              </section>\n              <footer class=\"modal-card-foot\">\n                <button class=\"button is-primary\">Login</button>\n              </footer>\n            </div>\n          </div>\n            <div class=\"buttons is-centered\">\n              <center><span class=\"button is-small is-primary is-inverted is-small\" style=\"margin-right: 5px;\" id=\"login\">Login</span></center>\n              <center><span class=\"button is-small is-primary is-inverted is-small\" style=\"margin-left: 5px;\" id=\"register\">Register</span></center>\n            </div>\n          </div>\n        </div>\n      </div>\n    </section>\n  </body>\n</html>\n")
//line views/index.ego:51
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

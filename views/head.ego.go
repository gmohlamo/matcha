// Generated by ego.
// DO NOT EDIT

//line views/head.ego:1

package views

import "fmt"
import "html"
import "io"
import

//line views/head.ego:6
"context"

func ProduceHead(w io.Writer) {

	_, _ = io.WriteString(w, "\n<meta charset=\"utf-8\">\n<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\n<meta name=\"description\" content=\"a dating app\">\n<meta name=\"author\" content=\"blavkboy\">\n<title>My Site</title>\n\n<link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:8080/static/css/bulma.min.css\">\n<link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:8080/static/css/custom.css\">\n\n<script defer src=\"http://localhost:8080/static/js/all.js\"></script>\n\n<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->\n<!--[if lt IE 9]>\n  <script src=\"https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js\"></script>\n  <script src=\"https://oss.maxcdn.com/respond/1.4.2/respond.min.js\"></script>\n<![endif]-->\n")
//line views/head.ego:24
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

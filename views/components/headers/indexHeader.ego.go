// Generated by ego.
// DO NOT EDIT

//line indexHeader.ego:1

package headers

import "fmt"
import "html"
import "io"
import

//line indexHeader.ego:6
"context"

func IndexStyles(w io.Writer) {

	_, _ = io.WriteString(w, "\n<link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:8080/static/css/custom.css\">\n")
//line indexHeader.ego:7
}

func IndexScripts(w io.Writer) {

//line indexHeader.ego:11
	_, _ = io.WriteString(w, "\n<script src=\"http://localhost:8080/static/js/index.js\"></script>\n")
//line indexHeader.ego:12
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

// Generated by ego.
// DO NOT EDIT

//line views/components/homeSidePanel.ego:1
package components

import "fmt"
import "html"
import "io"
import "context"

import "github.com/blavkboy/matcha/models"

func HomeSidePanel(w io.Writer, u *models.User) {

//line views/components/homeSidePanel.ego:8
	_, _ = io.WriteString(w, "\n<nav class=\"panel\">\n  <p class=\"panel-heading\">\n    Welcome ")
//line views/components/homeSidePanel.ego:10
	fmt.Fprint(w, u.Username)
//line views/components/homeSidePanel.ego:11
	_, _ = io.WriteString(w, "\n  </p>\n  <p class=\"panel-block\">\n    <span class=\"panel-icon\" style=\"color: orangered;\">\n      <i class=\"fas fa-fire\"></i>\n    </span>\n    Fame rating: ")
//line views/components/homeSidePanel.ego:16
	fmt.Fprint(w, u.Profile.Fame)
//line views/components/homeSidePanel.ego:17
	_, _ = io.WriteString(w, "\n  </p>\n  <p class=\"panel-block\">\n    <span class=\"panel-icon\" style=\"color: red;\">\n      <i class=\"fas fa-heart\"></i>\n    </span>\n    Likes:  ")
//line views/components/homeSidePanel.ego:22
	fmt.Fprintf(w, "%d", u.Profile.Likes)
//line views/components/homeSidePanel.ego:23
	_, _ = io.WriteString(w, "\n  </p>\n  <p class=\"panel-block\">\n    <span class=\"panel-icon\">\n      <i class=\"far fa-eye\"></i>\n    </span>\n    Vists: ")
//line views/components/homeSidePanel.ego:28
	fmt.Fprintf(w, "%d", u.Profile.Visits)
//line views/components/homeSidePanel.ego:29
	_, _ = io.WriteString(w, "\n  </p>\n</nav>\n")
//line views/components/homeSidePanel.ego:31
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

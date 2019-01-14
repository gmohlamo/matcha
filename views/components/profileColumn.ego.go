// Generated by ego.
// DO NOT EDIT

//line views/components/profileColumn.ego:1
package components

import "fmt"
import "html"
import "io"
import "context"

import "github.com/blavkboy/matcha/models"

func RenderProfileColumn(w io.Writer, u *models.User) {

//line views/components/profileColumn.ego:8
	_, _ = io.WriteString(w, "\n<div class=\"container\" id=\"main_column\" style=\"grid-row: 1; grid-column: 1;\">\n  <!-- home component -->\n  <div class=\"container\" id=\"picture\" style=\"width: 150px;\">\n  ")
//line views/components/profileColumn.ego:11
	if u.Profile.Propic == "" {
//line views/components/profileColumn.ego:12
		_, _ = io.WriteString(w, "\n    <img src=\"http://localhost:8080/static/defaultpic.jpg\" alt=\"profile picture\">\n  ")
//line views/components/profileColumn.ego:13
	} else {
//line views/components/profileColumn.ego:14
		_, _ = io.WriteString(w, "\n    <img src=\"")
//line views/components/profileColumn.ego:14
		fmt.Fprint(w, u.Profile.Propic)
//line views/components/profileColumn.ego:14
		_, _ = io.WriteString(w, "\" id=\"propic\" style=\"width: 150px;\">\n  ")
//line views/components/profileColumn.ego:15
	}
//line views/components/profileColumn.ego:16
	_, _ = io.WriteString(w, "\n    <div class=\"field\">\n      <div class=\"file is-success is-small is-rounded\">\n        <label class=\"file-label\">\n          <input class=\"file-input\" type=\"file\" name=\"resume\" id=\"propic\">\n          <span class=\"file-cta\">\n            <span class=\"file-icon\">\n              <i class=\"fas fa-upload\"></i>\n            </span>\n            <span class=\"file-label\">\n              Profile Picture\n            </span>\n          </span>\n        </label>\n      </div>\n    </div>\n    <span class=\"button is-success is-small\" id=\"picsubmit\">Change Profile Picture</span>\n  </div>\n  <div style=\"grid-row:1/2; grid-column: 2;\">\n    <form>\n      <div class=\"field\">\n        <label class=\"label\">First Name</label>\n        <div class=\"control\">\n          <input id=\"fname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:38
	fmt.Fprint(w, u.Fname)
//line views/components/profileColumn.ego:38
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div style=\"field\">\n        <label class=\"label\">Last Name</label>\n        <div class=\"control\">\n          <input id=\"lname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:45
	fmt.Fprint(w, u.Lname)
//line views/components/profileColumn.ego:45
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div style=\"field\">\n        <label class=\"label\">Username</label>\n        <div class=\"control\">\n          <input id=\"uname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:52
	fmt.Fprint(w, u.Username)
//line views/components/profileColumn.ego:52
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Email</label>\n        <div class=\"control\">\n          <input id=\"email\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:59
	fmt.Fprint(w, u.Email)
//line views/components/profileColumn.ego:59
	_, _ = io.WriteString(w, "\">\n        </div>\n        <p class=\"help\">Changing your email will require you to confirm your email in order to use key features again</p>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Gender</label>\n        <div class=\"control\">\n          <div class=\"select\">\n            <select id=\"gender\" ")
//line views/components/profileColumn.ego:68
	if u.Sex == "" {
		fmt.Fprint(w, "required")
	}
//line views/components/profileColumn.ego:70
	_, _ = io.WriteString(w, ">\n            ")
//line views/components/profileColumn.ego:71
	sex := [4]string{"Select", "Male", "Female", "Other"}
	for _, val := range sex {
//line views/components/profileColumn.ego:74
		_, _ = io.WriteString(w, "\n                <option>")
//line views/components/profileColumn.ego:74
		fmt.Fprint(w, val)
//line views/components/profileColumn.ego:74
		_, _ = io.WriteString(w, "</option>\n            ")
//line views/components/profileColumn.ego:75
	}
//line views/components/profileColumn.ego:76
	_, _ = io.WriteString(w, "\n            </select>\n          </div>\n        </div>\n      </div>\n      <p class=\"subtitle is-5\" id=\"sexTitle\">")
//line views/components/profileColumn.ego:80
	fmt.Fprint(w, u.Sex)
//line views/components/profileColumn.ego:80
	_, _ = io.WriteString(w, "</p>\n      \n      <div class=\"field\">\n        <label class=\"label\">Who are interested in meeting?</label>\n        <div class=\"control\">\n          <div class=\"select\">\n            <select id=\"orientation\" ")
//line views/components/profileColumn.ego:86
	if u.Profile.Orientation == "" {
		fmt.Fprint(w, "required")
	}
//line views/components/profileColumn.ego:88
	_, _ = io.WriteString(w, ">\n              ")
//line views/components/profileColumn.ego:89
	orientation := [4]string{"Select", "Men", "Women", "Both"}
	for _, val := range orientation {
//line views/components/profileColumn.ego:92
		_, _ = io.WriteString(w, "\n                  <option>")
//line views/components/profileColumn.ego:92
		fmt.Fprint(w, val)
//line views/components/profileColumn.ego:92
		_, _ = io.WriteString(w, "</option>\n              ")
//line views/components/profileColumn.ego:93
	}
//line views/components/profileColumn.ego:94
	_, _ = io.WriteString(w, "\n            </select>\n          </div>\n        </div>\n      </div>\n      <p class=\"subtitle is-5\" id=\"orientationTitle\">")
//line views/components/profileColumn.ego:98
	fmt.Fprint(w, u.Profile.Orientation)
//line views/components/profileColumn.ego:98
	_, _ = io.WriteString(w, "</p>\n      \n      <div class=\"field\">\n        <label class=\"label\">Interests</label>\n        <div class=\"control\">\n          <textarea class=\"textarea is-medium\" placeholder=\"Enter your interests in here like you would a series of twitter hash tags\"></textarea>\n        </div>\n      </div>\n      <div class=\"buttons\">\n        <span class=\"button is-success\" id=\"submit\">Submit Changes</span>\n        <span class=\"button is-success\" id=\"pwChange\">Change Password</span>\n      </div>\n    </div>\n  </form>\n  <!-- home component -->\n</div>\n<script type=\"text/javascript\" src=\"/static/js/profileSelect.js\"></script>\n")
//line views/components/profileColumn.ego:115
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

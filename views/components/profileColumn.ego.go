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
	_, _ = io.WriteString(w, "\n<div class=\"container\" id=\"main_column\" style=\"grid-row: 1; grid-column: 1;\">\n  <!-- home component -->\n  <div class=\"container\" id=\"picture\" style=\"width: 150px;\">\n    <img src=\"http://localhost:8080/static/defaultpic.jpg\" alt=\"profile picture\">\n    <div class=\"field\">\n      <div class=\"file is-success is-small is-rounded\">\n        <label class=\"file-label\">\n          <input class=\"file-input\" type=\"file\" name=\"resume\" id=\"propic\">\n          <span class=\"file-cta\">\n            <span class=\"file-icon\">\n              <i class=\"fas fa-upload\"></i>\n            </span>\n            <span class=\"file-label\">\n              Profile Picture\n            </span>\n          </span>\n        </label>\n      </div>\n    </div>\n    <span class=\"button is-success is-small\">Change Profile Picture</span>\n  </div>\n  <div style=\"grid-row:1/2; grid-column: 2;\">\n    <form>\n      <div class=\"field\">\n        <label class=\"label\">First Name</label>\n        <div class=\"control\">\n          <input id=\"fname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:34
	fmt.Fprint(w, u.Fname)
//line views/components/profileColumn.ego:34
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div style=\"field\">\n        <label class=\"label\">Last Name</label>\n        <div class=\"control\">\n          <input id=\"lname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:41
	fmt.Fprint(w, u.Lname)
//line views/components/profileColumn.ego:41
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div style=\"field\">\n        <label class=\"label\">Username</label>\n        <div class=\"control\">\n          <input id=\"uname\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:48
	fmt.Fprint(w, u.Username)
//line views/components/profileColumn.ego:48
	_, _ = io.WriteString(w, "\">\n        </div>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Email</label>\n        <div class=\"control\">\n          <input id=\"email\" class=\"input is-medium\" type=\"text\" placeholder=\"")
//line views/components/profileColumn.ego:55
	fmt.Fprint(w, u.Email)
//line views/components/profileColumn.ego:55
	_, _ = io.WriteString(w, "\">\n        </div>\n        <p class=\"help\">Changing your email will require you to confirm your email in order to use key features again</p>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Gender</label>\n        <div class=\"control\">\n          <div class=\"select\">\n            <select id=\"gender\" ")
//line views/components/profileColumn.ego:64
	if u.Sex == "" {
		fmt.Fprint(w, "required")
	}
//line views/components/profileColumn.ego:66
	_, _ = io.WriteString(w, ">\n            ")
//line views/components/profileColumn.ego:67

	sex := [4]string{"Select", "Male", "Female", "Other"}
	for _, val := range sex {
//line views/components/profileColumn.ego:70
		_, _ = io.WriteString(w, "\n                <option>")
//line views/components/profileColumn.ego:70
		fmt.Fprint(w, val)
//line views/components/profileColumn.ego:70
		_, _ = io.WriteString(w, "</option>\n            ")
//line views/components/profileColumn.ego:71
	}
//line views/components/profileColumn.ego:72
	_, _ = io.WriteString(w, "\n            </select>\n            <p class=\"subtitle is-5\" id=\"sexTitle\">")
//line views/components/profileColumn.ego:73
	fmt.Fprint(w, u.Sex)
//line views/components/profileColumn.ego:73
	_, _ = io.WriteString(w, "</p>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Who are interested in meeting?</label>\n        <div class=\"control\">\n          <div class=\"select\">\n            <select id=\"orientation\" ")
//line views/components/profileColumn.ego:82
	if u.Profile.Orientation == "" {
		fmt.Fprint(w, "required")
	}
//line views/components/profileColumn.ego:84
	_, _ = io.WriteString(w, ">\n              ")
//line views/components/profileColumn.ego:85

	orientation := [4]string{"Select", "Men", "Women", "Both"}
	for _, val := range orientation {
//line views/components/profileColumn.ego:88
		_, _ = io.WriteString(w, "\n                  <option>")
//line views/components/profileColumn.ego:88
		fmt.Fprint(w, val)
//line views/components/profileColumn.ego:88
		_, _ = io.WriteString(w, "</option>\n              ")
//line views/components/profileColumn.ego:89
	}
//line views/components/profileColumn.ego:90
	_, _ = io.WriteString(w, "\n            </select>\n            <p class=\"subtitle is-5\" id=\"orientationTitle\">")
//line views/components/profileColumn.ego:91
	fmt.Fprint(w, u.Profile.Orientation)
//line views/components/profileColumn.ego:91
	_, _ = io.WriteString(w, "</p>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"field\">\n        <label class=\"label\">Interests</label>\n        <div class=\"control\">\n          <textarea class=\"textarea is-medium\" placeholder=\"Enter your interests in here like you would a series of twitter hash tags\"></textarea>\n        </div>\n      </div>\n      <div class=\"buttons\">\n        <span class=\"button is-success\" id=\"submit\">Submit Changes</span>\n        <span class=\"button is-success\" id=\"pwChange\">Change Password</span>\n      </div>\n    </div>\n  </form>\n  <!-- home component -->\n</div>\n<script type=\"text/javascript\" src=\"/static/js/profileSelect.js\"></script>\n")
//line views/components/profileColumn.ego:111
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

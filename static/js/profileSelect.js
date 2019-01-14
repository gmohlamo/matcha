let submit = document.getElementById("submit");
let form = document.querySelector("form");
var submission;
let reader = new FileReader();
let propic = document.getElementById("propic");
let newPic;
let picSubmit = document.getElementById("picsubmit");
//when clicking the submit button all the values from the inputs
//should be collected and packaged so that they can be sent to the server.
submit.onclick = function() {
  submission = {
    "type": "command",
    "commandType": "profile",
    "pform": {
      "fname": form.elements[0].value,
      "lname": form.elements[1].value,
      "uname": form.elements[2].value,
      "email": form.elements[3].value,
      "gender": form.elements[4].selectedOptions[0].value,
      "orientation": form.elements[5].selectedOptions[0].value,
      "interests": form.elements[6].value.split(" ")
    }
  }
  let subs = JSON.stringify(submission);
  ws.send(subs);
}

function resolveMessage() {
  if (submission["type"] == "command") {
    if (submission["commandType"] == "propic") {
      document.getElementById("propic").src = submission.command;
    } else if (submission["commandType"] == "profile") {
      if (submission["pform"]["fname"] != "" && submission["pform"]["fname"].length > 0) {
        form.elements[0].placeholder = submission["pform"]["fname"];
      }
      if (submission["pform"]["lname"] != "" && submission["pform"]["lname"].length > 0) {
        form.elements[1].placeholder = submission["pform"]["lname"];
      }
      if (submission["pform"]["uname"] != "" && submission["pform"]["uname"].length > 0) {
        form.elements[2].placeholder = submission["pform"]["uname"];
        let panel = document.querySelector("nav.panel").firstElementChild;
        panel.innerText = "Welcome " + submission["pform"]["uname"];
      }
      if (submission["pform"]["email"] != "" && submission["pform"]["email"].length > 0) {
        form.elements[3].placeholder = submission["pform"]["email"];
      }
      if (submission["pform"]["gender"] != "" && submission["pform"]["gender"].length > 0) {
        let gender = document.getElementById("genderSub");
        gender.innerText = submission["pform"]["gender"]
      }
      if (submission["pform"]["orientation"] != "" && submission["pform"]["orientation"].length > 0) {
        let orientation = document.getElementById("orientationSub");
        orientation.innerText = submission["pform"]["orientation"]
      }
      if (submission["pform"]["interests"] > 0 && submission["pform"]["interests"][0] != "") {
        form.elements[6].innerText = submission["pform"]["interests"];
      }
    }
  }
}

propic.onchange = function() {
  reader.readAsDataURL(propic.files[0]);
}

picsubmit.onclick = function() {
  submission = {
    "type": "command",
    "commandType": "propic",
    "command": reader.result
  }
  let subs = JSON.stringify(submission);
  ws.send(subs);
}

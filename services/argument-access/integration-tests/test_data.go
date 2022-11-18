package integration_tests

import protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"

var subject = &protoOut.CreateSubjectRequest{
	Subject: &protoOut.Subject{
		Body: "X1",
	},
}

var subject2 = &protoOut.CreateSubjectRequest{
	Subject: &protoOut.Subject{
		Id:   "potato",
		Body: "X2",
	},
}

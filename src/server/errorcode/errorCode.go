package errorCode

import "errors"

var USEREXISTERROR = errors.New("the user exist")
var USERNOEXISTERROR = errors.New("the user no exist")
var USERLOGINERROR = errors.New("the user login fail")

var TCPCREATEERROR = errors.New("tcp connect error")
var TCPCLOSEERROR = errors.New("tcp close error")
var TCPACCEPTERROR = errors.New("tcp accept error")
var TCPREADERROR = errors.New("tcp read error")

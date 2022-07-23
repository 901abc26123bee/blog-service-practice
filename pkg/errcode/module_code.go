package errcode

import ()

var (
	ErrorGetTagListFail = NewError(20010001, "Get Tag List Fail")
	ErrorCreateTagFail  = NewError(20010002, "Create Tag Fail")
	ErrorUpdateTagFail  = NewError(20010003, "Update Tag Fail")
	ErrorDeleteTagFail  = NewError(20010004, "Delete Tag Fail")
	ErrorCountTagFail   = NewError(20010005, "Count Tag Fail")
)

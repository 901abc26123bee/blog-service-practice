package errcode

import ()

var (
	ErrorGetTagListFail = NewError(20010001, "Fail to Get Tag List")
	ErrorCreateTagFail  = NewError(20010002, "Fail to Create Tag")
	ErrorUpdateTagFail  = NewError(20010003, "Fail to Update Tag")
	ErrorDeleteTagFail  = NewError(20010004, "Fail to Delete Tag")
	ErrorCountTagFail   = NewError(20010005, "Fail to Count Tag")
	ErrorUploadFileFail = NewError(20030001, "Fail to Upload File")
)

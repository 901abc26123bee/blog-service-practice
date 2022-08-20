package errcode

import ()

var (
	ErrorGetTagListFail = NewError(20010001, "Fail to Get Tag List")
	ErrorCreateTagFail  = NewError(20010002, "Fail to Create Tag")
	ErrorUpdateTagFail  = NewError(20010003, "Fail to Update Tag")
	ErrorDeleteTagFail  = NewError(20010004, "Fail to Delete Tag")
	ErrorCountTagFail   = NewError(20010005, "Fail to Count Tag")

	ErrorGetArticleFail    = NewError(20020001, "Fail to get an Article")
	ErrorGetArticlesFail   = NewError(20020002, "Fail to get Articles")
	ErrorCreateArticleFail = NewError(20020003, "Fail to create an Article")
	ErrorUpdateArticleFail = NewError(20020004, "Fail to update an Article")
	ErrorDeleteArticleFail = NewError(20020005, "Fail to delete an Article")

	ErrorUploadFileFail = NewError(20030001, "Fail to Upload File")
)

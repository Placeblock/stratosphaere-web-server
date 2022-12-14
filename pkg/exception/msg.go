package exception

var Messages = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "Invalid Parameters",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Invalid token",
	ERROR_AUTH_INVALID_CREDENTIALS: "Invalid credentials",
	ERROR_AUTH_TOKEN_FAIL:          "Error occured while trying to generate Token",
	ERROR_ARTICLE_NOT_EXIST:        "Article doesn't exist",
	ERROR_ARTICLE_FAIL_CREATE:      "Error occured while trying to create article",
	ERROR_ARTICLE_FAIL_DELETE:      "Error occured while trying to delete article",
	ERROR_ARTICLE_FAIL_EDIT:        "Error occured while trying to edit article",
	ERROR_ARTICLE_FAIL_PUBLISH:     "Error occured while trying to publish article",
	ERROR_ARTICLE_FAIL_REVERT:      "Error occured while trying to revert article",
	ERROR_ARTICLES_FAIL_GET:        "Error occured while trying to get articles",
	ERROR_ARTICLE_FAIL_GET:         "Error occured while trying to get article",
	ERROR_ARTICLE_FAIL_CHECK_EXIST: "Error occured while trying to check if article exists",
	ERROR_ARTICLE_FAIL_COUNT:       "Error occured while trying to count articles",
	ERROR_EMAIL_ALREADY_EXIST:      "This email already exists",
	ERROR_EMAIL_NOT_EXIST:          "This email doesn't exist",
	ERROR_AUTH_TOKEN_MISSING:       "No token provided in Authorization Header",
	ERROR_ARTICLE_DELETE_NOT_OWNER: "This is not your article",
	ERROR_IMAGE_OPEN:               "Error occured while trying to open file",
	ERROR_IMAGE_READ:               "Error occurde while trying to read file",
	ERROR_IMAGE_TOO_LARGE:          "The provided file is too large",
	ERROR_IMAGE_INVALID_TYPE:       "Only Images are supported",
	ERROR_IMAGE_DECODE:             "Error occured while trying to decode file",
	ERROR_IMAGE_SAVE:               "Error occured while trying to save file",
	ERROR_IMAGE_DELETE:             "Error occured while trying to delete image",
	ERROR_LIVE_DATA_GET:            "Error occured while trying to get live data",
}

func GetMsg(code int) string {
	msg, ok := Messages[code]
	if ok {
		return msg
	}
	return Messages[ERROR]
}

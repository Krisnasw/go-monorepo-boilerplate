package exceptions

import "errors"

var InternalServerError = errors.New("Internal Server Error")
var PasswordNotMatch = errors.New("Password not match")
var DataAlreadyExist = errors.New("Data already exists")
var DataCreateFailed = errors.New("Create data failed")
var DataUpdateFailed = errors.New("Update data failed")
var DataDeleteFailed = errors.New("Delete data failed")
var DataNotFound = errors.New("Data not found")

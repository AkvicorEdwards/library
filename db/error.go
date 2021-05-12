package db

import "errors"

var ErrorNoUser = errors.New("no user")
var ErrorNoBook = errors.New("no book")

var ErrorUnfinished = errors.New("last read unfinished")
var ErrorFinished = errors.New("last read finished")

var ErrorBidNotValid = errors.New("bid not valid")

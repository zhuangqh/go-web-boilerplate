package db

// database operation error
type OPError struct {
    Code int
    Err error
}

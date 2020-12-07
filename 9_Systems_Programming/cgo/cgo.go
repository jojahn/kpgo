package main

// // #cgo LDFLAGS: -lsqlite3
// #include <stdio.h>
// #include <stdlib.h>
// #include "./sqlite-amalgamation-3330000/sqlite3.h"
// #include "./sqlite-amalgamation-3330000/sqlite3ext.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(C.SQLITE_VERSION)

	var db *C.sqlite3
	name := C.CString("./db.sqlite")
	defer C.free(unsafe.Pointer(name))

	rc := C.sqlite3_open(name, &db)
	defer C.sqlite3_close(db)

	if rc != C.SQLITE_OK {
		panic(rc)
	}

	// statement := ""
	// exec(statement, db)
}

func exec(statement string, db *C.sqlite3) error {
	cStatement := C.CString(statement)
	defer C.free(unsafe.Pointer(cStatement))

	rc := C.sqlite3_exec(db, cStatement, nil, nil, nil)

	if rc != C.SQLITE_OK {
		return fmt.Errorf("exec error")
	}

	return nil
}
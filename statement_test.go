// +build go1.6

package sqlmock

import (
	"errors"
	"fmt"
	"testing"
)

func TestExpectedPreparedStatemtCloseError(t *testing.T) {
	conn, mock, err := New()
	if err != nil {
		t.Fatal("failed to open sqlmock database:", err)
	}

	mock.ExpectBegin()
	want := errors.New("STMT ERROR")
	mock.ExpectPrepare("SELECT").WillReturnCloseError(want)

	txn, err := conn.Begin()
	if err != nil {
		t.Fatal("unexpected error while opening transaction:", err)
	}

	stmt, err := txn.Prepare("SELECT")
	if err != nil {
		t.Fatal("unexpected error while preparing a statement:", err)
	}

	if err := stmt.Close(); err != want {
		t.Fatalf("Got = %v, want = %v", err, want)
	}
}

func TestMatchOrderedExpectPrepare(t *testing.T) {
	//firstQuery := `select x from A`
	//thirdQuery := `delete from A where x = ?`
	secondQuery := `select x from A where y = ?`

	//t.Run("succeed when prepared in the right order", func(t *testing.T) {
	//	db, mock, err := New()
	//	if err != nil {
	//		t.Fatal("failed to open sqlmock database:", err)
	//	}
	//
	//	// Ensure that this is set to true
	//	mock.MatchExpectationsInOrder(true)
	//
	//	mock.ExpectPrepare(fmt.Sprintf("select .+ from A"))
	//	mock.ExpectPrepare(fmt.Sprintf("select .+ from A where .+"))
	//	//mock.ExpectPrepare(fmt.Sprintf("delete from A where .+"))
	//
	//	_, err = db.Prepare(firstQuery)
	//	if err != nil {
	//		t.Fatal("failed to prepare the statement regardless of being in the right order:", err)
	//	}
	//
	//	_, err = db.Prepare(secondQuery)
	//	if err != nil {
	//		t.Fatal("failed to prepare the statement regardless of being in the right order:", err)
	//	}
	//})

	t.Run("fail when prepared in the wrong order", func(t *testing.T) {
		db, mock, err := New()
		if err != nil {
			t.Fatal("failed to open sqlmock database:", err)
		}

		// Ensure that this is set to true
		mock.MatchExpectationsInOrder(true)

		mock.ExpectPrepare(fmt.Sprintf("select .+ from A"))
		mock.ExpectPrepare(fmt.Sprintf("select .+ from A where .+"))
		//mock.ExpectPrepare(fmt.Sprintf("delete from A where .+"))

		_, err = db.Prepare(secondQuery)
		if err == nil {
			t.Fatal("successfully prepared a statement out of order:")
		}
	})

}

//func TestMatchUnorderedExpectPrepare(t *testing.T) {
//	db, mock, err := New()
//	if err != nil {
//		t.Fatal("failed to open sqlmock database:", err)
//	}
//
//	mock.MatchExpectationsInOrder(false)
//
//	mock.ExpectPrepare(fmt.Sprintf("select (.+) from A"))
//	mock.ExpectPrepare(fmt.Sprintf("select (.+) from A where (.+)"))
//
//	firstQuery := `select x from A`
//	secondQuery := `select x from A where y = ?`
//
//	_, err = db.Prepare(firstQuery)
//	if err != nil {
//		t.Fatal("failed to prepare the statement regardless the order:", err)
//	}
//
//	_, err = db.Prepare(secondQuery)
//	if err != nil {
//		t.Fatal("failed to prepare the statement regardless the order:", err)
//	}
//}
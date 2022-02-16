package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ansalamdaniel/library-api/pkg/models"
)

func TestPing(t *testing.T) {
	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal()
	}

	ping(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	var pong struct {
		Message string
	}
	json.Unmarshal(body, &pong)
	if pong.Message != "pong!!!!!" {
		t.Errorf("Expected pong!!!! but got %+v", pong)
	}
}

func TestGetBooks(t *testing.T) {
	models.SetupTestDb()

	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal()
	}

	getBooks(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	var books []models.Book

	json.Unmarshal(body, &books)
	if len(books) != 2 {
		t.Errorf("Expected records retrieved should be 2 but got %d", len(books))
	}

	if books[0].Title != "Nuclear and Particle Physics" {
		t.Errorf("Expected first record from db is wrong!!")
	}

	models.TeardownTestDb()
}

func TestGetBooksMethodNotAllowed(t *testing.T) {

	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/books", nil)
	if err != nil {
		t.Fatal()
	}

	getBooks(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d", http.StatusMethodNotAllowed, rs.StatusCode)
	}

	if rs.Header.Get("Allow") != "GET, OPTIONS" {
		t.Errorf("Expected headers GET, OPTIONS but got %v", rs.Header.Get("Allow"))
	}
}

func TestGetBooksOptionsMethod(t *testing.T) {
	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("OPTIONS", "/books", nil)
	if err != nil {
		t.Fatal()
	}

	getBooks(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusNoContent {
		t.Errorf("Expected %d but got %d", http.StatusNoContent, rs.StatusCode)
	}

	if rs.Header.Get("Allow") != "GET, OPTIONS" {
		t.Errorf("Expected headers GET, OPTIONS but got %v", rs.Header.Get("Allow"))
	}
}

func TestCreateBook(t *testing.T) {
	models.SetupTestDb()

	rrecord := httptest.NewRecorder()

	fp, err := os.Open("testdata/post.json")
	if err != nil {
		t.Fatal(err)
	}

	defer fp.Close()

	req, err := http.NewRequest("POST", "/books/create", fp)
	if err != nil {
		t.Fatal()
	}

	req.Header.Set("Content-Type", "application/json")
	createBook(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusCreated {
		t.Errorf("Expected %d but got %d", http.StatusCreated, rs.StatusCode)
	}

}

func TestCreateBookInternalServerError(t *testing.T) {
	rrecord := httptest.NewRecorder()

	fp, err := os.Open("testdata/post.json")
	if err != nil {
		t.Fatal(err)
	}

	defer fp.Close()

	req, err := http.NewRequest("POST", "/books/create", fp)
	if err != nil {
		t.Fatal()
	}

	req.Header.Set("Content-Type", "application/json")
	createBook(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d", http.StatusInternalServerError, rs.StatusCode)
	}

	models.TeardownTestDb()
}

func TestCreateBookMethodNotAllowed(t *testing.T) {
	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/books/create", nil)
	if err != nil {
		t.Fatal()
	}

	createBook(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d", http.StatusMethodNotAllowed, rs.StatusCode)
	}

	if rs.Header.Get("Allow") != "POST, OPTIONS" {
		t.Errorf("Expected headers POST, OPTIONS but got %v", rs.Header.Get("Allow"))
	}
}

func TestCreateBookOptionsMethod(t *testing.T) {
	rrecord := httptest.NewRecorder()

	req, err := http.NewRequest("OPTIONS", "/books/create", nil)
	if err != nil {
		t.Fatal()
	}

	createBook(rrecord, req)

	rs := rrecord.Result()

	if rs.StatusCode != http.StatusNoContent {
		t.Errorf("Expected %d but got %d", http.StatusNoContent, rs.StatusCode)
	}

	if rs.Header.Get("Allow") != "POST, OPTIONS" {
		t.Errorf("Expected headers POST, OPTIONS but got %v", rs.Header.Get("Allow"))
	}
}

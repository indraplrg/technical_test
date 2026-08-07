package validator

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

type sampleRequest struct {
	Nama         string `json:"nama" validate:"required"`
	Umur         int    `json:"umur" validate:"gt=0"`
	Email        string `json:"email" validate:"email"`
	NamaLengkap  string `json:"nama_lengkap" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"date"`
}

func TestValidateErrorsMissingFields(t *testing.T) {
	v := validator.New()
	_ = v.RegisterValidation("date", dateValidation)

	err := v.Struct(sampleRequest{})
	if err == nil {
		t.Fatal("expected validation errors")
	}

	messages := ValidateErrors(err)
	if len(messages) == 0 {
		t.Fatal("expected error messages")
	}

	for _, want := range []string{"nama", "umur", "email", "nama_lengkap", "tanggal_lahir"} {
		if !containsAny(messages, want) {
			t.Errorf("expected a message mentioning %q, got %v", want, messages)
		}
	}
}

func TestValidateErrorsInvalidDate(t *testing.T) {
	v := validator.New()
	_ = v.RegisterValidation("date", dateValidation)

	err := v.Struct(sampleRequest{
		Nama:         "Budi",
		Umur:         20,
		Email:        "budi@mail.com",
		NamaLengkap:  "Budi Santoso",
		TanggalLahir: "20-01-2000",
	})
	if err == nil {
		t.Fatal("expected date validation error")
	}

	messages := ValidateErrors(err)
	if !containsAny(messages, "tanggal_lahir") {
		t.Errorf("expected date format message, got %v", messages)
	}
}

func TestSnakeCase(t *testing.T) {
	cases := map[string]string{
		"Nama":         "nama",
		"NamaJurusan":  "nama_jurusan",
		"TanggalLahir": "tanggal_lahir",
		"IDJurusan":    "i_d_jurusan",
		"Alamat":       "alamat",
	}
	for input, want := range cases {
		if got := snakeCase(input); got != want {
			t.Errorf("snakeCase(%q) = %q, want %q", input, got, want)
		}
	}
}

func containsAny(messages []string, substr string) bool {
	for _, m := range messages {
		if strings.Contains(m, substr) {
			return true
		}
	}
	return false
}

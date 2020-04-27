package main

import (
	"github.com/ayupov-ayaz/marshalers/models"
	"testing"
)

func BenchmarkSimpleMarshal(b *testing.B) {
	user := models.NewUser()
	b.ReportAllocs()
	var err error
	for i := 0; i < b.N; i++ {
		_, err = SimpleMarshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSimpleUnmarshal(b *testing.B) {
	data, err := SimpleMarshal(models.NewUser())
	if err != nil {
		b.Fatal(err)
	}

	user := &models.User{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = SimpleUnmarshal(data, user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyMarshal(b *testing.B) {
	user := models.NewUser()
	b.ReportAllocs()

	var err error
	for i := 0; i < b.N; i++ {
		_, err = EasyMarshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyUnmarshal(b *testing.B) {
	data, err := EasyMarshal(models.NewUser())
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	user := &models.User{}
	for i := 0; i < b.N; i++ {
		 err = EasyUnmarshal(data, user)
		 if err != nil {
		 	b.Fatal(err)
		 }
	}
}

func BenchmarkFastMarshal(b *testing.B) {
	user := models.NewUser()
	b.ReportAllocs()

	var err error
	for i := 0; i < b.N; i++ {
		_, err = FastMarshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFastUnmarshal(b *testing.B) {
	data, err := FastMarshal(models.NewUser())
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	user := &models.User{}
	for i := 0; i < b.N; i++ {
		err = FastUnmarshal(data, user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFfJsonMarshal(b *testing.B) {
	user := models.NewUser()
	b.ReportAllocs()

	var err error
	for i := 0; i < b.N; i++ {
		_, err = FfJsonMarshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFfJsonUnmarshal(b *testing.B) {
	data, err := FfJsonMarshal(models.NewUser())
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	user := &models.User{}
	for i := 0; i < b.N; i++ {
		err = FfJsonUnmarshal(data, user)
		if err != nil {
			b.Fatal(err)
		}
	}
}


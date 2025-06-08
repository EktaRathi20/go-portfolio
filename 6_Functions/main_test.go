package main

import "testing"

func TestFunctionWithoutParams(t *testing.T) {
	got := functionWithoutParams()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFunctionWithParams(t *testing.T) {
	got := functionWithParams("World")
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFunctionWithMultipleParams(t *testing.T) {
	got := functionWithMultipleParams("Test", "Spanish")
	want := "Hola, Test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = functionWithMultipleParams("", "Spanish")
	want = "Hola, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFunctionWithVariable(t *testing.T) {
	got := functionWithVariable()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFunctionWithIfElse(t *testing.T) {
	got := functionWithIfElse("Test", "Spanish")
	want := "Hola, Test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = functionWithIfElse("", "Spanish")
	want = "Hola, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = functionWithIfElse("Test", "French")
	want = "Holaa, Test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = functionWithIfElse("Test", "")
	want = "HelloTest"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

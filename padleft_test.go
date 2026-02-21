package main

import (
	"strings"
	"testing"
)

func TestPad(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		n       int
		want    string
		wantErr error
	}{
		// --- Negative counts: always error ---
		{name: "negative_one", s: "", n: -1, want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_one_nonempty", s: "a", n: -1, want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_five", s: "hello", n: -5, want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_large", s: "hello", n: -100, want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_very_large", s: "", n: -1000000, want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_one_spaces", s: "   ", n: -1, want: "", wantErr: ErrInvalidPadLength},

		// --- Empty string ---
		{name: "empty_zero", s: "", n: 0, want: "", wantErr: nil},
		{name: "empty_one", s: "", n: 1, want: " ", wantErr: nil},
		{name: "empty_two", s: "", n: 2, want: "  ", wantErr: nil},
		{name: "empty_five", s: "", n: 5, want: "     ", wantErr: nil},
		{name: "empty_ten", s: "", n: 10, want: "          ", wantErr: nil},
		{name: "empty_large", s: "", n: 100, want: strings.Repeat(" ", 100), wantErr: nil},

		// --- String longer than count: returned unchanged ---
		{name: "longer_count_zero", s: "a", n: 0, want: "a", wantErr: nil},
		{name: "longer_by_one", s: "ab", n: 1, want: "ab", wantErr: nil},
		{name: "longer_by_many", s: "hello", n: 3, want: "hello", wantErr: nil},
		{name: "longer_short_count", s: "hello world", n: 5, want: "hello world", wantErr: nil},
		{name: "longer_long_string", s: strings.Repeat("a", 1000), n: 100, want: strings.Repeat("a", 1000), wantErr: nil},
		{name: "longer_spaces_string", s: "   ", n: 1, want: "   ", wantErr: nil},

		// --- String exactly equal to count: returned unchanged ---
		{name: "exact_single_char", s: "a", n: 1, want: "a", wantErr: nil},
		{name: "exact_five_chars", s: "hello", n: 5, want: "hello", wantErr: nil},
		{name: "exact_with_space", s: " ", n: 1, want: " ", wantErr: nil},
		{name: "exact_with_spaces", s: "     ", n: 5, want: "     ", wantErr: nil},
		{name: "exact_eleven", s: "hello world", n: 11, want: "hello world", wantErr: nil},
		{name: "exact_100_chars", s: strings.Repeat("x", 100), n: 100, want: strings.Repeat("x", 100), wantErr: nil},

		// --- Padding needed (string shorter than count) ---
		{name: "pad_one_space", s: "a", n: 2, want: " a", wantErr: nil},
		{name: "pad_four_spaces", s: "a", n: 5, want: "    a", wantErr: nil},
		{name: "pad_one_to_hello", s: "hello", n: 6, want: " hello", wantErr: nil},
		{name: "pad_five_to_hello", s: "hello", n: 10, want: "     hello", wantErr: nil},
		{name: "pad_four_to_sentence", s: "hello world", n: 15, want: "    hello world", wantErr: nil},
		{name: "pad_nine_to_sentence", s: "hello world", n: 20, want: "         hello world", wantErr: nil},
		{name: "pad_single_digit", s: "7", n: 5, want: "    7", wantErr: nil},
		{name: "pad_two_digits", s: "42", n: 5, want: "   42", wantErr: nil},
		{name: "pad_three_digits", s: "100", n: 5, want: "  100", wantErr: nil},
		{name: "pad_large_count", s: "hi", n: 100, want: strings.Repeat(" ", 98) + "hi", wantErr: nil},
		{name: "pad_50_of_100", s: strings.Repeat("a", 50), n: 100, want: strings.Repeat(" ", 50) + strings.Repeat("a", 50), wantErr: nil},

		// --- Strings with special characters (counted as bytes) ---
		{name: "tab_in_string_exact", s: "\t", n: 1, want: "\t", wantErr: nil},
		{name: "tab_in_string_pad", s: "\t", n: 3, want: "  \t", wantErr: nil},
		{name: "newline_in_string_exact", s: "\n", n: 1, want: "\n", wantErr: nil},
		{name: "newline_in_string_pad", s: "\n", n: 3, want: "  \n", wantErr: nil},
		{name: "mixed_whitespace_exact", s: "foo\tbar", n: 7, want: "foo\tbar", wantErr: nil},
		{name: "mixed_whitespace_pad", s: "foo\tbar", n: 10, want: "   foo\tbar", wantErr: nil},
		{name: "newline_in_middle_exact", s: "hello\nworld", n: 11, want: "hello\nworld", wantErr: nil},
		{name: "newline_in_middle_pad", s: "hello\nworld", n: 15, want: "    hello\nworld", wantErr: nil},
		{name: "leading_spaces_exact", s: "  hello  ", n: 9, want: "  hello  ", wantErr: nil},
		{name: "leading_spaces_pad", s: "  hello  ", n: 15, want: "      " + "  hello  ", wantErr: nil},
		{name: "negative_sign_string", s: "-1", n: 5, want: "   -1", wantErr: nil},
		{name: "positive_sign_string", s: "+1", n: 5, want: "   +1", wantErr: nil},
		{name: "zero_padded_number_string", s: "007", n: 5, want: "  007", wantErr: nil},

		// --- Unicode string content (len() counts bytes, not runes) ---
		// "é" is 2 bytes in UTF-8
		{name: "unicode_e_acute_longer", s: "é", n: 1, want: "é", wantErr: nil},
		{name: "unicode_e_acute_exact", s: "é", n: 2, want: "é", wantErr: nil},
		{name: "unicode_e_acute_pad_one", s: "é", n: 3, want: " é", wantErr: nil},
		// "café" = c(1)+a(1)+f(1)+é(2) = 5 bytes
		{name: "unicode_cafe_exact", s: "café", n: 5, want: "café", wantErr: nil},
		{name: "unicode_cafe_longer", s: "café", n: 3, want: "café", wantErr: nil},
		{name: "unicode_cafe_pad_one", s: "café", n: 6, want: " café", wantErr: nil},
		// "日本語" = 3 CJK chars × 3 bytes each = 9 bytes
		{name: "unicode_cjk_longer", s: "日本語", n: 3, want: "日本語", wantErr: nil},
		{name: "unicode_cjk_exact", s: "日本語", n: 9, want: "日本語", wantErr: nil},
		{name: "unicode_cjk_pad_one", s: "日本語", n: 10, want: " 日本語", wantErr: nil},
		{name: "unicode_cjk_pad_three", s: "日本語", n: 12, want: "   日本語", wantErr: nil},
		// 🎉 emoji = 4 bytes in UTF-8
		{name: "unicode_emoji_longer", s: "🎉", n: 3, want: "🎉", wantErr: nil},
		{name: "unicode_emoji_exact", s: "🎉", n: 4, want: "🎉", wantErr: nil},
		{name: "unicode_emoji_pad_one", s: "🎉", n: 5, want: " 🎉", wantErr: nil},
		{name: "unicode_emoji_pad_four", s: "🎉", n: 8, want: "    🎉", wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pad(tt.s, tt.n)
			if err != tt.wantErr {
				t.Errorf("Pad(%q, %d) error = %v, wantErr %v", tt.s, tt.n, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pad(%q, %d) = %q, want %q", tt.s, tt.n, got, tt.want)
			}
		})
	}
}

func TestPadCharacter(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		count   int
		char    rune
		want    string
		wantErr error
	}{
		// --- Negative counts: always error ---
		{name: "negative_space", s: "", count: -1, char: ' ', want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_zero_char", s: "a", count: -1, char: '0', want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_star", s: "hello", count: -5, char: '*', want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_large_unicode_char", s: "hello", count: -100, char: '★', want: "", wantErr: ErrInvalidPadLength},
		{name: "negative_very_large", s: "", count: -1000000, char: 'x', want: "", wantErr: ErrInvalidPadLength},

		// --- Space character (mirrors Pad) ---
		{name: "space_empty_zero", s: "", count: 0, char: ' ', want: "", wantErr: nil},
		{name: "space_empty_five", s: "", count: 5, char: ' ', want: "     ", wantErr: nil},
		{name: "space_exact", s: "hello", count: 5, char: ' ', want: "hello", wantErr: nil},
		{name: "space_longer", s: "hello", count: 3, char: ' ', want: "hello", wantErr: nil},
		{name: "space_pad_five", s: "hello", count: 10, char: ' ', want: "     hello", wantErr: nil},

		// --- Zero character (numeric zero-filling) ---
		{name: "zero_empty_zero_count", s: "", count: 0, char: '0', want: "", wantErr: nil},
		{name: "zero_empty_five", s: "", count: 5, char: '0', want: "00000", wantErr: nil},
		{name: "zero_single_digit_longer", s: "1", count: 0, char: '0', want: "1", wantErr: nil},
		{name: "zero_single_digit_exact", s: "1", count: 1, char: '0', want: "1", wantErr: nil},
		{name: "zero_single_digit_pad", s: "1", count: 3, char: '0', want: "001", wantErr: nil},
		{name: "zero_two_digits", s: "42", count: 5, char: '0', want: "00042", wantErr: nil},
		{name: "zero_three_digits_exact", s: "100", count: 3, char: '0', want: "100", wantErr: nil},
		{name: "zero_three_digits_longer", s: "1000", count: 3, char: '0', want: "1000", wantErr: nil},
		{name: "zero_007_style", s: "7", count: 3, char: '0', want: "007", wantErr: nil},
		{name: "zero_042_style", s: "42", count: 3, char: '0', want: "042", wantErr: nil},
		{name: "zero_zero_digit", s: "0", count: 5, char: '0', want: "00000", wantErr: nil},
		{name: "zero_negative_number_string", s: "-5", count: 5, char: '0', want: "000-5", wantErr: nil},
		{name: "zero_double_digit_99", s: "99", count: 5, char: '0', want: "00099", wantErr: nil},

		// --- Asterisk character ---
		{name: "star_empty_five", s: "", count: 5, char: '*', want: "*****", wantErr: nil},
		{name: "star_pad_hello", s: "hello", count: 10, char: '*', want: "*****hello", wantErr: nil},
		{name: "star_exact", s: "hello", count: 5, char: '*', want: "hello", wantErr: nil},
		{name: "star_longer", s: "hello", count: 3, char: '*', want: "hello", wantErr: nil},

		// --- Dash character ---
		{name: "dash_empty_five", s: "", count: 5, char: '-', want: "-----", wantErr: nil},
		{name: "dash_pad_hello", s: "hello", count: 10, char: '-', want: "-----hello", wantErr: nil},

		// --- Equals character ---
		{name: "equals_pad_hello", s: "hello", count: 10, char: '=', want: "=====hello", wantErr: nil},

		// --- Dot character ---
		{name: "dot_pad_hello", s: "hello", count: 10, char: '.', want: ".....hello", wantErr: nil},

		// --- Underscore character ---
		{name: "underscore_pad_hello", s: "hello", count: 10, char: '_', want: "_____hello", wantErr: nil},

		// --- Hash character ---
		{name: "hash_pad_hello", s: "hello", count: 10, char: '#', want: "#####hello", wantErr: nil},

		// --- Slash character ---
		{name: "slash_pad_hello", s: "hello", count: 10, char: '/', want: "/////hello", wantErr: nil},

		// --- Alphabetic pad character ---
		{name: "alpha_empty", s: "", count: 3, char: 'a', want: "aaa", wantErr: nil},
		{name: "alpha_pad_one", s: "b", count: 3, char: 'a', want: "aab", wantErr: nil},
		{name: "alpha_pad_two", s: "bb", count: 3, char: 'a', want: "abb", wantErr: nil},
		{name: "alpha_exact", s: "bbb", count: 3, char: 'a', want: "bbb", wantErr: nil},
		{name: "alpha_longer", s: "bbbb", count: 3, char: 'a', want: "bbbb", wantErr: nil},

		// --- Tab character as pad ---
		{name: "tab_pad_char_empty", s: "", count: 3, char: '\t', want: "\t\t\t", wantErr: nil},
		{name: "tab_pad_char_hi", s: "hi", count: 5, char: '\t', want: "\t\t\thi", wantErr: nil},
		{name: "tab_pad_char_exact", s: "hi", count: 2, char: '\t', want: "hi", wantErr: nil},
		{name: "tab_pad_char_longer", s: "hi", count: 1, char: '\t', want: "hi", wantErr: nil},

		// --- Newline character as pad ---
		{name: "newline_pad_char_empty", s: "", count: 3, char: '\n', want: "\n\n\n", wantErr: nil},
		{name: "newline_pad_char_hi", s: "hi", count: 5, char: '\n', want: "\n\n\nhi", wantErr: nil},

		// --- Null character as pad ---
		{name: "null_pad_char_empty", s: "", count: 3, char: 0, want: "\x00\x00\x00", wantErr: nil},
		{name: "null_pad_char_hi", s: "hi", count: 5, char: 0, want: "\x00\x00\x00hi", wantErr: nil},
		{name: "null_pad_char_zero_count_empty", s: "", count: 0, char: 0, want: "", wantErr: nil},
		{name: "null_pad_char_exact", s: "hi", count: 2, char: 0, want: "hi", wantErr: nil},

		// --- Unicode (multi-byte) pad character ---
		// '★' (U+2605) = 3 bytes in UTF-8
		// count - len(s) repetitions of the character are prepended
		{name: "unicode_star_empty_three", s: "", count: 3, char: '★', want: "★★★", wantErr: nil},
		{name: "unicode_star_pad_hi", s: "hi", count: 4, char: '★', want: "★★hi", wantErr: nil},
		{name: "unicode_star_pad_hello", s: "hello", count: 10, char: '★', want: "★★★★★hello", wantErr: nil},
		{name: "unicode_star_exact", s: "hello", count: 5, char: '★', want: "hello", wantErr: nil},
		{name: "unicode_star_longer", s: "hello", count: 3, char: '★', want: "hello", wantErr: nil},
		{name: "unicode_star_negative", s: "hello", count: -1, char: '★', want: "", wantErr: ErrInvalidPadLength},

		// '中' (U+4E2D) = 3 bytes in UTF-8
		{name: "unicode_cjk_char_pad", s: "hi", count: 4, char: '中', want: "中中hi", wantErr: nil},
		{name: "unicode_cjk_char_empty", s: "", count: 2, char: '中', want: "中中", wantErr: nil},

		// '€' (U+20AC) = 3 bytes in UTF-8
		{name: "unicode_euro_pad", s: "10", count: 5, char: '€', want: "€€€10", wantErr: nil},

		// --- count = 0 with various chars ---
		{name: "zero_count_empty_zero", s: "", count: 0, char: '0', want: "", wantErr: nil},
		{name: "zero_count_nonempty_zero", s: "abc", count: 0, char: '0', want: "abc", wantErr: nil},
		{name: "zero_count_empty_star", s: "", count: 0, char: '*', want: "", wantErr: nil},
		{name: "zero_count_nonempty_star", s: "abc", count: 0, char: '*', want: "abc", wantErr: nil},

		// --- count = 1 boundary ---
		{name: "count_one_empty_zero", s: "", count: 1, char: '0', want: "0", wantErr: nil},
		{name: "count_one_single_char_exact", s: "a", count: 1, char: '0', want: "a", wantErr: nil},
		{name: "count_one_two_chars_longer", s: "ab", count: 1, char: '0', want: "ab", wantErr: nil},

		// --- Large count ---
		{name: "large_count_empty", s: "", count: 100, char: '0', want: strings.Repeat("0", 100), wantErr: nil},
		{name: "large_count_hi", s: "hi", count: 100, char: '-', want: strings.Repeat("-", 98) + "hi", wantErr: nil},

		// --- Strings with existing pad character ---
		{name: "string_contains_pad_char", s: "0001", count: 6, char: '0', want: "000001", wantErr: nil},
		{name: "string_is_all_pad_char", s: "000", count: 5, char: '0', want: "00000", wantErr: nil},

		// --- Long strings at exact boundary ---
		{name: "long_string_exact", s: strings.Repeat("a", 100), count: 100, char: '*', want: strings.Repeat("a", 100), wantErr: nil},
		{name: "long_string_longer", s: strings.Repeat("a", 101), count: 100, char: '*', want: strings.Repeat("a", 101), wantErr: nil},
		{name: "long_string_pad_one", s: strings.Repeat("a", 99), count: 100, char: '*', want: "*" + strings.Repeat("a", 99), wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PadCharacter(tt.s, tt.count, tt.char)
			if err != tt.wantErr {
				t.Errorf("PadCharacter(%q, %d, %q) error = %v, wantErr %v", tt.s, tt.count, tt.char, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PadCharacter(%q, %d, %q) = %q, want %q", tt.s, tt.count, tt.char, got, tt.want)
			}
		})
	}
}

// TestPadDelegates verifies that Pad is exactly equivalent to PadCharacter with ' '.
func TestPadDelegates(t *testing.T) {
	cases := []struct {
		s string
		n int
	}{
		{"", 0}, {"", 5}, {"", -1},
		{"hello", 0}, {"hello", 5}, {"hello", 10}, {"hello", -1},
		{"a", 1}, {"ab", 1}, {"abc", 100},
	}
	for _, c := range cases {
		padResult, padErr := Pad(c.s, c.n)
		charResult, charErr := PadCharacter(c.s, c.n, ' ')
		if padResult != charResult || padErr != charErr {
			t.Errorf("Pad(%q, %d) = (%q, %v); PadCharacter(%q, %d, ' ') = (%q, %v): mismatch",
				c.s, c.n, padResult, padErr,
				c.s, c.n, charResult, charErr)
		}
	}
}

// TestErrInvalidPadLength verifies the error value is stable.
func TestErrInvalidPadLength(t *testing.T) {
	_, err := Pad("", -1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != ErrInvalidPadLength {
		t.Errorf("got error %v, want ErrInvalidPadLength", err)
	}
	if err.Error() != "Invalid pad length, must be 0 or greater" {
		t.Errorf("unexpected error message: %q", err.Error())
	}
}

palindrome_list(X) :- reverse(X, X).
% atom_codes, name
palindrome(Str) :- string_codes(Str, X), palindrome_list(X).
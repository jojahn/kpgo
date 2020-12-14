male(albert).
male(bob).
male(bill).
male(carl).
male(charlie).
male(dan).
male(edward).

female(alice).
female(betsy).
female(diana).

parent(albert, bob).
parent(albert, betsy).
parent(albert, bill).

parent(alice, bob).
parent(alice, betsy).
parent(alice, bill).

parent(bob, carl).
parent(bob, charlie).
parent(diana, charlie).


grandparent(X) :- parent(X, Y), parent(Y, charlie).

grandmothers(X) :- grandparent(X), female(X).

siblings(X) :- parent(Y, bob), parent(Y, X).


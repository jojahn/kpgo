cross(a,b).
cross(b,a).

% single moves
move([X, Goat, Cabbage, Wolf], farmer, [Y, Goat, Cabbage, Wolf]) :- cross(X, Y).

% partner moves
move([X, X, Cabbage, Wolf], goat, [Y, Y, Cabbage, Wolf]) :- cross(X, Y).
move([X, Goat, X, Wolf], cabbage, [Y, Goat, Y, Wolf]) :- cross(X, Y).
move([X, Goat, Cabbage, X], wolf, [Y, Goat, Cabbage, Y]) :- cross(X, Y).

safe([Farmer, Goat, _, _]) :- Farmer = Goat.
safe([Farmer, _, Cabbage, Wolf]) :- Farmer = Wolf, Farmer = Cabbage.

solution([b,b,b,b], []).
solution(State, [Move|OtherMoves]) :- move(State, Move, NextState),
    safe(NextState),
    solution(NextState, OtherMoves).

run :- length(X, 7), solution([a,a,a,a], X), write(X).

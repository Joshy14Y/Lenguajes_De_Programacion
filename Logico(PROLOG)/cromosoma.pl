% Ejercicio Cromosomas
cromosoma1(['A','A','B','C','D','A','C','B','C','A']).
cromosoma2(['A','A','C','C','D','A','B','A','D','B']).
cromosoma3(['B','A','D','C','D','A','B','A','D','B']).
cromosoma4(['C','A','D','C','D','C','B','C','D','A']).

cromosoma(X) :-
    is_list(X),
    length(X, 10),
    forall(member(G, X), (G = 'A'; G = 'B'; G = 'C'; G = 'D')).

candidato(X) :- 
    atom(X),
    atom_concat(cromosoma, _, X),
    call(X, L),
    cromosoma(L).


mapPosicion([], [], []).
mapPosicion([X|Xs], [Y|Ys], [Z|Zs]) :-
    (X = Y -> Z = 1; Z = 0), 
    mapPosicion(Xs, Ys, Zs). 

filtrarCeros([], []). 
filtrarCeros([X|Xs], Ys) :- 
    (
    X = 0 -> Ys = Zs; Ys = [X|Zs]
    ),
    filtrarCeros(Xs, Zs). 

sum([], 0). 
sum([X|Xs], Y) :- 
    sum(Xs, Z), 
    Y is Z + X * 10. 

cromosoma(S, C, V) :- 
    call(S, L1), 
    call(C, L2),  
    mapPosicion(L1, L2, L3), 
    filtrarCeros(L3, L4), 
    sum(L4, V). 

% Test: ?- cromosoma(cromosoma1, cromosoma2, Valor).  



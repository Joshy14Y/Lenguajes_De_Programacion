% Defina sumlist(L, S) que es verdadero si S es la suma de los elementos de L.
sumlist([],0).  
sumlist([_|R],S) :- sumlist(R,S1), S is S1 + 1.

% Defina la relación subconj(S, S1), donde S y S1 son listas representando conjuntos, que es verdadera si S1 es subconjunto de S.
subconj([], []).
subconj([S|SR], [S|S1R]) :- 
    subconj(SR, S1R).
subconj(SR, [_|S1R]) :- 
    subconj(SR,S1R).

/*
Defina la función aplanar(L,L2). Esta función recibe una lista con múltiples listas anidadas como elementos otra lista que contendría los mismos 
elementos de manera lineal (sin listas). 
*/
flatten(L1, FL) :-
    flatten(L1, [], FL0),
    !,
    FL = FL0.
flatten(V, T, [V|T]) :-
    var(V),
    !.
flatten([], T, T) :- !.
  flatten([H|T], T1, L) :-
    !,
    flatten(H, FHT, L),
    flatten(T, T1, FHT).
    flatten(NL, T, [NL|T]).

% Defina un predicado llamado partir(Lista, Umbral, Menores, Mayores)
partir([], _, [], []).
partir([H|T], Umbral, [H|Menores], Mayores) :-
    H < Umbral,
    partir(T, Umbral, Menores, Mayores).
partir([H|T], Umbral, Menores, [H|Mayores]) :-
    H >= Umbral,
    partir(T, Umbral, Menores, Mayores).


/* Implemente un predicado que, a partir de una lista de cadenas string, filtre aquellas que contengan una subcadena
 que el usuario indique en otro argumento.
*/
contiene(S, C) :- sub_atom(C, _, _, _, S).
sub_cadenas(S, LC, F) :-
    include(contiene(S), LC, F).

%TESTS
main :-
    writeln("Ejecutando consultas..."),
    
    % Ejemplo de consulta 1
    sumlist([1, 2, 3, 4, 5], Suma),
    writeln("La suma de la lista es:"),
    writeln(Suma),
    
    % Ejemplo de consulta 2: Debe realizarse manualmente.
    subconj([1, 2], [1, 3, 2, 4]),
    
    % Ejemplo de consulta 3
    flatten([1, [2, [3, 4], 5], [6]], ListaAplanada),
    writeln("Lista aplanada:"),
    writeln(ListaAplanada),
    
    % Ejemplo de consulta 4
    partir([2, 7, 4, 8, 9, 1], 6, Menores, Mayores),
    writeln("Menores:"),
    writeln(Menores),
    writeln("Mayores:"),
    writeln(Mayores),
    
    % Ejemplo de consulta 5
    sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"], Filtradas),
    writeln("Cadenas que contienen 'la':"),
    writeln(Filtradas).

% Iniciar el programa
:- initialization(main).
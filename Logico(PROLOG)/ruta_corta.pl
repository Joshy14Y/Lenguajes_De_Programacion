%Implementacion de Dijkstra
grafo(a,c,60).
grafo(a,e,80).
grafo(b,a,30).
grafo(b,d,20).
grafo(d,e,10).
grafo(d,a,60).
grafo(c,e,50).
grafo(c,x,40).
grafo(e,f,20).
grafo(f,x,50).
grafo(x,g,30).
grafo(g,c,20).

dijkstra(_,_,_,[],L):- L=[].
dijkstra(I,F,S,_,L):- 
    menorArista(I,A),
    A==F,
    append(S,A,Lista),
    L = Lista.
dijkstra(I,F,S,C,L):- 
    menorArista(I,A),
    A\=F,
    append(S,A,Lista1),
    eliminar(A,C,Lista2),
    dijkstra(A,F,Lista1,Lista2,L).

menorArista(I,A):- 
    findall(P,grafo(I,_,P),L),
    menorlista(L,Menor),
    grafo(I,A,Menor).

menorlista([],0).
menorlista([X],X).
menorlista([X|L],X):- 
    menorlista(L,Y),
    menor(X,Y).
menorlista([X|L],Y):- 
    menorlista(L,Y),
    menor(Y,X).
    menor(X,Y):- X =< Y.

append([],L,[L]).
append([H|T],N,[H|S]):-append(T,N,S).

eliminar(_,[],[]):-fail.
eliminar(X,[X|R],R).
eliminar(X,[C|R],[C|R1]):- eliminar(X,R,R1).

% Test: ?- dijkstra(a, g, [a], [b, c, d, e, f, x], Ruta). 
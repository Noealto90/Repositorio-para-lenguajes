%!  Ejercicio 1
sub_conjunto([], _).
sub_conjunto([X|Xs], Ys) :-
    member(X, Ys),
    sub_conjunto(Xs, Ys).


%!  Ejercicio 2

% Predicado para aplanar una lista sin utilizar append
aplanar([], []). % Caso base: aplanar una lista vacía devuelve una lista vacía
aplanar([X|Resto], Aplanado) :-
    X = [_|_], % Verifica si X es una lista
    !,         % Corta para evitar backtracking innecesario
    aplanar(X, AplanadoPrimer), % Aplana el primer elemento de la lista
    aplanar(Resto, AplanadoResto), % Aplana el resto de la lista
    concatenar(AplanadoPrimer, AplanadoResto, Aplanado). % Concatena las listas resultantes
aplanar([X|Resto], [X|AplanadoResto]) :- % Si X no es una lista
    aplanar(Resto, AplanadoResto). % Continúa aplanando el resto de la lista

% Predicado para concatenar dos listas sin utilizar append
concatenar([], L, L). % Caso base: concatenar una lista vacía con cualquier otra lista devuelve la lista original
concatenar([X|Resto], L, [X|Resultado]) :- % Concatenar dos listas
    concatenar(Resto, L, Resultado). % Continúa concatenando el resto de la lista


%!   Ejercicio 3
% Predicado para calcular la distancia de Hamming entre dos cadenas
distanciaH(Str1, Str2, Distancia) :-
    atom_chars(Str1, Lista1),    % Convierte la primera cadena en una lista de caracteres
    atom_chars(Str2, Lista2),    % Convierte la segunda cadena en una lista de caracteres
    length(Lista1, Len1),        % Obtiene la longitud de la primera lista
    length(Lista2, Len2),        % Obtiene la longitud de la segunda lista
    min(Len1, Len2, MinLen),     % Obtiene la longitud mínima entre las dos cadenas
    calcular_distancia(Lista1, Lista2, MinLen, 0, Distancia), % Calcula la distancia de Hamming
    !. % Corte para evitar backtracking

% Predicado auxiliar para calcular la distancia de Hamming recursivamente
calcular_distancia(_, _, 0, Distancia, Distancia) :- !.
calcular_distancia([X|Resto1], [X|Resto2], N, Acumulador, Distancia) :-
    N > 0,
    N1 is N - 1,
    calcular_distancia(Resto1, Resto2, N1, Acumulador, Distancia),
    !.
calcular_distancia([_|Resto1], [_|Resto2], N, Acumulador, Distancia) :-
    N > 0,
    N1 is N - 1,
    Acumulador1 is Acumulador + 1,
    calcular_distancia(Resto1, Resto2, N1, Acumulador1, Distancia),
    !.

% Predicado para obtener el mínimo entre dos números
min(X, Y, X) :- X =< Y, !.
min(_, Y, Y).


%!  Ejercicio 4

% Hechos de los platos
plato(entrada, guacamole, 200).
plato(entrada, ensalada, 150).
plato(entrada, consome, 300).
plato(entrada, tostadas_caprese, 250).

plato(carne, filete_de_cerdo, 400).
plato(carne, pollo_al_horno, 280).
plato(carne, carne_en_salsa, 320).

plato(pescado, tilapia, 160).
plato(pescado, salmon, 300).
plato(pescado, trucha, 225).

plato(postre, flan, 200).
plato(postre, nueces_con_miel, 500).
plato(postre, naranja_confitada, 450).
plato(postre, flan_de_coco, 375).

% Regla para determinar si un plato es principal
plato_principal(Plato) :-
    plato(carne, Plato, _).
plato_principal(Plato) :-
    plato(pescado, Plato, _).

% Regla para calcular las calorías totales de una lista de platos
calorias_totales([], 0).
calorias_totales([Plato|Platos], TotalCalorias) :-
    plato(_, Plato, Calorias),
    calorias_totales(Platos, RestoCalorias),
    TotalCalorias is Calorias + RestoCalorias.


% Regla final para generar las primeras 5 combinaciones de comidas
combinaciones_de_comidas(MaxCalorias, Combinaciones) :-
    findall((Entrada, Principal, Postre), (
        plato(entrada, Entrada, _),
        plato_principal(Principal),
        plato(postre, Postre, _),
        calorias_totales([Entrada, Principal, Postre], TotalCalorias),
        TotalCalorias =< MaxCalorias
    ), TodasCombinaciones),
    length(TodasCombinaciones, Length),
    Length >= 5,
    length(Combinaciones, 5),
    append(Combinaciones, _, TodasCombinaciones).

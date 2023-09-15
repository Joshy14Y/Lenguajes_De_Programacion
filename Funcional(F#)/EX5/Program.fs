// 5) Construya un programa que resuelva el problema de encontrar la o las rutas de un laberinto con datos quemados. 

let maze = [
    ("INICIO", ["2"]);
    ("1", ["7"]);
    ("2", ["INICIO"; "8"; "3"]);
    ("3", ["2"; "9"; "4"]);
    ("4", ["10"]);
    ("5", ["11"; "6"]);
    ("6", ["5"]);
    ("7", ["1";"13"]);
    ("8", ["2";"9"]);
    ("9", ["8";"3"]);
    ("10", ["4";"16"]);
    ("11", ["5";"17"]);
    ("12", ["18"]);
    ("13", ["7";"14"]);
    ("14", ["13";"15";"20"]);
    ("15", ["14";"21"]);
    ("16", ["10";"22"]);
    ("17", ["11";"23"]);
    ("18", ["12";"24"]);
    ("19", ["25"]);
    ("20", ["14";"26"]);
    ("21", ["15";"22"]);
    ("22", ["16";"21"]);
    ("23", ["17";"29"]);
    ("24", ["18";"30"]);
    ("25", ["19";"31"]);
    ("26", ["20";"27"]);
    ("27", ["26";"28"]);
    ("28", ["27";"34";"29"]);
    ("29", ["23";"28"]);
    ("30", ["24";"36"]);
    ("31", ["25";"32"]);
    ("32", ["31";"33";"FIN"]);
    ("33", ["32";"34"]);
    ("34", ["33";"28";"35"]);
    ("35", ["34";"36"]);
    ("36", ["35";"30"]);
    ("FIN", ["32"]);
]

let miembro elem lista =
    List.exists (fun x -> x = elem) lista

// Función para generar vecinos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

// Función para extender una ruta
let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 

// Función principal de búsqueda en profundidad
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)  ) grafo
    prof_aux fin [[ini]] grafo

let rec getShortestRoute list =
    match list with
    | [] -> [] 
    | [singleList] -> singleList 
    | head :: rest ->
        let shortestRest = getShortestRoute rest 
        if List.length head <= List.length shortestRest then head
        else shortestRest
        

let ruta = prof2 "INICIO" "FIN" maze
printfn "Rutas encontradas: %A" ruta
0 // Salida del programa

let shortestList = getShortestRoute ruta
printfn "Ruta mas corta: %A" shortestList
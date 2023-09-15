// 1) Realizar ejercicio para desplazar (SHIFT) una lista de elementos n posiciones a la izquierda o la derecha según se
// indique por argumento

let rec shift direccion n lista =
    
    let len = List.length lista
    
    let shiftLeft list n =
        let n = if n > len then len else n
        List.append (List.skip n list) (List.replicate n 0)

    let shiftRight list n =
        let n = if n > len then len else n
        List.append (List.replicate (len - n) 0) (List.take n list)
    
    if n >= len then List.replicate len 0 
    else  
        match direccion with
        | "izq" when n > 0 -> shiftLeft lista n
        | "der" when n > 0 -> shiftRight lista n
        | _ -> List.replicate len 0


let resultado1 = shift "izq" 3 [1; 2; 3; 4; 5]
let resultado2 = shift "der" 2 [1; 2; 3; 4; 5]
let resultado3 = shift "izq" 6 [1; 2; 3; 4; 5]

printfn "Resultado 1: %A" resultado1
printfn "Resultado 2: %A" resultado2
printfn "Resultado 3: %A" resultado3
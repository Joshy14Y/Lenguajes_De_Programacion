// 2) Haciendo uso de la función filter, implemente una función que, a partir de una lista de cadenas de parámetro, filtre aquellas
// que contengan una subcadena que el usuario indique en otro argumento.

let filtrarPorSubcadena(subcadena: string) (listaCadenas: string list) =
    listaCadenas
    |> List.filter (fun cadena -> cadena.Contains(subcadena))

// Ejemplo de uso:
let subcadena = "la" 
let listaCadenas = ["la casa"; "el perro"; "pintando la cerca"]

let resultado = filtrarPorSubcadena subcadena listaCadenas
printfn "%A" resultado
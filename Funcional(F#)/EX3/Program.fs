// 3) Realizar una función que obtenga el n-esimo elemento de una lista pero utilizando solamente map (sin recursión).

let mapPosition position list =
    List.mapi (fun index x -> if index = position then x else -1) list

let filterNotMinusOne list =
    List.filter (fun x -> x <> -1) list

let n_esimo position list = 
    if position >= List.length list then -1
    else
       let filteredList = mapPosition position list
       let filteredResult = filterNotMinusOne filteredList
       filteredResult.[0]

let numbers = [1; 2; 3; 4; 5]
let result = n_esimo 0 numbers

printfn "Result: %d" result
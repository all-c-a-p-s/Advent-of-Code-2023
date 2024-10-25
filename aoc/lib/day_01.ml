let read_lines filename =
  let channel = open_in filename in
  let rec read_lines_acc acc =
    try
      let line = input_line channel in
      read_lines_acc (line :: acc) (* :: operator appends to the FRONT of a list so we need to call .rev later*)
    with End_of_file ->
      close_in channel;
      List.rev acc
  in
  read_lines_acc []
;;

let is_digit a =
  match a with
    | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '0' -> true
    | _ -> false
;;

let char_to_digit c =
  if c >= '0' && c <= '9' then
    int_of_char c - int_of_char '0'
  else
    failwith "tried to convert non-digit to int"
;;


let print_char_list lst =
  List.iter (fun c -> Printf.printf "%c " c) lst;
  Printf.printf "\n";  (* Print a newline after printing all elements *)
;;

let line_score line =
  let l = line |> String.to_seq |> List.of_seq in
  let digits = List.filter is_digit l in
  ((List.nth digits 0) |> char_to_digit) * 10 + ((List.nth (List.rev digits) 0) |> char_to_digit)
;;

let rec total_score lines =
  match lines with
    | [] -> 0
    | x :: xs -> line_score x + total_score xs 
    (*I think this syntax is pretty nice - if the list is now empty then it is
    equivalent to some first element x :: some rest of list xs which may or may not be empty*)
;;

let answer filename =
  let lines = read_lines filename in
  total_score lines



let rec scan_file channel =
  match input_line channel with
  | line -> line :: scan_file channel
  | exception End_of_file -> []

let read_lines filename =
  let ic = open_in filename in
  let lines = scan_file ic in
  close_in ic;
  List.rev lines

let is_digit a = match a with '0' .. '9' -> true | _ -> false

let char_to_digit c =
  if c >= '0' && c <= '9' then int_of_char c - int_of_char '0'
  else failwith "tried to convert non-digit to int"

let line_score line =
  let l = line |> String.to_seq |> List.of_seq in
  let digits = List.filter is_digit l in
  ((List.nth digits 0 |> char_to_digit) * 10)
  + (List.nth (List.rev digits) 0 |> char_to_digit)

let rec total_score lines =
  match lines with [] -> 0 | x :: xs -> line_score x + total_score xs
(*I think this syntax is pretty nice - if the list is now empty then it is
  equivalent to some first element x :: some rest of list xs which may or may not be empty*)

let answer filename =
  let lines = read_lines filename in
  total_score lines

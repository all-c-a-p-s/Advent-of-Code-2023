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
  match c with
  | '0' .. '9' -> int_of_char c - int_of_char '0'
  | _ -> failwith "tried to convert non-digit to int"

let rec last_number reversed_line =
  if is_digit reversed_line.[0] then char_to_digit reversed_line.[0]
  else
    match reversed_line with
    | s when String.starts_with s ~prefix:"eno" -> 1
    | s when String.starts_with s ~prefix:"owt" -> 2
    | s when String.starts_with s ~prefix:"eerht" -> 3
    | s when String.starts_with s ~prefix:"ruof" -> 4
    | s when String.starts_with s ~prefix:"evif" -> 5
    | s when String.starts_with s ~prefix:"xis" -> 6
    | s when String.starts_with s ~prefix:"neves" -> 7
    | s when String.starts_with s ~prefix:"thgie" -> 8
    | s when String.starts_with s ~prefix:"enin" -> 9
    | s when String.starts_with s ~prefix:"orez" -> 0
    | u when String.length u > 0 ->
        last_number (String.sub u 1 (String.length u - 1))
    | _ -> failwith "failde to find last number"

let rec first_number line =
  if is_digit line.[0] then char_to_digit line.[0]
  else
    match line with
    | s when String.starts_with s ~prefix:"one" -> 1
    | s when String.starts_with s ~prefix:"two" -> 2
    | s when String.starts_with s ~prefix:"three" -> 3
    | s when String.starts_with s ~prefix:"four" -> 4
    | s when String.starts_with s ~prefix:"five" -> 5
    | s when String.starts_with s ~prefix:"six" -> 6
    | s when String.starts_with s ~prefix:"seven" -> 7
    | s when String.starts_with s ~prefix:"eight" -> 8
    | s when String.starts_with s ~prefix:"nine" -> 9
    | s when String.starts_with s ~prefix:"zero" -> 0
    | u when String.length u > 0 ->
        first_number (String.sub u 1 (String.length u - 1))
    | _ -> failwith "failed to find first number"

let line_score line =
  let l = line |> String.to_seq |> List.of_seq in
  let digits = List.filter is_digit l in
  ((List.nth digits 0 |> char_to_digit) * 10)
  + (List.nth (List.rev digits) 0 |> char_to_digit)

let rec total_score f lines =
  match lines with [] -> 0 | x :: xs -> f x + total_score f xs

let rev x =
  let len = String.length x in
  String.init len (fun n -> String.get x (len - n - 1))

let line_score_part_two line = (first_number line * 10) + last_number (rev line)

let part_one filename =
  let lines = read_lines filename in
  total_score line_score lines

let part_two filename =
  let lines = read_lines filename in
  total_score line_score_part_two lines

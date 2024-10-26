let rec scan_file channel =
  match input_line channel with
  | line -> line :: scan_file channel
  | exception End_of_file -> []

let read_lines filename =
  let ic = open_in filename in
  let lines = scan_file ic in
  close_in ic;
  List.rev lines

let in_filter c =
  match c with
  | '0' .. '9' -> true
  | 'd' | 'g' | 'b' | ':' | ';' ->
      true (*use d for red because theres an r in green*)
  | _ -> false

let is_digit c = match c with '0' .. '9' -> true | _ -> false

let char_to_digit c =
  match c with
  | '0' .. '9' -> int_of_char c - int_of_char '0'
  | _ -> failwith "tried to convert non-digit to int"

let get_id filtered =
  match filtered with
  | a :: ':' :: _ when a != ' ' -> char_to_digit a
  | a :: b :: ':' :: _ when a != ' ' && b != ' ' ->
      (10 * char_to_digit a) + char_to_digit b
  | _ -> 100 (*there are exactly 100 games*)

let rec get_sublists filtered =
  match filtered with
  | _ :: ':' :: xs | _ :: _ :: ':' :: xs | _ :: _ :: _ :: ':' :: xs ->
      get_sublists xs
  | [] -> []
  | a :: b :: ';' :: xs -> [ a; b ] :: get_sublists xs
  | a :: b :: c :: ';' :: xs -> [ a; b; c ] :: get_sublists xs
  | a :: b :: c :: d :: ';' :: xs -> [ a; b; c; d ] :: get_sublists xs
  | a :: b :: c :: d :: e :: ';' :: xs -> [ a; b; c; d; e ] :: get_sublists xs
  | a :: b :: c :: d :: e :: f :: ';' :: xs ->
      [ a; b; c; d; e; f ] :: get_sublists xs
  | a :: b :: c :: d :: e :: f :: g :: ';' :: xs ->
      [ a; b; c; d; e; f; g ] :: get_sublists xs
  | a :: b :: c :: d :: e :: f :: g :: h :: ';' :: xs ->
      [ a; b; c; d; e; f; g; h ] :: get_sublists xs
  | a :: b :: c :: d :: e :: f :: g :: h :: i :: ';' :: xs ->
      [ a; b; c; d; e; f; g; h; i ] :: get_sublists xs
  | _ ->
      failwith
        ("failed to get sublists from" ^ String.of_seq (List.to_seq filtered))

let is_colour c = match c with 'd' | 'g' | 'b' -> true | _ -> false

let rec get_colour colour list =
  match list with
  | [] -> 0
  | c :: a :: x :: _ when is_colour x && c == colour -> char_to_digit a
  | c :: a :: b :: x :: _ when is_colour x && c == colour ->
      (char_to_digit b * 10) + char_to_digit a (*reversed*)
  | [ c; a ] when c == colour -> char_to_digit a
  | [ c; a; b ] when c == colour ->
      (char_to_digit b * 10) + char_to_digit a (*reversed*)
  | _ :: xs -> get_colour colour xs

let parse_sublist sl =
  (*reverse sublist so that colour precedes number so its easier*)
  let red = get_colour 'd' (List.rev sl) in
  let green = get_colour 'g' (List.rev sl) in
  let blue = get_colour 'b' (List.rev sl) in
  [ red; green; blue ]

let rec highest_nth n lists highest_so_far =
  match lists with
  | [] -> highest_so_far
  | x :: xs -> highest_nth n xs (max highest_so_far (List.nth x n))

let line_score line =
  let filtered =
    List.of_seq (Seq.filter in_filter (String.to_seq (line ^ ";")))
  in
  (*add semicolon so my parser works*)
  let id = get_id filtered in
  let sublists = List.map parse_sublist (get_sublists filtered) in
  let max_red = highest_nth 0 sublists 0 in
  let max_green = highest_nth 1 sublists 0 in
  let max_blue = highest_nth 2 sublists 0 in

  if max_red <= 12 && max_green <= 13 && max_blue <= 14 then id else 0

let line_score_part_two line =
  let filtered =
    List.of_seq (Seq.filter in_filter (String.to_seq (line ^ ";")))
  in
  let sublists = List.map parse_sublist (get_sublists filtered) in
  highest_nth 0 sublists 0 * highest_nth 1 sublists 0 * highest_nth 2 sublists 0

let rec total_score f lines =
  match lines with [] -> 0 | x :: xs -> f x + total_score f xs

let part_one filename =
  let lines = read_lines filename in
  total_score line_score lines

let part_two filename =
  let lines = read_lines filename in
  total_score line_score_part_two lines

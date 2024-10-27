let rec scan_file channel =
  match input_line channel with
  | line -> line :: scan_file channel
  | exception End_of_file -> []

let read_lines filename =
  let ic = open_in filename in
  let lines = scan_file ic in
  close_in ic;
  List.rev lines

let numbers s =
  s |> String.split_on_char ' '
  |> List.filter (fun x -> String.length x > 0)
  |> List.map int_of_string

let rec count_repeats l1 l2 =
  match l1 with
  | [] -> 0
  | x :: xs -> (if List.mem x l2 then 1 else 0) + count_repeats xs l2

let rec pow a n = match n with 0 -> 1 | _ -> a * pow a (n - 1)

let line_score_part_one l =
  let parts = l |> String.split_on_char ':' in
  let number_sets = List.nth parts 1 |> String.split_on_char '|' in
  let winning_numbers = numbers (List.nth number_sets 0) in
  let card_numbers = numbers (List.nth number_sets 1) in
  let repeats = count_repeats winning_numbers card_numbers in
  match repeats with 0 -> 0 | x -> pow 2 (x - 1)

let rec total_score f lines =
  match lines with [] -> 0 | x :: xs -> f x + total_score f xs

let part_one filename =
  let lines = read_lines filename in
  total_score line_score_part_one lines

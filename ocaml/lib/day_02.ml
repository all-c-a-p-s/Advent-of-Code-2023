let rec scan_file channel =
  match channel with
  | line -> line :: scan_file channel
  | exception End_of_file -> []

let read_lines filename =
  let ic = open_in filename in
  List.rev (scan_file ic)

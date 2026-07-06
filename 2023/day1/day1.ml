let read_lines filename =
  let ic = open_in filename in
  let rec aux acc =
    try
      let line = input_line ic in
      aux (line :: acc)
    with End_of_file ->
      close_in ic;
      List.rev acc
  in
  aux []

let calibration_value s =
  let first = ref None in
  let last = ref None in

  String.iter (fun ch ->
    if ch >= '0' && ch <= '9' then
      let d = Char.code ch - Char.code '0' in
      begin
        if !first = None then first := Some d;
        last := Some d
      end
  ) s;

  match !first, !last with
  | Some f, Some l -> f * 10 + l
  | _ -> 0

let () =
  let lines = read_lines "inputs.txt" in
  let ans =
    List.fold_left (fun acc line -> acc + calibration_value line) 0 lines
  in
  Printf.printf "%d\n" ans

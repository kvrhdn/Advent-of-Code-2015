use aoc_runner_derive::aoc;
use itertools::Itertools;

type Seat = (u32, u32);

fn parse_seat(input: &str) -> Seat {
    let (row, col) = input.split_at(7);

    (parse_binary(row), parse_binary(col))
}

fn parse_binary(input: &str) -> u32 {
    let mut weight = 2u32.pow(input.len() as u32 - 1);
    let mut value = 0;

    for c in input.chars() {
        if c == 'B' || c == 'R' {
            value += weight;
        }
        weight /= 2;
    }

    value
}

fn seat_id(seat: Seat) -> u32 {
    (seat.0 * 8) + seat.1
}

fn process_input(input: &str) -> impl Iterator<Item = Seat> + '_ {
    input.lines().map(parse_seat)
}

#[aoc(day5, part1)]
fn solve_part1(input: &str) -> u32 {
    process_input(input).map(seat_id).max().unwrap()
}

#[aoc(day5, part2)]
fn solve_part2(input: &str) -> u32 {
    let mut seats = process_input(input).collect::<Vec<_>>();

    // sort by row
    seats.sort_unstable_by_key(|seat| seat.0);

    let min_row = seats[0].0;
    let max_row = seats[seats.len() - 1].0;

    for (row, row_seats) in &seats.into_iter().group_by(|seat| seat.0) {
        if row == min_row || row == max_row {
            continue;
        }

        let mut row_seats = row_seats.collect::<Vec<_>>();

        if row_seats.len() != 8 {
            // sort by col
            row_seats.sort_unstable_by_key(|seat| seat.1);

            let missing_seat_id = row_seats
                .iter()
                .enumerate()
                .find(|&(i, seat)| i as u32 != seat.1)
                .map(|(i, seat)| seat_id((seat.0, i as u32)))
                .unwrap();

            return missing_seat_id;
        }
    }

    panic!("could not find a missing seat");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn examples_part1() {
        let examples = vec![
            ("BFFFBBFRRR", (70u32, 7u32), 567u32),
            ("FFFBBBFRRR", (14, 7), 119),
            ("BBFFBBFRLL", (102, 4), 820),
        ];

        for (input, expected_seat, expected_seat_id) in examples {
            let seat = parse_seat(input);

            assert_eq!(seat, expected_seat);
            assert_eq!(seat_id(seat), expected_seat_id);
        }
    }

    #[test]
    fn real_input() {
        let input = include_str!("../input/2020/day5.txt");

        assert_eq!(solve_part1(input), 828);
        assert_eq!(solve_part2(input), 565);
    }
}

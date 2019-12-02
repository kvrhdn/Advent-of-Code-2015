use wasm_bindgen::prelude::*;

#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

/// Calculate the sum of the fuel requirements for all of the modules on the spacecraft.
/// See: https://adventofcode.com/2019/day/1
#[wasm_bindgen]
pub fn part1(input: &str) -> Result<i32, JsValue> {
    let sum = parse_input(input)?
        .iter()
        .map(|&mass| fuel_required(mass))
        .sum();

    Ok(sum)
}

/// Calculate the sum of the fuel requirements for all of the modules on the spacecraft, also
/// taking into account the mass of the added fuel.
/// See: https://adventofcode.com/2019/day/1
#[wasm_bindgen]
pub fn part2(input: &str) -> Result<i32, JsValue> {
    let sum = parse_input(input)?
        .iter()
        .map(|&mass| total_fuel_required(mass))
        .sum();

    Ok(sum)
}

/// Parse the input (a list of modules their masses) as a list of integers.
fn parse_input(input: &str) -> Result<Vec<i32>, &'static str> {
    input
        .lines()
        .map(|l| {
            l.parse::<i32>()
                .map_err(|_| "could not parse input as integers")
        })
        .collect()
}

/// Fuel required to carry the given mass.
fn fuel_required(mass: i32) -> i32 {
    // integer dision already rounds down
    (mass / 3) - 2
}

/// Total fuel required to carry the given mass, including the mass of the fuel itself.
fn total_fuel_required(mass: i32) -> i32 {
    let fuel = fuel_required(mass);
    if fuel <= 0 {
        return 0;
    }

    fuel + total_fuel_required(fuel)
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn table_test_fuel_required() {
        assert_eq!(fuel_required(12), 2);
        assert_eq!(fuel_required(14), 2);
        assert_eq!(fuel_required(1969), 654);
        assert_eq!(fuel_required(100_756), 33_583);
    }

    #[test]
    fn table_test_total_fuel_required() {
        assert_eq!(total_fuel_required(14), 2);
        assert_eq!(total_fuel_required(1969), 966);
        assert_eq!(total_fuel_required(100_756), 50_346);
    }
}

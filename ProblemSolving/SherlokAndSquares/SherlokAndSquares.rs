use std::io;

/*
  FIRST ATTEPMT

  This function
    1- Gets the square root of each number in the input range (which can be between 1 and 10^9)
    2- subsctract the integer part from the result obtained in the previous step
    3- if the result is not 0 the number is not an square number

  This method is to slow

*/
// fn squares(s: i64, e: i64) -> i64 {
//     let mut squares: i64 = 0;
//     let upperbound = ((e as f64).sqrt() as i64) + 2;
//     for i in s..upperbound {
//         let current: f64 = (i as f64).sqrt();
//         let current_as_int: i64 = current as i64;
//         if current - (current_as_int as f64) == 0.0 {
//             squares = squares + 1;
//         }
//     }
//     squares
// }

/*
 SECOND ATTEPMT

 lets imagine we precalculate in an array all perfect squares up to 400
   indices   0|1|2| 3| 4| 5| 6| 7| 8|  9| 10| 11| 12| 13| 14| 15| 16| 17| 18|
   valeurs   1|4|9|16|25|36|49|64|81|100|121|144|169|196|225|256|289|324|361|

   we notice that the square root of a perfect square -1 is equal to the index of its
   corresponding square root ex sqrt(4) = 2 and 2-1 = 1 which is the index of 4 in our
   generated square roots array

   we notice that if we take the square root of any number that is not a perfect square
   the result will a decimal number which integer part is equal to the sqrt of the perfect square
   that  precedes it EX sqrt(17) = 4.1231 (the preceding perfect square was 16 with a sqrt of 4)

   we realize that we don't need to precalculate the square roots array we can simply take the
   square root of the lower and upper boundaries and substract them (upper -lower) to get the
   number of squares in that range. Since the upper and lower bundaries are included we have to know
   if the lower boundary is itself a perfect square because if we just simply return the substraction
   and the lower boundary is a perfect square we won't be including it, so we have to correct for it.
*/

fn squares(s: i64, e: i64) -> i64 {
    let sidx: i64 = (s as f64).sqrt() as i64;
    let eidx: i64 = (e as f64).sqrt() as i64;

    let result: i64 = eidx - sidx;

    // if lower boundary is a perfect square return the result + 1
    if is_perfect_square(s) {
        return result + 1;
    }

    result
}

// returns true if a number is a perfect square, false otherwise
fn is_perfect_square(n: i64) -> bool {
    let n_sqrt: f64 = (n as f64).sqrt();
    let n_sqrt_int: i64 = n_sqrt as i64;

    return n_sqrt - (n_sqrt_int as f64) == 0.0;
}

fn main() {
    let mut tc_string: String = String::new();
    io::stdin().read_line(&mut tc_string).unwrap();
    let mut tc: i32 = tc_string.trim().parse().unwrap();
    while tc > 0 {
        let mut numbers: String = String::new();
        io::stdin().read_line(&mut numbers).unwrap();
        let numbers_arr: Vec<&str> = numbers.split(" ").collect();
        let start: &str = numbers_arr[0].trim();
        let end: &str = numbers_arr[1].trim();
        println!("{}", squares(start.parse().unwrap(), end.parse().unwrap()));
        tc = tc - 1;
    }
}


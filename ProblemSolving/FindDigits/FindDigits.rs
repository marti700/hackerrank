use std::io;

/*
  a number % 10 will yied the number last digit
  a number / 10 will be the number without its last digit
  This function basically loops over the digits of a number
  asking if the reminder of the number by the current digit is 0
 */
fn find_digits(n: i64) -> u32 {
    let mut divisors: u32 = 0;
    let mut temp = n;
    while temp > 0 {
      let reminder: i64 = temp % 10;
        // first condition avoids a devide by zero panic
        if (reminder != 0) && (n % reminder == 0) {
            divisors = divisors + 1;
        }
        temp = temp / 10;
    }
    divisors
}

fn main() {
    let mut tc_string: String = String::new();
    io::stdin().read_line(&mut tc_string).unwrap();
    let mut tc: i32 = tc_string.trim().parse().unwrap();
    while tc > 0 {
        let mut number: String = String::new();
        io::stdin().read_line(&mut number).unwrap();
        let n: i64 = number.trim().parse().unwrap();
        println!("{}", find_digits(n));
        tc = tc - 1;
    }
}


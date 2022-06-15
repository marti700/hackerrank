use std::io;

// Hanldes the cases where the number of candies are less the the number of prisoners from the get go!
fn last_candy(prisoner: i64, candy: i64, start: i64) -> i64 {
    let til_end: i64 = prisoner - start;

    // when the guard have to restart the candies distrbution from the prisoner wiht label 1
    // remember that at this point we can say that distribution start as one and that there are
    // less candies than prisoners
    if candy > (til_end + 1) {
        return candy - til_end - 1;
    } else {
        return (start - 1) + candy;
    }
}

// Let | be a candy and let the numbers be the label of the prisoners (e.x. 1,2,3 are the prisoner 1 to 3)
// the difference from the start (the prisoner from where the distribution starts) to the last prisoner are subsctracted
// to the number of candies. Now the start index is on prisoner 1
// E.X: for prisoners = 6 and candies = 5 and distribution starts at 4
// (4 - 6) + 1 == 3 (the plus 1 is because a candy were also given to prisoner 4)
// the handout looks like this
// 1 2 3 4 5 6
//       | | |
// so, the are 5-3 = 2 candies left to be distributed.
// We can also think that now the start distribution index start at prisoner 1 for the distribution of the remaining candies
// In this case we can just return 2 which is the number of candies left to distribute (5-3 = 2)
//
// when the number of candies is less than the number of prisoners and the candy distribution
// starts at an index from which the guard will runout of candies before reaching the end
// E.X: for prisoners = 6 and candies = 4 and distribution starts at 3
// the handout looks like this
// 1 2 3 4 5 6
//     | | | |
// In this case we return the number of candies - 1 (the minus 1 is because a candy is given to the prisoner 3)
fn save_the_prisoner(prisoner: i64, candy: i64, start: i64) -> i64 {
    // if the number of candies is less than the number of prisoners
    if candy - prisoner < 0 {
        return last_candy(prisoner, candy, start);
    }

    // gets the number of candies distributed from the starting prisoner to the last one
    let candy_left: i64 = candy - ((prisoner - start) + 1);

    // substract the candies distributed in the previous line from the total number of candies
    let result: i64 = candy_left - ((candy_left / prisoner) * prisoner);
    // if the result is zero means that all the candies have been distribuited and that the last candy was given to the last prisoner
    // so the numer prisoner is returned
    if result == 0 {
        return prisoner;
    }
    return result;
}

fn main() {
    let mut tc = String::new();

    io::stdin().read_line(&mut tc).expect("a number");

    let mut _tc: u32 = tc.trim().parse().unwrap();

    while _tc > 0 {
        let mut input = String::new();

        io::stdin().read_line(&mut input).expect("a number");

        let _input: Vec<i64> = input
            .split(" ")
            .map(|e: &str| e.trim().parse::<i64>().unwrap())
            .collect();

        println!("{}", save_the_prisoner(_input[0], _input[1], _input[2]));
        // The same result can be obtained using this formula: (M - 1 + S - 1) % N + 1)
        // Where M is the number of sweets S the starting position and N the number of prisoners
        // this formula sums the start + the number of candies (the -1(s) are to take into account the fact that in this problem prisoner indexes are inclusive)
        // taking the modulo N have the effect of looping back to the first prisoner (the actual first not where the guard started). this calculation
        // returns the index before at which the last candy was handedout and because of this 1 is added to correct for this error
        // println!("{}", ((_input[1]-1) + (_input[2]-1)) % _input[0] +1);


        _tc -= 1;
    }
}

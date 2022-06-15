object Solution {
  def getIntervals(lower: Int, upper: Int): List[Double] =
    (lower.toDouble to upper.toDouble by 0.001).toList.tail

  /** calculates the area of a rectangle in a given X of a polynomial function
    *
    * @param coefficients a list of number representing the function coefficients
    * @param powers a list of numbers representing the x variable powers
    * @param x the x on wich the function will be evaluated
    */
  def f(coefficients: List[Int], powers: List[Int], x: Double): Double = {
    coefficients.zip(powers).map(n => (n._1 * scala.math.pow(x,n._2))).sum * 0.001
  }

  /**
    * Calculates the volume of a cylinder in a given X of a polynomial function
    *
    * @param coefficients a list of number representing the function coefficients
    * @param powers a list of numbers representing the x variable powers
    * @param x the x on wich the function will be evaluated
    * @return a double representing the value of a
    */
  def area(coefficients: List[Int], powers: List[Int], x: Double): Double = {
    val funcValueOnx = coefficients.zip(powers).map(n => n._1 * scala.math.pow(x,n._2)).sum
    math.pow(funcValueOnx,2)* 0.001 * scala.math.Pi
  }

  /**
    * integrates a polinomial function using the riemman sum method
    *
    * @param func a fuction to evaluate the x values of a polinomial function
    * @param upper upper limit
    * @param lower lower limit
    * @param coefficients the coefficients of the polinomial
    * @param powers the powers of the polinomial
    * @return a double representing the result of the integration
    */
  def summation(
      func: (List[Int], List[Int], Double) => Double,
      upper: Int,
      lower: Int,
      coefficients: List[Int],
      powers: List[Int]
  ): Double = {
    getIntervals(lower, upper).map(x => func(coefficients, powers, x)).sum
  }
}

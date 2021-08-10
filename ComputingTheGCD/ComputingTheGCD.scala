object Solution {
  def gcd(x: Int, y: Int): Int = {
    if (x < y) return gcd(y,x)
    if (x % y == 0) return y
    return gcd(y, (x % y))
  }
}

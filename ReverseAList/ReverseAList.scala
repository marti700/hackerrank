object ReverseAList {
  def f(arr: List[Int]): List[Int] = {
    def go(l: Int) : List[Int] = {
      l match {
        case -1 => Nil
        case _ => arr(l)::go(l-1)
      }
    }
    go(arr.length -1)
  }
}
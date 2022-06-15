object Solution {
  def mingle (s1: String, s2:String, result: String) : String = {
    s1.zip(s2).toList.map(x => x._1.toString + x._2.toString).foldLeft("") {(acc, e) => acc + e}
    //the below recursive solution do not pass the last test case because of a timeout apparently it is slow
    // if (s1 == "") return result
    // mingle (s1.tail, s2.tail, result + s1.head.toString + s2.head.toString)
  }

  def main(args: Array[String]) ={
    val input = io.Source.stdin.getLines().toList
    print(mingle(input(0), input(1), ""))
  }
}
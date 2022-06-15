object Solution {
  //failed recursive solution
  // def prefix (s1: List[Char], s2: List[Char]): String = {
  //   (s1, s2) match {
  //     case (Nil, _) => ""
  //     case (_, Nil) => ""
  //     case (x::xs, y::ys) if x == y => x.toString + prefix(xs, ys)
  //     case (_, _) => ""
  //   }
  // }

def prefix (s1: String, s2: String): String = {
  s1.zip(s2).toList.takeWhile(x => x._1 == x._2).map(x=> x._1).mkString("")
}
  def getStrPostfix(str: String, prefLength: Int) : String = {
    str.drop(prefLength).length.toString + " " + str.drop(prefLength) + "\n"
  }


  def buildOutput(str1: String, str2: String) : String ={
    val pref = prefix(str1, str2)
    pref.length().toString() + " " + pref + "\n" + getStrPostfix(str1, pref.length()) + getStrPostfix(str2, pref.length())
  }

  def main(args: Array[String]) : Unit = {
    val input = io.Source.stdin.getLines().toList
    var str1 = input(0)
    var str2 = input(1)
    print(buildOutput(str1,str2))

  }
}
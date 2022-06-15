-- String-o-permutation
stringOPermutation :: String -> String
stringOPermutation [] = []
stringOPermutation (x:xs) = head xs:x:stringOPermutation (tail xs)

main = interact $ unlines . map stringOPermutation . tail . lines
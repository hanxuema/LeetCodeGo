package main

func addBinary(a string, b string) string {
    if a == "0" {
        return b
    } else if b == "0"{
        return a
    }
    res := ""
    lenA := len(a) -1 
    lenB := len(b) -1
    
    carry := 0
    for lenA>= 0 || lenB >= 0  {
        na := 0
        if (lenA >=0 ){
            na , _= strconv.Atoi(string(a[lenA]))
            fmt.Printf("ba %+v\n", na)
        }
        nb := 0
        if lenB >= 0{
            nb , _= strconv.Atoi(string(b[lenB]))
            fmt.Printf("nb %+v\n", nb)
        }
        temp := na + nb + carry
      
        if temp == 2{
            carry = 1
            temp = 0
        }else if temp == 3{
            carry = 1
            temp = 1
        } else {
            carry = 0 
        }
      
        res = strconv.Itoa(temp) + res
        lenA--
        lenB--
    }
    if carry == 1 {
        res = "1"+ res
    }
    return res
}
 
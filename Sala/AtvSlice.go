package main
import "fmt"

type Pessoa struct{
    Peso float64
    Altura float64
    IMC float64
}

func main() {
    
  pessoas := make([]Pessoa, 0)
  
  for{
      
      p := Pessoa{}
      fmt.Println("Digite o peso da pessoa: ")
      fmt.Scan(&p.Peso)
      
      fmt.Println("Digite a altura da pessoa: ")
      fmt.Scan(&p.Altura)
      
      p.IMC = p.Peso / (p.Altura * p.Altura)
      
      pessoas = append(pessoas, p)
      
      var op string
      
      fmt.Println("Inserir mais pessoas?: y -> sim")
      fmt.Scan(&op)
      
      if(op != "y"){
        break;
      }
  }
  
  fmt.Println(pessoas)
}

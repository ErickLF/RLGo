package main

import( 
	"fmt"
	"math"
	"os"
	"bufio"
)

const(
	goal = 100
	ph = 0.55
	gamma =1-ph
	theta = 0.0000000001
)

type gambler struct{
	states [] float64
	policys [] float64
}

func main (){	
	gambler := gambler{}
	gambler.states =  make([]float64, goal+1) 	
	gambler.policys = make([]float64, goal+1)
	gambler.states[goal] = 1.0
	gambler.value_iteration()
	
	
	gambler.file()
	gambler.fileP()
}

func (g *gambler) value_iteration(){
	delta := 1.0
	for delta > theta{ 
		delta = 0
		for s:=1;s<goal;s++{
			v:=g.states[s]
			g.states[s],_=max_actions(s,g.states)
			delta+=math.Abs(v-g.states[s])
		}
	}
	
	for s := 1; s < goal; s++ {
		_, g.policys[s] = max_actions(s, g.states)
	}
}

func max_actions(s int,sv []float64) (float64,float64){
	n:=int(math.Min(float64(s),float64(goal-s)))
	outcome:=ph*sv[s]+gamma*sv[s]
	optValue:=0
	for action:=1;action<=n;action++{
		aux:=ph*sv[s+action]+gamma*sv[s-action]
		if aux > outcome{
			outcome = aux
			optValue = action
		}
	}
	return outcome,float64(optValue)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func(g*gambler) file(){
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
    }
    defer file.Close()

    w := bufio.NewWriter(file)
	
	for i:=1;i<len(g.states);i++{
    	fmt.Fprintf(w,"%v \t %e \n", i, g.states[i])
	}
	
    w.Flush()
}

func(g*gambler) fileP(){
	file, err := os.OpenFile("outputP.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
    }
    defer file.Close()

    w := bufio.NewWriter(file)
	
	for i:=1;i<len(g.states);i++{
    	fmt.Fprintf(w,"%v \t %e \n", i,g.policys[i])
	}
	
    w.Flush()
}
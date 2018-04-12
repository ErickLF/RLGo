package main
//----------------------------------------------------------------------
import (
    "fmt"
    "math/rand"
    "time"
)
//----------------------------------------------------------------------
const (
      up int = iota
      down
      left
      right
      worldheight int = 4 //world height
      worldwidth int = 12 //world width
      numactions int = 4 //number of actions
  )
//----------------------------------------------------------------------
type QLearning struct{
    Q [][]float64
    alpha float64  //exange rate
    gamma float64  //how much does the past matter?
    epsilon float64 
    startX int
    startY int
    goalX int 
    goalY int 
}
//----------------------------------------------------------------------
func main() {
    rand.Seed(time.Now().Unix())
    ql:=QLearning{}
    //states
    ql.startX,ql.startY = 0,0
    ql.goalX,ql.goalY = 11,0
    ql.alpha,ql.gamma,ql.epsilon=0.5,1.0,0.1
    
    ql.Q = make([][]float64,numactions) //4 numero de acciones
    
    for i:=0;i<numactions;i++{ //initialize the size of the matrix
        ql.Q[i] = make([]float64,worldheight*worldwidth)
    }
    
    for i:=0;i<numactions;i++{  
        for j:=0;j<worldheight*worldwidth;j++{
            ql.Q[i][j]=rand.Float64()
        }
    }
    
    for a:=0;a<4;a++{ //we fill our matrix with zeros
        ql.Q[a][ql.goalY*worldheight+ql.goalX] = 0
    }
    ql.IniciarQ()
    ql.Print()
}
//----------------------------------------------------------------------
func (ql* QLearning) IniciarQ(){
    episode:=1000 
    for i:=0;i<episode;i++{//Repeart (for each step of episode )
        //Initialize S
        sx,sy:=ql.startX,ql.startY
        //loop until reaches goal
        for sx!=ql.goalX || sy!=ql.goalY{ 
            //choose A from S using policy derived from Q
            action:=ql.chooseAction(sx,sy)
            //take action A,observe R,S' //sq -> single quote
            R,sqsx,sqsy:=ql.move(sx,sy,action)
            QSA:=ql.Q[action][sy*worldwidth+sx]
            maxAction:=ql.getAction(sx,sy)
            sqQSA:=ql.Q[maxAction][sqsy*worldwidth+sqsx]
            //Q(S,A)<-Q(S,A)+alpha*[R+gammaMax v(a) Q(S',a)-Q(S,A)]
            Q:=QSA+ql.alpha*(R+ql.gamma*(sqQSA-QSA))
            ql.Q[action][sy*worldwidth+sx]=Q
            //S<-S'
            sx=sqsx
            sy=sqsy
            //until S is terminal
        }
    }
}
//----------------------------------------------------------------------
func (q *QLearning) chooseAction(x,y int) int{ //epsilon greedy
	if rand.Float64() < 1-q.epsilon {
		return q.getAction(x, y)
	}
  
	return rand.Intn(numactions)
}
//----------------------------------------------------------------------
func (ql* QLearning) getAction(x,y int) int {//both parameters are int and returns int
  ids := y*worldwidth+x //id of square
	max := ql.Q[0][ids]
	action := 0
	for i := 0; i < numactions; i++ {
		if max < ql.Q[i][ids] {
			max = ql.Q[i][ids]
			action = i
		}
	}
  
	return action
}
//----------------------------------------------------------------------
func (ql* QLearning) move(x,y,action int) (float64,int,int){
  qx,qy:=x,y
  
  switch action{
    case up:
      if y!=worldheight-1{ qy=y+1 }
    case down:
      if y!=0{ qy=y-1 }
    case left:
       if x!=0{ qx=x-1 }
    case right:
      if x!=worldwidth-1{ qx=x+1 }
  }
  
  if qy == 0 && qx >= 1 && qx < worldwidth-1 { //if it falls in the cliff
		return -100.0, 0, 0
  }else if qy==ql.goalY && qx==ql.goalX{ //If reaches goal
		return 0.0, ql.goalX, ql.goalY
	}
  
  return -1.0,qx,qy//if doesnt reach goal OR doesnt fall on THE CLIFF just continue on board
}
//----------------------------------------------------------------------
/**Funcion super bonis hecha originalmente por Github:@JoseGurrola**/
func (q *QLearning) Print() {
	for i := worldheight - 1; i >= 0; i-- {
		for j := 0; j < worldwidth; j++ {
			if i != q.goalY || j != q.goalX {
				if j > 0 && j < worldwidth-1 && i == 0 {
					fmt.Print("<- ")
				} else {
					switch q.getAction(j, i) {
					case up:
						fmt.Print("U  ")
					case down:
						fmt.Print("D  ")
					case left:
						fmt.Print("L  ")
					case right:
						fmt.Print("R  ")
					}
				}
			} else {
				fmt.Print("G ")
			}
		}
		fmt.Println("")
	}
}
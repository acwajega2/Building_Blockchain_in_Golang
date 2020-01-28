package main
import(
	"fmt"
	"bytes"
	"crypto/sha256"
)
//----creating  a block struct
type Block struct{
	Hash []byte
	Data []byte
	PrevHash []byte

}


//---Blockchain struct
type BlockChain struct{
	blocks []*Block
}

//-----Method to implement a block chain----
func (b *Block) DeriveHash(){
	info :=bytes.Join([][]byte{b.Data,b.PrevHash},[]byte{})
	hash :=sha256.Sum256(info)
	b.Hash = hash[:]
}


//-----creating a block
func CreateBlock(data string, prevHash []byte) *Block{
	block :=&Block{[]byte{},[]byte(data),prevHash}
	block.DeriveHash()
	return block
	  
}

//----Method to add Block
func (chain *BlockChain) AddBlock (data string){
	prevBlock :=chain.blocks[len(chain.blocks) -1]
	new :=CreateBlock(data,prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}


//----fucntion to create the first block
func Genenis() *Block{
	return CreateBlock("Genesis",[]byte{})

}



//----Creating the first blockchain
func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genenis()}}
}

func main() {
	chain :=InitBlockChain()

	chain.AddBlock("First Block Chain after Genesis")
	chain.AddBlock("Second Block Chain after Genesis")
	chain.AddBlock("Third Block Chain after Genesis")
	chain.AddBlock("Fourth Block Chain after Genesis")
	

	//---preating out fiels from each block
	for _,block :=range chain.blocks{

		fmt.Printf("Previous Hash: %x \n",block.PrevHash)
		fmt.Printf("Data in Block %s \n",block.Data)
		fmt.Printf("Has: %x\n",block.Hash)
	}
}
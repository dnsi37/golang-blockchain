package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	//fmt.Printf("target: %d", target)
	//fmt.Println(uint(256 - Difficulty))

	pow := &ProofOfWork{b, target}

	return pow

}

// 난수를 넣어 임의 data 생성
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	fmt.Println("Run func begins")
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data) // 임의 데이터를 해쉬화

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		//fmt.Println("hash")
		//fmt.Printf("\r%d", &intHash)
		//fmt.Println("inthash")
		//fmt.Printf("\r%d", pow.Target)
		//fmt.Println("target")
		if intHash.Cmp(pow.Target) == -1 {
			//해쉬화 된 데이터가 작업증명 타겟보다 작을때
			// 해쉬 데이터는 256bit 이며 타겟이 256-난이도 로 왼쪽 쉬프트 하기때문에 결과 값은 무조건 난이도 만큼 0을 가질 수 밖에 없다.
			break
		} else {
			nonce++
		}
	}
	fmt.Println("Run func ends")

	return nonce, hash[:]

}

func (pow *ProofOfWork) Validate() bool {
	var initHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])

	return initHash.Cmp(pow.Target) == -1

}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic((err))
	}
	return buff.Bytes()
}

package main

import (
	"os"
	"fmt"	
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

func main() {
	action := os.Args[1]
	 key := os.Args[2]
	 input := os.Args[3]
	 aesBlock,err := aes.NewCipher([]byte(key))
	 if err != nil{
		 panic(err)
	 }
	 aesgcm,err := cipher.NewGCM(aesBlock)
	 if err != nil{
		 panic(err)
	 }
	 nonce := []byte("012345678912")
	 if action == "en"{
		 encrypted := aesgcm.Seal(nil,nonce,[]byte(input),nil)
		 fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
	 }else if action == "de"{
		 cipherText,err := base64.StdEncoding.DecodeString(input)
		 if err != nil{
			 panic(err)
		 }
		 text,err := aesgcm.Open(nil,nonce,cipherText,nil)
		 if err != nil{
			 fmt.Printf("Decrypt failed: %v\n",err)
			 os.Exit(1)
		 }
		 fmt.Println(string(text))
	 }else{
		 fmt.Println("UnDefined action")
		 os.Exit(1)
	 }
}

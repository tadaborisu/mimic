package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var key string = "1234123412341234"
var massage string
var inputText string

func hide(key string) string {

	massage = "Wrong len of key: " + fmt.Sprint(len(key))

	//Key have a wrong len
	if len(key) != 16 {
		return massage
	}

	//Folder files didnt exist
	massage = "Missing `files` folder"
	if _, err := os.Stat("files//"); os.IsNotExist(err) {
		return massage
	}

	//Zip data
	zipit("files\\", "filesToEncrypt.zip")

	//Open file
	file, _ := os.Open("filesToEncrypt.zip")

	//Read file data
	read, _ := ioutil.ReadAll(file)

	//Encrypt file
	ciphertext := encrypt(string(read), key)

	//Write encypted data to txt
	writeToFile(ciphertext, "encryptedData.txt")

	//Zip encrypted data
	zipit("W:\\Mimic\\encryptedData.txt", "encryptedData.zip")

	//Absolute path to hide.cmd
	path, _ := filepath.Abs("hide.cmd")

	//Set filepath to cmd command
	cmd := exec.Command(path)

	//Start cmd command
	cmd.Start()

	//Close file
	file.Close()

	//All is ok
	massage = "Hided in normal.png"
	return massage
}

func dcrypt() string {
	massage = "Wrong len of key: " + fmt.Sprint(len(key))
	if len(key) != 16 {
		return massage
	}
	cryptFile, _ := readFromFile("encryptedData.txt")
	decriptedData := decrypt(string(cryptFile), key)
	writeToFile(decriptedData, "decryptedData.zip")
	massage = "Decrypted"
	return massage
}

func clean() {
	os.Remove("encryptedData.zip")
	os.Remove("encryptedData.txt")
	os.Remove("filesToEncrypt.zip")
	os.Remove("normal.png")
	os.Remove("decryptedData.zip")
	dir, _ := ioutil.ReadDir("files//")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"files", d.Name()}...))
	}
}

func main() {
	gui()
}

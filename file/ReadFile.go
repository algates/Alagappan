package file

import (
	"Alagappan/compare"
	"bitbucket.org/swigy/hudor/genericUtilities"
	"bufio"
	"bytes"
	"io"
	"os"
	"errors"
)

func ParseFile(filename string) ([]string, string, error) {

	var lines []string
	// Opening a file
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return lines, "Not able to open the file " + filename, err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			if !isPrefix {
				break
			}

			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		// converting buffer to string
		line := buffer.String()
		lines = append(lines, line)
	}

	return lines, "Successfully read and parsed", nil
}

func LoadFileInChannels(filePath1 string, filePath2 string) (chan compare.Compare, int64, error) {

		values1,message1,err1 := genericUtilities.ParseFile(filePath1)
		values2,message2,err2 := genericUtilities.ParseFile(filePath2)
		if err1 != nil {
			return nil, int64(0), errors.New(message1)
		}
		if err2 != nil {
			return nil, int64(0), errors.New(message2)
		}
		length:=0
		if len(values1) >= len(values2) {
			length = len(values1)
		}

		if len(values1) <= len(values2) {
			length = len(values2)
		}

		channel := make(chan compare.Compare, length)

		// pushing urls values in async mode
		go func() {
			for i:=0; i<length; i++ {
				url1:=""
				if i < len(values1) {
					url1 = values1[i]
				}
				url2:=""
				if i < len(values2) {
					url2 = values2[i]
				}


				compObj:=compare.Compare{
					URL1: url1,
					URL2: url2,
					Resp1:[]byte{},
					Resp2:[]byte{},
				}
				channel <- compObj
			}
		}()

	return channel, int64(length), nil
}



/*
* Echotron-GO
* Copyright (C) 2018  Nicolò Santamaria, Alessandro Ianne
*/

package plugins

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	)

/*
*	It returns a string array containing the paths of the matching files
*/

func SearchFile(dirName string, file string) []string {
    files, err := ioutil.ReadDir(dirName)
    if err != nil {
        log.Fatal(err)
    }
    result := make([]string, 0)
    for _, f := range files {
		if !f.IsDir() {
    		if strings.Contains(f.Name(), file) {
        		result = append(result, fmt.Sprintf(dirName + f.Name()))
        	}
        } else {
          	result = append(result, SearchFile(fmt.Sprintf(dirName + f.Name() + "/"), file)...)
        }
    }
    return result
}


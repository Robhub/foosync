package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ReadTSV(path string) (map[string]int64, error) {
	tracksFirstPlayed := map[string]int64{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "\t")
		if len(values) != 4 {
			continue
		}
		firstPlayed, err := time.Parse(time.RFC3339, values[3])
		if err != nil {
			return nil, err
		}
		//fmt.Printf("%v\n", values)
		trackid := strings.ToLower(values[0] + "\t" + values[1] + "\t" + values[2])
		tracksFirstPlayed[trackid] = firstPlayed.Unix() - 2*60*60
	}
	return tracksFirstPlayed, nil
}

func UpdateXML(pathin string, pathout string, first2last map[int64]int64, first2track map[int64]string) error {
	re := regexp.MustCompile("<Entry (.*?)FirstPlayed=\"(.+?)\" LastPlayed=\"(.+?)\"")
	replace := "<Entry ${1}FirstPlayed=\"${2}\" LastPlayed=\""
	fileIn, err := os.Open(pathin)
	if err != nil {
		return err
	}
	defer fileIn.Close()
	fileOut, err := os.Create(pathout)
	if err != nil {
		return err
	}
	defer fileOut.Close()
	scanner := bufio.NewScanner(fileIn)
	for scanner.Scan() {
		text := scanner.Text()
		res := re.FindStringSubmatch(text)
		if len(res) != 4 {
			fileOut.WriteString(text + "\n")
			continue
		}
		xmlFirstPlayed, err := strconv.ParseInt(res[2], 10, 64)
		if err != nil {
			return err
		}
		xmlLastPlayed, err := strconv.ParseInt(res[3], 10, 64)
		if err != nil {
			return err
		}
		xmlFirstPlayed = Win2unix(xmlFirstPlayed)
		xmlLastPlayed = Win2unix(xmlLastPlayed)
		format := "02/01/2006 15:04:05"
		if lastPlayed, ok := first2last[xmlFirstPlayed]; ok {
			//first2last[firstPlayed] = int64(scrobble.Timestamp.UnixTimestamp)
			if lastPlayed > xmlLastPlayed {
				tBefore := time.Unix(xmlLastPlayed, 0).Format(format)
				tAfter := time.Unix(lastPlayed, 0).Format(format)
				fmt.Printf("[%s]\n%s > %s\n\n", first2track[xmlFirstPlayed], tBefore, tAfter)
				xmlNewLastPlayed := strconv.FormatInt(Unix2win(lastPlayed), 10)
				fileOut.WriteString(re.ReplaceAllString(text, replace+xmlNewLastPlayed+"\"") + "\n")
				continue
			}
		}
		fileOut.WriteString(text + "\n")
	}
	return nil
}

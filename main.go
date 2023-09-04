package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

var (
	err error
	// The flags that can be used by users.
	update bool
	search string
	// All of the application's waitgroups.
	scrapeWaitGroup  sync.WaitGroup
	cleanUpWaitGroup sync.WaitGroup
	// All of the slices that were utilized in the application.
	exclusionIPRange []string
	// The files current location.
	assetsLocation            = "assets/"
	exclusionFileLocation     = assetsLocation + "exclusion"
	abuseFileLocation         = assetsLocation + "abuse"
	anonymizersFileLocation   = assetsLocation + "anonymizers"
	attacksFileLocation       = assetsLocation + "attacks"
	malwareFileLocation       = assetsLocation + "malware"
	organizationsFileLocation = assetsLocation + "organizations"
	reputationFileLocation    = assetsLocation + "reputation"
	spamFileLocation          = assetsLocation + "spam"
	unroutableFileLocation    = assetsLocation + "unroutable"
)

func init() {
	// Make sure you have provided the correct flags, or else the program will exit.
	if len(os.Args) > 1 {
		tempUpdate := flag.Bool("update", false, "Make any necessary changes to the listings.")
		tempSearch := flag.String("search", "example.example", "Check to see if a specific IP is on a list.")
		flag.Parse()
		update = *tempUpdate
		search = *tempSearch
	} else {
		log.Fatal("Error: No flags provided. Please use -help for more information.")
	}
}

func main() {
	// Lists should be updated.
	if update {
		updateTheLists()
	}
	// Look through all of the listings to check if the IP address appears on any of them.
	if len(search) > 1 && search != "example.example" {
		findAllMatchingDomains(search)
	}
}

func updateTheLists() {
	// Scrape the data and save it in the system.
	getAllURlsAndStartScraping()
	// We may begin cleaning up the files when the scraping is completed.
	if fileExists(exclusionFileLocation) {
		exclusionIPRange = readAndAppend(exclusionFileLocation, exclusionIPRange)
	}
	cleanUPLocation(abuseFileLocation)
	cleanUPLocation(anonymizersFileLocation)
	cleanUPLocation(attacksFileLocation)
	cleanUPLocation(malwareFileLocation)
	cleanUPLocation(organizationsFileLocation)
	cleanUPLocation(reputationFileLocation)
	cleanUPLocation(spamFileLocation)
	cleanUPLocation(unroutableFileLocation)
	// Wait for the cleaning to be completed.
	cleanUpWaitGroup.Wait()
}

func getAllURlsAndStartScraping() {
	// Abuse
	abuseIPList := []string{
		"https://myip.ms/files/blacklist/csf/latest_blacklist.txt",
		"https://raw.githubusercontent.com/complexorganizations/proxy-registry/main/assets/hosts",
	}
	for _, content := range abuseIPList {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, abuseFileLocation)
		}
	}
	// Anonymizers
	anonymizersIPLists := []string{
		"https://check.torproject.org/torbulkexitlist",
	}
	for _, content := range anonymizersIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, anonymizersFileLocation)
		}
	}
	// Attacks
	attacksIPLists := []string{
		"https://lists.blocklist.de/lists/all.txt",
	}
	for _, content := range attacksIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, attacksFileLocation)
		}
	}
	// Malware
	malwareIPLists := []string{
		"https://dronebl.org/activity_log",
		"https://www.cloudflare.com/ips-v4",
	}
	for _, content := range malwareIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, malwareFileLocation)
		}
	}
	// Organizations
	organizationsIPLists := []string{
		"https://docs.oracle.com/en-us/iaas/tools/public_ip_ranges.json",
		"https://www.gstatic.com/ipranges/goog.json",
		"https://ip-ranges.amazonaws.com/ip-ranges.json",
		"https://www.cloudflare.com/ips-v4",
		"https://www.cloudflare.com/ips-v6",
		"https://api.fastly.com/public-ip-list",
		"https://digitalocean.com/geo/google.csv",
	}
	for _, content := range organizationsIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, organizationsFileLocation)
		}
	}
	// Reputation
	reputationIPLists := []string{
		"https://www.spamhaus.org/drop/drop.txt",
	}
	for _, content := range reputationIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, reputationFileLocation)
		}
	}
	// Spam
	spamIPLists := []string{
		"https://www.stopforumspam.com/downloads/toxic_ip_cidr.txt",
	}
	for _, content := range spamIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, spamFileLocation)
		}
	}
	// Unroutable
	unroutableIPLists := []string{
		"https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/unroutable",
	}
	for _, content := range unroutableIPLists {
		if validURL(content) {
			scrapeWaitGroup.Add(1)
			go scrapeAllIP(content, unroutableFileLocation)
		}
	}
	// Wait until the content has been scraped before proceeding.
	scrapeWaitGroup.Wait()
}

func scrapeAllIP(url string, saveLocation string) {
	// If you find any old files, delete them from your system.
	removeThisFile(saveLocation)
	// Send a request for all of the information you want.
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	// Read the whole body of the document.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	err = response.Body.Close()
	if err != nil {
		log.Println(err)
	}
	// If the url's response code is invalid, inform the users.
	if response.StatusCode != 200 {
		log.Println("Sorry, but we were unable to scrape the page you requested due to a error.", url)
	}
	// Scraped data is read and appended to an array.
	scanner := bufio.NewScanner(bytes.NewReader(body))
	scanner.Split(bufio.ScanLines)
	var returnContent []string
	for scanner.Scan() {
		returnContent = append(returnContent, scanner.Text())
	}
	for _, content := range returnContent {
		content = strings.TrimSpace(content)
		// This is a list of all the content discovered using the regex.
		foundIpv4 := regexp.MustCompile(`\b([0-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?\b`).Find([]byte(content))
		foundIPv6 := regexp.MustCompile(`\bs*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:)))(%.+)?s*(\/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8]))?\b`).Find([]byte(content))
		// Check if all the content is valid.
		if len(foundIpv4) >= 7 {
			_, _, err = net.ParseCIDR(string(foundIpv4))
			if net.ParseIP(string(foundIpv4)) != nil || err == nil {
				writeToFile(saveLocation, string(foundIpv4))
			}
		}
		if len(foundIPv6) >= 7 {
			_, _, err = net.ParseCIDR(string(foundIPv6))
			if net.ParseIP(string(foundIPv6)) != nil || err == nil {
				writeToFile(saveLocation, string(foundIPv6))
			}
		}
	}
	// Once everything is done, close the wait group.
	scrapeWaitGroup.Done()
}

// Look through all of the files to see if any include an IP address.
func findAllMatchingDomains(search string) {
	//
}

// Clean up the exclusions because users may have altered them.
func finalCleanup(filePath string) {
	var tempCleanupContent []string
	var finalCleanupContent []string
	tempCleanupContent = readAndAppend(filePath, tempCleanupContent)
	sort.Strings(tempCleanupContent)
	// Make each domain one-of-a-kind.
	uniqueExclusionContent := makeUnique(tempCleanupContent)
	// Remove and Add the contents.
	for _, content := range uniqueExclusionContent {
		if checkIfIPInRange(content, exclusionIPRange) {
			content = ""
		}
		if arrayContains(content, exclusionIPRange) {
			content = ""
		}
		content = strings.TrimSpace(content)
		finalCleanupContent = append(finalCleanupContent, addCidr(content))
	}
	// Remove the original file before rewriting it.
	err = os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}
	for _, content := range finalCleanupContent {
		if len(content) > 2 && validateIPWhileParsingCidr(content) && validateCIDR(content) {
			writeToFile(filePath, content)
		}
	}
	// Close the waitgroup.
	cleanUpWaitGroup.Done()
}

// Line by line, read and add the file to an array.
func readAndAppend(fileLocation string, arrayName []string) []string {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		arrayName = append(arrayName, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
	}
	return arrayName
}

// Create a specfic file with certain content.
func writeToFile(pathInSystem string, content string) {
	filePath, err := os.OpenFile(pathInSystem, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	_, err = filePath.WriteString(content + "\n")
	if err != nil {
		log.Println(err)
	}
	err = filePath.Close()
	if err != nil {
		log.Println(err)
	}
}

// Create a one-of-a-kind string array.
func makeUnique(randomStrings []string) []string {
	var uniqueString []string
	for _, value := range randomStrings {
		if !arrayContains(value, uniqueString) {
			uniqueString = append(uniqueString, value)
		}
	}
	return uniqueString
}

// Verify that the value is in the array.
func arrayContains(cointains string, originalArray []string) bool {
	for _, value := range originalArray {
		if value == cointains {
			return true
		}
	}
	return false
}

// Make sure the url is legitimate by verifying it.
func validURL(uri string) bool {
	_, err = url.ParseRequestURI(uri)
	return err == nil
}

// Check to determine whether the file you're looking for already exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Determine whether a particular cdir range contains a specific IP address.
func checkIfIPInRange(ip string, blacklist []string) bool {
	for _, cidr := range blacklist {
		if strings.Contains(cidr, "/") {
			_, ipnet, _ := net.ParseCIDR(cidr)
			if ipnet.Contains(net.ParseIP(ip)) {
				return true
			}
		}
	}
	return false
}

// If no cidr is discovered, add one.
func addCidr(ipToAddCidr string) string {
	if strings.Contains(ipToAddCidr, "/") {
		return ipToAddCidr
	} else if strings.Contains(ipToAddCidr, ".") {
		return ipToAddCidr + "/32"
	} else if strings.Contains(ipToAddCidr, ":") {
		return ipToAddCidr + "/128"
	}
	return ipToAddCidr
}

// Check if a file exists and if it does remove the file
func removeThisFile(filePath string) {
	if fileExists(filePath) {
		err = os.Remove(filePath)
		if err != nil {
			log.Println(err)
		}
	}
}

// Cleanup waitgroup/goroutines.
func cleanUPLocation(fileLocation string) {
	if fileExists(fileLocation) {
		cleanUpWaitGroup.Add(1)
		go finalCleanup(fileLocation)
	}
}

// Ensure that the CIDR block is valid.
func validateCIDR(ipWithCidr string) bool {
	_, _, err := net.ParseCIDR(ipWithCidr)
	return err == nil
}

// Check if the given IP address is valid with a given cidr
func validateIPWhileParsingCidr(ipWithCidr string) bool {
	ipAddress, _, _ := net.ParseCIDR(ipWithCidr)
	return net.ParseIP(ipAddress.String()) != nil
}

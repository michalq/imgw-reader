package synop

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func SynopDataDownloader(debug bool, baseUrl, outFolder string) error {
	errs := make([]error, 0)
	c := colly.NewCollector()
	c.OnHTML(`td>a`, func(e *colly.HTMLElement) {
		fileUrl := e.Request.AbsoluteURL(e.Attr("href"))
		if IsYearFolder(fileUrl) {
			data := colly.NewCollector()
			data.OnHTML("a", func(e *colly.HTMLElement) {
				fileUrl := e.Request.AbsoluteURL(e.Attr("href"))
				if isZip(fileUrl) {
					if debug {
						fmt.Println("Downloading", fileUrl)
					}
					if err := downloadFile(fileUrl, outFolder); err != nil {
						errs = append(errs, fmt.Errorf("download file %s error: %v", fileUrl, err))
					}
				}
			})
			if err := data.Visit(fileUrl); err != nil {
				errs = append(errs, fmt.Errorf("visit %s error: %v", fileUrl, err))
			}
		}
	})
	if err := c.Visit(baseUrl); err != nil {
		errs = append(errs, fmt.Errorf("visit %s error: %v", baseUrl, err))
	}

	return errors.Join(errs...)
}

func isZip(fileUrl string) bool {
	splitted := strings.Split(fileUrl, ".")
	return splitted[len(splitted)-1] == "zip"
}

func IsYearFolder(fileUrl string) bool {
	if fileUrl[len(fileUrl)-1] == '/' {
		splitted := strings.Split(fileUrl, "/")
		rg := regexp.MustCompile(`^\d{4}(_\d{4})?$`)
		if rg.MatchString(splitted[len(splitted)-2]) {
			return true
		}
	}
	return false
}

func downloadFile(fileUrl, outFolder string) error {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("getting file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	filename := filepath.Base(fileUrl)
	f, err := os.Create(outFolder + filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("copying file: %w", err)
	}
	return nil
}

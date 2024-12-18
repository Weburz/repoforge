package template

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

// The base directory where the templates will be stored locally
var templateDir = filepath.Join(xdg.DataHome, "terox")

/**
 * Scaffold - Scaffold the project by download the template (if necessary)
 *
 * Parameters:
 * None
 *
 * Returns:
 * A wrapped error if any is raised by the underlying wrapping function.
 */
func Scaffold(repo string) error {
	// Check if the template already exists locally
	parts := strings.Split(repo, "/")
	owner := parts[0]
	repository := parts[1]
	dir := filepath.Join(templateDir, owner, repository)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Template not found locally...downloading\n")

		// Download the template from GitHub
		f, err := downloadTemplate(repo)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		defer os.Remove(f)

		// Extract the downloaded zipball to the expected destination
		if err := extractTemplate(f, templateDir); err != nil {
			return fmt.Errorf("%w", err)
		}

		return nil
	} else if err != nil {
		return fmt.Errorf("Error checking the template path: %w", err)
	}

	fmt.Printf("Template found locally at: %s\n", dir)

	return nil
}

/**
 * downloadTemplate - Download the template contents from the GitHub repository.
 *
 * Parameters:
 * repo (string): The GitHub repository's URL to fetch the template from.
 *
 * Returns:
 * A wrapped error if any were raised during the downloading process.
 */
func downloadTemplate(repo string) (string, error) {
	// Create the URL to download the zipball from
	url := fmt.Sprintf("https://api.github.com/repos/%s/zipball", repo)

	// TODO: Configure the client to make authenticated requests too to fetch
	// templates from private repositories
	// Create the client to make a GET request with
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to download the template to: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad server response: %d", resp.StatusCode)
	}

	tempFile, err := os.CreateTemp("", "terox-*.zip")
	if err != nil {
		return "", fmt.Errorf("Failed to create the template zipball: %w", err)
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, resp.Body); err != nil {
		return "", fmt.Errorf("Failed to write the zipball: %w", err)
	}

	return tempFile.Name(), nil
}

/**
 * Extract the downloaded zipfile.
 *
 * Parameters:
 * zipfile (string): The path to the (downloaded) zipfile to download it from.
 * dest (string): The destination path to extract the zipped contents to.
 *
 * Returns:
 * A wrapped error (if any was raised) during the extraction procedure.
 */
func extractTemplate(zipfile, dest string) error {
	// Read the zipfile and close it when the function completes execution
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return fmt.Errorf("Failed to open zip file: %w", err)
	}
	defer r.Close()

	// Store the name of the top-level folder for further string processing
	var topLevelFolder string
	for _, f := range r.File {
		parts := strings.Split(f.Name, "/")
		if len(parts) > 1 && topLevelFolder == "" {
			topLevelFolder = parts[0]
		}
	}

	if topLevelFolder == "" {
		return fmt.Errorf(
			"Failed to detect the top-level directory in the archive.",
		)
	}

	// Split the top-level folder by the "-" character and store them in
	// variables for further processing
	parts := strings.Split(topLevelFolder, "-")
	if len(parts) < 2 {
		return fmt.Errorf("Unexpected folder structure: %s", topLevelFolder)
	}
	owner := parts[0]
	repo := parts[1]

	// Ensure the destination directory before the zipfile can be extracted
	finalDest := filepath.Join(dest, owner, repo)
	if err := os.MkdirAll(finalDest, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create destination directory: %w", err)
	}

	// Extract each file
	for _, f := range r.File {
		// Create the correct file path
		relativePath := strings.TrimPrefix(f.Name, topLevelFolder+"/")
		filePath := filepath.Join(finalDest, relativePath)

		// If the file is a directory, create it
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, f.Mode()); err != nil {
				return fmt.Errorf("Failed to create directory: %w", err)
			}
			continue
		}

		// Ensure parent directories exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("Failed to create parent directories: %w", err)
		}

		// Extract the file
		srcFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("Failed to open file in archive: %w", err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Failed to create file: %w", err)
		}

		if _, err := io.Copy(destFile, srcFile); err != nil {
			destFile.Close()
			return fmt.Errorf("Failed to copy file contents: %w", err)
		}

		destFile.Close()
	}

	return nil
}

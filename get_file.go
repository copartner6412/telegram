package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

// GetFileRequest represents parameters for getFile method.
//
// See https://core.telegram.org/bots/api#getfile
type GetFileRequest struct {
	// (Required) File identifier to get information about
	FileID string `json:"file_id"`
}

// GetFile retrieves basic information about a file and prepares it for downloading.
//
// Bots can download files of up to 20MB in size.
//
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// See https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileID string) (*File, error) {
	request := GetFileRequest{
		FileID: fileID,
	}
	result := new(File)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodGetFile, request); err != nil {
		return nil, err
	}

	return result, nil
}

// DownloadFile downloads a file by its file ID on Telegram servers.
//
// Parameters:
//   - fileID: File identifier to get information about
//
// See "getFile" https://core.telegram.org/bots/api#getfile
func (b *Bot) DownloadFile(fileID string) (*os.File, error) {
	telegramFile, err := b.GetFile(fileID)
	if err != nil {
		return nil, fmt.Errorf("error getting file ID from Telegram: %w", err)
	}

	if telegramFile.FilePath == "" {
		return nil, fmt.Errorf("no filepath received")
	}

	fileURL, err := url.JoinPath("https://api.telegram.org/file/", "bot"+b.Token, telegramFile.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error joining elements of file URL: %w", err)
	}

	response, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", response.Status)
	}

	out, err := os.Create(path.Join("/tmp"))
	if err != nil {
		return nil, fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return nil, fmt.Errorf("error saving file: %w", err)
	}

	return out, nil
}

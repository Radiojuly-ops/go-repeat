package file

import "os"

func WriterFile(path string, content string) error {
	data := []byte(content)

	err := os.WriteFile(path, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}

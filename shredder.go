package shredder

import( 	
	"crypto/rand"
	"os"
)

type Config struct {
	Iterations int
	Remove bool
}

func (config Config) File(Path string) error{
	for i := 0 ; i < config.Iterations ; i++{
		if err := Shred(Path); err != nil {
			return err
		}
	}

	if config.Remove {
		if err := os.Remove(Path); err != nil {
			return err
		}
	}
	return nil
}

func Shred(Path string) error{
	file, err := os.OpenFile(Path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	fileSize, err := file.Stat()
	if err != nil {
		return err
	}
	buff := make([]byte, fileSize.Size())
	rand.Read(buff)
	file.Write(buff)

	_, err = file.Seek(0, 0)
    if err != nil {
        return err
    }
	file.Close()
	
	return nil
}
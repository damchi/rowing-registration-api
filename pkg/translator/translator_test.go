package translator

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestInitTranslator(t *testing.T) {
	translationFolder := "translations"

	createTranslationFile(translationFolder, "translations/fr.toml")

	cfg := Config{
		TranslationFolder: translationFolder,
	}

	InitTranslator(cfg)

	msg := Trans("hello", "fr", nil)

	if msg != "salut" {
		t.Errorf("Fail to initialize translation")
	}

	deleteTranslationFiles(translationFolder)
}

func createTranslationFile(translationFolder string, translationFile string) {

	if _, err := os.Stat(translationFolder); os.IsNotExist(err) {
		err = os.Mkdir(translationFolder, os.ModePerm)

		if err != nil {
			log.Println(err)
		}
	}

	f, err := os.OpenFile(translationFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	_, err = f.WriteString("hello=\"salut\" \n")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = f.Sync()

	if err != nil {
		fmt.Println(err.Error())
	}

}

func deleteTranslationFiles(translationFolder string) {
	err := os.RemoveAll(translationFolder)

	if err != nil {
		log.Println(err)
	}
}

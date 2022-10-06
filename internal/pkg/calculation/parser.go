package calculation

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/LiliyaD/Test_task_GasPrice/internal/journal"
	"github.com/LiliyaD/Test_task_GasPrice/internal/pkg/models"
)

var source models.SourceJSON

func Parse(body io.ReadCloser) {
	err := json.NewDecoder(body).Decode(&source)
	if err != nil {
		journal.LogFatal(err)
	}
}

func SaveJSON(responce *models.ResponceJSON) {
	data, err := json.MarshalIndent(responce, "", "    ")
	if err != nil {
		journal.LogFatal(err)
	}

	err = os.MkdirAll("responce_jsons", os.ModePerm)
	if err != nil {
		journal.LogError(err)
	}

	fileName := "gas_price_" + time.Now().Format("2006-01-02 15.04.05") + ".json"
	filePath := filepath.Join("responce_jsons", fileName)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		journal.LogError(err)
		journal.LogWarn("Unsaved json data: ", responce)
	} else {
		journal.LogInfo("File ", fileName, " is created")
	}
}

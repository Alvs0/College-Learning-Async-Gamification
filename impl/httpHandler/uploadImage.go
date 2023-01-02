package httpHandler

import (
	"bytes"
	"college-learning-asynchronous-gamification/gcs"
	"college-learning-asynchronous-gamification/http"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	stdHttp "net/http"
)

func UploadImagePostHandler(generalData GeneralData, fileUploader gcs.UploaderClient) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		if err := req.ParseMultipartForm(0); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		file, _, err := req.FormFile("file")
		defer file.Close()
		if err != nil {
			fmt.Println("[UploadImage] File not found")
			return nil, nil
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		bytesWriter := Writer{Bytes: buf.Bytes()}

		imageId := uuid.New()
		fileUrl, err := fileUploader.Upload(imageId.String(), &bytesWriter)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		imageMap := map[string]string{
			"imageUrl": fileUrl,
		}

		outputByte, err := json.Marshal(imageMap)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(stdHttp.StatusOK)
		res.Write(outputByte)

		return nil, nil
	}
}

type Writer struct {
	Bytes []byte
}

func (ths *Writer) Write(writer io.Writer) error {
	_, err := writer.Write(ths.Bytes)
	if err != nil {
		return err
	}

	return nil
}

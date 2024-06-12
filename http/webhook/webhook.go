package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

// admitFunc 定义了处理Webhook请求的函数类型
type admitFunc func(admissionReview *v1.AdmissionReview) *v1.AdmissionResponse

// admitForDebug 是一个示例admitFunc，仅打印接收到的请求并允许所有操作
func admitForDebug(ar *v1.AdmissionReview) *v1.AdmissionResponse {
	fmt.Printf("Received AdmissionReview: %+v\n", ar.Request)
	return &v1.AdmissionResponse{
		Allowed: true,
	}
}

// toAdmissionReview 解析HTTP请求体为AdmissionReview对象
func toAdmissionReview(r io.ReadCloser) (*v1.AdmissionReview, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("can't read body: %v", err)
	}
	ar := v1.AdmissionReview{}
	if err := json.Unmarshal(body, &ar); err != nil {
		return nil, fmt.Errorf("can't unmarshal body: %v", err)
	}
	return &ar, nil
}

// serveHTTP 处理HTTP请求，调用admitFunc处理Webhook
func serveHTTP(w http.ResponseWriter, r *http.Request, admit admitFunc) {
	var reviewResponse *v1.AdmissionReview
	if r.URL.Path != "/validate" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method type.", http.StatusMethodNotAllowed)
		return
	}

	ar, err := toAdmissionReview(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing AdmissionReview: %v", err), http.StatusBadRequest)
		return
	}

	reviewResponse = &v1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
		Response: &v1.AdmissionResponse{
			UID:     ar.Request.UID,
			Allowed: false, // 默认拒绝，admitFunc将根据实际情况修改
		},
	}

	if ar.Request != nil {
		reviewResponse.Response = admit(ar)
	}

	resp, err := json.Marshal(reviewResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't encode response: %v", err), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(resp); err != nil {
		http.Error(w, fmt.Sprintf("Can't write response: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		serveHTTP(w, r, admitForDebug)
	})
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

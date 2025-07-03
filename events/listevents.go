package events

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func eventInterface(con *kubernetes.Clientset, namespace string) typev1.EventInterface {
	return con.CoreV1().Events(namespace)
}

func ListEvent(con *kubernetes.Clientset, ctx context.Context, namespace string) (*corev1.EventList, error) {
	eventinterface := eventInterface(con, namespace)
	getEventList, err := eventinterface.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return getEventList, nil
}

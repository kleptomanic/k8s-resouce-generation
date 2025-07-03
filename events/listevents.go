package events

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Create Event Connection Interface for the namespace as provided.
func eventInterface(con *kubernetes.Clientset, namespace string) typev1.EventInterface {
	return con.CoreV1().Events(namespace)
}

// Returns the slices of Events
func ListEvent(con *kubernetes.Clientset, ctx context.Context, namespace string) ([]corev1.Event, error) {
	eventinterface := eventInterface(con, namespace)
	getEventList, err := eventinterface.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return getEventList.Items, nil
}

package jobs

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typev1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func JobsClient(connection *kubernetes.Clientset, namespace string) typev1.JobInterface {
	return connection.BatchV1().Jobs(namespace)
}

func ListJobs(connection *kubernetes.Clientset, ctx context.Context, namespace string) ([]corev1.Job, error) {
	listinterface := JobsClient(connection, namespace)
	jobs, err := listinterface.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return jobs.Items, nil
}

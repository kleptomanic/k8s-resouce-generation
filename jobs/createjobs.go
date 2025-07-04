package jobs

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func NewJobObject(name string, namespace string) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "TestinitContainer",
							Image:   "busybox",
							Command: []string{"ping", "-c 4", "google.com"},
						},
					},
				},
			},
		},
	}
}

func CreateJobs(jobClient typev1.JobInterface, ctx context.Context) (*batchv1.Job, error) {
	job, err := jobClient.Create(ctx, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return job, nil
}

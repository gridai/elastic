package controllers

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

func TestInsertTorchArgs(t *testing.T) {
	RegisterFailHandler(Fail)

	container := corev1.Container{
		Args: []string{"run.py", "--arg1", "val1"},
	}
	torchArgs := []string{"--rdvz", "etcd"}
	InsertTorchArgs(&container, torchArgs)
	Expect(container.Args).To(Equal([]string{"--rdvz", "etcd", "run.py", "--arg1", "val1"}))

	container = corev1.Container{
		Args: []string{"python", "run.py", "python", "-m", "torchelastic.distributed.launch", "script.py", "--arg1", "val1"},
	}
	InsertTorchArgs(&container, torchArgs)
	Expect(container.Args).To(Equal(
		[]string{"python", "run.py", "python", "-m", "torchelastic.distributed.launch", "--rdvz", "etcd", "script.py", "--arg1", "val1"}))
}

package mocks

import (
	"fmt"

	v1beta1extensions "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	api "k8s.io/client-go/pkg/api"
	unversioned "k8s.io/client-go/pkg/api/unversioned"
	v1 "k8s.io/client-go/pkg/api/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/pkg/watch"
)

type dClient struct {
}

func (d dClient) Get(name string) (*extensions.DaemonSet, error) {
	if name != "lgtm" {
		return nil, fmt.Errorf("Mock daemonset didnt work")
	}
	ds := &extensions.DaemonSet{
		ObjectMeta: v1.ObjectMeta{Name: name},
		Spec: extensions.DaemonSetSpec{
			Selector: &unversioned.LabelSelector{
				MatchLabels: map[string]string{"name": "test"},
			},
		},
	}
	return ds, nil
}
func (d dClient) Create(ds *extensions.DaemonSet) (*extensions.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Delete(name string, options *v1.DeleteOptions) error {
	return fmt.Errorf("Not implemented")
}
func (d dClient) List(options v1.ListOptions) (*extensions.DaemonSetList, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Update(ds *extensions.DaemonSet) (*extensions.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) UpdateStatus(ds *extensions.DaemonSet) (*extensions.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return fmt.Errorf("Not implemented")
}

func (d dClient) Watch(options v1.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *extensions.DaemonSet, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func NewDSClient() v1beta1extensions.DaemonSetInterface {
	return dClient{}
}

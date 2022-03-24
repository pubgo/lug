package k8s

import (
	"context"
	"fmt"
	"github.com/pubgo/lava/event"
	"github.com/pubgo/lava/plugins/registry/registry_type"
	"github.com/pubgo/xerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"

	"github.com/pubgo/lava/pkg/k8s"
	"github.com/pubgo/lava/plugins/registry"
)

var _ registry_type.Watcher = (*Watcher)(nil)

// Watcher performs the conversion from channel to iterator
// It reads the latest changes from the `chan []*registry.ServiceInstance`
// And the outside can sense the closure of Watcher through stopCh
type Watcher struct {
	service string
	watcher watch.Interface
	client  *kubernetes.Clientset
}

// newWatcher is used to initialize Watcher
func newWatcher(s *Registry, service string) *Watcher {
	watcher, err := s.client.
		CoreV1().
		Endpoints(k8s.Namespace()).
		Watch(context.Background(),
			metav1.ListOptions{FieldSelector: fmt.Sprintf("%s=%s", "metadata.name", service)})
	xerror.Panic(err)
	return &Watcher{watcher: watcher, client: s.client, service: service}
}

// Next will block until ServiceInstance changes
func (t *Watcher) Next() (*registry_type.Result, error) {
	select {
	case _, ok := <-t.watcher.ResultChan():
		if ok {
			endpoints, err := t.client.
				CoreV1().
				Endpoints(k8s.Namespace()).
				List(context.Background(),
					metav1.ListOptions{FieldSelector: fmt.Sprintf("%s=%s", "metadata.name", t.service)})
			xerror.Panic(err)

			var resp = &registry_type.Result{
				Action: event.EventType_UPDATE,
				Service: &registry_type.Service{
					Name: t.service,
				},
			}

			for _, endpoint := range endpoints.Items {
				for _, subset := range endpoint.Subsets {
					realPort := ""
					for _, p := range subset.Ports {
						realPort = fmt.Sprint(p.Port)
						break
					}

					for _, addr := range subset.Addresses {
						resp.Service.Nodes = append(resp.Service.Nodes, &registry_type.Node{
							Id:      string(addr.TargetRef.UID),
							Address: fmt.Sprintf("%s:%s", addr.IP, realPort),
						})
					}
				}
			}

			return resp, err
		}

		return nil, registry.ErrWatcherStopped
	}
}

// Stop is used to close the iterator
func (t *Watcher) Stop() error {
	t.watcher.Stop()
	return nil
}

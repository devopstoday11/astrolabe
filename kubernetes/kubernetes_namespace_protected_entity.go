package kubernetes

import (
	"context"
	"github.com/vmware/arachne"
	v1 "k8s.io/api/core/v1"
)

type KubernetesNamespaceProtectedEntity struct {
	knpetm    *KubernetesNamespaceProtectedEntityTypeManager
	id        arachne.ProtectedEntityID
	namespace *v1.Namespace
}

func NewKubernetesNamespaceProtectedEntity(knpetm *KubernetesNamespaceProtectedEntityTypeManager,
	namespace *v1.Namespace) (*KubernetesNamespaceProtectedEntity, error) {
	nsPEID := arachne.NewProtectedEntityID("k8sns", namespace.Name)
	returnPE := KubernetesNamespaceProtectedEntity{
		knpetm:    knpetm,
		id:        nsPEID,
		namespace: namespace,
	}
	return &returnPE, nil
}

func (this *KubernetesNamespaceProtectedEntity) GetInfo(ctx context.Context) (arachne.ProtectedEntityInfo, error) {
	return nil, nil
}
func (this *KubernetesNamespaceProtectedEntity) GetCombinedInfo(ctx context.Context) ([]arachne.ProtectedEntityInfo, error) {
	return nil, nil

}

func (this *KubernetesNamespaceProtectedEntity) Snapshot(ctx context.Context) (*arachne.ProtectedEntitySnapshotID, error) {
	return nil, nil

}
func (this *KubernetesNamespaceProtectedEntity) ListSnapshots(ctx context.Context) ([]arachne.ProtectedEntitySnapshotID, error) {
	return nil, nil

}
func (this *KubernetesNamespaceProtectedEntity) DeleteSnapshot(ctx context.Context,
	snapshotToDelete arachne.ProtectedEntitySnapshotID) (bool, error) {
	return false, nil

}
func (this *KubernetesNamespaceProtectedEntity) GetInfoForSnapshot(ctx context.Context,
	snapshotID arachne.ProtectedEntitySnapshotID) (*arachne.ProtectedEntityInfo, error) {
	return nil, nil

}

func (this *KubernetesNamespaceProtectedEntity) GetComponents(ctx context.Context) ([]arachne.ProtectedEntity, error) {
	return nil, nil

}

func (this *KubernetesNamespaceProtectedEntity) GetID() arachne.ProtectedEntityID {
	return this.id
}
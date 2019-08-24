package ivd

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/arachne/pkg/arachne"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
	"github.com/vmware/gvddk/gDiskLib"
	"io"
	"io/ioutil"
	"strings"

	//	"github.com/vmware/govmomi/vslm"
	"context"
	"time"
)

type IVDProtectedEntity struct {
	ipetm    *IVDProtectedEntityTypeManager
	id       arachne.ProtectedEntityID
	data     []arachne.DataTransport
	metadata []arachne.DataTransport
	combined []arachne.DataTransport
}

type metadata struct {
	VirtualStorageObject vim.VStorageObject `xml:"virtualStorageObject"`
	Datastore vim.ManagedObjectReference `xml:"datastore""`
}

func (this IVDProtectedEntity) GetDataReader(ctx context.Context) (io.Reader, error) {

	diskHandle, err := this.getDiskHandle(ctx)
	if err != nil {
		return nil, err
	}
	return arachne.NewReaderAtReader(diskHandle), nil
}

func (this IVDProtectedEntity) copy(ctx context.Context, dataReader io.Reader,
	metadata metadata) error {
	// TODO - restore metadata
	dataWriter, err := this.getDataWriter(ctx)
	if err == nil {
		_, err = io.Copy(dataWriter, dataReader) // TODO - add a copy routine that we can interrupt via context
	}
	return err
}

func (this IVDProtectedEntity) getDataWriter(ctx context.Context) (io.Writer, error) {
	diskHandle, err := this.getDiskHandle(ctx)
	if err != nil {
		return nil, err
	}
	return arachne.NewWriterAtWriter(diskHandle), nil
}

func (this IVDProtectedEntity) getDiskHandle(ctx context.Context) (gDiskLib.DiskHandle, error) {
	url := this.ipetm.client.URL()
	serverName := url.Hostname()
	userName := this.ipetm.user
	password := this.ipetm.password
	/*
		thumbprint := this.ipetm.client.Thumbprint(serverName)
		thumbprint = "3D:62:45:37:88:36:3E:03:7A:6C:5A:63:D6:D6:AB:85:F7:DE:A3:AB"
		if thumbprint == "" {
			return nil, errors.New("Thumbprint was not set in client")
		}*/
	fcdid := this.id.GetID()
	vso, err := this.ipetm.vsom.Retrieve(context.Background(), NewVimIDFromPEID(this.id))
	if err != nil {
		return gDiskLib.DiskHandle{}, err
	}
	datastore := vso.Config.Backing.GetBaseConfigInfoBackingInfo().Datastore.String()
	datastore = strings.TrimPrefix(datastore, "Datastore:")

	fcdssid := ""
	if this.id.HasSnapshot() {
		fcdssid = this.id.GetSnapshotID().String()
	}
	params := gDiskLib.NewConnectParams("",
		serverName,
		"3D:62:45:37:88:36:3E:03:7A:6C:5A:63:D6:D6:AB:85:F7:DE:A3:AB",
		userName,
		password,
		fcdid,
		datastore,
		fcdssid,
		"",
		"vm-example")
	conn, err := gDiskLib.Connect(params)
	if err != nil {
		return gDiskLib.DiskHandle{}, errors.Wrap(err, "Connect failed")
	}
	err = gDiskLib.PrepareForAccess(params)
	if err != nil {
		return gDiskLib.DiskHandle{}, errors.Wrap(err, "PrepareForAccess failed")
	}
	diskHandle, err := gDiskLib.Open(conn, "", 1 /*C.VIXDISKLIB_FLAG_OPEN_UNBUFFERED*/)
	if err != nil {
		return gDiskLib.DiskHandle{}, err
	}
	return diskHandle, nil
}

func (this IVDProtectedEntity) GetMetadataReader(ctx context.Context) (io.Reader, error) {
	infoBuf, err := this.getMetadataBuf(ctx)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(infoBuf), nil
}

func (this IVDProtectedEntity) getMetadataBuf(ctx context.Context) ([]byte, error) {
	md, err := this.getMetadata(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Retrieve failed")
	}
	retBuf, err := xml.MarshalIndent(md, "  ", "    ")
	if err == nil {
		fmt.Println(retBuf)
	}
	return retBuf, err
}

func (this IVDProtectedEntity) getMetadata(ctx context.Context) (metadata, error) {
	vsoID := vim.ID{
		Id: this.id.GetID(),
	}
	vso, err := this.ipetm.vsom.Retrieve(ctx, vsoID)
	if err != nil {
		return metadata{}, err
	}
	datastore := vso.Config.BaseConfigInfo.GetBaseConfigInfo().Backing.GetBaseConfigInfoBackingInfo().Datastore
	retVal := metadata{
		VirtualStorageObject: *vso,
		Datastore: datastore,
	}
	return retVal, nil
}

func readMetadataFromReader(ctx context.Context, metadataReader io.Reader) (metadata, error) {
	mdBuf, err := ioutil.ReadAll(metadataReader) // TODO - limit this so it can't run us out of memory here
	if err != nil {
		return metadata{}, err
	}
	return readMetadataFromBuf(ctx, mdBuf)
}

func readMetadataFromBuf(ctx context.Context, buf []byte) (metadata, error) {
	var retVal = metadata{}
	err := xml.Unmarshal(buf, &retVal)
	return retVal, err
}

func newProtectedEntityID(id vim.ID) arachne.ProtectedEntityID {
	return arachne.NewProtectedEntityID("ivd", id.Id)
}

func newProtectedEntityIDWithSnapshotID(id vim.ID, snapshotID arachne.ProtectedEntitySnapshotID) arachne.ProtectedEntityID {
	return arachne.NewProtectedEntityIDWithSnapshotID("ivd", id.Id, snapshotID)
}

func newIVDProtectedEntity(ipetm *IVDProtectedEntityTypeManager, id arachne.ProtectedEntityID) (IVDProtectedEntity, error) {
	data, metadata, combined, err := ipetm.getDataTransports(id)
	if err != nil {
		return IVDProtectedEntity{}, err
	}
	newIPE := IVDProtectedEntity{
		ipetm:    ipetm,
		id:       id,
		data:     data,
		metadata: metadata,
		combined: combined,
	}
	return newIPE, nil
}
func (this IVDProtectedEntity) GetInfo(ctx context.Context) (arachne.ProtectedEntityInfo, error) {
	vsoID := vim.ID{
		Id: this.id.GetID(),
	}
	vso, err := this.ipetm.vsom.Retrieve(ctx, vsoID)
	if err != nil {
		return nil, errors.Wrap(err, "Retrieve failed")
	}

	retVal := arachne.NewProtectedEntityInfo(
		this.id,
		vso.Config.Name,
		this.data,
		this.metadata,
		this.combined,
		[]arachne.ProtectedEntityID{})
	return retVal, nil
}

func (this IVDProtectedEntity) GetCombinedInfo(ctx context.Context) ([]arachne.ProtectedEntityInfo, error) {
	ivdIPE, err := this.GetInfo(ctx)
	if err != nil {
		return nil, err
	}
	return []arachne.ProtectedEntityInfo{ivdIPE}, nil
}

const waitTime = 3600 * time.Second

/*
 * Snapshot APIs
 */
func (this IVDProtectedEntity) Snapshot(ctx context.Context) (*arachne.ProtectedEntitySnapshotID, error) {
	vslmTask, err := this.ipetm.vsom.CreateSnapshot(ctx, NewVimIDFromPEID(this.GetID()), "ArachneSnapshot")
	if err != nil {
		return nil, errors.Wrap(err, "Snapshot failed")
	}
	ivdSnapshotIDAny, err := vslmTask.Wait(ctx, waitTime)
	if err != nil {
		return nil, errors.Wrap(err, "Wait failed")
	}
	ivdSnapshotID := ivdSnapshotIDAny.(vim.ID)
	/*
		ivdSnapshotStr := ivdSnapshotIDAny.(string)
		ivdSnapshotID := vim.ID{
			id: ivdSnapshotStr,
		}
	*/
	retVal := arachne.NewProtectedEntitySnapshotID(ivdSnapshotID.Id)
	return &retVal, nil
}

func (this IVDProtectedEntity) ListSnapshots(ctx context.Context) ([]arachne.ProtectedEntitySnapshotID, error) {
	snapshotInfo, err := this.ipetm.vsom.RetrieveSnapshotInfo(ctx, NewVimIDFromPEID(this.GetID()))
	if err != nil {
		return nil, errors.Wrap(err, "RetrieveSnapshotInfo failed")
	}
	peSnapshotIDs := []arachne.ProtectedEntitySnapshotID{}
	for _, curSnapshotInfo := range snapshotInfo {
		peSnapshotIDs = append(peSnapshotIDs, arachne.NewProtectedEntitySnapshotID(curSnapshotInfo.Id.Id))
	}
	return peSnapshotIDs, nil
}
func (this IVDProtectedEntity) DeleteSnapshot(ctx context.Context, snapshotToDelete arachne.ProtectedEntitySnapshotID) (bool, error) {
	vslmTask, err := this.ipetm.vsom.DeleteSnapshot(ctx, NewVimIDFromPEID(this.id), NewVimSnapshotIDFromPESnapshotID(snapshotToDelete))
	if err != nil {
		return false, errors.Wrap(err, "DeleteSnapshot failed")
	}
	_, err = vslmTask.Wait(ctx, waitTime)
	if err != nil {
		return false, errors.Wrap(err, "Wait failed")
	}
	return true, nil
}

func (this IVDProtectedEntity) GetInfoForSnapshot(ctx context.Context, snapshotID arachne.ProtectedEntitySnapshotID) (*arachne.ProtectedEntityInfo, error) {
	return nil, nil
}

func (this IVDProtectedEntity) GetComponents(ctx context.Context) ([]arachne.ProtectedEntity, error) {
	return make([]arachne.ProtectedEntity, 0), nil
}

func (this IVDProtectedEntity) GetID() arachne.ProtectedEntityID {
	return this.id
}

func NewIDFromString(idStr string) vim.ID {
	return vim.ID{
		Id: idStr,
	}
}

func NewVimIDFromPEID(peid arachne.ProtectedEntityID) vim.ID {
	if peid.GetPeType() == "ivd" {
		return vim.ID{
			Id: peid.GetID(),
		}
	} else {
		return vim.ID{}
	}
}

func NewVimSnapshotIDFromPEID(peid arachne.ProtectedEntityID) vim.ID {
	if peid.HasSnapshot() {
		return vim.ID{
			Id: peid.GetSnapshotID().GetID(),
		}
	} else {
		return vim.ID{}
	}
}

func NewVimSnapshotIDFromPESnapshotID(peSnapshotID arachne.ProtectedEntitySnapshotID) vim.ID {
	return vim.ID {
		Id: peSnapshotID.GetID(),
	}
}
package testing

import (
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/snapshots"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	page, err := snapshots.List(client.ServiceClient(), &snapshots.ListOpts{Limit: 1}).AllPages()
	actual, err := snapshots.ExtractSnapshots(page)
	if err != nil {
		t.Errorf("Failed to extract snapshots: %v", err)

	}

	expected := []snapshots.Snapshot{
		{
			ID:          "289da7f8-6440-407c-9fb4-7db01ec49164",
			Name:        "snapshot-001",
			VolumeID:    "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
			Status:      "available",
			Size:        30,
			CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
			Description: "Daily Backup",
		},
		{
			ID:          "96c3bda7-c82a-4f50-be73-ca7621794835",
			Name:        "snapshot-002",
			VolumeID:    "76b8950a-8594-4e5b-8dce-0dfa9c696358",
			Status:      "available",
			Size:        25,
			CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
			Description: "Weekly Backup",
		},
	}

	th.CheckDeepEquals(t, expected, actual.Snapshots)

}

func TestDetail(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDetailResponse(t)

	allPages, err := snapshots.Detail(client.ServiceClient(), &snapshots.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := snapshots.ExtractSnapshots(allPages)
	th.AssertNoErr(t, err)

	expected := []snapshots.Snapshot{
		{
			ID:          "6cd26877-3ca3-4f4e-ae2a-38cc3d6183fa",
			Name:        "name_xx2-snap",
			Size:        5,
			Description: "",
			Status:      "available",
			VolumeID:    "ae11e59c-bd56-434a-a00c-04757e1c066d",
			Metadata:    map[string]string{},
			ProjectID:   "d6c277ba8820452e83df36f33c9fa561",
			Progress:    "100%",
			CreatedAt:   time.Date(2013, 6, 19, 7, 15, 29, 0000000000, time.UTC),
		},
		{
			ID:          "b3253e26-5c37-48dd-8bf2-8795dd1e848f",
			Name:        "name_xx2-snap",
			Size:        5,
			Description: "",
			Status:      "available",
			VolumeID:    "ae11e59c-bd56-434a-a00c-04757e1c066d",
			Metadata:    map[string]string{},
			ProjectID:   "d6c277ba8820452e83df36f33c9fa561",
			Progress:    "100%",
			CreatedAt:   time.Date(2013, 6, 19, 9, 8, 8, 000000000, time.UTC),
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := snapshots.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "snapshot-001")
	th.AssertEquals(t, v.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := snapshots.CreateOpts{VolumeID: "1234", Name: "snapshot-001"}
	n, err := snapshots.Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.VolumeID, "1234")
	th.AssertEquals(t, n.Name, "snapshot-001")
	th.AssertEquals(t, n.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataResponse(t)

	expected := map[string]interface{}{"key": "v1"}

	options := &snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "v1",
		},
	}

	actual, err := snapshots.UpdateMetadata(client.ServiceClient(), "123", options).ExtractMetadata()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, actual, expected)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := snapshots.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := snapshots.UpdateOpts{Name: "name_xx3", Description: "hello"}
	v, err := snapshots.Update(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "snap-001", v.Name)
}

func TestRollback(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockRollbackResponse(t)

	options := snapshots.RollbackOpts{Name: "test-001", VolumeID: "5aa119a8-d25b-45a7-8d1b-88e127885635"}
	v, err := snapshots.Rollback(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).ExtractVolumeID()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "5aa119a8-d25b-45a7-8d1b-88e127885635", v)
}

func TestGetMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataResponse(t)

	v, err := snapshots.GetMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestCreateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateMetadataResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.CreateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.UpdateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestGetMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataKeyResponse(t)

	v, err := snapshots.GetMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataKeyResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.UpdateMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestDeleteMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteMetadataKeyResponse(t)

	res := snapshots.DeleteMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key")
	th.AssertNoErr(t, res.Err)
}

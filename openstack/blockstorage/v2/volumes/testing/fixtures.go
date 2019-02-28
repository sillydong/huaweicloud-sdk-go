package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
  {
  "volumes": [
    {
      "volume_type": "lvmdriver-1",
      "created_at": "2015-09-17T03:35:03.000000",
      "bootable": "false",
      "name": "vol-001",
      "os-vol-mig-status-attr:name_id": null,
      "consistencygroup_id": null,
      "source_volid": null,
      "os-volume-replication:driver_data": null,
      "multiattach": false,
      "snapshot_id": null,
      "replication_status": "disabled",
      "os-volume-replication:extended_status": null,
      "encrypted": false,
      "os-vol-host-attr:host": null,
      "availability_zone": "nova",
      "attachments": [{
        "server_id": "83ec2e3b-4321-422b-8706-a84185f52a0a",
        "attachment_id": "05551600-a936-4d4a-ba42-79a037c1-c91a",
        "attached_at": "2016-08-06T14:48:20.000000",
        "host_name": "foobar",
        "volume_id": "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
        "device": "/dev/vdc",
        "id": "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75"
      }],
      "id": "289da7f8-6440-407c-9fb4-7db01ec49164",
      "size": 75,
      "user_id": "ff1ce52c03ab433aaba9108c2e3ef541",
      "os-vol-tenant-attr:tenant_id": "",
      "os-vol-mig-status-attr:migstat": null,
      "metadata": {"foo": "bar"},
      "status": "available",
      "description": null
    },
    {
      "volume_type": "lvmdriver-1",
      "created_at": "2015-09-17T03:32:29.000000",
      "bootable": "false",
      "name": "vol-002",
      "os-vol-mig-status-attr:name_id": null,
      "consistencygroup_id": null,
      "source_volid": null,
      "os-volume-replication:driver_data": null,
      "multiattach": false,
      "snapshot_id": null,
      "replication_status": "disabled",
      "os-volume-replication:extended_status": null,
      "encrypted": false,
      "os-vol-host-attr:host": null,
      "availability_zone": "nova",
      "attachments": [],
      "id": "96c3bda7-c82a-4f50-be73-ca7621794835",
      "size": 75,
      "user_id": "ff1ce52c03ab433aaba9108c2e3ef541",
      "os-vol-tenant-attr:tenant_id": "",
      "os-vol-mig-status-attr:migstat": null,
      "metadata": {},
      "status": "available",
      "description": null
    }
  ]
}
  `)
	})
}


func MockDetailResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "volumes": [
        {
            "attachments": [],
            "availability_zone": "xxx",
            "bootable": "false",
            "consistencygroup_id": null,
            "created_at": "2016-05-25T02:42:10.856332",
            "description": null,
            "encrypted": false,
            "id": "b104b8db-170d-441b-897a-3c8ba9c5a214",
            "links": [
                {
                    "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
                    "rel": "self"
                },
                {
                    "href": "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
                    "rel": "bookmark"
                }
            ],
            "metadata": {
                "__openstack_region_name": "pod01.xxx",
                "a": "b",
                "quantityGB": "1",
                "volInfoUrl": "fusionstorage://172.30.64.10/0/FEFEEB07D3924CDEA93C612D4E16882D"
            },
            "name": "zjb_u25_test",
            "os-vol-host-attr:host": "pod01.xxx#SATA",
            "volume_image_metadata": {},
            "os-vol-mig-status-attr:migstat": null,
            "os-vol-mig-status-attr:name_id": null,
            "os-vol-tenant-attr:tenant_id": "dd14c6ac581f40059e27f5320b60bf2f",
            "os-volume-replication:driver_data": null,
            "os-volume-replication:extended_status": null,
            "replication_status": "disabled",
            "shareable": false,
            "multiattach": false,
            "size": 1,
            "snapshot_id": null,
            "source_volid": null,
            "status": "available",
            "updated_at": "2016-05-25T02:42:22.341984",
            "user_id": "b0524e8342084ef5b74f158f78fc3049",
            "volume_type": "SATA"
        }
    ],
    "volumes_links": [
        {
            "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/detail?limit=1&marker=b104b8db-170d-441b-897a-3c8ba9c5a214",
            "rel": "next"
        }
    ]
}
  `)
	})
}


func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
  "volume": {
    "volume_type": "lvmdriver-1",
    "created_at": "2015-09-17T03:32:29.000000",
    "bootable": "false",
    "name": "vol-001",
    "os-vol-mig-status-attr:name_id": null,
    "consistencygroup_id": null,
    "source_volid": null,
    "os-volume-replication:driver_data": null,
    "multiattach": false,
    "snapshot_id": null,
    "replication_status": "disabled",
    "os-volume-replication:extended_status": null,
    "encrypted": false,
    "os-vol-host-attr:host": null,
    "availability_zone": "nova",
    "attachments": [{
      "server_id": "83ec2e3b-4321-422b-8706-a84185f52a0a",
      "attachment_id": "05551600-a936-4d4a-ba42-79a037c1-c91a",
      "attached_at": "2016-08-06T14:48:20.000000",
      "host_name": "foobar",
      "volume_id": "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
      "device": "/dev/vdc",
      "id": "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75"
    }],
    "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
    "size": 75,
    "user_id": "ff1ce52c03ab433aaba9108c2e3ef541",
    "os-vol-tenant-attr:tenant_id": "304dc00909ac4d0da6c62d816bcb3459",
    "os-vol-mig-status-attr:migstat": null,
    "metadata": {},
    "status": "available",
    "description": null
  }
}
      `)
	})
}

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "volume": {
    	"name": "vol-001",
        "size": 75,
		"availability_zone": "nova"
	}
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
  "volume": {
    "size": 75,
    "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
    "metadata": {},
    "created_at": "2015-09-17T03:32:29.044216",
    "encrypted": false,
    "bootable": "false",
    "availability_zone": "nova",
    "attachments": [],
    "user_id": "ff1ce52c03ab433aaba9108c2e3ef541",
    "status": "creating",
    "description": null,
    "volume_type": "lvmdriver-1",
    "name": "vol-001",
    "replication_status": "disabled",
    "consistencygroup_id": null,
    "source_volid": null,
    "snapshot_id": null,
    "multiattach": false
  }
}
    `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockUpdateResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
  "volume": {
    "name": "vol-002"
  }
}
        `)
	})
}

func MockCreateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockUpdateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockGetMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockGetMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockUpdateMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockDeleteMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockExtendSizeResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-extend":
    {
        "new_size": 3
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockSetBootableResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-set_bootable":
    {
        "bootable": true
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockSetReadOnlyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-update_readonly_flag":
    {
        "readonly": true
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

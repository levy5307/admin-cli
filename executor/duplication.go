package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/XiaoMi/pegasus-go-client/idl/admin"
)

// QueryDuplication command
func QueryDuplication(c *Client, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.meta.QueryDuplication(ctx, &admin.DuplicationQueryRequest{
		AppName: tableName,
	})
	if err != nil {
		return err
	}
	// formats into JSON
	outputBytes, err := json.MarshalIndent(resp.EntryList, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintln(c, string(outputBytes))
	return nil
}

// AddDuplication command
func AddDuplication(c *Client, tableName string, remoteCluster string, freezed bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.meta.AddDuplication(ctx, &admin.DuplicationAddRequest{
		AppName:           tableName,
		RemoteClusterName: remoteCluster,
		Freezed:           freezed,
	})
	if err != nil {
		// TODO(wutao): print error hints if it has.
		return err
	}
	fmt.Fprintf(c, "successfully add duplication [dupid: %d]\n", resp.Dupid)
	return nil
}

// ModifyDuplication command
func ModifyDuplication(c *Client, tableName string, dupid int, status admin.DuplicationStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := c.meta.ModifyDuplication(ctx, &admin.DuplicationModifyRequest{
		AppName: tableName,
		Dupid:   int32(dupid),
		Status:  &status,
	})
	if err != nil {
		return err
	}
	return nil
}
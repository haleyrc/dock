package dock

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Client struct {
c *client.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Client{c: c}, nil
}

func (c *Client) Clean(ctx context.Context) error {
	containers, err := c.c.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}

	for _, container := range containers {
		log.Printf("Removing container %s...\n", container.ID)
		if err := c.c.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
			return err
		}
	}

	images, err := c.c.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return err
	}

	for _, image := range images {
		log.Printf("Removing image %s...\n", image.ID)
		if _, err := c.c.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{Force: true, PruneChildren: true}); err != nil {
			return err
		}
	}

	return nil
}

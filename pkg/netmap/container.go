package netmap

// ContainerNodes represents nodes in the container.
type ContainerNodes interface {
	Replicas() []Nodes
	Flatten() Nodes
}

type containerNodes []Nodes

// Flatten returns list of all nodes from the container.
func (c containerNodes) Flatten() Nodes {
	return flattenNodes(c)
}

// Replicas return list of container replicas.
func (c containerNodes) Replicas() []Nodes {
	return c
}

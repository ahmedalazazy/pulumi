package engine

import (
	"io"

	"github.com/pulumi/pulumi/pkg/resource/deploy"
	"github.com/pulumi/pulumi/pkg/workspace"
)

// SnapshotManager is responsible for maintaining the in-memory representation
// of the current state of the resource world.
type SnapshotManager interface {
	io.Closer

	// BeginMutation signals to the SnapshotManager that the planner intends to mutate the global
	// snapshot. It provides the step that it intends to execute. Based on that step, BeginMutation
	// will record this intent in the global snapshot and return a `SnapshotMutation` that, when ended,
	// will complete the transaction.
	BeginMutation(step deploy.Step) (SnapshotMutation, error)

	// RegisterResourceOutputs registers the set of resource outputs generated by performing the
	// given step. These outputs are persisted in the snapshot.
	RegisterResourceOutputs(step deploy.Step) error

	// RecordPlugin records that the current plan loaded a plugin and saves it in the snapshot.
	RecordPlugin(plugin workspace.PluginInfo) error
}

// SnapshotMutation represents an outstanding mutation that is yet to be completed. When the engine completes
// a mutation, it must call `End` in order to record the successful completion of the mutation.
type SnapshotMutation interface {
	// End terminates the transaction and commits the results to the snapshot, returning an error if this
	// failed to complete.
	End(step deploy.Step, successful bool) error
}

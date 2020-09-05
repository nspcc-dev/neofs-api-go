/*
Package netmap provides routines for working with netmap and placement policy.
Work is done in 4 steps:
1. Create context containing results shared between steps.
2. Processing filters.
3. Processing selectors.
4. Processing replicas.

Each step depends only on previous ones.
*/
package netmap

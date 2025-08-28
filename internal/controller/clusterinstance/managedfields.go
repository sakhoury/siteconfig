/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusterinstance

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ManagedFieldsHelper provides utilities for working with Kubernetes managedFields metadata.
// This replaces the previous approach of using a custom "last-applied" annotation with
// Kubernetes' native field ownership tracking mechanism.
type ManagedFieldsHelper struct {
	log *zap.Logger
}

// NewManagedFieldsHelper creates a new instance of ManagedFieldsHelper.
func NewManagedFieldsHelper(log *zap.Logger) *ManagedFieldsHelper {
	return &ManagedFieldsHelper{
		log: log.Named("ManagedFieldsHelper"),
	}
}

// FieldOwnership represents which field manager owns specific fields in an object.
type FieldOwnership struct {
	Manager   string
	Operation string
	Fields    map[string]interface{}
}

// GetFieldOwnership extracts field ownership information from an object's managedFields.
// This function parses the managedFields metadata to understand which controllers or users
// manage specific fields, particularly useful for labels and annotations.
func (mfh *ManagedFieldsHelper) GetFieldOwnership(obj *unstructured.Unstructured) ([]FieldOwnership, error) {
	managedFields := obj.GetManagedFields()
	var ownerships []FieldOwnership

	for _, field := range managedFields {
		if field.FieldsV1 == nil || field.FieldsV1.Raw == nil {
			continue
		}

		// Parse the FieldsV1 JSON to understand field structure
		var fieldsMap map[string]interface{}
		if err := json.Unmarshal(field.FieldsV1.Raw, &fieldsMap); err != nil {
			mfh.log.Warn("Failed to unmarshal FieldsV1",
				zap.String("manager", field.Manager),
				zap.Error(err),
			)
			continue
		}

		ownership := FieldOwnership{
			Manager:   field.Manager,
			Operation: string(field.Operation),
			Fields:    fieldsMap,
		}

		ownerships = append(ownerships, ownership)
	}

	return ownerships, nil
}

// IsFieldManagedBy checks if a specific field path is managed by the given field manager.
// This is useful for determining if the ClusterInstance controller should manage a particular
// label or annotation, or if it's managed by another controller.
func (mfh *ManagedFieldsHelper) IsFieldManagedBy(obj *unstructured.Unstructured, fieldPath string, manager string) bool {
	ownerships, err := mfh.GetFieldOwnership(obj)
	if err != nil {
		mfh.log.Warn("Failed to get field ownership", zap.Error(err))
		return false
	}

	for _, ownership := range ownerships {
		if ownership.Manager == manager {
			// Check if the field path exists in this manager's fields
			// This is a simplified check - a more sophisticated implementation
			// would parse the field path structure more carefully
			if mfh.fieldExists(ownership.Fields, fieldPath) {
				return true
			}
		}
	}

	return false
}

// GetManagedMetadataKeys returns the metadata keys (labels and annotations) that are
// managed by the specified field manager. This helps understand what the ClusterInstance
// controller is currently responsible for managing.
func (mfh *ManagedFieldsHelper) GetManagedMetadataKeys(obj *unstructured.Unstructured, manager string) (labels, annotations []string) {
	ownerships, err := mfh.GetFieldOwnership(obj)
	if err != nil {
		mfh.log.Warn("Failed to get field ownership", zap.Error(err))
		return nil, nil
	}

	for _, ownership := range ownerships {
		if ownership.Manager == manager {
			labels = append(labels, mfh.extractMetadataKeys(ownership.Fields, "f:metadata", "f:labels")...)
			annotations = append(annotations, mfh.extractMetadataKeys(ownership.Fields, "f:metadata", "f:annotations")...)
		}
	}

	return labels, annotations
}

// LogFieldOwnership provides debug logging of field ownership information.
// This is useful for troubleshooting field management issues.
func (mfh *ManagedFieldsHelper) LogFieldOwnership(obj *unstructured.Unstructured) {
	ownerships, err := mfh.GetFieldOwnership(obj)
	if err != nil {
		mfh.log.Error("Failed to get field ownership for logging", zap.Error(err))
		return
	}

	mfh.log.Debug("Object field ownership",
		zap.String("name", obj.GetName()),
		zap.String("namespace", obj.GetNamespace()),
		zap.String("kind", obj.GetKind()),
		zap.Int("managerCount", len(ownerships)),
	)

	for _, ownership := range ownerships {
		mfh.log.Debug("Field manager details",
			zap.String("manager", ownership.Manager),
			zap.String("operation", ownership.Operation),
		)
	}
}

// fieldExists is a helper function to check if a field path exists in the fields map.
// This is a simplified implementation that could be enhanced based on specific needs.
func (mfh *ManagedFieldsHelper) fieldExists(fields map[string]interface{}, fieldPath string) bool {
	// This is a basic implementation - in practice, you might want to implement
	// a more sophisticated field path parser based on the Kubernetes field path format
	_, exists := fields[fieldPath]
	return exists
}

// extractMetadataKeys extracts keys from nested field maps for metadata fields.
// This helper function navigates the FieldsV1 structure to find label and annotation keys.
func (mfh *ManagedFieldsHelper) extractMetadataKeys(fields map[string]interface{}, path ...string) []string {
	var keys []string

	current := fields
	for _, segment := range path {
		if next, ok := current[segment].(map[string]interface{}); ok {
			current = next
		} else {
			return keys // Path doesn't exist
		}
	}

	// Extract keys from the final map level
	for key := range current {
		if key != "." { // Skip the special "." key used in FieldsV1
			// Remove the "f:" prefix that Kubernetes uses in field paths
			if len(key) > 2 && key[:2] == "f:" {
				keys = append(keys, key[2:])
			} else {
				keys = append(keys, key)
			}
		}
	}

	return keys
}

// ValidateFieldManager checks if the given field manager name is valid and unique enough
// to avoid conflicts with other controllers.
func ValidateFieldManager(manager string) error {
	if manager == "" {
		return fmt.Errorf("field manager name cannot be empty")
	}

	// Check for common field manager names that might cause conflicts
	conflictingManagers := []string{
		"kubectl",
		"kubectl-client-side-apply",
		"argocd-controller",
		"controller-manager",
	}

	for _, conflicting := range conflictingManagers {
		if manager == conflicting {
			return fmt.Errorf("field manager name '%s' conflicts with well-known manager", manager)
		}
	}

	return nil
}
